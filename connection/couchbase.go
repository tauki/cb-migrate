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

	data.Buckets = &buckets

	return &Server{
		Cluster: cluster,
		Data:    data,
	}, nil
}

func getBucketSettings(cluster *gocb.Cluster, cred *models.Data) ([]*gocb.BucketSettings, error) {
	mngr := cluster.Manager(cred.DBUser, cred.DBPassword)
	return mngr.GetBuckets()
}

func (s *Server) CreateBucket(name string, settings *gocb.BucketSettings) (error) {
	mngr := s.Cluster.Manager(s.Data.DBUser, s.Data.DBPassword)
	return mngr.InsertBucket(settings)
}

func (s *Server) Copy(settings *gocb.BucketSettings) (error) {
	return nil // todo
}

func (s *Server) BucketExists(name string) (bool) {
	for _, bucket := range *s.Data.Buckets {
		if name == bucket.Name {
			return true
		}
	}
	return false
}