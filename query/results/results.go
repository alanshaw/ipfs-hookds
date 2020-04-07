package results

import (
	resopts "github.com/alanshaw/ipfs-hookds/query/results/opts"
	"github.com/ipfs/go-datastore/query"
	"github.com/jbenet/goprocess"
)

type Results struct {
	res     query.Results
	options resopts.Options
}

func NewResults(res query.Results, options ...resopts.Option) *Results {
	opts := resopts.Options{}
	opts.Apply(options...)
	return &Results{res: res, options: opts}
}

func (hres *Results) Query() query.Query {
	return hres.res.Query()
}

func (hres *Results) Next() <-chan query.Result {
	if hres.options.OnBeforeNext != nil {
		hres.options.OnBeforeNext()
	}
	c := hres.res.Next()
	if hres.options.OnAfterNext != nil {
		c = hres.options.OnAfterNext(c)
	}
	return c
}

func (hres *Results) NextSync() (query.Result, bool) {
	if hres.options.OnBeforeNextSync != nil {
		hres.options.OnBeforeNextSync()
	}
	r, ok := hres.res.NextSync()
	if hres.options.OnAfterNextSync != nil {
		r, ok = hres.options.OnAfterNextSync(r, ok)
	}
	return r, ok
}

func (hres *Results) Rest() ([]query.Entry, error) {
	if hres.options.OnBeforeRest != nil {
		hres.options.OnBeforeRest()
	}
	es, err := hres.res.Rest()
	if hres.options.OnAfterRest != nil {
		es, err = hres.options.OnAfterRest(es, err)
	}
	return es, err
}

func (hres *Results) Close() error {
	if hres.options.OnBeforeClose != nil {
		hres.options.OnBeforeClose()
	}
	err := hres.res.Close()
	if hres.options.OnAfterClose != nil {
		err = hres.options.OnAfterClose(err)
	}
	return err
}

func (hres *Results) Process() goprocess.Process {
	return hres.res.Process()
}
