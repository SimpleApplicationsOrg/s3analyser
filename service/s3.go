package service

import (
	"github.com/SimpleApplicationsOrg/s3analyser/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	service "github.com/aws/aws-sdk-go-v2/service/s3"
)

type svc struct {
	*service.S3
}

// S3 is used to access all s3 objects
type S3 interface {
	Objects(filter model.FilterMap) ([]*model.ObjectData, error)
}

// S3Factory creates a S3 service using aws configuration. ~/.aws/credentials, environment variables, ...
func S3Factory(config aws.Config) S3 {
	return &svc{service.New(config)}
}

// List all objects from S3 using the filter
func (svc *svc) Objects(filter model.FilterMap) ([]*model.ObjectData, error) {

	req := svc.ListBucketsRequest(&service.ListBucketsInput{})

	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	var objects []*model.ObjectData

	for _, bucket := range resp.Buckets {

		if _, ok := filter[*bucket.Name]; !ok && len(filter) > 0 {
			continue
		}

		objs, err := svc.bucketObjects(&bucket, filter[*bucket.Name])
		if err != nil {
			return nil, err
		}
		objects = append(objects, objs...)
	}

	return objects, nil
}

func (svc *svc) bucketObjects(bucket *service.Bucket, prefix string) ([]*model.ObjectData, error) {

	region := svc.getRegion(bucket.Name)

	svc.Config.Region = region

	var objects []*model.ObjectData

	err := svc.ListObjectsPages(&service.ListObjectsInput{Bucket: bucket.Name, Prefix: &prefix},
		func(page *service.ListObjectsOutput, morePages bool) bool {

			if len(page.Contents) == 0 {
				return false
			}

			slice := slice(page, bucket, region)

			objects = append(objects, slice...)

			return true
		})

	if err != nil {
		return nil, err
	}

	return objects, nil
}

func (svc *svc) getRegion(bucketName *string) string {

	req := svc.GetBucketLocationRequest(&service.GetBucketLocationInput{Bucket: bucketName})
	req.Handlers.Unmarshal.PushBackNamed(service.NormalizeBucketLocationHandler)
	resp, err := req.Send()

	if err != nil {
		return ""
	}

	return string(resp.LocationConstraint)

}

func slice(page *service.ListObjectsOutput, bucket *service.Bucket, region string) []*model.ObjectData {

	slice := make([]*model.ObjectData, len(page.Contents))
	for i, obj := range page.Contents {
		storage := string(obj.StorageClass)
		slice[i] = &model.ObjectData{Bucket: bucket.Name, CreationDate: bucket.CreationDate, Region: &region, Key: obj.Key,
			LastModified: obj.LastModified, Size: obj.Size, StorageClass: &storage}
	}

	return slice
}
