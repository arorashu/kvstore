// a simple key value store

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
		fmt.Printf("[%s: %s], ", store.keys[i], store.values[i])
	}
	fmt.Printf(" }\n")
}


func (store *KV) ToString() string{
	if len(store.keys) == 0 {
		return ""
	}
    res := ""
	for i, _ := range store.keys {
		res += fmt.Sprintf("%s,%s\n", store.keys[i], store.values[i])
	}
    return res
}

func (store *KV) LoadFromFile(db string) {
    data, err := ioutil.ReadFile(db + ".kv")
    if err != nil {
        fmt.Fprintf(os.Stderr, "load: %v\n", err)
        return
    }
    // each line is a key-value pair, separated by \n
    // each line is key,value

    for _, line := range strings.Split(string(data), "\n") {
       if line != "" {
           if words := strings.Split(line, ","); len(words)==2 {
            store.Put(words[0], words[1])
           } else {
                fmt.Fprintf(os.Stderr, "load, unreadable line: %v, length: %d\n", line, len(words))
           }
       } // ignores empty lines
    }
}


func (store *KV) WriteToFile(db string) {
    data := store.ToString()
    err := ioutil.WriteFile(db + ".kv", []byte(data), 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "write to file: %v\n", err)
    }
}


func deleteHandler(w http.ResponseWriter, r *http.Request, store *KV) {
	q := r.URL.Query()
	keys, ok := q["key"]
	if !ok {
		fmt.Fprintf(w, "Invalid PUT request! No key parameter")
	}
	key := keys[0]
	store.Delete(key)
	fmt.Fprintf(w, "DELETE key: %s success", key)
}

func getHandler(w http.ResponseWriter, r *http.Request, store *KV) {
	q := r.URL.Query()
	keys, ok := q["key"]
	if !ok {
		fmt.Fprintf(w, "Invalid PUT request! No key parameter")
	}
	key := keys[0]
	value, ok := store.Get(key)
	if !ok {
		fmt.Fprintf(w, "No value for key: %s", key)
		return
	}
	fmt.Fprintf(w, "key: %s, value: %s", key, value)
}

func putHandler(w http.ResponseWriter, r *http.Request, store *KV) {
	q := r.URL.Query()
	key, ok := q["key"]
	if !ok {
		fmt.Fprintf(w, "Invalid PUT request! No key parameter")
	}
	value, ok := q["value"]
	if !ok {
		fmt.Fprintf(w, "Invalid PUT request! No value parameter")
	}
	store.Put(key[0], value[0])
	fmt.Fprintf(w, "PUT [key: %s, value: %s] success!", key, value)
}

func closeConnection(w http.ResponseWriter, r *http.Request, store *KV) {
	store.WriteToFile("db")
	fmt.Fprintf(w, "closing kv server. Goodbye!")
	fmt.Fprintf(os.Stdout, "closing kv server. Goodbye!")
	os.Exit(0)
}


func main() {
	fmt.Println("Hello, kvstore")
    store := KV{}
    store.LoadFromFile("db")
	store.Put("hello", "world")
	store.Put("cake", "walk")

    // get all
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, store.ToString())
    })

    http.HandleFunc("/put", func (w http.ResponseWriter, r *http.Request) {
    	putHandler(w, r, &store)
    })

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		getHandler(w, r, &store)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		deleteHandler(w, r, &store)
	})

	http.HandleFunc("/close", func (w http.ResponseWriter, r *http.Request) {
		closeConnection(w, r, &store)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
