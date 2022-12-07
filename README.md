# gradecryptDB
A simple key/value store written in Go.    
Note; We will not activly maintain this project, unless we need to. But feel free to make a pull request, we'll look over it

Documentation:
Every key you add is stored locally, and can be accessed and modified from the user itself without using gradecryptDB.

```
addKey(key, value, dir) //Add a key and it's Value
readKey(key, dir) //Get the keys value, returns String
hasKey(key, dir) //check if a key exists, returns Boolean
removeKey(key, dir) //Delete a key from the DB
addKeyUnsafe(key, dir //Add a key and it's Value without encryption
readKey(key, dir) //Get the value from an unsafe Key, returns String 
//Dir is the directory the file is stored in (e.g. addKey("joemama", "pswd132", "Users/UnfunnyUserNames")
```

Assign an AES Key (Random Bytes) to the "crypto" variable. Only the unsafe functions will work otherways
