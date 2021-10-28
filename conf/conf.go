package conf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
)

type ConfType int

const (
	CONF = "application"

	JSON ConfType = iota
	YAML
	TOML
)

type ExtMetadata struct {
	Name string
	Ext  string
}

var Extensions = map[ConfType]*ExtMetadata{
	JSON: {
		Name: "json",
		Ext:  ".json",
	},
	YAML: {
		Name: "yaml",
		Ext:  ".yml",
	},
	TOML: {
		Name: "toml",
		Ext:  ".toml",
	},
}

func ParseConfType(t string) ConfType {
	for k, v := range Extensions {
		if strings.EqualFold(t, v.Name) {
			return k
		}
	}

	panic("type err!")
}

var logger = golog.New()

type Conf struct {
	iris.Configuration `yaml:"Configuration"`

	AutoMigrate bool   `json:"autoMigrate" yaml:"AutoMigrate" toml:"AutoMigrate"`
	Level       string `json:"level" yaml:"Level" toml:"Level"`
	Port        int    `json:"port" yaml:"Port" toml:"Port"`
	MySQL       string `json:"mysql,omitempty" yaml:"Mysql" toml:"Mysql"`
	Redis       string `json:"redis,omitempty" yaml:"Redis" toml:"Redis"`
}

func ReadConf(t ConfType) *Conf {
	c := Conf{
		Port:  9000,
		Level: "warn",
	}
	c.Configuration = iris.DefaultConfiguration()
	js, err := os.ReadFile(CONF + Extensions[t].Ext)
	if err != nil {
		logger.Error("Open file err:" + err.Error())
		panic(err)
	}

	switch t {
	case JSON:
		if err = json.Unmarshal(js, &c); err != nil {
			logger.Error("read config error.")
			panic(err)
		}
	case YAML:
		if err = yaml.Unmarshal(js, &c); err != nil {
			logger.Error("read config error.")
			panic(err)
		}
	case TOML:
		if _, err = toml.Decode(string(js), &c); err != nil {
			logger.Error("read config error.")
			panic(err)
		}
	}

	return &c
}

func (cfg *Conf) Save(t ConfType, path string) {
	f := filepath.Join(path, CONF+Extensions[t].Ext)
	switch t {
	case JSON:
		if js, err := json.MarshalIndent(cfg, "", "	"); err == nil {
			os.WriteFile(f, js, 0666)
		} else {
			logger.Error(err.Error())
		}
	case YAML:
		if js, err := yaml.Marshal(cfg); err == nil {
			os.WriteFile(f, js, 0666)
		} else {
			logger.Error(err.Error())
		}
	case TOML:
		if fs, err := os.Create(f); err == nil {
			defer fs.Close()

			err = toml.NewEncoder(fs).Encode(cfg)
			if err != nil {
				logger.Error(err)
			}
		}
	default:
		panic("not support type.")
	}
}
