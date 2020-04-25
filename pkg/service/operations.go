package service

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Operations interface {
	listBuckets() ([]s3.Bucket, error)
	listObjects(bucketName string, prefix string) ([]s3.Object, error)
	getRegion(bucketName string) (string, error)
}

type operation struct {
	*s3.Client
}

func (o *operation) listBuckets() ([]s3.Bucket, error) {
	req := o.ListBucketsRequest(&s3.ListBucketsInput{})
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	return resp.Buckets, nil
}

func (o *operation) listObjects(bucketName string, prefix string) ([]s3.Object, error) {
	req := o.ListObjectsRequest(&s3.ListObjectsInput{Bucket: &bucketName, Prefix: &prefix})
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	return resp.Contents, err
}

func (o *operation) getRegion(bucketName string) (string, error) {
	req := o.GetBucketLocationRequest(&s3.GetBucketLocationInput{Bucket: &bucketName})
	req.Handlers.Unmarshal.PushBackNamed(s3.NormalizeBucketLocationHandler)
	resp, err := req.Send(context.Background())
	if err != nil {
		return "", err
	}
	return string(resp.LocationConstraint), nil
}
