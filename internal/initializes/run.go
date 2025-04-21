package initializes

func Run() {
	loadConfig()
	InitLogger()
	InitDatabase()
	InitRedis()
	r := InitRouter()

	err := r.Run()
	if err != nil {
		return
	}
}
