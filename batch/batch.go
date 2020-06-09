package batch

import (
	"github.com/ipfs/go-datastore"
)

// Batch is a group of Puts and Deletes to be performed (Committed) in one shot.
type Batch struct {
	bch     datastore.Batch
	options Options
}

// NewBatch wraps a datastore.Batch and adds optional before and after hooks into it's methods.
func NewBatch(bch datastore.Batch, options ...Option) *Batch {
	opts := Options{}
	opts.Apply(options...)
	return &Batch{bch: bch, options: opts}
}

// Put stores the object `value` named by `key`, it calls OnBeforePut and OnAfterPut hooks.
func (hbh *Batch) Put(key datastore.Key, value []byte) error {
	if hbh.options.BeforePut != nil {
		key, value = hbh.options.BeforePut(key, value)
	}
	err := hbh.bch.Put(key, value)
	if hbh.options.AfterPut != nil {
		err = hbh.options.AfterPut(key, value, err)
	}
	return err
}

// Delete removes the value for given `key`, it calls OnBeforeDelete and OnAfterDelete hooks.
func (hbh *Batch) Delete(key datastore.Key) error {
	if hbh.options.BeforeDelete != nil {
		key = hbh.options.BeforeDelete(key)
	}
	err := hbh.bch.Delete(key)
	if hbh.options.AfterDelete != nil {
		err = hbh.options.AfterDelete(key, err)
	}
	return err
}

// Commit submits the batch to the datastore for processing, it calls OnBeforeCommit and OnAfterCommit hooks.
func (hbh *Batch) Commit() error {
	if hbh.options.BeforeCommit != nil {
		hbh.options.BeforeCommit()
	}
	err := hbh.bch.Commit()
	if hbh.options.AfterCommit != nil {
		err = hbh.options.AfterCommit(err)
	}
	return err
}
