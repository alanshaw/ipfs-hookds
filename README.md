# ipfs-hookds

[![Build Status](https://travis-ci.org/alanshaw/ipfs-hookds.svg?branch=master)](https://travis-ci.org/alanshaw/ipfs-hookds)
[![Coverage](https://codecov.io/gh/alanshaw/ipfs-hookds/branch/master/graph/badge.svg)](https://codecov.io/gh/alanshaw/ipfs-hookds)
[![Standard README](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/alanshaw/ipfs-hookds)
[![golang version](https://img.shields.io/badge/golang-%3E%3D1.14.0-orange.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanshaw/ipfs-hookds)](https://goreportcard.com/report/github.com/alanshaw/ipfs-hookds)

> A wrapper for an [IPFS datastore](https://github.com/ipfs/go-datastore) that adds optional before and after hooks to it's methods.

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
	"github.com/alanshaw/ipfs-hookds"
)

func main() {
	ds := datastore.NewMapDatastore()
	hds := hook.NewDatastore(ds, hook.WithAfterPut(func(k datastore.Key, v []byte, err error) error {
		fmt.Printf("key: %v value: %s was put to the datastore\n", k, v)
		return err
	}))
	defer hds.Close()

	key := datastore.NewKey("test")
	value := []byte("test")

	hds.Put(key, value)

	// Output:
	// key: /test value: test was put to the datastore
}
```

Hook into a batch `Put`:

```go
package main

import (
	"fmt"

	"github.com/ipfs/go-datastore"
	"github.com/alanshaw/ipfs-hookds/batch"
	"github.com/alanshaw/ipfs-hookds"
)

func main() {
	ds := datastore.NewMapDatastore()
	hds := hook.NewDatastore(ds, hook.WithAfterBatch(func(b datastore.Batch, err error) (datastore.Batch, error) {
		return batch.NewBatch(b, batch.WithAfterPut(func(datastore.Key, []byte, error) error {
			fmt.Printf("key: %v value: %s was put to a batch\n", k, v)
			return err
		})), err
	}))
	defer hds.Close()

	key := datastore.NewKey("test")
	value := []byte("test")

	bch := hds.Batch()

	bch.Put(key, value)
	bch.Commit()

	// Output:
	// key: /test value: test was put to a batch
}
```

Hook into a query `NextSync`:

```go
package main

import (
	"fmt"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/alanshaw/ipfs-hookds/query/results"
	"github.com/alanshaw/ipfs-hookds"
)

func main() {
	ds := datastore.NewMapDatastore()
	hds := hook.NewDatastore(ds, hook.WithAfterQuery(func(q query.Query, res query.Results, err error) (query.Results, error) {
		return results.NewResults(res, results.WithAfterNextSync(func(r query.Result, ok bool) (query.Result, bool) {
			fmt.Printf("result: %v ok: %s was next\n", r, ok)
			return r, ok
		})), err
	}))
	defer hds.Close()

	key := datastore.NewKey("test")
	value := []byte("test")
	hds.Put(key, value)

	res := hds.Query(query.Query{
		Prefix: "/test"
	})

	res.NextSync()
}
```

## API

[GoDoc Reference](https://godoc.org/github.com/alanshaw/ipfs-hookds)

## Contribute

Feel free to dive in! [Open an issue](https://github.com/alanshaw/ipfs-hookds/issues/new) or submit PRs.

## License

[MIT](LICENSE) © Alan Shaw

