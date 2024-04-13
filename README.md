# Cacher

Cacher is a simple in-memory database replication tool developed in Go. It provides methods for managing key-value data in memory, including `List`, `Get`, `Set`, and `Delete`.

## Features

- **In-Memory Database**: Cacher stores key-value pairs in memory for fast access.
- **Expiration**: Records inserted into the memory have an expiration date. The software automatically removes expired records from memory at the appropriate time.


## Usage

1. Import the Cacher package in your Go code:

```go
import "github.com/cmparrela/go-db-inmemory/cache"
```


2. Create a new Cacher instance:
```go
cacher := cache.NewCacher()
```

3. Use the provided methods to manage data:
```go
// Set a key-value pair with an expiration time of 10 minutes
expiration := time.Duration(rand.Intn(10)) * time.Second
cacher.Set("key", "value", expiration)

// Get the value associated with a key
value := cache.Get("key")

// Delete a key-value pair
cache.Delete("key")

// List all key-value pairs in the cache
allData := cache.List()
```

## Example
```go
package main

import (
	"context"
	"github.com/cmparrela/go-db-inmemory/cache"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	cacher := cache.NewCacher()

	for i := 1; i < 1000000; i++ {
		expiration := time.Duration(rand.Intn(10)) * time.Second

		key := uuid.New().String()
		cacher.Set(key, "teste", expiration)
	}

	<-ctx.Done()
}
```
