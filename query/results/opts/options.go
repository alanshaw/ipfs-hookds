package opts

import (
	"fmt"

	"github.com/ipfs/go-datastore/query"
)

// Options are results options.
type Options struct {
	OnBeforeNextSync func()
	OnAfterNextSync  func(query.Result, bool) (query.Result, bool)
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

// OnBeforeNextSync configures a hook that is called _before_ NextSync.
// Defaults to noop.
func OnBeforeNextSync(f func()) Option {
	return func(o *Options) error {
		o.OnBeforeNextSync = f
		return nil
	}
}

// OnAfterNextSync configures a hook that is called _after_ NextSync.
// Defaults to noop.
func OnAfterNextSync(f func(query.Result, bool) (query.Result, bool)) Option {
	return func(o *Options) error {
		o.OnAfterNextSync = f
		return nil
	}
}
