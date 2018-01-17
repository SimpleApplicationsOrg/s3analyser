package analyser

import (
	"github.com/SimpleApplicationsOrg/s3analyser/model"
	"fmt"
)

type Result struct {
	Buckets []*model.Bucket
}

func Analyse(s3 model.S3) (*Result, error) {

	buckets, err := s3.ListBuckets()
	if err != nil {
		return nil, err
	}

	return &Result{buckets}, nil
}

func Print(result *Result)  {
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name, *bucket.CreationDate)
	}
}
