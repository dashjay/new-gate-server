package main

import (
	"NewGateServer/WeHub"
	"NewGateServer/configs"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"log"
)

func main() {

	log.SetFlags(log.Lshortfile)
	app := iris.New()
	crs := cors.AllowAll()
	app.Any("/we/{suffix}", WeHub.MainHandler)
	app.Any("/api/create", crs, WeHub.Create)
	app.Any("/api/del", crs, WeHub.Del)
	app.Any("/api/monitor", crs, WeHub.Monitor)
	app.Any("/api/stop", crs, WeHub.Stop)
	app.Any("/api/start", crs, WeHub.Start)
	app.Any("/api/update", crs, WeHub.Update)
	app.Any("/api/userinfo", crs, WeHub.GetUsersInfoTest)
	app.Run(iris.Addr(fmt.Sprintf(":%d", configs.C.Port)), iris.WithOptimizations)
}
