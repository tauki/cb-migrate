package models

import "github.com/couchbase/gocb"

type Data struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	Buckets    *[]*gocb.BucketSettings
}
