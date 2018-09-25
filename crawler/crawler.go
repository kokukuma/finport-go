package crawler

import (
	"fmt"
	"sync"
)

// Source interface is implimented by ...
type Source interface {
	GetIDs() ([]string, error)
	Update(id string) error
}

// Crawler executed by main
type Crawler struct {
	sources []Source
}

// Start get some data from source.
func (c *Crawler) Start() error {
	for _, source := range c.sources {
		ids, err := source.GetIDs()
		if err != nil {
			// TODO: Notify by slack
			fmt.Println(err)
		}

		// goroutine is restrict to 5
		wg := sync.WaitGroup{}
		semaphore := make(chan int, 5)

		for _, id := range ids {
			wg.Add(1)
			go func(id string, wg *sync.WaitGroup, semaphore chan int) {
				defer wg.Done()
				semaphore <- 1
				err := source.Update(id)
				if err != nil {
					// TODO: Notify by slack
					fmt.Println(err)
				}
				<-semaphore
			}(id, &wg, semaphore)
		}
		wg.Wait()
	}
	return nil
}

// NewCrawler creates a new Crawler.
func NewCrawler() *Crawler {
	return &Crawler{
		sources: []Source{},
	}
}
