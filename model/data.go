package model

import (
	"time"
)

// ObjectData is used for keep s3 objetcs data and the result of analysis
type ObjectData struct {
	Bucket       *string
	CreationDate *time.Time
	Region       *string
	Key          *string
	LastModified *time.Time
	Count        *int
	Size         *int64
	StorageClass *string
}
