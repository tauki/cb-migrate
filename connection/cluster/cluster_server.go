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

	return &Server{
		Cluster: cluster,
		Cred:    cred,
	}, nil
}

func (c *Server) GetBucketNames() (*[]models.Bucket, error) {
	mngr := c.Cluster.Manager(c.Cred.DBUser, c.Cred.DBPassword)
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
