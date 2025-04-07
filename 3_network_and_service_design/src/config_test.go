package inout

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name   string                      `yaml:"name"`
	Source map[string]DataSourceConfig `yaml:"source"`
}
type DataSourceConfig struct {
	Type          string `yaml:"type"`
	KafkaDsConfig `yaml:",inline"`
	DbConfig      `yaml:",inline"`
}

type KafkaDsConfig struct {
	ReadTimeout time.Duration `yaml:"read_timeout,omitempty"`
	Brokers     []string      `yaml:"brokers,omitempty"`
}

type DbConfig struct {
	Host string `yaml:"host,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

func TestYAML(t *testing.T) {
	conf := Config{
		Name: "app_name",
		Source: map[string]DataSourceConfig{
			"kafka1": {
				Type: "kafka",
				KafkaDsConfig: KafkaDsConfig{
					ReadTimeout: time.Duration(10 * time.Second),
					Brokers:     []string{"host1", "host2"},
				},
			},
			"db1": {
				Type:     "db",
				DbConfig: DbConfig{Host: "localhost"},
			},
		},
	}

	// marshall
	buf, err := yaml.Marshal(conf)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(buf))

	// viper.SetConfigType("yaml")
	// viper.ReadConfig(bytes.NewBuffer(buf))

	// unmarshal
	var conf1 Config
	if err := yaml.Unmarshal(buf, &conf1); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(conf, conf1) {
		t.Errorf("exp %v, got %v", conf, conf1)
	}

}
