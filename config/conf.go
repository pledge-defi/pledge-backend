package config

var Config *Conf

type Conf struct {
	Mysql        MysqlConfig
	Redis        RedisConfig
	TestNet      TestNetConfig
	MainNet      MainNetConfig
	Token        TokenConfig
	DefaultAdmin DefaultAdminConfig
	Jwt          JwtConfig
	Env          EnvConfig
}

type EnvConfig struct {
	Port       string `toml:"port"`
	Version    string `toml:"version"`
	DomainName string `toml:"domain_name"`
	Protocol   string `toml:"protocol"`
}
type DefaultAdminConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`
	ExpireTime int    `toml:"expire_time"` // duration, s
}

type TokenConfig struct {
	LogoUrl string `toml:"logo_url"`
}

type MysqlConfig struct {
	Address      string `toml:"address"`
	Port         string `toml:"port"`
	DbName       string `toml:"db_name"`
	UserName     string `toml:"user_name"`
	Password     string `toml:"password"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxLifeTime  int    `toml:"max_life_time"`
}

type TestNetConfig struct {
	ChainId              string `toml:"chain_id"`
	NetUrl               string `toml:"net_url"`
	PledgePoolToken      string `toml:"pledge_pool_token"`
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"`
}

type MainNetConfig struct {
	ChainId              string `toml:"chain_id"`
	NetUrl               string `toml:"net_url"`
	PledgePoolToken      string `toml:"pledge_pool_token"`
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"`
}

type RedisConfig struct {
	Address     string `toml:"address"`
	Port        string `toml:"port"`
	Db          string `toml:"db"`
	Password    string `toml:"password"`
	MaxIdle     int    `toml:"max_idle"`
	MaxActive   int    `toml:"max_active"`
	IdleTimeout int    `toml:"idle_timeout"`
}
