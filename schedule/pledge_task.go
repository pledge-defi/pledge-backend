package main

import (
	"pledge-backend/db"
	"pledge-backend/schedule/tasks"
)

func main() {
	// init mysql
	db.InitMysql()

	// init redis
	db.InitRedis()

	// pool task
	tasks.Task()

}

/*
 If you change the version, you need to modify the following files'
 config/init.go
*/
