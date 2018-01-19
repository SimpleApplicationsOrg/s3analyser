package model

import (
	"time"
)

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
