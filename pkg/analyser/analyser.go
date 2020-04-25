package analyser

import (
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/service"
	"io"
)

var zero = 0
var zero64 = int64(zero)
var blank = ""

// Result has the analysis result
type Result struct {
	Objects map[string]*model.ObjectData
}

// S3Analyser is used to analyze s3 objects and print the result
type S3Analyser interface {
	Analyse(s3 service.S3) (*Result, error)
	Print(writer io.Writer, result *Result)
}

type sat struct {
	byRegion    bool
	withStorage bool
	filter      model.FilterMap
	size        string
}

// Factory creates the analyzer with the configuration flags
func Factory(byRegion bool, withStorage bool, filter model.FilterMap, size string) S3Analyser {
	return &sat{byRegion, withStorage, filter, size}
}

// Analyze s3 buckets
func (sat *sat) Analyse(s3 service.S3) (*Result, error) {

	objects, err := s3.Objects(sat.filter)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*model.ObjectData)
	for _, obj := range objects {

		key := sat.key(*obj)

		if _, ok := result[key]; !ok {
			result[key] = &model.ObjectData{Bucket: obj.Bucket, CreationDate: obj.CreationDate, Region: obj.Region,
				LastModified: obj.LastModified, Count: &zero, Size: &zero64, StorageClass: obj.StorageClass}
		}

		object := result[key]
		size := *object.Size + *obj.Size
		count := *object.Count + 1

		result[key].Size = &size
		result[key].Count = &count

		if obj.LastModified.After(*object.LastModified) {
			result[key].LastModified = obj.LastModified
		}

		if !sat.withStorage {
			result[key].StorageClass = &blank
		}
	}

	return &Result{result}, nil
}

func (sat *sat) key(object model.ObjectData) string {

	key := *object.Bucket
	if sat.byRegion {
		key = *object.Region
	}

	if sat.withStorage {
		key += *object.StorageClass
	}

	return key
}
