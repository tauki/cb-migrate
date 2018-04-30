package cluster

import (
	"github.com/cb-migrate/models"
	"github.com/couchbase/gocb"
)

type Server struct {
	Cluster *gocb.Cluster
	Cred    *models.Cluster
}

func GetServer(cred *models.Cluster) (*Server, error) {
	cluster, err := gocb.Connect(cred.DBHost + ":" + cred.DBPort)
	if err != nil {
		return nil, err
	}

	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: cred.DBUser,
		Password: cred.DBPassword,
	})
	if err != nil {
		return nil, err
	}

	buckets, err := getBuckets(cluster, cred)
	if err != nil {
	} // todo: error handle

	cred.Buckets = *buckets

	return &Server{
		Cluster: cluster,
		Cred:    cred,
	}, nil
}

func getBuckets(cluster *gocb.Cluster, cred *models.Cluster) (*[]models.Bucket, error) {
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
