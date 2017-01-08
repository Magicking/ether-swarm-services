package internal

import (
	"fmt"
	"net/url"

	"github.com/Magicking/ether-swarm-services/restapi/operations"
)

func RegisterNodeHandler(ctx *Context, params operations.RegisterNodeParams) (bool, error) {
	u, err := url.Parse(params.EnodeURL)
	if err != nil {
		return false, fmt.Errorf("RegisterNodeHandler: %v", err)
	}
	//TODO register / deregister / ...
	err = InsertBootnode(ctx, u.String())
	if err != nil {
		return false, fmt.Errorf("RegisterNodeHandler: %v", err)
	}
	return true, nil
}

func StartNodeHandler(ctx *Context, params operations.StartNodeParams) (bool, error) {
	address := params.Etherbase
	//TODO validate address
	var err error
	if err != nil {
		return false, fmt.Errorf("StartNodeHandler: %v", err)
	}
	//TODO register / deregister / ...
	err = InsertCoinbase(ctx, address)
	//TODO start new geth miner
	if err != nil {
		return false, fmt.Errorf("StartNodeHandler: %v", err)
	}
	return true, nil
}
