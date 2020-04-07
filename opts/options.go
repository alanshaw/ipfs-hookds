package opts

import (
	"fmt"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// OnBeforeGetFunc is a handler for the before Get hook
type OnBeforeGetFunc func(datastore.Key) datastore.Key

// OnAfterGetFunc is a handler for the after Get hook
type OnAfterGetFunc func(datastore.Key, []byte, error) ([]byte, error)

// OnBeforePutFunc is a handler for the before Put hook
type OnBeforePutFunc func(datastore.Key, []byte) (datastore.Key, []byte)

// OnAfterPutFunc is a handler for the after Put hook
type OnAfterPutFunc func(datastore.Key, []byte, error) error

// OnBeforeDeleteFunc is a handler for the before Delete hook
type OnBeforeDeleteFunc func(datastore.Key) datastore.Key

// OnAfterDeleteFunc is a handler for the after Delete hook
type OnAfterDeleteFunc func(datastore.Key, error) error

// OnBeforeBatchFunc is a handler for the before Batch hook
type OnBeforeBatchFunc func()

// OnAfterBatchFunc is a handler for the after Batch hook
type OnAfterBatchFunc func(datastore.Batch, error) (datastore.Batch, error)

// OnBeforeHasFunc is a handler for the before Has hook
type OnBeforeHasFunc func(datastore.Key) datastore.Key

// OnAfterHasFunc is a handler for the after Has hook
type OnAfterHasFunc func(datastore.Key, bool, error) (bool, error)

// OnBeforeQueryFunc is a handler for the before Query hook
type OnBeforeQueryFunc func(query.Query) query.Query

// OnAfterQueryFunc is a handler for the after Query hook
type OnAfterQueryFunc func(query.Query, query.Results, error) (query.Results, error)

// Options are hook datastore options.
type Options struct {
	OnBeforeGet    OnBeforeGetFunc
	OnAfterGet     OnAfterGetFunc
	OnBeforePut    OnBeforePutFunc
	OnAfterPut     OnAfterPutFunc
	OnBeforeDelete OnBeforeDeleteFunc
	OnAfterDelete  OnAfterDeleteFunc
	OnBeforeBatch  OnBeforeBatchFunc
	OnAfterBatch   OnAfterBatchFunc
	OnBeforeHas    OnBeforeHasFunc
	OnAfterHas     OnAfterHasFunc
	OnBeforeQuery  OnBeforeQueryFunc
	OnAfterQuery   OnAfterQueryFunc
}

// Option is the hook datastore option type.
type Option func(*Options) error

// Apply applies the given options to this Option.
func (o *Options) Apply(opts ...Option) error {
	for i, opt := range opts {
		if err := opt(o); err != nil {
			return fmt.Errorf("hook datastore option %d failed: %s", i, err)
		}
	}
	return nil
}

// OnBeforeGet configures a hook that is called _before_ Get.
// Defaults to noop.
func OnBeforeGet(f OnBeforeGetFunc) Option {
	return func(o *Options) error {
		o.OnBeforeGet = f
		return nil
	}
}

// OnAfterGet configures a hook that is called _after_ Get.
// Defaults to noop.
func OnAfterGet(f OnAfterGetFunc) Option {
	return func(o *Options) error {
		o.OnAfterGet = f
		return nil
	}
}

// OnBeforePut configures a hook that is called _before_ Put.
// Defaults to noop.
func OnBeforePut(f OnBeforePutFunc) Option {
	return func(o *Options) error {
		o.OnBeforePut = f
		return nil
	}
}

// OnAfterPut configures a hook that is called _after_ Put.
// Defaults to noop.
func OnAfterPut(f OnAfterPutFunc) Option {
	return func(o *Options) error {
		o.OnAfterPut = f
		return nil
	}
}

// OnBeforeDelete configures a hook that is called _before_ Delete.
// Defaults to noop.
func OnBeforeDelete(f OnBeforeDeleteFunc) Option {
	return func(o *Options) error {
		o.OnBeforeDelete = f
		return nil
	}
}

// OnAfterDelete configures a hook that is called _after_ Delete.
// Defaults to noop.
func OnAfterDelete(f OnAfterDeleteFunc) Option {
	return func(o *Options) error {
		o.OnAfterDelete = f
		return nil
	}
}

// OnBeforeBatch configures a hook that is called _before_ Batch.
// Defaults to noop.
func OnBeforeBatch(f OnBeforeBatchFunc) Option {
	return func(o *Options) error {
		o.OnBeforeBatch = f
		return nil
	}
}

// OnAfterBatch configures a hook that is called _after_ Batch.
// Defaults to noop.
func OnAfterBatch(f OnAfterBatchFunc) Option {
	return func(o *Options) error {
		o.OnAfterBatch = f
		return nil
	}
}

// OnBeforeHas configures a hook that is called _before_ Has.
// Defaults to noop.
func OnBeforeHas(f OnBeforeHasFunc) Option {
	return func(o *Options) error {
		o.OnBeforeHas = f
		return nil
	}
}

// OnAfterHas configures a hook that is called _after_ Has.
// Defaults to noop.
func OnAfterHas(f OnAfterHasFunc) Option {
	return func(o *Options) error {
		o.OnAfterHas = f
		return nil
	}
}

// OnBeforeQuery configures a hook that is called _before_ Query.
// Defaults to noop.
func OnBeforeQuery(f OnBeforeQueryFunc) Option {
	return func(o *Options) error {
		o.OnBeforeQuery = f
		return nil
	}
}

// OnAfterQuery configures a hook that is called _after_ Query.
// Defaults to noop.
func OnAfterQuery(f OnAfterQueryFunc) Option {
	return func(o *Options) error {
		o.OnAfterQuery = f
		return nil
	}
}
