1 项目定时任务

2 github.com/jasonlvhit/gocron 该定时任务库会在协程里执行任务，在执行耗时任务时，会按照定时时间开始执行，结束时间根据具体任务而定

3 代码示例

    //可以执行带参数的方法（任务），下面有示例
    //gocron.Every(1).Second().Do(tasks.Task)
    //gocron.Every(1).Second().Do(tasks.Task2)
    //gocron.Every(2).Seconds().Do(tasks.Task)
    //gocron.Every(1).Minute().Do(tasks.Task)
    //gocron.Every(2).Minutes().Do(tasks.Task)
    //gocron.Every(1).Hour().Do(tasks.Task)
    //gocron.Every(2).Hours().Do(tasks.Task)
    //gocron.Every(1).Day().Do(tasks.Task)
    //gocron.Every(2).Days().Do(tasks.Task)
    //gocron.Every(1).Week().Do(tasks.Task)
    //gocron.Every(2).Weeks().Do(tasks.Task)

	// Do jobs with params
	//gocron.Every(1).Second().Do(tasks.TaskWithParams, "hello")

	// Do jobs on specific weekday
	//gocron.Every(1).Monday().Do(tasks.Task)
	//gocron.Every(1).Thursday().Do(tasks.Task)
	//
	//// Do a job at a specific time - 'hour:min:sec' - seconds optional
	//gocron.Every(1).Day().At("10:30").Do(tasks.Task)
	//gocron.Every(1).Monday().At("18:30").Do(tasks.Task)
	//gocron.Every(1).Tuesday().At("18:30:59").Do(tasks.Task)
	//
	//// Begin job immediately upon start
	//gocron.Every(1).Hour().From(gocron.NextTick()).Do(tasks.Task)
	//
	//// Begin job at a specific date/time
	//t := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	//gocron.Every(1).Hour().From(&t).Do(tasks.Task)

	// NextRun gets the next running time
	_, time := gocron.NextRun()
	fmt.Println(time)

	// Remove a specific job
	//gocron.Remove(tasks.task)

	// Clear all scheduled jobs
	//gocron.Clear()

	// Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(tasks.Task)
	<-s.Start()