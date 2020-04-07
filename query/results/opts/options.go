package opts

import (
	"fmt"

	"github.com/ipfs/go-datastore/query"
)

// OnBeforeNextFunc is a handler for the before Next hook
type OnBeforeNextFunc func()

// OnAfterNextFunc is a handler for the after Next hook
type OnAfterNextFunc func(<-chan query.Result) <-chan query.Result

// OnBeforeNextSyncFunc is a handler for the before NextSync hook
type OnBeforeNextSyncFunc func()

// OnAfterNextSyncFunc is a handler for the after NextSync hook
type OnAfterNextSyncFunc func(query.Result, bool) (query.Result, bool)

// OnBeforeRestFunc is a handler for the before Rest hook
type OnBeforeRestFunc func() ([]query.Entry, error)

// OnAfterRestFunc is a handler for the after Rest hook
type OnAfterRestFunc func([]query.Entry, error) ([]query.Entry, error)

// OnBeforeCloseFunc is a handler for the before Close hook
type OnBeforeCloseFunc func()

// OnAfterCloseFunc is a handler for the after Close hook
type OnAfterCloseFunc func(error) error

// Options are results options.
type Options struct {
	OnBeforeNext     OnBeforeNextFunc
	OnAfterNext      OnAfterNextFunc
	OnBeforeNextSync OnBeforeNextSyncFunc
	OnAfterNextSync  OnAfterNextSyncFunc
	OnBeforeRest     OnBeforeRestFunc
	OnAfterRest      OnAfterRestFunc
	OnBeforeClose    OnBeforeCloseFunc
	OnAfterClose     OnAfterCloseFunc
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

// OnBeforeNext configures a hook that is called _before_ Next.
// Defaults to noop.
func OnBeforeNext(f OnBeforeNextFunc) Option {
	return func(o *Options) error {
		o.OnBeforeNext = f
		return nil
	}
}

// OnAfterNext configures a hook that is called _after_ Next.
// Defaults to noop.
func OnAfterNext(f OnAfterNextFunc) Option {
	return func(o *Options) error {
		o.OnAfterNext = f
		return nil
	}
}

// OnBeforeNextSync configures a hook that is called _before_ NextSync.
// Defaults to noop.
func OnBeforeNextSync(f OnBeforeNextSyncFunc) Option {
	return func(o *Options) error {
		o.OnBeforeNextSync = f
		return nil
	}
}

// OnAfterNextSync configures a hook that is called _after_ NextSync.
// Defaults to noop.
func OnAfterNextSync(f OnAfterNextSyncFunc) Option {
	return func(o *Options) error {
		o.OnAfterNextSync = f
		return nil
	}
}

// OnBeforeRest configures a hook that is called _before_ Rest.
// Defaults to noop.
func OnBeforeRest(f OnBeforeRestFunc) Option {
	return func(o *Options) error {
		o.OnBeforeRest = f
		return nil
	}
}

// OnAfterRest configures a hook that is called _after_ Rest.
// Defaults to noop.
func OnAfterRest(f OnAfterRestFunc) Option {
	return func(o *Options) error {
		o.OnAfterRest = f
		return nil
	}
}

// OnBeforeClose configures a hook that is called _before_ Close.
// Defaults to noop.
func OnBeforeClose(f OnBeforeCloseFunc) Option {
	return func(o *Options) error {
		o.OnBeforeClose = f
		return nil
	}
}

// OnAfterClose configures a hook that is called _after_ Rest.
// Defaults to noop.
func OnAfterClose(f OnAfterCloseFunc) Option {
	return func(o *Options) error {
		o.OnAfterClose = f
		return nil
	}
}
