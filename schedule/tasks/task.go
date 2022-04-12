package tasks

import (
	"github.com/jasonlvhit/gocron"
	"pledge-backend/db"
	"pledge-backend/schedule/services"
	"time"
)

func Task() {

	// flush redis db
	err := db.RedisFlushDB()
	if err != nil {
		panic("clear redis error " + err.Error())
	}

	//run pool task v21
	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	_ = s.Every(2).Minutes().From(gocron.NextTick()).Do(services.NewPool().UpdateAllPoolInfo)
	_ = s.Every(1).Minute().From(gocron.NextTick()).Do(services.NewTokenPrice().UpdateContractPrice)
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(services.NewTokenSymbol().UpdateContractSymbol)
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(services.NewTokenLogo().UpdateTokenLogo)
	_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(services.NewBalanceMonitor().Monitor)
	<-s.Start() // Start all the pending jobs

}
