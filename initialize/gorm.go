package initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
	dsn := "postgres://liguanjie:Ike_19930623@localhost:5432/shuxiangge?sslmode=disable"
	postgresConfig := postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	GormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	db, err := gorm.Open(postgres.New(postgresConfig), GormConfig); if err != nil{
		panic(err)
	} else {
		return db
	}
}
