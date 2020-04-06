package hook

import (
	hookopts "github.com/alanshaw/ipfs-hookds/opts"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// Datastore is a wrapper for a datastore that adds optional before and after hooks into it's methods.
type Datastore struct {
	ds      datastore.Datastore
	options hookopts.Options
}

// NewDatastore wraps a datastore.Datastore datastore and adds optional before and after hooks into it's methods.
func NewDatastore(ds datastore.Datastore, options ...hookopts.Option) *Datastore {
	opts := hookopts.Options{}
	opts.Apply(options...)
	return &Datastore{ds: ds, options: opts}
}

// Put stores the object `value` named by `key`, it calls OnBeforePut and OnAfterPut hooks.
func (hds *Datastore) Put(key datastore.Key, value []byte) error {
	if hds.options.OnBeforePut != nil {
		key, value = hds.options.OnBeforePut(key, value)
	}
	err := hds.ds.Put(key, value)
	if hds.options.OnAfterPut != nil {
		err = hds.options.OnAfterPut(key, value, err)
	}
	return err
}

// Delete removes the value for given `key`, it calls OnBeforeDelete and OnAfterDelete hooks.
func (hds *Datastore) Delete(key datastore.Key) error {
	if hds.options.OnBeforeDelete != nil {
		key = hds.options.OnBeforeDelete(key)
	}
	err := hds.ds.Delete(key)
	if hds.options.OnAfterDelete != nil {
		err = hds.options.OnAfterDelete(key, err)
	}
	return err
}

// Get retrieves the object `value` named by `key`, it calls OnBeforeGet and OnAfterGet hooks.
func (hds *Datastore) Get(key datastore.Key) ([]byte, error) {
	if hds.options.OnBeforeGet != nil {
		key = hds.options.OnBeforeGet(key)
	}
	value, err := hds.ds.Get(key)
	if hds.options.OnAfterGet != nil {
		value, err = hds.options.OnAfterGet(key, value, err)
	}
	return value, err
}

// Has returns whether the `key` is mapped to a `value`.
func (hds *Datastore) Has(key datastore.Key) (bool, error) {
	if hds.options.OnBeforeHas != nil {
		key = hds.options.OnBeforeHas(key)
	}
	exists, err := hds.ds.Has(key)
	if hds.options.OnAfterHas != nil {
		exists, err = hds.options.OnAfterHas(key, exists, err)
	}
	return exists, err
}

// GetSize returns the size of the `value` named by `key`.
func (hds *Datastore) GetSize(key datastore.Key) (int, error) {
	return hds.ds.GetSize(key)
}

// Query searches the datastore and returns a query result, it calls OnBeforeQuery and OnAfterQuery hooks.
func (hds *Datastore) Query(q query.Query) (query.Results, error) {
	if hds.options.OnBeforeQuery != nil {
		q = hds.options.OnBeforeQuery(q)
	}
	res, err := hds.ds.Query(q)
	if hds.options.OnAfterQuery != nil {
		res, err = hds.options.OnAfterQuery(q, res, err)
	}
	return res, err
}

// Sync guarantees that any Put or Delete calls under prefix that returned
// before Sync(prefix) was called will be observed after Sync(prefix)
// returns, even if the program crashes.
func (hds *Datastore) Sync(prefix datastore.Key) error {
	return hds.ds.Sync(prefix)
}

// Close closes the underlying datastore
func (hds *Datastore) Close() error {
	return hds.ds.Close()
}
