package service

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Operations interface {
	listBuckets(scv svc) ([]s3.Bucket, error)
	listObjects(svc svc, bucketName string, prefix string) ([]s3.Object, error)
	getRegion(svc svc, bucketName string) (string, error)
}

type operations struct {}

func (o *operations) listBuckets(svc svc) ([]s3.Bucket, error) {
	req := svc.ListBucketsRequest(&s3.ListBucketsInput{})
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}
	return resp.Buckets, nil
}

func (o *operations) listObjects(svc svc, bucketName string, prefix string) ([]s3.Object, error) {
	objects := make([]s3.Object,0)
	err := svc.ListObjectsPages(&s3.ListObjectsInput{Bucket: &bucketName, Prefix: &prefix},
		func(page *s3.ListObjectsOutput, morePages bool) bool {
			if len(page.Contents) == 0 {
				return false
			}
			objects = append(objects, page.Contents...)
			return true
		})
	return objects, err
}

func (o *operations) getRegion(svc svc, bucketName string) (string, error) {
	req := svc.GetBucketLocationRequest(&s3.GetBucketLocationInput{Bucket: &bucketName})
	req.Handlers.Unmarshal.PushBackNamed(s3.NormalizeBucketLocationHandler)
	resp, err := req.Send()
	if err != nil {
		return "", err
	}
	return string(resp.LocationConstraint), nil
}

