# Golang Uevent bindings

This repository hosts golang uevent bindings for receiving Linux kernel events.

## Usage

```
package main

import (
	"fmt"
	"log"

	"github.com/s-urbaniak/uevent"
)

func main() {
	r, err := uevent.NewReader()
	if err != nil {
		log.Fatal(err)
	}

	dec := uevent.NewDecoder(r)

	for {
		evt, err := dec.Decode()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(evt)
	}
}
```

## Prerequisites

- Linux