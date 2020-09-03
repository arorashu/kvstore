# A simple key value store


*Supported operations*
1. Put key, value
2. Get value for key
3. Delete key, value

*How to run*
1. Start http server
```
$ go run main.go &
```

2. Operations
```
# get all key, value pairs
$ curl 'http://localhost:8080/'
# get value for key
$ curl 'http://localhost:8080/get?key=someKey'
# put keys
$ curl 'http://localhost:8080/put?key=newKey&value=newValue'
# delete keys
$ curl 'http://localhost:8080/delete?key=someKey'
```

## Steps

1. commit: 71a5e584d415663ea3f06b4f0c4da1e21db589e3

    We have a basic data structure KV, that supports Put, Get and Delete.
    The Put is an O(1) operation, and Get and Delete are O(n) operations.
    This is not a persistent store. i.e. the store is only available while your program is running.

    Next steps, let's make this a persistent store.

2. commit: 96f5f7f88165186b97091ad32fda1d85defa573b

    The store is now persistent. If closed gracefully using the /close command.
    GET, PUT and DELETE work fine.
    Next steps: handle concurrent access
    
    
