package internal

import (
	"fmt"
	"encoding/json"

	"github.com/Magicking/ether-swarm-services/models"
	"github.com/Magicking/ether-swarm-services/restapi/operations"

	"github.com/Magicking/ether-swarm-services/internal/blockchain"
)

func CreateGenesisHandler(ctx *Context, gen_default *GenesisConf, params operations.CreateGenesisParams) (*models.Genesis, error) {
	if g, _ := GetGenesis(ctx); g != nil {
		return nil, fmt.Errorf("CreateGenesis: Genesis already set")
	}

	var allocator int

	if params.NewAllocator == nil {
		if params.Genesis == nil {
			allocator = 1
		} else {
			allocator = 0
		}
	} else {
		allocator = int(*params.NewAllocator)
	}

	var genesis *models.Genesis
	if params.Genesis == nil {
		genesis = &models.Genesis{
			Coinbase: &gen_default.Coinbase,
			Difficulty: &gen_default.Difficulty,
			ExtraData: &gen_default.ExtraData,
			GasLimit: &gen_default.GasLimit,
			Mixhash: &gen_default.Mixhash,
			Nonce: &gen_default.Nonce,
			ParentHash: &gen_default.ParentHash,
			Timestamp: &gen_default.Timestamp,
		}
	} else {
		genesis = params.Genesis
	}

	if genesis.Alloc == nil {
		genesis.Alloc = make(map[string]models.Allocator)
	}
	for ; allocator > 0; allocator-- {
		address, alloc, err := blockchain.NewAllocator(gen_default.Balance)
		if err != nil {
			return nil, fmt.Errorf("CreateGenesis: %v", err)
		}
		_, ok := genesis.Alloc[address]
		if ok {
			allocator++
			continue
		}
		genesis.Alloc[address] = *alloc
	}
	for _, e := range genesis.Alloc {
		e.PrivateKey = nil
	}
	jsonData, err := json.Marshal(genesis)
	if err != nil {
		return nil, fmt.Errorf("CreateGenesis: %v", err)
	}
	if err := InsertGenesis(ctx, &Genesis{JSONData: jsonData}); err != nil {
		return nil, fmt.Errorf("CreateGenesis: %v", err)
	}
	return genesis, nil
}
