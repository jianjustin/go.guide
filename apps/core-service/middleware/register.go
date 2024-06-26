package middleware

import (
	"context"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

const ServiceKey string = "/services/core/"
const ServiceName string = "core-service"

func GetEtcdRegister() etcdv3.Client {
	client, _ := etcdv3.NewClient(
		context.Background(),
		[]string{"http://etcd:2379"},
		etcdv3.ClientOptions{
			DialTimeout:   3 * time.Second,
			DialKeepAlive: 3 * time.Second,
		},
	)
	return client
}
