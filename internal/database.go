package internal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Magicking/ether-swarm-services/models"
)

type Genesis struct {
	gorm.Model
	Genesis *models.Genesis
}

type Bootnode struct {
	gorm.Model
	Url string
}

type Coinbase struct {
	gorm.Model
	Address string
	Used bool
}

func InsertGenesis(ctx *Context, g *Genesis) error {
	if err := ctx.Db.Create(g).Error; err != nil {
		return fmt.Errorf("InsertGenesis: %v", err)
	}
	return nil
}

func GetGenesis(ctx *Context) (*Genesis, error) {
	var g Genesis
	if err := ctx.Db.Last(&g).Error; err != nil {
		return nil, fmt.Errorf("GetGenesis: %v", err)
	}

	return &g, nil
}

func InsertBootnode(ctx *Context, url string) error {
	bn := Bootnode{Url: url}
	if err := ctx.Db.Create(&bn).Error; err != nil {
		return fmt.Errorf("InsertBootnode: %v", err)
	}
	return nil
}

func GetBootnodes(ctx *Context) ([]Bootnode, error) {
	var bns []Bootnode

	if err := ctx.Db.Find(&bns).Error; err != nil {
		return nil, fmt.Errorf("GetBootnodes: %v", err)
	}
	return bns, nil
}

func InsertCoinbase(ctx *Context, address string) error {
	cb := Coinbase{Address: address}
	if err := ctx.Db.Create(&cb).Error; err != nil {
		return fmt.Errorf("InsertCoinbase: %v", err)
	}
	return nil
}

func GetCoinbases(ctx *Context, used bool) ([]Coinbase, error) {
	var cbs []Coinbase

	if err := ctx.Db.Where(Coinbase{Used: used}).Find(&cbs).Error; err != nil {
		return nil, fmt.Errorf("GetCoinbase: %v", err)
	}
	return cbs, nil
}

func SaveCoinbase(ctx *Context, cb *Coinbase) error {
	if err := ctx.Db.Save(cb).Error; err != nil {
		return fmt.Errorf("SaveCoinbase: %v", err)
	}
	return nil
}

func InitDatabase(ctx *Context) error {
	if err := ctx.Db.AutoMigrate(&Genesis{}).Error; err != nil {
		ctx.Db.Close()
		return fmt.Errorf("InitDatabase: %v", err)
	}
	if err := ctx.Db.AutoMigrate(&Bootnode{}).Error; err != nil {
		ctx.Db.Close()
		return fmt.Errorf("InitDatabase: %v", err)
	}
	if err := ctx.Db.AutoMigrate(&Coinbase{}).Error; err != nil {
		ctx.Db.Close()
		return fmt.Errorf("InitDatabase: %v", err)
	}
	return nil
}
