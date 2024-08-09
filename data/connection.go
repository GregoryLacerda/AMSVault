package data

import (
	"context"
	"database/sql"

	"github.com.br/GregoryLacerda/AMSVault/config"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connection struct {
	Mongo *mongo.Client
	Mysql *sql.DB
}

func Connect(cfg *config.Config) (connection, error) {

	/*connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	dbMysql, err := sql.Open("mysql", connectionString)
	if err != nil {
		return connection{}, err
	}*/

	dbMongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return connection{}, err
	}

	dbMongo.Database(cfg.MongoDB)

	return connection{
		Mongo: dbMongo,
		//Mysql: dbMysql,
	}, nil
}
