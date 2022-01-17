package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var Config *config

func init() {
	time.Local = time.FixedZone("CST", 8*3600) // 东八
	fmt.Println("设置时区", time.Now())

	confUrl := os.Getenv("CONFIG_URI")
	Config = getConfig(confUrl)
}

func getConfig(url string) *config {
	e := "请求配置出错"
	resp, err := http.Get(url)
	FatalHandler(err, e)

	body, err := io.ReadAll(resp.Body)
	FatalHandler(err, e)

	FatalHandler(resp.Body.Close(), e)
	fmt.Println(string(body))

	var c *config
	FatalHandler(yaml.Unmarshal(body, c), e)
	return c
}

//config 设置结构体
type config struct {
	Semester          string    `yaml:"semester"`
	SemesterLast      string    `yaml:"semester_last"`
	FirstDay          time.Time `yaml:"first_day"`
	MysqlDsnLocal     string    `yaml:"mysql_dsn_local"`
	MysqlDsn          string    `yaml:"mysql_dsn"`
	RedisAddr         string    `yaml:"redis_addr"`
	PwdKey            string    `yaml:"pwd_key"`
	PwdIv             string    `yaml:"pwd_iv"`
	RpcAuth           string    `yaml:"rpc_auth"`
	MiniAppId         string    `yaml:"mini_app_id"`
	MiniAppSecret     string    `yaml:"mini_app_secret"`
	OffAppId          string    `yaml:"off_app_id"`
	OffAppSecret      string    `yaml:"off_app_secret"`
	OffToken          string    `yaml:"off_token"`
	OffEncodingAESKey string    `yaml:"off_encoding_AESKey"`
	EtcdAddr          string    `yaml:"etcd_addr"`
	ElasticAddr       string    `yaml:"elastic_addr"`
	ElasticUser       string    `yaml:"elastic_user"`
	ElasticPass       string    `yaml:"elastic_pass"`
	BanID             []struct {
		ID     string `yaml:"ID"`
		Reason string `yaml:"reason"`
	} `yaml:"banID"`
}
