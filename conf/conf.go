package conf

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
)

type ConfType int

const (
	CONF = "application.json"

	JSON ConfType = iota
	YAML
	TOML
)

type Conf struct {
	iris.Configuration `yaml:"Configuration"`

	Level string `json:"level" yaml:"Level" toml:"Level"`
	Port  int    `json:"port" yaml:"Port" toml:"Port"`
	MySQL string `json:"mysql" yaml:"Mysql" toml:"Mysql"`
	Redis string `json:"redis" yaml:"Redis" toml:"Redis"`
}

func ReadConf() *Conf {
	c := Conf{
		Port:  9000,
		Level: "warn",
	}
	c.Configuration = iris.DefaultConfiguration()

	js, err := os.ReadFile(CONF)
	if err != nil {
		fmt.Println("Open file err:" + err.Error())
		panic(err)
	}

	if err = json.Unmarshal(js, &c); err != nil {
		fmt.Println("read config error.")
		panic(err)
	}

	return &c
}

func (cfg *Conf) Save(path string, t ConfType) {
	switch t {
	case JSON:
		if js, err := json.MarshalIndent(cfg, "", "	"); err == nil {
			os.WriteFile(path, js, 0666)
		} else {
			fmt.Println(err.Error())
		}
	case YAML:
		if js, err := yaml.Marshal(cfg); err == nil {
			os.WriteFile(path, js, 0666)
		} else {
			fmt.Println(err.Error())
		}
	case TOML:
		if fs, err := os.Create(path); err == nil {
			defer fs.Close()

			err = toml.NewEncoder(fs).Encode(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	default:
		panic("not support type.")
	}
}
