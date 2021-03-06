package utils

//App : Struct for App (dns server) configuration in the config.ini file
type App struct {
	Port    int
	IP      string `ini:"IP"`
	Logdir  string
	Logfile bool
}

//Database : Struct for SQL Database configuration in the config.ini file
type Database struct {
	IP       string `ini:"IP"`
	Port     string
	Username string
	Password string
	Db       string `ini:"DB"`
	Type     string
}

//Redis : Struct for Redis  Database configuration in the config.ini file
type Redis struct {
	IP       string `ini:"IP"`
	Port     int
	Password string
	Db       int `ini:"DB"`
	TTL      int `ini:"TTL"`
}

//Conf : Struct for the whole config.ini file when it will be parsed by go-ini
type Conf struct {
	AppMode string `ini:"app_mode"`
	App
	Database
	Redis
}

//Record : Struct for a domain record
//Defined by it's ID, DomainID (parent domain), Fqdn (or name), Content (value of the record), Type (as Qtype/int), TTL (used only for the DNS response and not the Redis TTL)
type Record struct {
	ID       uint `gorm:"primaryKey"`
	DomainID int
	Fqdn     string
	Content  string
	Type     int
	Qtype    uint16 `gorm:"-"`
	TTL      int
}
