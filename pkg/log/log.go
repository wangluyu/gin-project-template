package log

import (
	"gin-project-template/pkg/util"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log = logrus.New()
var timeStr = time.Now().Format("20060102")
var fileName string
var wfFileName string
var Info logrus.Fields

func init() {
	Info = make(logrus.Fields)
	logrus.SetFormatter(&logrus.TextFormatter{})
	conf, err := util.FetchConf()
	if err == nil {
		fileName = conf.LogPath + timeStr + ".log"
		wfFileName = fileName + ".wf"
	}
}

func WriteInfo() error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	log.Out = file
	log.WithFields(Info).Info()
	return nil
}

func Error(msg string) error {
	file, err := os.OpenFile(wfFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	log.Out = file
	log.WithFields(logrus.Fields{
		"logId": "xxx",
	}).Error(msg)
	return nil
}

func Warn(msg string) error {
	file, err := os.OpenFile(wfFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	log.Out = file
	log.WithFields(logrus.Fields{
		"logId": "xxx",
	}).Warn(msg)
	return nil
}
