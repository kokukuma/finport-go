package crawler

import (
	"fmt"
	"testing"
)

type DummySource struct{}

func (d *DummySource) GetIDs() ([]string, error) {
	return []string{
		"1",
		"2",
		"3",
	}, nil
}

func (d *DummySource) Update(id string) error {
	fmt.Println(id)
	return nil
}

func TestServerHandle(t *testing.T) {
	// create crawler
	crawler := NewCrawler()
	crawler.sources = []Source{
		Source(&DummySource{}),
	}

	err := crawler.Start()
	if err != nil {
		t.Fatal(err.Error())
	}
}
