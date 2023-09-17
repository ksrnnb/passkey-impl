package kvs

import "fmt"

type KVS struct {
	store map[string]string
}

var kvs KVS

func init() {
	kvs = KVS{}
}

func Add(key string, value string) {
	kvs.store[key] = value
}

func Get(key string) (string, error) {
	v, ok := kvs.store[key]
	if !ok {
		return "", fmt.Errorf("%s not found in kvs", key)
	}
	return v, nil
}
