package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		// 初始化失败
		return err
	}

	// 初始化日志包
	c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 指定配置文件，捷信指定的文件
		viper.SetConfigFile(c.Name)
	} else {
		// 默认文件
		viper.AddConfigPath("conf")   // 目录
		viper.SetConfigName("config") // 文件名
	}
	// 设置配置文件格式为yaml格式
	viper.SetConfigType("yaml")
	// 自动匹配环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为APISERVER
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		// viper解析文件错误
		return err
	}
	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("config file changed : %s", in.Name)
	})
}
