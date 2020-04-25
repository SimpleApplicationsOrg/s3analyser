package analyser

import (
	"errors"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

var bucketNameMock = "test"
var creationDateMock = time.Date(2018, time.January, 1, 1, 1, 1, 1, time.UTC)
var sizeMock = int64(1000)
var lastModifiedMock = time.Date(2018, time.January, 1, 1, 1, 1, 1, time.UTC)
var regionMock = "us-east-1"
var storageMock = s3.ObjectStorageClassStandard
var storageStringMock = string(storageMock)
var countMock = 1

var objectMock = &model.ObjectData{Bucket: &bucketNameMock,
	Region:       &regionMock,
	Size:         &sizeMock,
	CreationDate: &creationDateMock,
	LastModified: &lastModifiedMock,
	StorageClass: &storageStringMock}

var objectResultMock = &model.ObjectData{Bucket: &bucketNameMock,
	Region:       &regionMock,
	Size:         &sizeMock,
	CreationDate: &creationDateMock,
	LastModified: &lastModifiedMock,
	StorageClass: &blank,
	Count:        &countMock}

var objectResultWithStorageMock = &model.ObjectData{Bucket: &bucketNameMock,
	Region:       &regionMock,
	Size:         &sizeMock,
	CreationDate: &creationDateMock,
	LastModified: &lastModifiedMock,
	StorageClass: &storageStringMock,
	Count:        &countMock}

var resultBucketMock = &Result{map[string]*model.ObjectData{bucketNameMock: objectResultMock}}
var resultBucketWithStorageMock = &Result{map[string]*model.ObjectData{bucketNameMock: objectResultWithStorageMock}}

type serviceMock struct{}

func (s *serviceMock) Objects(filter model.FilterMap) ([]*model.ObjectData, error) {
	return []*model.ObjectData{objectMock}, nil
}

type serviceErrorMock struct{}

func (s *serviceErrorMock) Objects(filter model.FilterMap) ([]*model.ObjectData, error) {
	return nil, errors.New("func Objects error message")
}