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

func Test_sat_formatRegion(t *testing.T) {
	type fields struct {
		byRegion    bool
		withStorage bool
		filter      model.FilterMap
		size        string
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
			fields{true, false, model.FilterMap{}, "KB"},
			args{*objectResultMock},
			bucketNameMock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &sat{
				byRegion:    tt.fields.byRegion,
				withStorage: tt.fields.withStorage,
				filter:      tt.fields.filter,
				size:        tt.fields.size,
			}
			if got := sat.formatRegion(tt.args.data); strings.Contains(got, tt.without) {
				t.Errorf("sat.formatRegion() = %v, has %v", got, tt.without)
			}
		})
	}
}

func Test_sat_formatBucket(t *testing.T) {
	type fields struct {
		byRegion    bool
		withStorage bool
		filter      model.FilterMap
		size        string
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
			fields{true, false, model.FilterMap{}, "KB"},
			args{*objectResultMock},
			bucketNameMock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &sat{
				byRegion:    tt.fields.byRegion,
				withStorage: tt.fields.withStorage,
				filter:      tt.fields.filter,
				size:        tt.fields.size,
			}
			if got := sat.formatBucket(tt.args.data); !strings.Contains(got, tt.expect) {
				t.Errorf("sat.formatBucket() = %v, does not have %v", got, tt.expect)
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

func Test_sat_Print(t *testing.T) {
	type fields struct {
		byRegion    bool
		withStorage bool
		filter      model.FilterMap
		size        string
	}
	type args struct {
		result *Result
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"When Print without flags set, it should print the expected output",
			fields{false, false, model.FilterMap{}, "KB"},
			args{resultBucketMock},
			expectedBucketOutput},
		{"When Print with byRegion set, it should print the expected output",
			fields{true, false, model.FilterMap{}, "KB"},
			args{resultBucketMock},
			expectedRegionOutput},
		{"When Print with withStorage set, it should print the expected output",
			fields{false, true, model.FilterMap{}, "KB"},
			args{resultBucketWithStorageMock},
			expectedBucketWithStorageOutput},
		{"When Print with byRegion and withStorage set, it should print the expected output",
			fields{true, true, model.FilterMap{}, "KB"},
			args{resultBucketWithStorageMock},
			expectedRegionWithStorageOutput},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &sat{
				byRegion:    tt.fields.byRegion,
				withStorage: tt.fields.withStorage,
				filter:      tt.fields.filter,
				size:        tt.fields.size,
			}
			output := &bytes.Buffer{}
			sat.Print(output, tt.args.result)
			got := output.String()
			if got != tt.want {
				t.Errorf("Print() = \n%v \nwant: \n%v", got, tt.want)
			}
		})
	}
}
