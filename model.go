package mysqllib

import (
	_ "github.com/pelletier/go-toml/v2"
)

type MySQLCfg struct {
	HostAddr     string `toml:"host"`
	Port         uint   `toml:"port"`
	UserName     string `toml:"user-name"`
	Password     string `toml:"password"`
	DatabaseName string `toml:"database-name"`
}
