package hook

import (
	"github.com/alanshaw/ipfs-hookds/opts"
	hookopts "github.com/alanshaw/ipfs-hookds/opts"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// Batching is a datastore with hooks that also supports batching
type Batching struct {
	ds      datastore.Batching
	hds     *Datastore
	options opts.Options
}

// NewBatching wraps a datastore.Batching datastore and adds optional before and after hooks into it's methods
func NewBatching(ds datastore.Batching, options ...hookopts.Option) *Batching {
	return &Batching{ds: ds, hds: NewDatastore(ds, options...)}
}

// Put stores the object `value` named by `key`, it calls OnBeforePut and OnAfterPut hooks.
func (bds *Batching) Put(key datastore.Key, value []byte) error {
	return bds.hds.Put(key, value)
}

// Delete removes the value for given `key`, it calls OnBeforeDelete and OnAfterDelete hooks.
func (bds *Batching) Delete(key datastore.Key) error {
	return bds.hds.Delete(key)
}

// Get retrieves the object `value` named by `key`, it calls OnBeforeGet and OnAfterGet hooks.
func (bds *Batching) Get(key datastore.Key) ([]byte, error) {
	return bds.hds.Get(key)
}

// Has returns whether the `key` is mapped to a `value`.
func (bds *Batching) Has(key datastore.Key) (bool, error) {
	return bds.hds.Has(key)
}

// GetSize returns the size of the `value` named by `key`.
func (bds *Batching) GetSize(key datastore.Key) (int, error) {
	return bds.hds.GetSize(key)
}

// Query searches the datastore and returns a query result.
func (bds *Batching) Query(q query.Query) (query.Results, error) {
	return bds.hds.Query(q)
}

// Batch creates a container for a group of updates, it calls OnBeforeBatch and OnAfterBatch hooks.
func (bds *Batching) Batch() (datastore.Batch, error) {
	if bds.hds.options.OnBeforeBatch != nil {
		bds.hds.options.OnBeforeBatch()
	}
	bch, err := bds.ds.Batch()
	if bds.hds.options.OnAfterBatch != nil {
		bch, err = bds.hds.options.OnAfterBatch(bch, err)
	}
	return bch, err
}

// Sync guarantees that any Put or Delete calls under prefix that returned
// before Sync(prefix) was called will be observed after Sync(prefix)
// returns, even if the program crashes.
func (bds *Batching) Sync(prefix datastore.Key) error {
	return bds.hds.Sync(prefix)
}

// Close closes the underlying datastore
func (bds *Batching) Close() error {
	return bds.hds.Close()
}
