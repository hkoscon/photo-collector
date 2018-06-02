package exec

import (
	"context"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"hkoscon.org/photos/pkg/meta"
	"path"
	"sync"
	"time"
)

type sendObject struct {
	fullPath string
	rootPath string
}

type Sender struct {
	Client *minio.Client
	Bucket string
	Logger logrus.FieldLogger
}

func (s *Sender) Send(files []sendObject, bucket string, recorder meta.Recorder) {
	var wait sync.WaitGroup
	for _, file := range files {
		wait.Add(1)
		go s.fork(&wait, bucket, file, recorder)
	}
	wait.Wait()
	return
}

func (s *Sender) fork(group *sync.WaitGroup, bucket string, file sendObject, recorder meta.Recorder) {
	defer group.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	_, err := s.Client.FPutObjectWithContext(ctx, bucket, path.Base(file.fullPath), file.fullPath, minio.PutObjectOptions{})
	if err != nil {
		s.Logger.Error(err)
		return
	}
	recorder.MarkAsCopied(file.fullPath[len(file.rootPath):])
}
