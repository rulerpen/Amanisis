package config

type Config struct{
	MysqlServers MysqlServer `toml:"mysql"`
	LogServers LogServer `toml:"log"`
}

type MysqlServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type LogServer struct {
	LogUrl string `toml:"url"`
	DebugLogUrl string `toml:"debug_url"`
}