package storage

import (
	"fmt"

	"github.com/ilibs/gosql/v2"
)

type MySQL struct {
	config *gosql.Config
}

func NewMySqlConfig(enable bool, driver string, dsn string, showSql bool) *MySQL {
	return &MySQL{
		config: &gosql.Config{
			Enable:  enable,
			Driver:  driver,
			Dsn:     dsn,
			ShowSql: showSql,
		},
	}
}

func (c *MySQL) Connect() bool {
	_config := make(map[string]*gosql.Config)
	_config["default"] = c.config

	err := gosql.Connect(_config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false
	}

	return true
}

func (c *MySQL) Count(num *int64, sql string, args ...interface{}) error {
	return gosql.QueryRowx(sql, args...).Scan(num)
}

func (c *MySQL) Find(sql string, args ...interface{}) {
	gosql.Exec(sql, args...)
}

func (c *MySQL) Get(dest interface{}, sql string, args ...interface{}) error {
	return gosql.Get(dest, sql, args...)
}

func (c *MySQL) Query(dest interface{}, sql string, args ...interface{}) error {
	return gosql.Select(dest, sql, args...)
}

func (c *MySQL) Drop(sql string) (int64, error) {
	r, e := gosql.Exec(sql)
	if e != nil {
		return 0, e
	}

	return r.RowsAffected()
}

func (c *MySQL) Delete(sql string) (int64, error) {
	r, e := gosql.Exec(sql)
	if e != nil {
		return 0, e
	}

	return r.RowsAffected()
}

func (c *MySQL) Insert(sql string) (int64, error) {
	r, e := gosql.Exec(sql)
	if e != nil {
		return 0, e
	}

	return r.RowsAffected()
}

func (c *MySQL) Update(table string, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	a, e := gosql.Table(table).Where(where, args...).Update(data)
	if e != nil {
		return 0, e
	}

	return a, nil
}
