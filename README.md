# In-memory cache
In-memory cache golang library for storing data

## Example
```
package main

import (
	"fmt"
	"github.com/troxanna/In-memory-cache"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cache := cache.New(ctx)

	_, err := cache.Set("userId", 22, time.Second * 2)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 3)
	userId, err := cache.Get("userId")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}

	_, err = cache.Set("userId", 24, time.Second * 2)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second)
	userId, err = cache.Get("userId")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}
	cancel()
}
```
