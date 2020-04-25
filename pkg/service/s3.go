package service

import (
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Service is responsible by the connection to s3 and the operations performed on s3
type Service struct {
	*s3.Client
	s3Operations
}

// New creates a S3 service using aws configuration. ~/.aws/credentials, environment variables, ...
func New(config aws.Config) *Service {
	s := s3.New(config)
	o := &operation{s}
	return &Service{s, o}
}

// List all objects from S3 using the filter
func (s *Service) Objects(filter model.FilterMap) ([]*model.ObjectData, error) {

	buckets, err := s.listBuckets()
	if err != nil {
		return nil, err
	}

	var objects []*model.ObjectData
	for _, bucket := range buckets {
		if _, ok := filter[*bucket.Name]; !ok && len(filter) > 0 {
			continue
		}
		objs, err := s.bucketObjects(&bucket, filter[*bucket.Name])
		if err != nil {
			return nil, err
		}
		objects = append(objects, objs...)
	}

	return objects, nil
}

func (s *Service) bucketObjects(bucket *s3.Bucket, prefix string) ([]*model.ObjectData, error) {

	region, err := s.getRegion(*bucket.Name)
	if err != nil {
		return nil, err
	}
	s.Config.Region = region

	s3Objects, err := s.listObjects(*bucket.Name, prefix)
	if err != nil {
		return nil, err
	}

	return convert(s3Objects, bucket, region), nil
}

func convert(objects []s3.Object, bucket *s3.Bucket, region string) []*model.ObjectData {

	objDatas := make([]*model.ObjectData, len(objects))
	for i, obj := range objects {
		storage := string(obj.StorageClass)
		objDatas[i] = &model.ObjectData{Bucket: bucket.Name, CreationDate: bucket.CreationDate, Region: &region, Key: obj.Key,
			LastModified: obj.LastModified, Size: obj.Size, StorageClass: &storage}
	}

	return objDatas
}
