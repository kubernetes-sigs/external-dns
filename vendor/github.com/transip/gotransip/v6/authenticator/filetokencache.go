package authenticator

import (
	"encoding/json"
	"fmt"
	"github.com/transip/gotransip/v6/jwt"
	"io/ioutil"
	"os"
	"sync"
)

// cacheItem is one named item inside the filesystem cache
type cacheItem struct {
	// Key of the cache item, containing
	Key string `json:"key"`
	// Data containing the content of the cache item
	Data []byte `json:"data"`
}

// FileTokenCache is a cache that takes a path and writes a json marshalled File to it,
// it decodes it when created with the NewFileTokenCache method.
// It has a Set method to save a token by name as jwt.Token
// and a Get method one to get a previously acquired token by name returned as jwt.Token
type FileTokenCache struct {
	// File contains the cache file
	File *os.File
	// CacheItems contains a list of cache items, all of them have a key
	CacheItems []cacheItem `json:"items"`
	// prevent simultaneous cache writes
	writeLock sync.RWMutex
}

// NewFileTokenCache opens or creates a filesystem cache File on the specified path
func NewFileTokenCache(path string) (*FileTokenCache, error) {
	// open the File or create a new one on the given location
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return &FileTokenCache{}, fmt.Errorf("error opening cache File: %w", err)
	}

	cache := FileTokenCache{File: file}

	// try to read the File
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return &FileTokenCache{}, fmt.Errorf("error reading cache File: %w", err)
	}

	if len(fileContent) > 0 {
		// read the cached File data as json
		if err := json.Unmarshal(fileContent, &cache); err != nil {
			return &FileTokenCache{}, fmt.Errorf("error unmarshalling cache File: %w", err)
		}
	}

	return &cache, nil
}

// Set will save a token by name as jwt.Token
func (f *FileTokenCache) Set(key string, token jwt.Token) error {
	for idx, item := range f.CacheItems {
		if item.Key == key {
			f.CacheItems[idx].Data = []byte(token.String())

			// persist this change to the cache File
			return f.writeCacheToFile()
		}
	}

	// if the key did not exist before, we append a new item to the cache item list
	f.CacheItems = append(f.CacheItems, cacheItem{Key: key, Data: []byte(token.String())})

	return f.writeCacheToFile()
}

func (f *FileTokenCache) writeCacheToFile() error {
	// try to convert the cache to json, so we can write it to File
	cacheData, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("error marshalling cache File: %w", err)
	}

	f.writeLock.Lock()
	defer f.writeLock.Unlock()
	// write the cache data to the File cache
	if err := f.File.Truncate(0); err != nil {
		return fmt.Errorf("error while truncating cache File: %w", err)
	}
	_, err = f.File.WriteAt(cacheData, 0)

	return err
}

// Get a previously acquired token by name returned as jwt.Token
func (f *FileTokenCache) Get(key string) (jwt.Token, error) {
	for _, item := range f.CacheItems {
		if item.Key == key {
			dataAsString := string(item.Data)

			return jwt.New(dataAsString)
		}
	}

	return jwt.Token{}, nil
}
