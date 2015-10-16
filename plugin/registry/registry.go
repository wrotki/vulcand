// This file will be generated to include all customer specific middlewares
package registry

import (
	"github.com/polyverse-security/vulcand/plugin"
	"github.com/polyverse-security/vulcand/plugin/cbreaker"
	"github.com/polyverse-security/vulcand/plugin/connlimit"
	"github.com/polyverse-security/vulcand/plugin/ratelimit"
	"github.com/polyverse-security/vulcand/plugin/rewrite"
	"github.com/polyverse-security/vulcand/plugin/trace"
)

func GetRegistry() *plugin.Registry {
	r := plugin.NewRegistry()

	specs := []*plugin.MiddlewareSpec{
		ratelimit.GetSpec(),
		connlimit.GetSpec(),
		rewrite.GetSpec(),
		cbreaker.GetSpec(),
		trace.GetSpec(),
	}

	for _, spec := range specs {
		if err := r.AddSpec(spec); err != nil {
			panic(err)
		}
	}

	return r
}
