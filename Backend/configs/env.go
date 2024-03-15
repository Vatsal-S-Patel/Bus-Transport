package configs

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

var envMap map[string]string

func ReadEnv() error {
	m, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	envMap = m

	return nil
}

func GetEnv(detail string) (string, error) {
	data, ok := envMap[detail]
	if ok {
		return data, nil
	}

	return "", errors.New("ERROR: Detail not exist in env file")
}
