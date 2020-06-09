package batch

import (
	"fmt"

	"github.com/ipfs/go-datastore"
)

// BeforePutFunc is a handler for the before Put hook
type BeforePutFunc func(datastore.Key, []byte) (datastore.Key, []byte)

// AfterPutFunc is a handler for the after Put hook
type AfterPutFunc func(datastore.Key, []byte, error) error

// BeforeDeleteFunc is a handler for the before Delete hook
type BeforeDeleteFunc func(datastore.Key) datastore.Key

// AfterDeleteFunc is a handler for the after Delete hook
type AfterDeleteFunc func(datastore.Key, error) error

// BeforeCommitFunc is a handler for the before Commit hook
type BeforeCommitFunc func()

// AfterCommitFunc is a handler for the after Commit hook
type AfterCommitFunc func(error) error

// Options are batch options.
type Options struct {
	BeforePut    BeforePutFunc
	AfterPut     AfterPutFunc
	BeforeDelete BeforeDeleteFunc
	AfterDelete  AfterDeleteFunc
	BeforeCommit BeforeCommitFunc
	AfterCommit  AfterCommitFunc
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

// WithBeforeCommit configures a hook that is called _before_ Commit.
// Defaults to noop.
func WithBeforeCommit(f BeforeCommitFunc) Option {
	return func(o *Options) error {
		o.BeforeCommit = f
		return nil
	}
}

// WithAfterCommit configures a hook that is called _after_ Commit.
// Defaults to noop.
func WithAfterCommit(f AfterCommitFunc) Option {
	return func(o *Options) error {
		o.AfterCommit = f
		return nil
	}
}
