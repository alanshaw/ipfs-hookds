package results

import (
	"fmt"

	"github.com/ipfs/go-datastore/query"
)

// BeforeNextFunc is a handler for the before Next hook
type BeforeNextFunc func()

// AfterNextFunc is a handler for the after Next hook
type AfterNextFunc func(<-chan query.Result) <-chan query.Result

// BeforeNextSyncFunc is a handler for the before NextSync hook
type BeforeNextSyncFunc func()

// AfterNextSyncFunc is a handler for the after NextSync hook
type AfterNextSyncFunc func(query.Result, bool) (query.Result, bool)

// BeforeRestFunc is a handler for the before Rest hook
type BeforeRestFunc func() ([]query.Entry, error)

// AfterRestFunc is a handler for the after Rest hook
type AfterRestFunc func([]query.Entry, error) ([]query.Entry, error)

// BeforeCloseFunc is a handler for the before Close hook
type BeforeCloseFunc func()

// AfterCloseFunc is a handler for the after Close hook
type AfterCloseFunc func(error) error

// Options are results options.
type Options struct {
	BeforeNext     BeforeNextFunc
	AfterNext      AfterNextFunc
	BeforeNextSync BeforeNextSyncFunc
	AfterNextSync  AfterNextSyncFunc
	BeforeRest     BeforeRestFunc
	AfterRest      AfterRestFunc
	BeforeClose    BeforeCloseFunc
	AfterClose     AfterCloseFunc
}

// Option is the results option type.
type Option func(*Options) error

// Apply applies the given options to this Option.
func (o *Options) Apply(opts ...Option) error {
	for i, opt := range opts {
		if err := opt(o); err != nil {
			return fmt.Errorf("results option %d failed: %s", i, err)
		}
	}
	return nil
}

// WithBeforeNext configures a hook that is called _before_ Next.
// Defaults to noop.
func WithBeforeNext(f BeforeNextFunc) Option {
	return func(o *Options) error {
		o.BeforeNext = f
		return nil
	}
}

// WithAfterNext configures a hook that is called _after_ Next.
// Defaults to noop.
func WithAfterNext(f AfterNextFunc) Option {
	return func(o *Options) error {
		o.AfterNext = f
		return nil
	}
}

// WithBeforeNextSync configures a hook that is called _before_ NextSync.
// Defaults to noop.
func WithBeforeNextSync(f BeforeNextSyncFunc) Option {
	return func(o *Options) error {
		o.BeforeNextSync = f
		return nil
	}
}

// WithAfterNextSync configures a hook that is called _after_ NextSync.
// Defaults to noop.
func WithAfterNextSync(f AfterNextSyncFunc) Option {
	return func(o *Options) error {
		o.AfterNextSync = f
		return nil
	}
}

// WithBeforeRest configures a hook that is called _before_ Rest.
// Defaults to noop.
func WithBeforeRest(f BeforeRestFunc) Option {
	return func(o *Options) error {
		o.BeforeRest = f
		return nil
	}
}

// WithAfterRest configures a hook that is called _after_ Rest.
// Defaults to noop.
func WithAfterRest(f AfterRestFunc) Option {
	return func(o *Options) error {
		o.AfterRest = f
		return nil
	}
}

// WithBeforeClose configures a hook that is called _before_ Close.
// Defaults to noop.
func WithBeforeClose(f BeforeCloseFunc) Option {
	return func(o *Options) error {
		o.BeforeClose = f
		return nil
	}
}

// WithAfterClose configures a hook that is called _after_ Rest.
// Defaults to noop.
func WithAfterClose(f AfterCloseFunc) Option {
	return func(o *Options) error {
		o.AfterClose = f
		return nil
	}
}
