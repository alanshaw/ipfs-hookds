package opts

import (
	"fmt"

	"github.com/ipfs/go-datastore"
)

// Options are hook datastore options.
type Options struct {
	OnBeforeGet    func(datastore.Key) datastore.Key
	OnAfterGet     func(datastore.Key, []byte, error) ([]byte, error)
	OnBeforePut    func(datastore.Key, []byte) (datastore.Key, []byte)
	OnAfterPut     func(datastore.Key, []byte, error) error
	OnBeforeDelete func(datastore.Key) datastore.Key
	OnAfterDelete  func(datastore.Key, error) error
	OnBeforeBatch  func()
	OnAfterBatch   func(datastore.Batch, error) (datastore.Batch, error)
	OnBeforeHas    func(datastore.Key) datastore.Key
	OnAfterHas     func(datastore.Key, bool, error) (bool, error)
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
func OnBeforeGet(f func(datastore.Key) datastore.Key) Option {
	return func(o *Options) error {
		o.OnBeforeGet = f
		return nil
	}
}

// OnAfterGet configures a hook that is called _after_ Get.
// Defaults to noop.
func OnAfterGet(f func(datastore.Key, []byte, error) ([]byte, error)) Option {
	return func(o *Options) error {
		o.OnAfterGet = f
		return nil
	}
}

// OnBeforePut configures a hook that is called _before_ Put.
// Defaults to noop.
func OnBeforePut(f func(datastore.Key, []byte) (datastore.Key, []byte)) Option {
	return func(o *Options) error {
		o.OnBeforePut = f
		return nil
	}
}

// OnAfterPut configures a hook that is called _after_ Put.
// Defaults to noop.
func OnAfterPut(f func(datastore.Key, []byte, error) error) Option {
	return func(o *Options) error {
		o.OnAfterPut = f
		return nil
	}
}

// OnBeforeDelete configures a hook that is called _before_ Delete.
// Defaults to noop.
func OnBeforeDelete(f func(datastore.Key) datastore.Key) Option {
	return func(o *Options) error {
		o.OnBeforeDelete = f
		return nil
	}
}

// OnAfterDelete configures a hook that is called _after_ Delete.
// Defaults to noop.
func OnAfterDelete(f func(datastore.Key, error) error) Option {
	return func(o *Options) error {
		o.OnAfterDelete = f
		return nil
	}
}

// OnBeforeBatch configures a hook that is called _before_ Batch.
// Defaults to noop.
func OnBeforeBatch(f func()) Option {
	return func(o *Options) error {
		o.OnBeforeBatch = f
		return nil
	}
}

// OnAfterBatch configures a hook that is called _after_ Batch.
// Defaults to noop.
func OnAfterBatch(f func(datastore.Batch, error) (datastore.Batch, error)) Option {
	return func(o *Options) error {
		o.OnAfterBatch = f
		return nil
	}
}

// OnBeforeHas configures a hook that is called _before_ Has.
// Defaults to noop.
func OnBeforeHas(f func(datastore.Key) datastore.Key) Option {
	return func(o *Options) error {
		o.OnBeforeHas = f
		return nil
	}
}

// OnAfterHas configures a hook that is called _after_ Has.
// Defaults to noop.
func OnAfterHas(f func(datastore.Key, bool, error) (bool, error)) Option {
	return func(o *Options) error {
		o.OnAfterHas = f
		return nil
	}
}
