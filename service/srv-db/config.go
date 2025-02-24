package srv_db

type DBConfig struct {
	Driver string `yaml:"driver"`
	Master string `yaml:"master"`
}
