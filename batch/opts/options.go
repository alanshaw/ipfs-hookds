package opts

import (
	"fmt"

	"github.com/ipfs/go-datastore"
)

// Options are batch options.
type Options struct {
	OnBeforePut    func(datastore.Key, []byte) (datastore.Key, []byte)
	OnAfterPut     func(datastore.Key, []byte, error) error
	OnBeforeDelete func(datastore.Key) datastore.Key
	OnAfterDelete  func(datastore.Key, error) error
	OnBeforeCommit func()
	OnAfterCommit  func(error) error
}

// Option is the hook datastore option type.
type Option func(*Options) error

// Apply applies the given options to this Option.
func (o *Options) Apply(opts ...Option) error {
	for i, opt := range opts {
		if err := opt(o); err != nil {
			return fmt.Errorf("batch option %d failed: %s", i, err)
		}
	}
	return nil
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

// OnAfterCommit configures a hook that is called _before_ Commit.
// Defaults to noop.
func OnBeforeCommit(f func()) Option {
	return func(o *Options) error {
		o.OnBeforeCommit = f
		return nil
	}
}

// OnAfterCommit configures a hook that is called _after_ Commit.
// Defaults to noop.
func OnAfterCommit(f func(error) error) Option {
	return func(o *Options) error {
		o.OnAfterCommit = f
		return nil
	}
}
