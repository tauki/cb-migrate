package models

type Cluster struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
}

type Bucket struct {
	BucketName     string
	BucketPassword string
}
