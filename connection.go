package adoquery

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Driver uint8

const (
	SQLServer Driver = 1 << (8 - 1 - iota)
	MySql
	Sqlite
	Postgresql
	Oracle
)

type Connection struct {
	Driver Driver
	Dsn    string

	db *gorm.DB
}

// 自定义SQL查询
func (c *Connection) Query(sql string, values ...interface{}) (tx *gorm.DB) {
	tx = c.db.Raw(sql, values...)
	return
}

// 连接数据库
func (c *Connection) Connect() (err error) {
	switch c.Driver {
	case SQLServer:
		c.db, err = gorm.Open(sqlserver.Open(c.Dsn), &gorm.Config{Logger: nil})
	case MySql:
		c.db, err = gorm.Open(mysql.Open(c.Dsn), &gorm.Config{Logger: nil})
	case Sqlite:
		c.db, err = gorm.Open(sqlite.Open(c.Dsn), &gorm.Config{Logger: nil})
	case Postgresql:
		c.db, err = gorm.Open(postgres.Open(c.Dsn), &gorm.Config{Logger: nil})
	case Oracle:
		c.db, err = gorm.Open(mysql.Open(c.Dsn), &gorm.Config{Logger: nil})
	default:
		err = errors.New("未知的Driver类型")
	}
	return
}

// 断开数据库连接
func (c *Connection) Disconnect() (err error) {
	if c.db != nil {
		db, _ := c.db.DB()
		err = db.Close()
	}
	return
}
