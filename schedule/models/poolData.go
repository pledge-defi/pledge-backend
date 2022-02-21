package models

import (
	"errors"
	"gorm.io/gorm"
	"pledge-backend/db"
	"pledge-backend/utils"
)

type PoolData struct {
	Id                     int    `json:"_" gorm:"column:id;primaryKey"`
	PoolId                 string `json:"pool_id" gorm:"column:pool_id"`
	ChainId                string `json:"chain_id" gorm:"column:chain_id"`
	FinishAmountBorrow     string `json:"finish_amount_borrow" gorm:"column:finish_amount_borrow"`
	FinishAmountLend       string `json:"finish_amount_lend" gorm:"column:finish_amount_lend"`
	LiquidationAmounBorrow string `json:"liquidation_amoun_borrow" gorm:"column:liquidation_amoun_borrow"`
	LiquidationAmounLend   string `json:"liquidation_amoun_lend" gorm:"column:liquidation_amoun_lend"`
	SettleAmountBorrow     string `json:"settle_amount_borrow" gorm:"column:settle_amount_borrow"`
	SettleAmountLend       string `json:"settle_amount_lend" gorm:"column:settle_amount_lend"`
	CreatedAt              string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt              string `json:"updated_at" gorm:"column:updated_at"`
}

func NewPoolData() *PoolData {
	return &PoolData{}
}

// Save poolData information
func (t *PoolData) SavePoolData(chainId, poolId string, poolData *PoolData) error {

	nowDateTime := utils.GetCurDateTimeFormat()
	poolData.UpdatedAt = nowDateTime
	err := db.Mysql.Table("pooldata").Where("chain_id=? and pool_id=?", chainId, poolId).First(&t).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			poolData.CreatedAt = nowDateTime
			err = db.Mysql.Table("pooldata").Create(poolData).Debug().Error
			if err != nil {
				return err
			}
		} else {
			return errors.New("record select err " + err.Error())
		}
	}

	err = db.Mysql.Table("pooldata").Where("chain_id=? and pool_id=?", chainId, poolId).Updates(poolData).Debug().Error
	if err != nil {
		return err
	}
	return nil
}
