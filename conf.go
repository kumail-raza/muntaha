package muntaha

type Configuration struct {
	DB       DBConf
	HttpPort string
	HttpHost string
}
type DBConf struct {
	Password string
	User     string
	Host     string
	Port     string
}
