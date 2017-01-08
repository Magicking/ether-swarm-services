package internal

import (
	"github.com/jinzhu/gorm"
)

type Context struct {
	Db *gorm.DB
	NetworkID int64
}

type GenesisConf struct {
	Nonce       string `long:"nonce" env:"GG_NONCE" default:"0x0000000000000042" description:"Nonce of the genesis block"`
	Timestamp   string `long:"timestamp" env:"GG_TIMESTAMP" default:"0x0" description:"Timestamp of the genesis block"`
	ParentHash  string `long:"parent-hash" env:"GG_PARENT_HASH" default:"0x0" description:"Parent hash of the genesis block"`
	ExtraData   string `long:"extra-data" env:"GG_EXTRA_DATA" description:"Extra data of the genesis block"`
	GasLimit    string `long:"gas-limit" env:"GG_GAS_LIMIT" default:"0x8000000" description:"Gas limit of the genesis block"`
	Difficulty  string `long:"difficulty" env:"GG_DIFFICULTY" default:"0x400" description:"Initial difficulty"`
	Mixhash     string `long:"mixhash" env:"GG_MIX_HASH" default:"0x0" description:"Mixhash of the genesis block"`
	Coinbase    string `long:"coinbase" env:"GG_COINBASE" default:"0x0" description:"Coinbase of the genesis block"`
	Balance     string `long:"balance" env:"GG_BALANCE" default:"1000000000000000000" description:"Default balance in Wei for new account when new_allocator is used"`
}
