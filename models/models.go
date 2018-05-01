package models

type Data struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	Buckets    []Bucket
}

type Bucket struct {
	BucketName     string
	BucketPassword string
}
