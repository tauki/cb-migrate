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

