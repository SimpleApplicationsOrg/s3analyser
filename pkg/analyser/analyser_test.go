package analyser

import (
	"reflect"
	"testing"

	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
)

func Test_Analyse(t *testing.T) {
	type fields struct {
		byRegion    bool
		withStorage bool
		size        string
	}
	type args struct {
		objects []*model.ObjectData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Result
		wantErr bool
	}{
		{"When Analyse is called it should return the expected result",
			fields{false, false, "KB"},
			args{objectsMock},
			resultBucketMock,
			false},
		{"When Analyse is called it should return an error",
			fields{false, false, "KB"},
			args{nil},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Analyser{
				byRegion:    tt.fields.byRegion,
				withStorage: tt.fields.withStorage,
				size:        tt.fields.size,
			}
			got, err := s.Analyse(tt.args.objects)
			if (err != nil) != tt.wantErr {
				t.Errorf("Analyser.Analyse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyser.Analyse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key(t *testing.T) {
	type fields struct {
		byRegion    bool
		withStorage bool
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
			fields{false, false,"KB"},
			args{*objectMock},
			bucketNameMock},
		{"When byRegion is passed, key should return the bucket name as key",
			fields{true, false,"KB"},
			args{*objectMock},
			regionMock},
		{"When only withStorage is passed, key should return the bucket name and storage as key",
			fields{false, true, "KB"},
			args{*objectMock},
			bucketNameMock + storageStringMock},
		{"When byRegion and withStorage are passed, key should return the region and storage as key",
			fields{true, true,"KB"},
			args{*objectMock},
			regionMock + storageStringMock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Analyser{
				byRegion:    tt.fields.byRegion,
				withStorage: tt.fields.withStorage,
				size:        tt.fields.size,
			}
			if got := s.key(tt.args.object); got != tt.want {
				t.Errorf("Analyser.key() = %v, want %v", got, tt.want)
			}
		})
	}
}
