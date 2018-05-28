package modals

import (
	"bytes"
	"io/ioutil"
	"testing"
)

const testPhotographerName = "The July Jasmine Chan"

func TestPhotographer_Marshal(t *testing.T) {
	photgrapher := &Photographer{
		Name: testPhotographerName,
	}

	data, err := photgrapher.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	content, _ := ioutil.ReadFile("test.code")
	if !bytes.Equal(content, data) {
		t.Fatal("incorrect code")
	}
}

func TestPhotographer_Unmarshal(t *testing.T) {
	content, _ := ioutil.ReadFile("test.code")
	photographer := new(Photographer)
	if err := photographer.Unmarshal(content); err != nil {
		t.Fatal(err)
	}
	if photographer.Name != testPhotographerName {
		t.Fatal("Wrong data")
	}
}
