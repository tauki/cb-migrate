package connection

import (
	"github.com/cb-migrate/models"
	"github.com/couchbase/gocb"
)

type ClusterServer struct {
	Cluster *gocb.Cluster
	Cred    *models.Cluster
}

func GetClusterServer(cred *models.Cluster) (*ClusterServer, error) {
	cluster, err := gocb.Connect(cred.DBHost + ":" + cred.DBPort)
	if err != nil {
		return nil, err
	}

	return &ClusterServer{
		Cluster: cluster,
		Cred:    cred,
	}, nil
}

func (c *ClusterServer) GetBucketNames() (*[]models.Bucket, error) {
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
