package config

import (
	"fmt"

	"github.com/tzrd/go-antdv-core/utils"
)

type ServerConf struct {
	Name    string
	Host    string
	Port    int
	Network string `json:",optional,default=tcp"`
	Debug   bool   `json:",optional,default=false"`
}

func (c ServerConf) GetServerAddr() string {
	if !utils.IsIPv4(c.Host) {
		panic("Error: not a valid ipv4")
	}

	if c.Port <= 0 {
		panic("Error: not a valid port")
	}

	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}
