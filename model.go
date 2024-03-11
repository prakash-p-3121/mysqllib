package mysqllib

type MySQLCfg struct {
	HostAddr     string `toml:"host-name"`
	Port         uint   `toml:"port"`
	UserName     string `toml:"user-name"`
	Password     string `toml:"password"`
	DatabaseName string `toml:"database-name"`
}
