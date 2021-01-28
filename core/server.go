package core

import (
	"fmt"
	irisCore "github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"video_storage/config"
	"video_storage/model"
	"video_storage/routes"
	"video_storage/tools"
)

var instance *iris.Application

func initCoreServer() {
	instance = iris.New()
	instance.Logger().SetLevel("info")
	if fileStream := tools.GeneratorLogStream(config.Instance.Common.LogPath, "app"); nil != fileStream {
		instance.Logger().SetOutput(fileStream)
	}
	instance.Logger().SetTimeFormat("2006-01-02 15:04:05")
	instance.Use(recover.New())
	instance.Use(logger.New())
	instance.Use(irisCore.New(irisCore.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
		MaxAge:           60,
		AllowCredentials: true,
	}))
}

func configCoreServer() {
	instance.OnAnyErrorCode(func(context iris.Context) {
		path := context.Path()
		if strings.Contains(path, "/") {
			_, _ = context.JSON(&model.ResponseBody{
				Code:    context.GetStatusCode(),
				Message: "",
			})
		}
	})

	instance.Any("/", func(context iris.Context) {
		_, _ = context.JSON(&model.ResponseBody{
			Code:    200,
			Message: "powered by afterloe <605728727@qq.com>",
		})
	})

	// 公用路由
	mvc.Configure(instance.Party("/"), func(pub *mvc.Application) {
		pub.Party("/").Handle(new(routes.PubRoute))
	})
	//
	//// 鉴权 api
	//mvc.Configure(instance.Party("/aip"), func(aip *mvc.Application) {
	//	aip.Router.Use(logic.AuthLogic)
	//	aip.Party("/device").Handle(new(routes.DeviceRoute))
	//	aip.Party("/app/{name:string}").Handle(new(routes.AppRoute))
	//})
}

func signalListener(server *http.Server) {
	pkg := make(chan os.Signal)
	signal.Notify(pkg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		code := <-pkg
		logrus.Infof("得到信号 [%s], 即将退出服务", code)
		logrus.Info("服务关闭")
		os.Exit(0)
	}()
}

func StartUpHttpServer() {
	initCoreServer()
	configCoreServer()
	instanceService := &http.Server{
		Addr: fmt.Sprintf("%s:%s", config.Instance.Backend.ServiceHost, config.Instance.Backend.Port),
	}
	signalListener(instanceService)
	if err := instance.Run(iris.Server(instanceService), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		RemoteAddrHeaders:                 []string{"X-Forwarded-For"},
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	})); nil != err {
		logrus.Error(err)
		os.Exit(-2)
	}
}
