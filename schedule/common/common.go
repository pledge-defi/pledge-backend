package common

import (
	"pledge-backend/config"
)

type TaskTimeControl struct {
	Task     string
	Duration int64
	OverTime int
	OnTime   int
}

type Scheduler struct {
	Task     string
	Duration int64
}

var PoolTask = TaskTimeControl{
	Task:     "pool",
	Duration: config.Config.Env.TaskDuration,
}

var PriceTask = TaskTimeControl{
	Task:     "price",
	Duration: config.Config.Env.TaskDuration,
}

var TaskDurationModifyChannel = make(chan Scheduler, 0)
