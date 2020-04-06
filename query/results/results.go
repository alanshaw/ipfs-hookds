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
	return hres.res.Next()
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
	return hres.res.Rest()
}

func (hres *Results) Close() error {
	return hres.res.Close()
}

func (hres *Results) Process() goprocess.Process {
	return hres.res.Process()
}
