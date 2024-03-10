package mysqllib

type MySQLCfg struct {
	HostName     string `toml:"host-name"`
	Port         uint   `toml:"port"`
	DatabaseName string `toml:"database-name"`
}
