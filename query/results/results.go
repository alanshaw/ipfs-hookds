package results

import (
	"github.com/ipfs/go-datastore/query"
	"github.com/jbenet/goprocess"
)

type Results struct {
	res     query.Results
	options Options
}

func NewResults(res query.Results, options ...Option) *Results {
	opts := Options{}
	opts.Apply(options...)
	return &Results{res: res, options: opts}
}

func (hres *Results) Query() query.Query {
	return hres.res.Query()
}

func (hres *Results) Next() <-chan query.Result {
	if hres.options.BeforeNext != nil {
		hres.options.BeforeNext()
	}
	c := hres.res.Next()
	if hres.options.AfterNext != nil {
		c = hres.options.AfterNext(c)
	}
	return c
}

func (hres *Results) NextSync() (query.Result, bool) {
	if hres.options.BeforeNextSync != nil {
		hres.options.BeforeNextSync()
	}
	r, ok := hres.res.NextSync()
	if hres.options.AfterNextSync != nil {
		r, ok = hres.options.AfterNextSync(r, ok)
	}
	return r, ok
}

func (hres *Results) Rest() ([]query.Entry, error) {
	if hres.options.BeforeRest != nil {
		hres.options.BeforeRest()
	}
	es, err := hres.res.Rest()
	if hres.options.AfterRest != nil {
		es, err = hres.options.AfterRest(es, err)
	}
	return es, err
}

func (hres *Results) Close() error {
	if hres.options.BeforeClose != nil {
		hres.options.BeforeClose()
	}
	err := hres.res.Close()
	if hres.options.AfterClose != nil {
		err = hres.options.AfterClose(err)
	}
	return err
}

func (hres *Results) Process() goprocess.Process {
	return hres.res.Process()
}
