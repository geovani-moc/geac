package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//Configuration struct
type Configuration struct {
	User     string
	Password string
	SMTPHost string
	SMTPPort string
}

//LoadConfig load configuration variables
func LoadConfig(path string) (*Configuration, error) {
	var config Configuration
	file, err := os.Open(path)
	if nil != err {
		log.Print("Erro ao abrir arquivo de configuração, ", err)
		return nil, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if nil != err {
		log.Print("Erro ao decodificar arquivo de configuração, ", err)
		return nil, err
	}

	return &config, nil
}

//StoreConfig store config file
func StoreConfig(path string, config *Configuration) error {
	file, err := json.Marshal(*config)
	if nil != err {
		log.Print("Erro ao salvar arquivo de configuraçoes, ", err)
		return err
	}

	err = ioutil.WriteFile(path, file, 0644)
	if nil != err {
		log.Print("erro ao escrever aquivo, ", err)
		return err
	}

	return nil
}
