// a simple key value store
/*
System Characteristics
 1. Supports put key, value
 2. get value for key

Eventually

 3. supports concurrent access
 4. distributed / sharded ?

*/

package main

import (
    "fmt"
    //"net/http"
    //"log"
  )

 type KV struct{
     keys []string
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
    for i,_ := range store.keys {
        if store.keys[i] == key {
            return store.values[i], true 
        }
    }
    return "", false
 }

//func handler(w http.ResponseWriter, r *http.Request) {
//  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}


func main() {
    fmt.Println("Hello, kvstore")
//    http.HandleFunc("/", handler)
//    log.Fatal(http.ListenAndServe(":8080", nil))
    store := KV{}
    store.Put("hello", "world")
    store.Put("cake", "walk")
    val, _ := store.Get("hello")
    fmt.Printf("key: hello, val: %s\n", val)
    val, err := store.Get("hi")
    if err {
        fmt.Printf("key: hi, val: %s\n", val)
    } else {
        fmt.Printf("key: hi not found\n")
    }
}

