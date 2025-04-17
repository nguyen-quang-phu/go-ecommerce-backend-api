package initializes

func Run() {
	loadConfig()
	InitLogger()
	InitDatabase()
	r := InitRouter()
	r.Run()
}
