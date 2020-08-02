package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
	"time"
)

const FILEPATH = "redis_config"
const FILETYPE = "yaml"

// Config client settings.
type Config struct {
	Name         string // redis name, for trace
	Proto        string // tcp
	Addr         string // "127.0.0.1:6379"
	Auth         string // passwd
	DB           string
	DialTimeout  time.Duration // 拨号超时
	ReadTimeout  time.Duration // 读取超时
	WriteTimeout time.Duration // 写入超时
}

func InitReadConfig(KeyName, FilePath, FileType string) *Config {
	if FilePath == "" {
		FilePath = FILEPATH
	}
	if FileType == "" {
		FileType = FILETYPE
	}
	viper.AddConfigPath(".")
	viper.SetConfigName(FilePath) // 读取文件名称为
	viper.SetConfigType(FileType)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		panic(err) // 读取配置文件失败致命错误
	}

	return &Config{
		viper.GetString(KeyName + ".Name"),
		viper.GetString(KeyName + ".Proto"),
		viper.GetString(KeyName + ".Addr"),
		viper.GetString(KeyName + ".Auth"),
		viper.GetString(KeyName + ".DB"),
		viper.GetDuration(KeyName + ".DialTimeout"),
		viper.GetDuration(KeyName + ".ReadTimeout"),
		viper.GetDuration(KeyName + ".WriteTimeout"),
	}
}

// 纯粹的执行命令版本,
// 将一些命令进行了封装的包 https://github.com/go-redis/redis
func ConnectRedis(keyName, filePath, fileType string) *redis.Pool {
	config := InitReadConfig(keyName, filePath, fileType)
	pool, err := PoolInitRedis(config)
	if err != nil {
		panic(err)
	}
	return pool
}

// redis pool
func PoolInitRedis(config *Config) (*redis.Pool, error) {

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(config.Proto, config.Addr)
			if err != nil {
				return nil, err
			}
			if config.Auth != "" {
				if _, err := c.Do("AUTH", config.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", config.DB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	return pool, nil

}
