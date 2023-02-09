package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/tzrd/go-antdv-core/storage"
)

type DatabaseConf struct {
	Host         string
	Port         int
	Username     string `json:",optional"`
	Password     string `json:",optional"`
	DBName       string `json:",optional"`
	SSLMode      string `json:",optional"`
	Type         string `json:",default=mysql,options=[mysql,postgres]"`
	MaxOpenConns *int   `json:",optional,default=100"`
	Debug        bool   `json:",optional,default=false"`
	CacheTime    int    `json:",optional,default=10"`
}

func (c DatabaseConf) NewDataBaseDriver() *storage.MySQL {
	driver := c.GetDSN()
	mysql := storage.NewMySqlConfig(true, "mysql", driver, true)
	mysql.Connect()

	return mysql
}

func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", c.Username, c.Password, c.Host, c.Port, c.DBName)
}

func (c DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

func (c DatabaseConf) GetDSN() string {
	switch c.Type {
	case "mysql":
		return c.MysqlDSN()
	case "postgres":
		return c.PostgresDSN()
	default:
		return "mysql"
	}
}
