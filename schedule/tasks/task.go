package tasks

import (
	"github.com/jasonlvhit/gocron"
	"pledge-backend/db"
	"pledge-backend/schedule/services/sv21"
	"pledge-backend/schedule/services/sv22"
)

func TaskV21() {

	// flush redis db
	err := db.RedisFlushDB()
	if err != nil {
		panic("clear redis error " + err.Error())
	}

	//Init all the jobs
	sv21.NewPool().UpdateAllPoolInfo()
	sv21.NewTokenPrice().UpdateContractPrice()
	sv21.NewTokenSymbol().UpdateContractSymbol()
	sv21.NewTokenLogo().UpdateTokenLogo()

	//run pool task v22
	s := gocron.NewScheduler()
	s.Every(2).Minutes().Do(sv21.NewPool().UpdateAllPoolInfo)
	s.Every(1).Minute().Do(sv21.NewTokenPrice().UpdateContractPrice)
	s.Every(2).Hours().Do(sv21.NewTokenSymbol().UpdateContractSymbol)
	s.Every(2).Hours().Do(sv21.NewTokenLogo().UpdateTokenLogo)
	<-s.Start() // Start all the pending jobs

}

func TaskV22() {

	// flush redis db
	err := db.RedisFlushDB()
	if err != nil {
		panic("clear redis error " + err.Error())
	}

	//Init all the jobs
	sv22.NewPool().UpdateAllPoolInfo()
	sv22.NewTokenPrice().UpdateContractPrice()
	sv22.NewTokenSymbol().UpdateContractSymbol()
	sv22.NewTokenLogo().UpdateTokenLogo()

	//run pool task v22
	s := gocron.NewScheduler()
	s.Every(2).Minutes().Do(sv22.NewPool().UpdateAllPoolInfo)
	s.Every(1).Minute().Do(sv22.NewTokenPrice().UpdateContractPrice)
	s.Every(2).Hours().Do(sv22.NewTokenSymbol().UpdateContractSymbol)
	s.Every(2).Hours().Do(sv22.NewTokenLogo().UpdateTokenLogo)
	<-s.Start() // Start all the pending jobs
}
