package env

import (
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"strings"
)

func Initialize() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func writeEnvVariable(key, value string) {
	envFile, err := godotenv.Read("./.env")
	envFile[key] = value

	content, err := godotenv.Marshal(envFile)
	if err != nil {
		log.Fatalf("Error writing %s to .env file", value)
	}

	formattedContent := strings.Replace(content, "\"", "", -1)
	err = ioutil.WriteFile("./env", []byte(formattedContent), 0644)

	//err = godotenv.Write(envFile, "./.env")
	if err != nil {
		log.Fatalf("Error writing %s to .env file", value)
	}
}

func SetDomain(domain string) {
	writeEnvVariable("DOMAIN_NAME", domain)
}

func SetEmail(email string) {
	writeEnvVariable("CERTBOT_EMAIL", email)
}

func SetOriginHost(host string) {
	writeEnvVariable("ORIGIN_HOST", host)
}
