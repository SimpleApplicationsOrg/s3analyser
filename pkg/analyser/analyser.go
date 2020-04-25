package analyser

import (
	"fmt"
	"github.com/SimpleApplicationsOrg/s3analyser/pkg/model"
)

var zero = 0
var zero64 = int64(zero)
var blank = ""

// Result has the analysis result
type Result struct {
	Objects map[string]model.ObjectData
}

type Analyser struct {
	byRegion    bool
	withStorage bool
	size        string
}

// New creates the analyzer with the configuration flags
func New(byRegion bool, withStorage bool, size string) *Analyser {
	return &Analyser{byRegion, withStorage, size}
}

// Analyze s3 buckets
func (a *Analyser) Analyse(objects []model.ObjectData) (*Result, error) {

	if objects == nil {
		return nil, fmt.Errorf("objects to be analysed should not be nil")
	}

	result := make(map[string]model.ObjectData)
	for _, obj := range objects {

		key := a.key(obj)

		if _, ok := result[key]; !ok {
			result[key] = model.ObjectData{Bucket: obj.Bucket, CreationDate: obj.CreationDate, Region: obj.Region,
				LastModified: obj.LastModified, Count: zero, Size: zero64, StorageClass: obj.StorageClass}
		}

		object := result[key]
		size := object.Size + obj.Size
		count := object.Count + 1

		r := result[key]
		r.Size = size
		r.Count = count

		if obj.LastModified.After(object.LastModified) {
			r.LastModified = obj.LastModified
		}

		if !a.withStorage {
			r.StorageClass = blank
		}

		result[key] = r
	}

	return &Result{result}, nil
}

func (a *Analyser) key(object model.ObjectData) string {

	key := object.Bucket
	if a.byRegion {
		key = object.Region
	}

	if a.withStorage {
		key += object.StorageClass
	}

	return key
}
