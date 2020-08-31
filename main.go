// a simple key value store

package main

import (
	"fmt"
)

type KV struct {
	keys   []string
	values []string
}

func (store *KV) Put(nKey string, nValue string) {
	for i, key := range store.keys {
		if key == nKey { // key already exists
			store.values[i] = nValue
			return
		}
	}
	store.keys = append(store.keys, nKey)
	store.values = append(store.values, nValue)
}

func (store *KV) Get(key string) (string, bool) {
	for i, _ := range store.keys {
		if store.keys[i] == key {
			return store.values[i], true
		}
	}
	return "", false
}

func (store *KV) Delete(key string) bool {
	for i, _ := range store.keys {
		if store.keys[i] == key {
			lastIdx := len(store.keys) - 1
			store.keys[i] = store.keys[lastIdx]
			store.values[i] = store.values[lastIdx]
			store.keys = store.keys[:lastIdx]
			store.values = store.values[:lastIdx]
			return true
		}
	}
	return false
}

// print all key, value pairs
func (store *KV) Print() {
	if len(store.keys) == 0 {
		return
	}

	fmt.Printf("All pairs: \n{ ")
	for i, _ := range store.keys {
		//fmt.Printf("[key: %s, value: %s], ", store.keys[i], store.values[i])
		fmt.Printf("[%s: %s], ", store.keys[i], store.values[i])
	}
	fmt.Printf(" }\n")
}

func main() {
	fmt.Println("Hello, kvstore")
	store := KV{}
	store.Put("hello", "world")
	store.Put("cake", "walk")
	store.Print()
	val, _ := store.Get("hello")
	fmt.Printf("key: hello, val: %s\n", val)
	val, err := store.Get("hi")
	if err {
		fmt.Printf("key: hi, val: %s\n", val)
	} else {
		fmt.Printf("key: hi not found\n")
	}
	ok := store.Delete("cake")
	if ok {
		fmt.Printf("Deleted key: cake\n")
	} else {
		fmt.Printf("Cannot delete key: cake\n")
	}
	store.Print()
}
