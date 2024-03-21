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
	ErrEmptyEnvFile = errors.New("env file exist but no content found")
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

func SetEnv(key, val string) error {
	if envMap == nil {
		return ErrEmptyEnvFile
	}
	envMap[key] = val
	return nil
}

func GetEnv(detail string) (string, error) {
	data, ok := envMap[detail]
	if ok {
		return data, nil
	}

	return "", ErrDataNotExist
}
