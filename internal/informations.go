package internal

import (
	"fmt"
	"encoding/json"

	"github.com/Magicking/ether-swarm-services/models"
)

func GetInformationHandler(ctx *Context) (*models.Information, error) {
	var infos models.Information

	infos.Networkid = ctx.NetworkID
	bns, err := GetBootnodes(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetInformation: %v", err)
	}
	for _, bn := range bns {
		infos.BootnodesUrls = append(infos.BootnodesUrls, bn.Url)
	}
	genesis, err := GetGenesis(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetInformation: %v", err)
	}
	if len(genesis.JSONData) > 0 {
		err = json.Unmarshal(genesis.JSONData, &infos.Genesis)
		if err != nil {
			return nil, fmt.Errorf("GetInformation: %v", err)
		}
	}
	return &infos, nil
}
