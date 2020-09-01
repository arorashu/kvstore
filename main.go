// a simple key value store

package main

import (
	"fmt"
    "net/http"
    "log"
    "io/ioutil"
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


//func showAll(w http.ResponseWriter, r *http.Request) {
//  fmt.Fprintf(w, store.Print())
//}


func main() {
	fmt.Println("Hello, kvstore")
    store := KV{}
    store.LoadFromFile("db")
	store.Put("hello", "world")
	store.Put("cake", "walk")

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, store.ToString())
    })

    http.HandleFunc("/close", func (w http.ResponseWriter, r *http.Request) {
        store.WriteToFile("db")
        fmt.Fprintf(w, "closing kv server. Goodbye!")
        fmt.Fprintf(os.Stdout, "closing kv server. Goodbye!")
        os.Exit(0)
    })


//    http.HandleFunc("/put", func (w http.ResponseWriter, r *http.Request) {
//        store.Put()
//        fmt.Fprintf(w, "closing kv server. Goodbye!")
//        fmt.Fprintf(os.Stdout, "closing kv server. Goodbye!")
//        os.Exit(0)
//})

    log.Fatal(http.ListenAndServe(":8080", nil))

//	store.Print()
//	val, _ := store.Get("hello")
//	fmt.Printf("key: hello, val: %s\n", val)
//	val, err := store.Get("hi")
//	if err {
//		fmt.Printf("key: hi, val: %s\n", val)
//	} else {
//		fmt.Printf("key: hi not found\n")
//	}
//	ok := store.Delete("cake")
//	if ok {
//		fmt.Printf("Deleted key: cake\n")
//	} else {
//		fmt.Printf("Cannot delete key: cake\n")
//	}
//	store.Print()
}
