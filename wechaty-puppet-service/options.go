package puppetservice

import (
	"time"

	wechatypuppet "github.com/leozeli/go-wechaty/wechaty-puppet"
)

// TLSConfig tls config
type TLSConfig struct {
	CaCert     string
	ServerName string

	Disable bool //  only for compatible with old clients/servers
}

// Options puppet-service options
type Options struct {
	wechatypuppet.Option

	GrpcReconnectInterval time.Duration
	Authority             string
	TLS                   TLSConfig
}
