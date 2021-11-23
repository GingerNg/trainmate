package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/revel/config"
)

var GlobalConfig *Config

// 常量
const (
	CONFIG_FILE_PATH = "conf/config.ini" //"/opt/project/DBM/conf/config.ini"
	DEFAULT_TCP_BIND = ":8080"
	DATA_PATH        = "/data/images"

	SQLDialectMysql    = "mysql"
	SQLDialectPostgres = "postgres"
	SQLDialectSqlite   = "sqlite3"
)

func IsDev(env string) bool {
	return env == "dev" || env == "development"
}

func (c *Config) IsDev() bool {
	return IsDev(c.Env)
}

type Config struct {
	Env     string `default:"dev" env:"ENVIRONMENT"`
	Version string `yaml:"-"`
	// App      appConfig
	// Security securityConfig
	Db       dbConfig
	UseMongo bool
	// Server   serverConfig
	// Sentry   sentryConfig
	// Mail     mailConfig

	// server config
	__tcp_bind     string //tcp监听地址和端口
	__server_name  string // 服务名
	__data_path    string
	__history_path string

	// log config
	__log_path          string // 日志目录
	__job_log_path      string // 日志目录
	__callback_log_path string // 日志目录
	__log_level         string // 日志级别
	__log_format        string // 日志format
	__log_expire_days   int    // 日志保存天数

	__store_mode string

	// ucloud - ufile
	__ufile_publickey  string
	__ufile_privatekey string
	__ufile_bucketname string
	__ufile_filehost   string
	__ufile_buckethost string
	__ufile_private    bool

	//MYSQL CONFIG
	__mysql_datasource        string //mysql dataSource
	__mysql_datasource_cyprex string //mysql dataSource
	__mysql_drivername        string //mysql driverName
	__mysql_addr              string //mysql ip地址
	__mysql_user              string //mysql 登陆用户名
	__mysql_password          string //mysql 登陆密码
	__mysql_dbname            string //mysql 数据库名称
	__mysql_max_idle_conn     int    //mysql 最大空闲连接数
	__mysql_max_conn          int    //mysql 最大连接数
}

func NewGlobalConfig() *Config {
	var err error
	GlobalConfig, err = newGlobalConfig()
	if err != nil {
		panic(err)
	}
	return GlobalConfig
}

func newGlobalConfig() (*Config, error) {

	GlobalConfig = new(Config)

	cfg, err := config.ReadDefault(path.Join(os.Getenv("DBM_PATH"), CONFIG_FILE_PATH))
	if err != nil {
		fmt.Println(err)
		return GlobalConfig, err
	}

	GlobalConfig.Db.Dialect = "sqlite3"
	GlobalConfig.Db.Name = "mate_db.db"

	GlobalConfig.UseMongo, _ = cfg.Bool("server", "use_mongo")
	fmt.Println(GlobalConfig.UseMongo)

	GlobalConfig.__server_name, _ = cfg.String("server", "name")
	GlobalConfig.__data_path, _ = cfg.String("server", "data_path")
	GlobalConfig.__history_path, _ = cfg.String("server", "history_path")
	// tcp 绑定端口
	GlobalConfig.__tcp_bind, err = cfg.String("server", "tcp_bind")
	if err != nil {
		GlobalConfig.__tcp_bind = DEFAULT_TCP_BIND
	}

	// 日志配置
	GlobalConfig.__log_path, _ = cfg.String("log", "log_path")
	GlobalConfig.__job_log_path, _ = cfg.String("log", "job_log_path")
	GlobalConfig.__callback_log_path, _ = cfg.String("log", "callback_log_path")
	GlobalConfig.__log_level, _ = cfg.String("log", "level")
	GlobalConfig.__log_format, _ = cfg.String("log", "format")
	GlobalConfig.__log_expire_days, _ = cfg.Int("log", "days")

	GlobalConfig.__store_mode, _ = cfg.String("store", "mode")

	// Ucloud ufile
	GlobalConfig.__ufile_publickey, _ = cfg.String("ufile", "public_key")
	GlobalConfig.__ufile_privatekey, _ = cfg.String("ufile", "private_key")
	GlobalConfig.__ufile_bucketname, _ = cfg.String("ufile", "user_source_bucket")
	GlobalConfig.__ufile_filehost, _ = cfg.String("ufile", "uploadsuffix")

	// MYSQL 配置信息
	GlobalConfig.__mysql_datasource, _ = cfg.String("mysql", "dataSource")
	GlobalConfig.__mysql_datasource_cyprex, _ = cfg.String("mysql", "dataSource_cyprex")
	GlobalConfig.__mysql_drivername, _ = cfg.String("mysql", "driverName")
	GlobalConfig.__mysql_max_conn, _ = cfg.Int("conf", "maxConn")
	GlobalConfig.__mysql_max_idle_conn, _ = cfg.Int("conf", "maxIdle")

	return GlobalConfig, nil
}

func GetGlobalConfig() *Config { // 单例
	if GlobalConfig != nil {
		return GlobalConfig
	}
	c, err := newGlobalConfig()
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func (c *Config) GetServerName() string {
	return c.__server_name
}

func (c *Config) GetTcpBind() string {
	return c.__tcp_bind
}

func (c *Config) GetDataPath() string {
	return c.__data_path
}

func (c *Config) GetHistoryPath() string {
	return c.__history_path
}

func (c *Config) GetLogPath() string {
	return c.__log_path
}

func (c *Config) GetJobLogPath() string {
	return c.__job_log_path
}

func (c *Config) GetCallbackLogPath() string {
	return c.__callback_log_path
}

func (c *Config) GetLogLevel() string {
	return c.__log_level
}

func (c *Config) GetLogFormat() string {
	return c.__log_format
}

func (c *Config) GetLogExpireDays() int {
	return c.__log_expire_days
}

func (c *Config) GetStoreMode() string {
	return c.__store_mode
}

func (c *Config) GetUfilePublickey() string {
	return c.__ufile_publickey
}

func (c *Config) GetUfilePrivateKey() string {
	return c.__ufile_privatekey
}

func (c *Config) GetUfileBucketName() string {
	return c.__ufile_bucketname
}

func (c *Config) GetUfileFileHost() string {
	return c.__ufile_filehost
}

func (c *Config) GetMysqlDatasource() string {
	return c.__mysql_datasource
}

func (c *Config) GetMysqlDatasourceCyprex() string {
	return c.__mysql_datasource_cyprex
}

func (c *Config) GetMysqlDrivername() string {
	return c.__mysql_drivername
}

func (c *Config) GetMysqlMaxConn() int {
	return c.__mysql_max_conn
}

func (c *Config) GetMysqlAddr() string {
	return c.__mysql_addr
}

func (c *Config) GetMysqlUser() string {
	return c.__mysql_user
}

func (c *Config) GetMysqlPassword() string {
	return c.__mysql_password
}

func (c *Config) GetMysqlDBName() string {
	return c.__mysql_dbname
}

func (c *Config) GetMysqlMaxIdleConn() int {
	return c.__mysql_max_idle_conn
}
