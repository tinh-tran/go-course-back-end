package common

import (
	. "PaymentService/config"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", timestamp, f.LevelDesc[entry.Level], entry.Message)), nil
}

func initLogger(config Config) {
	f, err := os.OpenFile(config.LogDir, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	plainFormatter := new(PlainFormatter)
	plainFormatter.TimestampFormat = "2006-01-02 15:04:05"
	plainFormatter.LevelDesc = []string{"PANIC", "FATAL", "ERRO", "WARN", "INFO", "DEBUG"}
	log.SetFormatter(plainFormatter)

	if err != nil {
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
	Level, err := log.ParseLevel(config.LogLevel)
	log.SetLevel(Level)
}
