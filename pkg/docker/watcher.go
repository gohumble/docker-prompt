package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"sync"
	"time"
)

// Watcher holds the docker client
type Watcher struct {
	client           *client.Client
	searchResultChan chan []registry.SearchResult
	mux              *sync.Mutex
	searching        bool
}

// creates a Watcher and starts watching docker for changes.
func NewWatcher(dockerClient *client.Client) Watcher {
	w := Watcher{client: dockerClient,
		searchResultChan: make(chan []registry.SearchResult, 1),
		mux:              &sync.Mutex{}}
	return w
}
func (w *Watcher) ClearSearch() {
	if len(w.searchResultChan) > 0 {
		for range <-w.searchResultChan {
		}
	}

}
func (w *Watcher) Search(q string) {
	if !w.searching {
		w.mux.Lock()
		w.searching = true
		w.mux.Unlock()
		//w.ClearSearch()
		res, _ := w.client.ImageSearch(context.Background(), q, types.ImageSearchOptions{Limit: 3})
		w.mux.Lock()
		w.searching = false
		w.searchResultChan <- res
		w.mux.Unlock()
	}
}
func (w Watcher) Start(completer *Completer) {
	go func() {
		for {
			w.watch(completer)
			time.Sleep(time.Second)
		}
	}()
}

// start watching docker containers
func (w Watcher) watch(completer *Completer) {
	w.updateContainers(completer)
	w.updateImages(completer)
}
func (w Watcher) updateContainers(completer *Completer) {
	containers, err := w.client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	completer.containers = containers
}

// start watching docker containers
func (w Watcher) updateImages(completer *Completer) {
	images, err := w.client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	completer.images = images
}
