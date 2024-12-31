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
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId) // Вывод: 42

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId) // Вывод: <nil>
}
```
