package config

const (
	Host       string = "http://localhost:1323"
	LimitQuery uint64 = 100

	MariaDBUser     string = "dnsfilter"
	MariaDBPassword string = "rahasiadns"
	MariaDBDB       string = "dnsfilter"
	MariaDBHost     string = "175.106.8.72"
	MariaDBPort     string = "3306"

	RedisHost string = "localhost"
	RedisPort string = "6379"

	DataBlokPath string = "/var/www/html/datablok/youtube.txt"
	BlokCategory string = "18"
)
