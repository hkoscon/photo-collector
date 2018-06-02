package meta

import (
	"hkoscon.org/photos/pkg/modals"
	"io/ioutil"
)

type Recorder map[string]bool

func ReadRecorder(filename string) (Recorder, error) {
	meta := new(modals.Photos)
	date, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := meta.Unmarshal(date); err != nil {
		return nil, err
	}

	recorder := make(Recorder)

	for _, photo := range meta.Photos {
		recorder[photo.Filename] = photo.Copied
	}

	return recorder, nil
}

func (r Recorder) Copied(filename string) bool {
	copied, exists := r[filename]
	if !exists {
		return true
	}

	return copied
}

func (r Recorder) MarkAsCopied(filename string) {
	r[filename] = true
}

func (r Recorder) Save(filename string) (err error) {
	meta := &modals.Photos{
		Photos: make([]*modals.Photo, 0, len(r)),
	}

	for name, copied := range r {
		meta.Photos = append(meta.Photos, &modals.Photo{
			Filename: name,
			Copied:   copied,
		})
	}

	content, err := meta.Marshal()
	if err != nil {
		return
	}

	return ioutil.WriteFile(filename, content, 0644)
}
