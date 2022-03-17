package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Dns struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
	TimeZone string
}

func NewProductionDB() *Dns {
	c := NewConfig()
	return &Dns{
		Host:     c.DB.Production.Host,
		User:     c.DB.Production.User,
		Password: c.DB.Production.Password,
		DbName:   c.DB.Production.DbName,
		Port:     c.DB.Production.Port,
		SslMode:  c.DB.Production.SslMode,
		TimeZone: c.DB.Production.TimeZone,
	}
}

func NewDevelopmentDB() *Dns {
	c := NewConfig()
	return &Dns{
		Host:     c.DB.Development.Host,
		User:     c.DB.Development.User,
		Password: c.DB.Development.Password,
		DbName:   c.DB.Development.DbName,
		Port:     c.DB.Development.Port,
		SslMode:  c.DB.Development.SslMode,
		TimeZone: c.DB.Development.TimeZone,
	}
}

func (d *Dns) NewDB() error {
	dsn := "host=" + d.Host + " " +
		"user=" + d.User + " " +
		"password=" + d.Password + " " +
		"dbname=" + d.DbName + " " +
		"port=" + d.Port + " " +
		"sslmode=" + d.SslMode + " " +
		"TimeZone=" + d.TimeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	return nil
}

func (db *Dns) Connect() *gorm.DB {
	return DB
}
