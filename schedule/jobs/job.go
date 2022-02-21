package jobs

//import (
//	"fmt"
//	"github.com/ethereum/go-ethereum/common"
//	"pledge-backend/api/models"
//	"pledge-backend/api/services"
//	"pledge-backend/api/services/pledge_pool_token"
//	"pledge-backend/config"
//	"pledge-backend/db"
//	"pledge-backend/log"
//	"time"
//)
//
//// abigen --abi abi/pledge_pool.abi --bin abi/pledge_pool.bin --pkg pledge_pool_token --out pledgePoolToken.go
//// update pool state
//func UpdatePoolState() {
//	err := services.NewContract().PoolLength(config.Config.TestNet.PledgePoolToken, config.Config.TestNet.NetUrl)
//	fmt.Println(err)
//	if err != nil {
//		log.Logger.Sugar().Error("get poolLength err:", err)
//	}
//
//	PledgePoolToken := pledge_pool_token.PledgePoolToken{}
//
//	tokengo, err := PledgePooll(common.HexToAddress(contractAddress), myCommon.EthConn)
//	tokengo, err := PledgePoolToken(common.HexToAddress(contractAddress), myCommon.EthConn)
//	if err != nil {
//		log.Panic("Failed to instantiate a Token contract: %v", err)
//	}
//
//	a, b := PledgePoolToken.PoolBaseInfo()
//	fmt.Println(a, b)
//	poolBases := models.Poolbases{}
//	db.Mysql.Find(&poolBases)
//	fmt.Println(poolBases)
//	log.Logger.Info("task")
//	time.Sleep(time.Second * 3)
//	fmt.Println(1)
//}
