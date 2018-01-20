package service

import (
	"github.com/SimpleApplicationsOrg/s3analyser/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	service "github.com/aws/aws-sdk-go-v2/service/s3"
)

type svc struct {
	*service.S3
	s3Operations
}

// S3 is used to access all s3 objects
type S3 interface {
	Objects(filter model.FilterMap) ([]*model.ObjectData, error)
}

// S3Factory creates a S3 service using aws configuration. ~/.aws/credentials, environment variables, ...
func S3Factory(config aws.Config) S3 {
	return &svc{service.New(config), &operations{}}
}

// List all objects from S3 using the filter
func (svc *svc) Objects(filter model.FilterMap) ([]*model.ObjectData, error) {
	o := svc.s3Operations

	buckets, err := o.listBuckets(*svc)
	if err != nil {
		return nil, err
	}

	var objects []*model.ObjectData
	for _, bucket := range buckets {
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
	o := svc.s3Operations

	region, err := o.getRegion(*svc, *bucket.Name)
	if err != nil {
		return nil, err
	}
	svc.Config.Region = region

	s3Objects, err := o.listObjects(*svc, *bucket.Name, prefix)
	if err != nil {
		return nil, err
	}

	return convert(s3Objects, bucket, region), nil
}

func convert(objects []service.Object, bucket *service.Bucket, region string) []*model.ObjectData {

	objDatas := make([]*model.ObjectData, len(objects))
	for i, obj := range objects {
		storage := string(obj.StorageClass)
		objDatas[i] = &model.ObjectData{Bucket: bucket.Name, CreationDate: bucket.CreationDate, Region: &region, Key: obj.Key,
			LastModified: obj.LastModified, Size: obj.Size, StorageClass: &storage}
	}

	return objDatas
}
