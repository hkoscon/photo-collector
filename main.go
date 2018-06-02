package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"hkoscon.org/photos/pkg/crypto"
	"hkoscon.org/photos/pkg/exec"
	"hkoscon.org/photos/pkg/user"
)

func execLoop(validator *user.Validator, logger logrus.FieldLogger) {
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

		name, err := validator.GetName(content)
		if err != nil {
			logger.Error(err)
			continue
		}

		if err := processor.Process(name); err != nil {
			logger.Error(err)
			continue
		}

		logger.Infoln("You may remove the card")
	}
}

func main() {
	logger := logrus.New()
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}

	cryptor := &crypto.RSACrypt{
		PriKeyPath: os.Getenv("PRIVATE_KEY_PATH"),
		KeyLabel:   []byte(os.Getenv("KEY_LABEL")),
		Logger:     logger.WithField("source", "crypto"),
	}

	cryptor.LoadKey(password)

	validator := &user.Validator{
		Decryptor: cryptor,
	}

	execLoop(validator)
}
