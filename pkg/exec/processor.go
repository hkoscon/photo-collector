package exec

import (
	"github.com/sirupsen/logrus"
	"path"
	"sync"
)

type Processor struct {
	Logger logrus.FieldLogger `inject:"processor logger"`
	Finder *Finder            `inject:""`
	Sender *Sender            `inject:""`
}

func (p *Processor) Process(username string) (err error) {
	cards, err := p.Finder.ScanSDCards()
	if err != nil {
		return
	}

	var wait sync.WaitGroup
	for _, card := range cards {
		wait.Add(1)
		go p.fork(path.Dir(card), &wait, username)
	}
	wait.Wait()

	return
}

func (p *Processor) fork(card string, wait *sync.WaitGroup, username string) {
	defer wait.Done()
	p.Logger.Infof("Scan SD Card %s", card)
	files, recorder, err := p.Finder.ScanSDCard(card)
	if err != nil {
		p.Logger.Error(err)
		return
	}
	p.Sender.Send(files, username, recorder)
	metapath := path.Join(card, metaFile)
	if err := recorder.Save(metapath); err != nil {
		p.Logger.Error(err)
	}
}
