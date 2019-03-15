package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func VerifyEnvironment() bool {
	err := godotenv.Load("./.env")
	if err != nil {
		println("No .env file found, please copy the sample file via `cp sample.env .env` or run the config command again with the necessary flags: -e -d -o")
		return false
	}

	if os.Getenv("CERTBOT_EMAIL") == "" {
		fmt.Print("Email is blank, please run 'accelerator config -e email@example.com' to set your email.")

		return false
	}

	if os.Getenv("DOMAIN_NAME") == "" {
		fmt.Print("Domain name is blank, please run 'accelerator config -d example.com' to set the domain name for your droplet. Be sure to have this set up in Digital Ocean prior. Certbot limits requests to 5 per week. Running this command without proper configuration can cause a delay in certificate retrieval. It is recommended to test this command with an A record first, like test.example.com.")

		return false
	}

	if os.Getenv("ORIGIN_HOST") == "" {
		fmt.Print("Origin Host is blank, please run 'accelerator config -o 206.45.X.X' to set your origin server address. This is the previous server you wish to enhance with Gladius Accelerator.")

		return false
	}

	return true
}

func writeEnvVariable(key, value string) {
	envFile, err := godotenv.Read("./.env")
	envFile[key] = value

	content, err := godotenv.Marshal(envFile)
	if err != nil {
		log.Fatalf("Error writing %s to .env file", value)
	}

	formattedContent := strings.Replace(content, "\"", "", -1)
	err = ioutil.WriteFile("./.env", []byte(formattedContent), 0644)

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
