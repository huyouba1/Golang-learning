package conf_test

import (
	"gitee.com/go-learn/restful-api-demo-g7/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("demo", conf.C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("MYSQL_DATABASE", "unit_test")
	should := assert.New(t)
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}
}
