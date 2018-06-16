package exec

import (
	"github.com/sirupsen/logrus"
	"hkoscon.org/photos/pkg/meta"
	"hkoscon.org/photos/pkg/modals"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const metaFile = "photos.meta"

type Finder struct {
	Logger logrus.FieldLogger `inject:"finder logger"`
}

func (f *Finder) ScanSDCards() (_ []string, err error) {
	f.Logger.Info("Start scanning SD card")
	sdCards, err := f.globSDCards()
	if err != nil {
		f.Logger.Error(err)
		return
	}
	f.Logger.Infof("Find %d SD Cards(s)", len(sdCards))

	return sdCards, nil
}

func (f *Finder) ScanSDCard(sdCard string) (_ []sendObject, _ meta.Recorder, err error) {
	metapath := path.Join(sdCard, metaFile)
	if _, err := os.Stat(metapath); os.IsNotExist(err) {
		f.Logger.Infof("Meta file not found in %s", sdCard)
		if err := f.touchMetaFile(sdCard); err != nil {
			f.Logger.Error(err)
			return nil, nil, err
		}
	}

	recorder, err := meta.ReadRecorder(metapath)
	if err != nil {
		return
	}

	imageRoot := path.Join(sdCard, "DCIM")

	files, err := f.globImageFiles(imageRoot)
	if err != nil {
		f.Logger.Error(err)
		return
	}

	f.Logger.Infof("Find %d image file(s)", len(files))

	result := make([]sendObject, 0, len(files))

	idx := len(sdCard)

	for _, file := range files {
		relativePath := file[idx:]
		if recorder.Copied(relativePath) {
			continue
		}

		stat, err := os.Stat(file)
		if err != nil {
			return nil, nil, err
		}

		if stat.Mode().Perm()&0222 != 0 {
			continue
		}

		result = append(result, sendObject{
			fullPath: file,
			rootPath: sdCard,
		})
	}

	return result, recorder, nil
}

func (f *Finder) touchMetaFile(root string) (err error) {
	filename := path.Join(root, metaFile)
	m := &modals.Photos{
		Photos: make([]*modals.Photo, 0),
	}
	content, err := m.Marshal()
	if err != nil {
		return
	}
	f.Logger.Infof("Create Meta file %s", filename)
	return ioutil.WriteFile(filename, content, 0644)
}

func (f *Finder) computeGlobPattern(root, filename string) string {
	return path.Join(root, "**", filename)
}

func (f *Finder) globImageFiles(root string) ([]string, error) {
	return filepath.Glob(path.Join(root, "**", "*.JPG"))
}

func (f *Finder) globSDCards() ([]string, error) {
	return filepath.Glob(path.Join(getMediaRoot(), "**", "DCIM"))
}
