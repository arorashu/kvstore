# A simple key value store


*Supported operations*
1. Put key, value
2. Get value for key
3. Delete key, value

*How to run*
```
go run main.go
```

## Steps

1. commit: 71a5e584d415663ea3f06b4f0c4da1e21db589e3

    We have a basic data structure KV, that supports Put, Get and Delete.
    The Put is an O(1) operation, and Get and Delete are O(n) operations.
    This is not a persistent store. i.e. the store is only available while your program is running.

    Next steps, let's make this a persistent store.



