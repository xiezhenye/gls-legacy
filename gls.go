package gls

import (
	"sync"
	"sync/atomic"
)


type Context struct {
	lock     sync.Mutex
	id       uint32
	ref      int32
	values   map[string]interface{}
}


var contexts = make(map[uint32]*Context)
var contextsLock = sync.Mutex{}

func WithContext(ctx *Context, f func()) {
	atomic.AddInt32(&ctx.ref, 1)
	start(ctx.id, f)
	newRef := atomic.AddInt32(&ctx.ref, -1)
	if newRef <= 0 {
		contextsLock.Lock()
		delete(contexts, ctx.id)
		contextsLock.Unlock()
		idPool.put(ctx.id)
	}
}

func newContext() *Context {
	ret := &Context {
		id: idPool.get(),
		ref: 0,
		values: make(map[string]interface{}),
	}
	contextsLock.Lock()
	contexts[ret.id] = ret
	contextsLock.Unlock()
	return ret
}

func WithNewContext(f func()) {
	WithContext(newContext(), f)
}

func (self *Context) Id() uint32 {
	return self.id
}

func (self *Context) Put(k string, v interface{}) {
	self.lock.Lock()
	self.values[k] = v
	self.lock.Unlock()
}

func (self *Context) Get(k string) interface{} {
	self.lock.Lock()
	ret := self.values[k]
	self.lock.Unlock()
	return ret
}

func GetContext() *Context {
	id := getContextId()
	contextsLock.Lock()
	ret := contexts[id]
	contextsLock.Unlock()
	return ret
}

func Go(f func()) {
	ctx := GetContext()
	atomic.AddInt32(&ctx.ref, 1)
	go WithContext(ctx, func(){
		atomic.AddInt32(&ctx.ref, -1)
		f()
	})

}
