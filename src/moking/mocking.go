package moking 

import (
    "errors"
    
)

var (
    ErrNotFound = errors.New("not found")
)

type HashTable interface {
    Get(key string) ([]byte, error)
    Set(key string, value []byte) error
}