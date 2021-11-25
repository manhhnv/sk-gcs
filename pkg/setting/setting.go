package setting

import (
	"github.com/go-ini/ini"
	"log"
)

type Storage struct {
	BucketName string
}

var StorageSetting = &Storage{}

type App struct {
	HttpPort int
}

var AppSetting = &App{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("storage", StorageSetting)
	mapTo("server", AppSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
