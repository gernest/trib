package trib

import (
	"net/http"
	"time"

	gorilla "github.com/gorilla/context"
	"golang.org/x/net/context"
)

const ctxKey = "req_ctx"

type Plan interface {
	ID() string
	Time
	Stages() ([]Stage, bool)
}

type Stage interface {
	ID() string
	Time
	Steps() ([]Step, bool)
}

type Step interface {
	ID() string
	Time
	Parallel() bool
	Exec(context.Context) error
}

type Time interface {
	Duration() (time.Duration, bool)
	SetDuration(time.Duration)
}

type Result interface {
	ID() string
	Dtata() interface{}
}

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func NewContext(r *http.Request) *Context {
	return &Context{Request: r}
}

func (c *Context) Clear() {
	gorilla.Clear(c.Request)
}

func (c *Context) Delete(key interface{}) {
	gorilla.Delete(c.Request, key)
}

func (c *Context) Get(key interface{}) interface{} {
	return gorilla.Get(c.Request, key)
}

func (c *Context) GetAll() map[interface{}]interface{} {
	return gorilla.GetAll(c.Request)
}

func (c *Context) GetAllOk() (map[interface{}]interface{}, bool) {
	return gorilla.GetAllOk(c.Request)
}

func (c *Context) GetOk(key interface{}) (interface{}, bool) {
	return gorilla.GetOk(c.Request, key)
}

func (c *Context) Set(key, val interface{}) {
	gorilla.Set(c.Request, key, val)
}

func Clear(ctx context.Context) {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			c.Clear()
		}
	}
}
func Delete(ctx context.Context, key interface{}) {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			c.Delete(key)
		}
	}
}
func Get(ctx context.Context, key interface{}) interface{} {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			return c.Get(key)
		}
	}
	return nil
}

func GetAll(ctx context.Context) map[interface{}]interface{} {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			return c.GetAll()
		}
	}
	return nil
}

func GetAllOk(ctx context.Context) (map[interface{}]interface{}, bool) {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			return c.GetAllOk()
		}
	}
	return nil, false
}
func GetOk(ctx context.Context, key interface{}) (interface{}, bool) {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			return c.GetOk(key)
		}
	}
	return nil, false
}

func Set(ctx context.Context, key, val interface{}) {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			c.Set(key, val)
		}
	}
}

func ActiveCtx(ctx context.Context) *Context {
	v := ctx.Value(ctxKey)
	if v != nil {
		if c, ok := v.(*Context); ok {
			return c
		}
	}
	return nil
}

type Frontend interface {
	Name() string
	Plan(*http.Request) (Plan, error)
	Commit(context.Context)
}
