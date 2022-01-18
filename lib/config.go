package lib

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"time"
)

var Config *config

const ConfEnvName = "CONFIG_URI"

func init() {
	time.Local = time.FixedZone("CST", 8*3600) // 东八
	logrus.Info("设置时区", time.Now())

	confUrl := os.Getenv(ConfEnvName)
	//logrus.Info(ConfEnvName, confUrl)

	b, err := getConfigFromUrl(confUrl)
	if err != nil {
		logrus.Info("请求网络配置出错", err)
		b, err = os.ReadFile("./config.yml")
		if err != nil {
			logrus.Fatal("读取本地配置出错")
		}
	}

	FatalHandler(yaml.Unmarshal(b, &Config), "解析yml出错")

}

func getConfigFromUrl(url string) (res []byte, err error) {
	if url == "" {
		err = errors.New("没设置CONFIG_URI环境变量")
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	res, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
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
}
