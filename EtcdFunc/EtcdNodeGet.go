package EtcdFunc

import (
	"go.etcd.io/etcd/clientv3"
	"time"
)

type Client struct {
	client     *clientv3.Client

}

func NewEtcdCl(addr []string)(*Client, error)  {

	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	if client, err := clientv3.New(conf); err == nil {
		return &Client{
			client:     client,
		}, nil
	} else {
		return nil, err
	}
}

