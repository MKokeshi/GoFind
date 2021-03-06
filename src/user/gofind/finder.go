package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
)

type Finder struct {
	pattern *regexp.Regexp
	mutex sync.Mutex
	work chan string // dal main thread ai workers
	newDirs chan string // dai workers al mainthread
	errors chan error // dai workers al mainthread
	matches []string
	dispatched int // contatore
	numWorkers int 
}

func NewFinder(pattern *regexp.Regexp) *Finder {
	numWorkers := runtime.NumCPU()
	return &Finder{
		pattern: pattern,
		work: make(chan string, numWorkers),
		newDirs: make(chan string, numWorkers),
		errors: make(chan error, numWorkers),
		numWorkers: numWorkers,
	}
}

func (finder *Finder) find(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			finder.newDirs <- filePath
		} else if finder.pattern.MatchString(filePath) {
			finder.mutex.Lock()
			finder.matches = append(finder.matches, filePath)
			finder.mutex.Unlock()
		}
	}

	return nil
}

func (finder *Finder) worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for dir := range finder.work {
		finder.errors <- finder.find(dir)
	}
}

func (finder *Finder) Find(startDir string) ([]string, error) {
	wg := &sync.WaitGroup{}

	defer close(finder.errors)
	defer close(finder.newDirs)
	defer wg.Wait()
	defer close(finder.work)

	fmt.Printf("Utilizzando %d workers !\n"), finder.numWorkers)
	for i := 0; i < finder.numWorkers; i++ {
		wg.Add(1)
		go finder.worker(wg)
	}

	forDispatch := StringQueue{}
	forDispatch.Push(startDir)

	for {
		work := finder.work
		var dir string
		var err error

		if forDispatch.Empty() {
			// Disabilita il secondo case quando la queue è vuota
			work = nil
		} else {
			dir, err = forDispatch.Front()
			if err != nil {
				return nil, err
			}
		}

		select {
		case dir := <-finder.newDirs:
			forDispatch.Push(dir)
		case work <- dir:
			_, err = forDispatch.Pop()
			if err != nil {
				return nil, err
			}
			finder.dispatched++
		case err = <-finder.errors:
			finder.dispatched--
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
			}
		default:
			if finder.dispatched == 0 && forDispatch.Empty() {
				return finder.matches, nil
			}
		}
	}
}