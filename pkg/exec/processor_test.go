package exec

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestProcessor_Process(t *testing.T) {
	logger := logrus.New()
	p := &Processor{
		Logger: logger.WithField("source", "processor"),
		Finder: &Finder{
			Logger: logger.WithField("source", "finder"),
		},
	}
	if err := p.Process(); err != nil {
		t.Fatal(err)
	}
}
