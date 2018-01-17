package model

import "time"

type Bucket struct {
	Name *string
	CreationDate *time.Time
}

type S3 interface {
	ListBuckets() ([]*Bucket, error)
}

