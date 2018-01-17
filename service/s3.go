package service

import (
	service "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/SimpleApplicationsOrg/s3analyser/model"
)

type svc struct {
	*service.S3
}

func (svc *svc) ListBuckets() (*[]model.Bucket, error) {

	req := svc.ListBucketsRequest(&service.ListBucketsInput{})

	resp, err := req.Send()
	if err != nil {
		return nil, err
	}

	buckets := make([]model.Bucket, len(resp.Buckets))
	for i, bucket := range resp.Buckets {
		buckets[i] = model.Bucket{Name: bucket.Name, CreationDate: bucket.CreationDate}
	}

	return &buckets, nil
}

func S3Factory(config aws.Config) model.S3 {
	return &svc{service.New(config)}
}
