package initial

import (
	"github.com/ZLinFeng/play-server/config"
	"github.com/ZLinFeng/play-server/global"
	"github.com/ZLinFeng/play-server/model/db/system"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDb(m *config.SysMysql) *gorm.DB {

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         250,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return nil
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	return db
}

func RegisterTables() {
	err := global.DB.AutoMigrate(
		system.SysUser{},
		system.SysRole{},
		system.SysUserRole{},
		system.SysMenu{},
		system.SysMenuBtn{},
		system.SysRoleBtn{},
		system.SysRoleMenu{},
		system.SysDept{},
	)

	if err != nil {
		global.Logger.Error("register system table error", zap.Error(err))
		os.Exit(1)
	}

	global.Logger.Info("register system table success")
}
