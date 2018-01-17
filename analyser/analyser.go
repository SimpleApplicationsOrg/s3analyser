package analyser

import (
	"github.com/SimpleApplicationsOrg/s3analyser/model"
	"fmt"
)

type Result struct {
	Buckets *[]model.Bucket
}

func Analyse(s3 model.S3) *Result {

	buckets, err := s3.ListBuckets()
	if err != nil {
		panic(err.Error())
	}

	return &Result{buckets}
}

func Print(result *Result)  {
	for _, bucket := range *result.Buckets {
		fmt.Println(*bucket.Name, *bucket.CreationDate)
	}
}
