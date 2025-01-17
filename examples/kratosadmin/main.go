package main

import (
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/quarkcloudio/quark-go/v3/pkg/adapter/kratosadapter"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/install"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/middleware"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	hs := http.NewServer(
		http.Address(":3000"),
	)

	// 数据库配置信息
	dsn := "root:fK7xPGJi1gJfIief@tcp(127.0.0.1:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 配置资源
	config := &builder.Config{
		AppKey:    "123456",
		Providers: service.Providers,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 创建对象
	b := builder.New(config)

	// 初始化安装
	install.Handle()

	// 中间件
	b.Use(middleware.Handle)

	// 适配kratos
	kratosadapter.Adapter(b, hs)

	// WEB根目录，只能放在后面，否则与其他路由有冲突
	hs.HandlePrefix("/", stdhttp.FileServer(stdhttp.Dir("./web/app")))

	// 创建服务
	app := kratos.New(
		kratos.Server(hs),
	)

	// 启动服务
	app.Run()
}
