package connection

import (
	"github.com/cb-migrate/models"
	"github.com/couchbase/gocb"
)

type Server struct {
	Cluster *gocb.Cluster
	Data    *models.Data
}

func GetServer(data *models.Data) (*Server, error) {
	cluster, err := gocb.Connect(data.DBHost + ":" + data.DBPort)
	if err != nil {
		return nil, err
	}

	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: data.DBUser,
		Password: data.DBPassword,
	})
	if err != nil {
		return nil, err
	}

	buckets, err := getBucketSettings(cluster, data)
	if err != nil {
	} // todo: error handle

	data.Buckets = *buckets

	return &Server{
		Cluster: cluster,
		Data:    data,
	}, nil
}

func getBucketSettings(cluster *gocb.Cluster, cred *models.Data) (*[]models.Bucket, error) {
	mngr := cluster.Manager(cred.DBUser, cred.DBPassword)
	bucketObj, _ := mngr.GetBuckets()

	var buckets []models.Bucket
	for _, bucket := range bucketObj {
		buckets = append(buckets, models.Bucket{
			BucketName:     bucket.Name,
			BucketPassword: bucket.Password,
		})
	}

	return &buckets, nil
}
