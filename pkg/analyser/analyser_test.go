package analyser

import (
	"reflect"
	"testing"

	"fmt"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
)

func Test_sat_Analyse(t *testing.T) {
	type fields struct {
		byRegion    bool
		withStorage bool
		filter      model.FilterMap
		size        string
	}
	type args struct {
		s3 S3
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Result
		wantErr bool
	}{
		{"When Analyse is called it should return the expected result",
			fields{false, false, model.FilterMap{}, "KB"},
			args{&serviceMock{}},
			resultBucketMock,
			false},
		{"When Analyse is called it should return the expected result",
			fields{false, false, model.FilterMap{}, "KB"},
			args{&serviceErrorMock{}},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &Analyser{
				byRegion:    tt.fields.byRegion,
				withStorage: tt.fields.withStorage,
				filter:      tt.fields.filter,
				size:        tt.fields.size,
			}
			got, err := sat.Analyse(tt.args.s3)
			if (err != nil) != tt.wantErr {
				t.Errorf("Analyser.Analyse() error = %v, wantErr %v", err, tt.wantErr)
				return
			} //func Objects error message
			if (err != nil) && tt.wantErr && fmt.Sprint(err) != "func Objects error message" {
				t.Errorf("Analyser.Analyse() error = %v, wantErr %v", err, "func Objects error message")
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyser.Analyse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sat_key(t *testing.T) {
	type fields struct {
		byRegion    bool
		withStorage bool
		filter      model.FilterMap
		size        string
	}
	type args struct {
		object model.ObjectData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"When no flag is passed, key should return the bucket name as key",
			fields{false, false, model.FilterMap{}, "KB"},
			args{*objectMock},
			bucketNameMock},
		{"When byRegion is passed, key should return the bucket name as key",
			fields{true, false, model.FilterMap{}, "KB"},
			args{*objectMock},
			regionMock},
		{"When only withStorage is passed, key should return the bucket name and storage as key",
			fields{false, true, model.FilterMap{}, "KB"},
			args{*objectMock},
			bucketNameMock + storageStringMock},
		{"When byRegion and withStorage are passed, key should return the region and storage as key",
			fields{true, true, model.FilterMap{}, "KB"},
			args{*objectMock},
			regionMock + storageStringMock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &Analyser{
				byRegion:    tt.fields.byRegion,
				withStorage: tt.fields.withStorage,
				filter:      tt.fields.filter,
				size:        tt.fields.size,
			}
			if got := sat.key(tt.args.object); got != tt.want {
				t.Errorf("Analyser.key() = %v, want %v", got, tt.want)
			}
		})
	}
}
