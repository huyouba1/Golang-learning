package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// 全局config实例对象
// 也就是我们程序在内存中配置对象
// 程序内部获取配置，都通过读取该对象来获取
// 该Config对象什么是被初始化？
//
//	  配置加载时
//		LoadConfigFromToml
//		LoadConfigFromEnv
//
// 为了不被程序在运行时额已修改，设置成私有变量
var config *Config

// 全局mysql客户端实例
var db *sql.DB

// 要想获取配置，单独提供一个函数
// 全局Config对象获取函数
func C() *Config {
	return config
}

// 初始化一个带有默认值的Config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMysql(),
	}
}

// Config 应用配置
// 通过封装为一个对象，来与外部进行对接
type Config struct {
	App   *App   `toml:"App"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	//Key  string `toml:"key" env:"APP_KEY"`
	//	EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	//	CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	//	KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

func NewDefaultLog() *Log {
	return &Log{
		// debug,info,error,warn
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// Log todo
// 用于配置全局logger对象
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
}

func NewDefaultMysql() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "demo",
		Password:    "123456",
		Database:    "demo",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

// MySQL todo
type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 因为使用的mysql的连接池，需要对池做一些规划配置
	// 控制当前程序的mysql打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制mysql连接的复用，比如5，最多运行5个来复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 连接的生命周期，这个和mysql server配置有关系，必须小于server配置
	// 一个连接用12h 换一个conn，保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle连接最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	// 小写，作为私有变量，用于控制db
	lock sync.Mutex
}

// 1. 第一种方式，使用LoadGloabl在加载时初始化全局db实例
// 2. 第二种方式，惰性加载，获取是动态判断再初始化
func (m *MySQL) GetDB() *sql.DB {
	// 直接加锁，锁住临界区
	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果实例不存在,就初始化一个新的实例
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		db = conn
	}
	// 全局变量db就一定存在了
	return db
}

// 连接池，driverConn具体的连接对象，它维护着一个Socket
// pool []*driverConn, 维护pool里面的连接都是可用的，定时检查conn的健康情况
// 某一个driverConn已经失效了，driverConn.ReSet()，清空该结构体的数据，Reconn获取一个新的连接，让该conn借壳存活
// 避免driverConn结构体的内存申请和释放的一个成本
func (m *MySQL) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}
