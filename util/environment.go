package util

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/geovani-moc/geac/entity"
	"github.com/joho/godotenv"
)

//LoadEnv load variables environment
func LoadEnv() {
	err := godotenv.Load(".env") // fazer tratamento de erro caso o arquivo nao exista
	if nil != err {
		log.Print("Erro ao carregar variaveis de ambiente, ", err)
	}
}

//StoreEnv store env file
func StoreEnv(env entity.User) {
	out := fmt.Sprintf("EMAIL_USER=%v\n", env.User)
	out += fmt.Sprintf("EMAIL_PASSWORD=%v\n", env.Password)

	err := ioutil.WriteFile(".env", []byte(out), 0644)
	if nil != err {
		log.Print("erro ao escrever aquivo, ", err)
	}
}
