package config

import (
	"io/ioutil"
	"fmt"
	"os"
	"regexp"
	"encoding/json"
	"github.com/ninorain22/gin_starter/util"
	"github.com/shen100/golang123/utils"
	"time"
)

type dbConfig struct {
	Dialect string
	Database string
	User string
	Password string
	Host string
	Port int
	Charset string
	URL string
	MaxIdleConns int
	MaxOpenConns int
	Prefix string
}

var DBConfig dbConfig

type redisConfig struct {
	Host      string
	Port      int
	Password  string
	URL       string
	MaxIdle   int
	MaxActive int
	Expire    time.Duration
}

var RedisConfig redisConfig

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

func initDB() {
	util.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}

func initRedis() {
	utils.SetStructByJSON(&RedisConfig, jsonData["redis"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port)
	RedisConfig.URL = url
	// default expire: 1 day
	RedisConfig.Expire = time.Hour * 12
}

func init()  {
	initJSON()
	initDB()
	initRedis()
}