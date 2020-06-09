package hook

import (
	"fmt"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// BeforeGetFunc is a handler for the before Get hook
type BeforeGetFunc func(datastore.Key) datastore.Key

// AfterGetFunc is a handler for the after Get hook
type AfterGetFunc func(datastore.Key, []byte, error) ([]byte, error)

// BeforePutFunc is a handler for the before Put hook
type BeforePutFunc func(datastore.Key, []byte) (datastore.Key, []byte)

// AfterPutFunc is a handler for the after Put hook
type AfterPutFunc func(datastore.Key, []byte, error) error

// BeforeDeleteFunc is a handler for the before Delete hook
type BeforeDeleteFunc func(datastore.Key) datastore.Key

// AfterDeleteFunc is a handler for the after Delete hook
type AfterDeleteFunc func(datastore.Key, error) error

// BeforeBatchFunc is a handler for the before Batch hook
type BeforeBatchFunc func()

// AfterBatchFunc is a handler for the after Batch hook
type AfterBatchFunc func(datastore.Batch, error) (datastore.Batch, error)

// BeforeHasFunc is a handler for the before Has hook
type BeforeHasFunc func(datastore.Key) datastore.Key

// AfterHasFunc is a handler for the after Has hook
type AfterHasFunc func(datastore.Key, bool, error) (bool, error)

// BeforeQueryFunc is a handler for the before Query hook
type BeforeQueryFunc func(query.Query) query.Query

// AfterQueryFunc is a handler for the after Query hook
type AfterQueryFunc func(query.Query, query.Results, error) (query.Results, error)

// Options are hook datastore options.
type Options struct {
	BeforeGet    BeforeGetFunc
	AfterGet     AfterGetFunc
	BeforePut    BeforePutFunc
	AfterPut     AfterPutFunc
	BeforeDelete BeforeDeleteFunc
	AfterDelete  AfterDeleteFunc
	BeforeBatch  BeforeBatchFunc
	AfterBatch   AfterBatchFunc
	BeforeHas    BeforeHasFunc
	AfterHas     AfterHasFunc
	BeforeQuery  BeforeQueryFunc
	AfterQuery   AfterQueryFunc
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

// WithBeforeGet configures a hook that is called _before_ Get.
// Defaults to noop.
func WithBeforeGet(f BeforeGetFunc) Option {
	return func(o *Options) error {
		o.BeforeGet = f
		return nil
	}
}

// WithAfterGet configures a hook that is called _after_ Get.
// Defaults to noop.
func WithAfterGet(f AfterGetFunc) Option {
	return func(o *Options) error {
		o.AfterGet = f
		return nil
	}
}

// WithBeforePut configures a hook that is called _before_ Put.
// Defaults to noop.
func WithBeforePut(f BeforePutFunc) Option {
	return func(o *Options) error {
		o.BeforePut = f
		return nil
	}
}

// WithAfterPut configures a hook that is called _after_ Put.
// Defaults to noop.
func WithAfterPut(f AfterPutFunc) Option {
	return func(o *Options) error {
		o.AfterPut = f
		return nil
	}
}

// WithBeforeDelete configures a hook that is called _before_ Delete.
// Defaults to noop.
func WithBeforeDelete(f BeforeDeleteFunc) Option {
	return func(o *Options) error {
		o.BeforeDelete = f
		return nil
	}
}

// WithAfterDelete configures a hook that is called _after_ Delete.
// Defaults to noop.
func WithAfterDelete(f AfterDeleteFunc) Option {
	return func(o *Options) error {
		o.AfterDelete = f
		return nil
	}
}

// WithBeforeBatch configures a hook that is called _before_ Batch.
// Defaults to noop.
func WithBeforeBatch(f BeforeBatchFunc) Option {
	return func(o *Options) error {
		o.BeforeBatch = f
		return nil
	}
}

// WithAfterBatch configures a hook that is called _after_ Batch.
// Defaults to noop.
func WithAfterBatch(f AfterBatchFunc) Option {
	return func(o *Options) error {
		o.AfterBatch = f
		return nil
	}
}

// WithBeforeHas configures a hook that is called _before_ Has.
// Defaults to noop.
func WithBeforeHas(f BeforeHasFunc) Option {
	return func(o *Options) error {
		o.BeforeHas = f
		return nil
	}
}

// WithAfterHas configures a hook that is called _after_ Has.
// Defaults to noop.
func WithAfterHas(f AfterHasFunc) Option {
	return func(o *Options) error {
		o.AfterHas = f
		return nil
	}
}

// WithBeforeQuery configures a hook that is called _before_ Query.
// Defaults to noop.
func WithBeforeQuery(f BeforeQueryFunc) Option {
	return func(o *Options) error {
		o.BeforeQuery = f
		return nil
	}
}

// WithAfterQuery configures a hook that is called _after_ Query.
// Defaults to noop.
func WithAfterQuery(f AfterQueryFunc) Option {
	return func(o *Options) error {
		o.AfterQuery = f
		return nil
	}
}
