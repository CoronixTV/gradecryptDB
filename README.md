# gradecryptDB
A simple key/value store written in Go
Note; We will not activly maintain this project, unless we need to. But feel free to make a merge/push request, we'll look over

Documentation:
Every key you add is stored locally, and can be accessed and modified from the user itself without using gradecryptDB.

```
addKey(key, value) //Add a key and it's Value
readKey(key) //Get the keys value, returns String
hasKey(key) //check if a key exists, returns Boolean
removeKey(key) //Delete a key from the DB
```


