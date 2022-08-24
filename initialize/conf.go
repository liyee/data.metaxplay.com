package initialize

import (
	"flag"
	"fmt"
	"os"

	"data.metaxplay.com/common"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConf(path ...string) {

	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose conf file.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv("CONFFILE"); configEnv == "" {
				config = common.ConfigFile
			} else {
				config = configEnv
			}
		}
	} else {
		config = path[0]
	}
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("conf file changed:", e.Name)
		if err = v.Unmarshal(&common.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&common.CONFIG); err != nil {
		fmt.Println(err)
	}
}
