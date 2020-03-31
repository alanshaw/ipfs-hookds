package batch

import (
	batchopts "github.com/alanshaw/ipfs-hookds/batch/opts"
	"github.com/ipfs/go-datastore"
)

// Batch is a group of Puts and Deletes to be performed (Committed) in one shot.
type Batch struct {
	bch     datastore.Batch
	options batchopts.Options
}

// NewBatch wraps a datastore.Batch and adds optional before and after hooks into it's methods.
func NewBatch(bch datastore.Batch, options ...batchopts.Option) *Batch {
	opts := batchopts.Options{}
	opts.Apply(options...)
	return &Batch{bch: bch, options: opts}
}

// Put stores the object `value` named by `key`, it calls OnBeforePut and OnAfterPut hooks.
func (hbh *Batch) Put(key datastore.Key, value []byte) error {
	if hbh.options.OnBeforePut != nil {
		key, value = hbh.options.OnBeforePut(key, value)
	}
	err := hbh.bch.Put(key, value)
	if hbh.options.OnAfterPut != nil {
		err = hbh.options.OnAfterPut(key, value, err)
	}
	return err
}

// Delete removes the value for given `key`, it calls OnBeforeDelete and OnAfterDelete hooks.
func (hbh *Batch) Delete(key datastore.Key) error {
	if hbh.options.OnBeforeDelete != nil {
		key = hbh.options.OnBeforeDelete(key)
	}
	err := hbh.bch.Delete(key)
	if hbh.options.OnAfterDelete != nil {
		err = hbh.options.OnAfterDelete(key, err)
	}
	return err
}

// Commit submits the batch to the datastore for processing, it calls OnBeforeCommit and OnAfterCommit hooks.
func (hbh *Batch) Commit() error {
	if hbh.options.OnBeforeCommit != nil {
		hbh.options.OnBeforeCommit()
	}
	err := hbh.bch.Commit()
	if hbh.options.OnAfterCommit != nil {
		err = hbh.options.OnAfterCommit(err)
	}
	return err
}
