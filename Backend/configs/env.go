package configs

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

var envMap map[string]string
var (
	ErrFileNotExist = errors.New("we can't found the env file")
	ErrDataNotExist = errors.New("env file is available but asked data is not there")
)

func ReadEnv() error {
	m, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err.Error())
		return ErrFileNotExist
	}
	envMap = m

	return nil
}

func GetEnv(detail string) (string, error) {
	data, ok := envMap[detail]
	if ok {
		return data, nil
	}

	return "", ErrDataNotExist
}
