package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"github.com/reaperhero/elasticsearch-alarm/pkg/utils"
	"github.com/sirupsen/logrus"
)

var (
	db *gorm.DB
)

func init() {
	connectDB()
	db.AutoMigrate(&model.AlarmConfig{}, &model.AlarmConfigInstance{}, &model.AlarmInstance{})
}

func connectDB() {
	gormDB, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
			utils.GetEnvWithDeafult("MYSQL_USER", "root"),
			utils.GetEnvWithDeafult("MYSQL_PASS", "Lzslov123!"),
			utils.GetEnvWithDeafult("MYSQL_ADDR", "127.0.0.1:3306"),
			utils.GetEnvWithDeafult("MYSQL_DB", "alarm"),
		),
	)
	if err != nil {
		logrus.Fatalf("[InitDb error ]%s", err)
	}
	db = gormDB
}

