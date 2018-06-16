package main

import (
	"bufio"
	"os"

	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"hkoscon.org/photos/pkg/exec"
)

func execLoop(logger logrus.FieldLogger) {
	scanner := bufio.NewScanner(os.Stdin)

	client, err := minio.New(os.Getenv("S3_ENDPOINT"), os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_SECRET_ACCESS_KEY"), false)
	if err != nil {
		panic(err)
	}
	processor := &exec.Processor{
		Logger: logger.WithField("source", "processor"),
		Finder: &exec.Finder{
			Logger: logger.WithField("source", "finder"),
		},
		Sender: &exec.Sender{
			Client: client,
			Logger: logger.WithField("source", "sender"),
		},
	}
	for scanner.Scan() {
		content := scanner.Text()
		if err := scanner.Err(); err != nil {
			logger.Error(err)
			continue
		}

		if err := processor.Process(content); err != nil {
			logger.Error(err)
			continue
		}

		logger.Infoln("You may remove the card")
	}
}

func main() {
	logger := logrus.New()
	logger.Infoln("Start service")
	execLoop(logger.WithField("source", "loop"))
}
