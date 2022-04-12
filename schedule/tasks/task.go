package tasks

import (
	"github.com/jasonlvhit/gocron"
	"pledge-backend/db"
	"pledge-backend/schedule/services/sv21"
	"pledge-backend/schedule/services/sv22"
	"time"
)

func TaskV21() {

	// flush redis db
	err := db.RedisFlushDB()
	if err != nil {
		panic("clear redis error " + err.Error())
	}

	//run pool task v21
	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	_ = s.Every(2).Minutes().From(gocron.NextTick()).Do(sv21.NewPool().UpdateAllPoolInfo)
	_ = s.Every(1).Minute().From(gocron.NextTick()).Do(sv21.NewTokenPrice().UpdateContractPrice)
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(sv21.NewTokenSymbol().UpdateContractSymbol)
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(sv21.NewTokenLogo().UpdateTokenLogo)
	_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(sv21.NewBalanceMonitor().Monitor)
	<-s.Start() // Start all the pending jobs

}

func TaskV22() {

	// flush redis db
	err := db.RedisFlushDB()
	if err != nil {
		panic("clear redis error " + err.Error())
	}

	//run pool task v22
	s := gocron.NewScheduler()
	_ = s.Every(2).Minutes().From(gocron.NextTick()).Do(sv22.NewPool().UpdateAllPoolInfo)
	_ = s.Every(1).Minute().From(gocron.NextTick()).Do(sv22.NewTokenPrice().UpdateContractPrice)
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(sv22.NewTokenSymbol().UpdateContractSymbol)
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(sv22.NewTokenLogo().UpdateTokenLogo)
	_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(sv22.NewBalanceMonitor().Monitor)
	<-s.Start() // Start all the pending jobs
}
