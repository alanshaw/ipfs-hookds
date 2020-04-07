package opts

import (
	"fmt"

	"github.com/ipfs/go-datastore"
)

// OnBeforePutFunc is a handler for the before Put hook
type OnBeforePutFunc func(datastore.Key, []byte) (datastore.Key, []byte)

// OnAfterPutFunc is a handler for the after Put hook
type OnAfterPutFunc func(datastore.Key, []byte, error) error

// OnBeforeDeleteFunc is a handler for the before Delete hook
type OnBeforeDeleteFunc func(datastore.Key) datastore.Key

// OnAfterDeleteFunc is a handler for the after Delete hook
type OnAfterDeleteFunc func(datastore.Key, error) error

// OnBeforeCommitFunc is a handler for the before Commit hook
type OnBeforeCommitFunc func()

// OnAfterCommitFunc is a handler for the after Commit hook
type OnAfterCommitFunc func(error) error

// Options are batch options.
type Options struct {
	OnBeforePut    OnBeforePutFunc
	OnAfterPut     OnAfterPutFunc
	OnBeforeDelete OnBeforeDeleteFunc
	OnAfterDelete  OnAfterDeleteFunc
	OnBeforeCommit OnBeforeCommitFunc
	OnAfterCommit  OnAfterCommitFunc
}

// Option is the batch option type.
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

// OnBeforeCommit configures a hook that is called _before_ Commit.
// Defaults to noop.
func OnBeforeCommit(f OnBeforeCommitFunc) Option {
	return func(o *Options) error {
		o.OnBeforeCommit = f
		return nil
	}
}

// OnAfterCommit configures a hook that is called _after_ Commit.
// Defaults to noop.
func OnAfterCommit(f OnAfterCommitFunc) Option {
	return func(o *Options) error {
		o.OnAfterCommit = f
		return nil
	}
}
