package main

import (
	"pledge-backend/db"
	"pledge-backend/schedule/tasks"
)

func main() {
	// init mysql
	db.InitMysql()

	//init redis
	db.InitRedis()

	//run pool task v21
	tasks.TaskV21()

	//run pool task v22
	//tasks.TaskV22()

}

/*
 If you change the version, you need to modify the following files
 pledge_task.go
 config/init.go
*/
