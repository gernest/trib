package trib

import (
	"net/http"

	"github.com/muesli/cache2go"
	"golang.org/x/net/context"
)

type Ulation struct {
	plans cache2go.CacheTable
	steps []Step
	e     Engine
	mux   http.ServeMux
}

func (u *Ulation) Handle(pattern string, front Frontend) {
	u.mux.Handle(pattern, u.newHandler(front))
}

func (u *Ulation) newHandler(front Frontend) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, err := front.Plan(r)
		if err != nil {
			// TODO (gernest): handle error
			return
		}
		ctx := NewContext(r)
		base := context.Background()
		bv := context.WithValue(base, ctxKey, ctx)
		err = u.e.ExecPlan(bv, p)
		if err != nil {
			// TODO (gernest): handle error
			return
		}
		front.Commit(bv)
	})
}

func (u *Ulation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.mux.ServeHTTP(w, r)
}
