// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/quarkcloudio/quark-go/v3/pkg/adapter/hertzadapter"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/install"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/middleware"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	h := server.Default(server.WithHostPorts(":3000"))

	// 注册路由
	register(h)

	// 静态文件
	h.StaticFile("/admin/", "./web/app/admin/index.html")

	// WEB根目录
	fs := &app.FS{Root: "./web/app", IndexNames: []string{"index.html"}}
	h.StaticFS("/", fs)

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

	// 适配hertz
	hertzadapter.Adapter(b, h)

	// 启动服务
	h.Spin()
}
