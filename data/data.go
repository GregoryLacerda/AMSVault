package data

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/data/mongo"
	"github.com.br/GregoryLacerda/AMSVault/data/mysql"
	"github.com.br/GregoryLacerda/AMSVault/database"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Data struct {
	Mongo      *mongo.Mongo
	Mysql      *mysql.Mysql
	UserGormDB *gorm.DB
}

func New(cfg *config.Config) (*Data, error) {

	service := new(Data)

	db, err := Connect(cfg)
	if err != nil {
		panic(err)
	}

	service.Mongo = mongo.NewMongo(db.Mongo, cfg)
	//service.Mysql = mysql.NewMysql(db.Mysql)

	gormDb, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDb.AutoMigrate(&entity.User{})

	userDB := database.NewUser(gormDb)
	service.UserGormDB = userDB.DB

	return service, nil
}
