package cache

import (
  "log"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Config struct {
  Cache struct {
    Urls []string `yaml:"urls"`
  }
}

func GetConfig() Config {
  config := Config{}
  yamlFile, err := ioutil.ReadFile("config.yml")
  if err != nil {
      log.Printf("yamlFile.Get err   #%v ", err)
  }
  err = yaml.Unmarshal(yamlFile, &config)
  if err != nil {
      log.Fatalf("Unmarshal: %v", err)
  }

  return config
}
