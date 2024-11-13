package fileloader

import (
	"errors"
	"sync"
)

type Loader interface {
	Load(b []byte, v interface{}) (string, error)
}

type LoaderRegistry struct {
	loaders map[string]Loader

	mu sync.RWMutex
}

var ErrLoaderNotFound = errors.New("loader not found for this format")
var ErrLoaderFormatAlreadyRegistered = errors.New("loader already registered for this format")

func NewLoaderRegistry() *LoaderRegistry {
	return &LoaderRegistry{
		loaders: make(map[string]Loader),
	}
}

func (e *LoaderRegistry) RegisterDecoder(ext string, enc Loader) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.loaders[ext]; ok {
		return ErrLoaderFormatAlreadyRegistered
	}
	e.loaders[ext] = enc
	return nil
}

func (e *LoaderRegistry) Load(ext string, b []byte, v interface{}) (string, error) {
	e.mu.RLock()
	loader, ok := e.loaders[ext]
	e.mu.RUnlock()
	if !ok {
		return "", ErrLoaderNotFound
	}
	return loader.Load(b, v)
}
