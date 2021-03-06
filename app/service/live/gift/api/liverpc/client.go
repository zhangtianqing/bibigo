// Code generated by liverpcgen, DO NOT EDIT.
// source: *.proto files under this directory
// If you want to change this file, Please see README in go-common/app/tool/liverpc/protoc-gen-liverpc/
package liverpc

import (
	"go-common/app/service/live/gift/api/liverpc/v0"
	"go-common/app/service/live/gift/api/liverpc/v1"
	"go-common/library/net/rpc/liverpc"
)

// Client that represents a liverpc gift service api
type Client struct {
	cli *liverpc.Client
	// V0Smalltv presents the controller in liverpc
	V0Smalltv v0.Smalltv
	// V1Gift presents the controller in liverpc
	V1Gift v1.Gift
}

// DiscoveryAppId the discovery id is not the tree name
var DiscoveryAppId = "live.gift"

// New a Client that represents a liverpc live.gift service api
// conf can be empty, and it will use discovery to find service by default
// conf.AppID will be overwrite by a fixed value DiscoveryAppId
// therefore is no need to set
func New(conf *liverpc.ClientConfig) *Client {
	if conf == nil {
		conf = &liverpc.ClientConfig{}
	}
	conf.AppID = DiscoveryAppId
	var realCli = liverpc.NewClient(conf)
	cli := &Client{cli: realCli}
	cli.clientInit(realCli)
	return cli
}

func (cli *Client) GetRawCli() *liverpc.Client {
	return cli.cli
}

func (cli *Client) clientInit(realCli *liverpc.Client) {
	cli.V0Smalltv = v0.NewSmalltvRpcClient(realCli)
	cli.V1Gift = v1.NewGiftRpcClient(realCli)
}
