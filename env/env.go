package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func VerifyEnvironment() bool {
	dirPath := viper.GetString("DIR_PATH")

	err := godotenv.Load(dirPath + "/.env")
	if err != nil {
		fmt.Println("No .env file found, run the config command again with the necessary flags: -e -d -o")
		return false
	}

	if os.Getenv("CERTBOT_EMAIL") == "" {
		fmt.Println("Email is blank, please run 'accelerator config -e email@example.com' to set your email.")

		return false
	}

	if os.Getenv("DOMAIN_NAME") == "" {
		fmt.Println("Domain name is blank, please run 'accelerator config -d example.com' to set the domain name for your droplet. Be sure to have this set up in Digital Ocean prior. Certbot limits requests to 5 per week. Running this command without proper configuration can cause a delay in certificate retrieval. It is recommended to test this command with an A record first, like test.example.com.")

		return false
	}

	if os.Getenv("ORIGIN_HOST") == "" {
		fmt.Println("Origin Host is blank, please run 'accelerator config -o 206.45.X.X' to set your origin server address. This is the previous server you wish to enhance with Gladius Accelerator.")

		return false
	}

	return true
}

func writeEnvVariable(key, value string) {
	dirPath := viper.GetString("DIR_PATH")
	envFile, err := godotenv.Read(dirPath + "/.env")
	envFile[key] = value

	content, err := godotenv.Marshal(envFile)
	if err != nil {
		log.Fatalf("Error writing %s to .env file", value)
	}

	formattedContent := strings.Replace(content, "\"", "", -1)
	err = ioutil.WriteFile(dirPath+"/.env", []byte(formattedContent), 0644)

	if err != nil {
		log.Fatalf("Error writing %s to %s/.env file", value, dirPath)
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
