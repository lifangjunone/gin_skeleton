package main

import "gin_skeleton/cmd"

func main() {
	//fmt.Println("hello gin")
	//engine := routers.InitEngine()
	//conf.InitConfig()
	//middlewares.InitMiddleware(engine)
	//routers.InitRouter(engine)
	//engine.Run("0.0.0.0:8080")
	cmd.Execute()
}
