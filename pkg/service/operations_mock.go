package service

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

var bucketNameMock = "test"
var creationDateMock = time.Now()
var sizeMock = int64(1000)
var lastModifiedMock = time.Now()
var storageMock = s3.ObjectStorageClassStandard
var storageStringMock = string(storageMock)
var regionMock = "us-east-1"
var prefixMock = "test"

var bucketMock = s3.Bucket{Name: &bucketNameMock, CreationDate: &creationDateMock}
var objectMock = s3.Object{Size: &sizeMock, LastModified: &lastModifiedMock, StorageClass: storageMock}

type operationsMock struct{}

func (o *operationsMock) listBuckets() ([]s3.Bucket, error) {
	bucket := s3.Bucket{Name: &bucketNameMock, CreationDate: &creationDateMock}
	return []s3.Bucket{bucket}, nil
}

func (o *operationsMock) listObjects(bucketName string, prefix string) ([]s3.Object, error) {
	object := s3.Object{Size: &sizeMock, LastModified: &lastModifiedMock, StorageClass: storageMock}
	return []s3.Object{object}, nil
}

func (o *operationsMock) getRegion(bucketName string) (string, error) {
	return regionMock, nil
}

type operationsErrorMock struct{}

func (o *operationsErrorMock) listBuckets() ([]s3.Bucket, error) {
	return nil, errors.New("list buckets error message")
}

func (o *operationsErrorMock) listObjects(bucketName string, prefix string) ([]s3.Object, error) {
	return nil, errors.New("list objects error message")
}

func (o *operationsErrorMock) getRegion(bucketName string) (string, error) {
	return "", errors.New("get region error message")
}
