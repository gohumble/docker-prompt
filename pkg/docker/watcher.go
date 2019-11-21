package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"time"
)

// Watcher holds the docker client
type Watcher struct {
	client           *client.Client
	searchResultChan chan []registry.SearchResult
}

// creates a Watcher and starts watching docker for changes.
func NewWatcher(dockerClient *client.Client) Watcher {
	w := Watcher{client: dockerClient,
		searchResultChan: make(chan []registry.SearchResult, 1)}

	return w
}

func (w *Watcher) ClearSearch() {
	if len(w.searchResultChan) > 0 {
		for range <-w.searchResultChan {
		}
	}

}
func (w *Watcher) Search(q string) {
	w.ClearSearch()
	res, err := w.client.ImageSearch(context.Background(), q, types.ImageSearchOptions{Limit: 5})
	if err != nil {
		fmt.Println(err)
	}
	w.searchResultChan <- res
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
