# ipfs-hookds

[![Build Status](https://img.shields.io/travis/com/libp2p/hydra-booster/master?style=flat-square)](https://travis-ci.org/alanshaw/ipfs-hookds)
[![Coverage](https://img.shields.io/codecov/c/github/alanshaw/ipfs-hookds?style=flat-square)](https://codecov.io/gh/alanshaw/ipfs-hookds)
[![Standard README](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/alanshaw/ipfs-hookds)
[![golang version](https://img.shields.io/badge/golang-%3E%3D1.14.0-orange.svg?style=flat-square)](https://golang.org/)

> A wrapper for an [IPFS datastore](https://github.com/ipfs/go-datastore) and adds optional before and after hooks into it's methods.

## Install

```console
go get github.com/alanshaw/ipfs-hookds
```

## Usage

### Example

Hook for after a `Put`:

```go
package main

import (
    "fmt"

	"github.com/ipfs/go-datastore"
    "github.com/alanshaw/ipfs-hookds/opts"
	"github.com/alanshaw/ipfs-hookds"
)

func main() {
    ds := datastore.NewMapDatastore()
    hds := hook.NewDatastore(ds, opts.OnAfterPut(func(k datastore.Key, v []byte, err error) error {
        fmt.Println("key: %v value: %v was put to the datastore")
		return err
    }))
    defer hds.Close()

    key := datastore.NewKey("test")
    value := []byte("test")

    hds.Put(key, value)
}
```

## API

[GoDoc Reference](https://godoc.org/github.com/alanshaw/ipfs-hookds)

## Contribute

Feel free to dive in! [Open an issue](https://github.com/alanshaw/ipfs-hookds/issues/new) or submit PRs.

## License

[MIT](LICENSE) © Alan Shaw
