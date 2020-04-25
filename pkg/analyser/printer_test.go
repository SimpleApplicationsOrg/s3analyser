package analyser

import (
	"bytes"
	"strings"
	"testing"

	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
)

const KB = int64(1024)
const MB = int64(1048576)
const GB = int64(1073741824)
const TB = int64(1099511627776)
const expected = float64(1)

const expectedBucketOutput = "Bucket  Region     Count  Total (KB)  Creation                                 Last Modified\ntest    us-east-1  1      1           2018-01-01 01:01:01.000000001 +0000 UTC  2018-01-01 01:01:01.000000001 +0000 UTC  \n\n"
const expectedRegionOutput = "Region     Count  Total (KB)  Creation                                 Last Modified\nus-east-1  1      1           2018-01-01 01:01:01.000000001 +0000 UTC  2018-01-01 01:01:01.000000001 +0000 UTC  \n\n"
const expectedBucketWithStorageOutput = "Bucket  Region     Count  Total (KB)  Creation                                 Last Modified                            Storage\ntest    us-east-1  1      1           2018-01-01 01:01:01.000000001 +0000 UTC  2018-01-01 01:01:01.000000001 +0000 UTC  STANDARD\n\n"
const expectedRegionWithStorageOutput = "Region     Count  Total (KB)  Creation                                 Last Modified                            Storage\nus-east-1  1      1           2018-01-01 01:01:01.000000001 +0000 UTC  2018-01-01 01:01:01.000000001 +0000 UTC  STANDARD\n\n"

func Test_formatRegion(t *testing.T) {
	type fields struct {
		result *Result
	}
	type args struct {
		data model.ObjectData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		without string
	}{
		{"When formatRegion is called the result should not have bucket name",
			fields{result: &Result{
				byRegion:    true,
				withStorage: false,
				size:        "KB",
				objects:     resultBucketMock.objects,
			}},
			args{objectResultMock},
			bucketNameMock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fields.result
			if got := result.formatRegion(tt.args.data); strings.Contains(got, tt.without) {
				t.Errorf("Analyser.formatRegion() = %v, has %v", got, tt.without)
			}
		})
	}
}

func Test_formatBucket(t *testing.T) {
	type fields struct {
		result *Result
	}
	type args struct {
		data model.ObjectData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expect string
	}{
		{"When formatRegion is called the result should not have bucket name",
			fields{result: &Result{
				byRegion:    true,
				withStorage: false,
				size:        "KB",
				objects:     resultBucketMock.objects,
			}},
			args{objectResultMock},
			bucketNameMock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fields.result
			if got := result.formatBucket(tt.args.data); !strings.Contains(got, tt.expect) {
				t.Errorf("Analyser.formatBucket() = %v, does not have %v", got, tt.expect)
			}
		})
	}
}

func Test_sizeCalc(t *testing.T) {
	type args struct {
		size   int64
		format string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"When converting KB should return the expected result",
			args{KB, "KB"},
			expected},
		{"When converting MB should return the expected result",
			args{MB, "MB"},
			expected},
		{"When converting GB should return the expected result",
			args{GB, "GB"},
			expected},
		{"When converting TB should return the expected result",
			args{TB, "TB"},
			expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sizeCalc(tt.args.size, tt.args.format); got != tt.want {
				t.Errorf("sizeCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Print(t *testing.T) {
	type fields struct {
		result *Result
	}
	type args struct {
		writer *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"When Print without flags set, it should print the expected output",
			fields{result: &Result{
				byRegion:    false,
				withStorage: false,
				size:        "KB",
				objects:     resultBucketMock.objects,
			}},
			args{&bytes.Buffer{}},
			expectedBucketOutput},
		{"When Print with byRegion set, it should print the expected output",
			fields{result: &Result{
				byRegion:    true,
				withStorage: false,
				size:        "KB",
				objects:     resultBucketMock.objects,
			}},
			args{&bytes.Buffer{}},
			expectedRegionOutput},
		{"When Print with withStorage set, it should print the expected output",
			fields{result: &Result{
				byRegion:    false,
				withStorage: true,
				size:        "KB",
				objects:     resultBucketWithStorageMock.objects,
			}},
			args{&bytes.Buffer{}},
			expectedBucketWithStorageOutput},
		{"When Print with byRegion and withStorage set, it should print the expected output",
			fields{result: &Result{
				byRegion:    true,
				withStorage: true,
				size:        "KB",
				objects:     resultBucketWithStorageMock.objects,
			}},
			args{&bytes.Buffer{}},
			expectedRegionWithStorageOutput},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fields.result
			writer := tt.args.writer
			result.PrintTo(writer)
			got := writer.String()
			if got != tt.want {
				t.Errorf("Print() = \n%v \nwant: \n%v", got, tt.want)
			}
		})
	}
}
