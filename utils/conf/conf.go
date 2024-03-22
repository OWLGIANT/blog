package conf

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type server struct {
	Port string `toml:"port"`
}

type redis struct {
	Server   string `toml:"server"`
	Pwd      string `toml:"pwd"`
	DB       int    `toml:"db"`
	PoolSize int    `toml:"poolSize"`
}

type influxDB struct {
	Server    string `toml:"server"`
	Token     string `toml:"token"`
	Org       string `toml:"org"`
	BucketRTT string `toml:"bucketRichRtt"`
}

type database struct {
	Host   string `toml:"host"`
	Port   string `toml:"port"`
	User   string `toml:"user"`
	Pwd    string `toml:"pwd"`
	DbName string `toml:"dbName"`
}

type jwt struct {
	Secret string
}
type commonCfg struct {
	Title   string `json:"title"`
	AppName string `json:"appName"`
	Type    string `json:"type"`
	AdminIp string `json:"adminIp"`
}

type config struct {
	Mutex    sync.RWMutex
	Server   server
	Redis    redis
	InfluxDB influxDB
	DB       database `toml:"database"`
	Email    email
	Jwt      jwt
	// common
	Common commonCfg
}

type email struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

var (
	etcClient  *clientv3.Client
	Conf       config
	configFile = "utils/conf/conf.toml" //服务器本地toml文件

	cPath    = "/local_server" //etcd 远程 配置地址
	etcdNode = "57.180.139.206:80"
	eName    = "root"
	ePwd     = "123456789"
)

func InitConfig() {
	if err := EtcInit(); err != nil { //1、获取etcd  配置
		logrus.Error(err)
		if _, err = toml.Decode(readLocalConfigFile(), &Conf); err != nil { //2、获取本地  配置
			panic("config decode err:" + err.Error())
		}
	}
}

func EtcInit() (err error) {
	// ETCD客户端连接信息
	etcConf := clientv3.Config{
		Endpoints:   []string{etcdNode}, // 节点信息
		DialTimeout: 5 * time.Second,    // 超时时间
		Username:    eName,
		Password:    ePwd,
	}

	if etcClient, err = clientv3.New(etcConf); err != nil {
		return fmt.Errorf("conect to etcd faild, err:%v\n", err)
	} else {
		logrus.Infoln(cPath, "etcd V3 connect to etcd success")
	}
	if err = readConfigFromEtcd(etcClient); err != nil { // 从etcd中读取初始配置

		return
	}
	go watchConfigChanges(etcClient) // 启动监视配置更改的协程
	return
}

func readLocalConfigFile() (content string) {
	if _, err := os.Stat(configFile); err != nil {
		panic("config file err:" + err.Error())
	} else {
		configBytes, rerr := os.ReadFile(configFile)
		if rerr != nil {
			panic("config load err:" + rerr.Error())
		}
		return string(configBytes)
	}
	return
}

func readConfigFromEtcd(client *clientv3.Client) (err error) {
	//根据本地文件设置配置
	if err = nodePutInit(client, cPath, readLocalConfigFile()); err != nil {
		return
	}
	// 获取配置
	resp, err := client.Get(context.Background(), cPath)
	if err != nil {
		return
	}
	// 解析配置
	Conf.Mutex.Lock()
	defer Conf.Mutex.Unlock()
	// 解码etcd中的值到结构体
	if err = toml.Unmarshal(resp.Kvs[0].Value, &Conf); err != nil {
		return err
	} else {
		if err = upduateTxt(configFile, string(resp.Kvs[0].Value)); err != nil { //读取到数据并且能够解析\再更新本地toml文件
			return err
		}
	}
	return
}

func watchConfigChanges(etcClient *clientv3.Client) {
	rch := etcClient.Watch(context.Background(), cPath)
	for wresp := range rch {
		// 处理事件...
		for _, ev := range wresp.Events {
			// 当检测到 Put 事件时
			if ev.Type == clientv3.EventTypePut {
				// 对全局配置加锁，确保并发安全
				Conf.Mutex.Lock()
				err := toml.Unmarshal(ev.Kv.Value, &Conf)
				if err != nil {
					fmt.Println("解码更新的配置时发生错误:", err)
				} else {
					fmt.Println("配置更新成功")
					// 尝试解码更新后的配置
					if err = upduateTxt(configFile, string(ev.Kv.Value)); err != nil {
						fmt.Println("更新本地配置失败", err)
					}
				}
				Conf.Mutex.Unlock()
			}
		}
	}
}

func upduateTxt(filePath, content string) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	if _, err = file.WriteString(content); err != nil {
		return
	}
	if err = file.Sync(); err != nil {
		return
	}
	return
}

// 在  etcd  创建配置
func nodePutInit(etcClient *clientv3.Client, keypath, value string) (err error) {
	_, err = etcClient.Put(context.Background(), keypath, value)
	return
}
