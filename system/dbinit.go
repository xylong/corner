package system

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// dbInit 初始化数据库
func dbInit(aliases ...string) {
	// 判断开发模式
	isDev := beego.AppConfig.String("runmode") == "dev"

	if len(aliases) > 0 {
		for _, alias := range aliases {
			registerDatabase(alias)
			if alias == "w" {
				orm.RunSyncdb("default", false, isDev)
			}
		}
	} else {
		registerDatabase("w")
		orm.RunSyncdb("default", false, isDev) // 自动建表
	}

	orm.Debug = isDev
}

func registerDatabase(alias string) {
	if len(alias) <= 0 {
		return
	}

	dbAlias := alias

	if alias == "w" || alias == "default" || len(alias) <= 0 {
		dbAlias = "default"
		alias = "w"
	}

	dbName := beego.AppConfig.String("db_" + alias + "_database")
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	dbPwd := beego.AppConfig.String("db_" + alias + "_password")
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	dbPort := beego.AppConfig.String("db_" + alias + "_port")

	orm.RegisterDataBase(dbAlias, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbUser,
		dbPwd,
		dbHost,
		dbPort,
		dbName))
}
