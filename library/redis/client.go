package redis

import (
	"fmt"
	"github.com/mediocregopher/radix/v3"
	"go.uber.org/zap"
	"net"
	"onionChat/library/logger"
)

var Client *radix.Pool

func init() {
	cname, servers, err := net.LookupSRV("redis", "tcp", "service.internal")
	if err != nil {
		logger.Instance.Panic("unable to resolve redis srv record", zap.Error(err))
	}

	if len(servers) == 0 {
		logger.Instance.Panic("unable to resolve redis srv record")
	}

	logger.Instance.Info("redis srv record",
		zap.String("cname", cname),
		zap.Uint16("pport", servers[0].Port),
		zap.String("target", servers[0].Target),
		zap.Uint16("weight", servers[0].Weight),
		zap.Uint16("priority", servers[0].Priority),
	)

	addr := net.JoinHostPort(servers[0].Target, fmt.Sprintf("%d", servers[0].Port))
	Client, err = radix.NewPool("tcp", addr, 1)
	if err != nil {
		logger.Instance.Panic("unable to create redis connection pool", zap.Error(err))
	}
}
