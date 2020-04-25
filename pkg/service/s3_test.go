package service

import (
	"reflect"
	"testing"

	"fmt"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var expectedOject = model.ObjectData{Bucket: &bucketNameMock,
	Region:       &regionMock,
	Size:         &sizeMock,
	CreationDate: &creationDateMock,
	LastModified: &lastModifiedMock,
	StorageClass: &storageStringMock}

func Test_svc_Objects(t *testing.T) {
	type fields struct {
		S3           *s3.Client
		s3Operations s3Operations
	}
	type args struct {
		filter model.FilterMap
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ObjectData
		wantErr bool
	}{
		//Test cases
		{"When Objects is called it should return the expected list of objects",
			fields{s3.New(aws.Config{}), &operationsMock{}}, // mocked s3 operations
			args{model.FilterMap{}},                         // no filter
			[]*model.ObjectData{&expectedOject},             // expected list of objects
			false}, // no error is expected
		{"When operation listBuckets fails Objects should return an error",
			fields{s3.New(aws.Config{}), &operationsErrorMock{}}, // mocked s3 operations with errors
			args{model.FilterMap{}}, // no filter
			[]*model.ObjectData{},   // no expected result
			true}, // error is expected
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				Client:           tt.fields.S3,
				s3Operations: tt.fields.s3Operations,
			}
			got, err := svc.Objects(tt.args.filter)
			if (err != nil) != tt.wantErr && err.Error() != "list buckets error message" {
				t.Errorf("Service.Objects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Objects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_svc_bucketObjects(t *testing.T) {
	type fields struct {
		S3           *s3.Client
		s3Operations s3Operations
	}
	type args struct {
		bucket *s3.Bucket
		prefix string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ObjectData
		wantErr bool
	}{
		// Test cases
		{"When bucketObjects is called it should return expected list of objects",
			fields{s3.New(aws.Config{}), &operationsMock{}}, // mocked s3 operations
			args{&bucketMock, prefixMock},                   // mocked filter
			[]*model.ObjectData{&expectedOject},             // expected list of objects
			false}, // no error is expected
		{"When operation getRegion returns an error bucketObjects should return an error",
			fields{s3.New(aws.Config{}), &operationsErrorMock{}}, // mocked s3 operations with error
			args{&bucketMock, prefixMock},                        // mocked filter
			[]*model.ObjectData{},                                // no expected return
			true}, // error is expected
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				Client:           tt.fields.S3,
				s3Operations: tt.fields.s3Operations,
			}
			got, err := svc.bucketObjects(tt.args.bucket, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.bucketObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr && fmt.Sprint(err) != "get region error message" {
				t.Errorf("Service.bucketObjects() error = %v, want %v", err, "get region error message")
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.bucketObjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		objects []s3.Object
		bucket  *s3.Bucket
		region  string
	}
	tests := []struct {
		name string
		args args
		want []*model.ObjectData
	}{
		{"convert should return expected list of objects",
			args{[]s3.Object{objectMock}, &bucketMock, regionMock},
			[]*model.ObjectData{&expectedOject}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.objects, tt.args.bucket, tt.args.region); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
