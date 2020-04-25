package analyser

import (
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

var objectMock = model.ObjectData{Bucket: bucketNameMock,
	Region:       regionMock,
	Size:         sizeMock,
	CreationDate: creationDateMock,
	LastModified: lastModifiedMock,
	StorageClass: storageStringMock}

var objectResultMock = model.ObjectData{Bucket: bucketNameMock,
	Region:       regionMock,
	Size:         sizeMock,
	CreationDate: creationDateMock,
	LastModified: lastModifiedMock,
	StorageClass: blank,
	Count:        countMock}

var objectResultWithStorageMock = model.ObjectData{Bucket: bucketNameMock,
	Region:       regionMock,
	Size:         sizeMock,
	CreationDate: creationDateMock,
	LastModified: lastModifiedMock,
	StorageClass: storageStringMock,
	Count:        countMock}

var objectsMock = []model.ObjectData{objectResultMock}
var resultBucketMock = &Result{false, false, "KB", map[string]model.ObjectData{bucketNameMock: objectResultMock}}

var resultBucketWithStorageMock = &Result{false, false, "KB",map[string]model.ObjectData{bucketNameMock: objectResultWithStorageMock}}
