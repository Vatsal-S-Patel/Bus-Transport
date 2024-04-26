package logging

import (
	"log"
	"os"
	"sync"
	"time"
)

type Logger struct {
	Mu   sync.Mutex
	File *os.File
}

var LogMe *Logger

func InitLogger() error {
	file, err := os.OpenFile("./logging/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	LogMe = &Logger{
		File: file,
	}
	return nil
}

func GetLogger() (*Logger, error) {
	if LogMe == nil {
		err := InitLogger()
		if err != nil {
			return nil, err
		}
		return LogMe, nil
	}
	return LogMe, nil
}

func (l *Logger) LogThis(data string) error {
	log.Println(data)
	if v, ok := os.LookupEnv("debug"); ok && v == "true" {
		l.Mu.Lock()
		defer l.Mu.Unlock()
		_, err := l.File.WriteString(time.Now().Local().Format("2015-02-25 11:06:39") + " " + data + "\n")
		return err
	}
	return nil
}
