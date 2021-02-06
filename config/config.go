package config

const (
	Host       string = "http://localhost:1323"
	LimitQuery uint64 = 100

	MariaDBUser     string = "root"
	MariaDBPassword string = "root@qdns"
	MariaDBDB       string = "quantumdns"
	MariaDBHost     string = "localhost"
	MariaDBPort     string = "3306"

	RedisHost string = "localhost"
	RedisPort string = "6379"

	DataBlockPath string = "/var/www/html/datablock/abuse.txt"
	BlockCategory string = "1"
)
