package initializes

func Run() {
	loadConfig()
	InitLogger()
	InitDatabase()
	InitRedis()
	r := InitRouter()
	r.Run()
}
