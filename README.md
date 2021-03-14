# go-tebex
A golang library for accessing tebex api easily 

### but why go?
because dragonfly uses golang and this would make it easier for dragonfly server owners to access the api

### TODO: Alot todo at the moment

### Examples
#### Server Information
```go
package main

import "github.com/moonlightpe/go-tebex"

func main() {
	session, err := go-tebex.New("API KEY")
	if err !- nil {
		panic(err)
	}
	info, err := session.GetInformation()
	if err != nil {
		panic(err)
	}

	fmt.Println(info)
}
```