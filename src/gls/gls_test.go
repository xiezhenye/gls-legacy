package gls

import (
	"testing"
	//"runtime/debug"
	"sync"
)

func test1() bool {
	ctx := GetContext()
	if ctx.Id() != 1 {
		return false
	}
	if ctx.Get("hello").(string) != "world" {
		return false
	}
	return true
}


func TestGoroutineId(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	WithNewContext(func(){
		ctx := GetContext()
		if ctx.Id() != 1 {
			t.Errorf("first ctx id != 1")
		}
		ctx.Put("hello", "world")
		ok := test1()
		if ! ok {
			t.Errorf("context failed")
		}
		Go(func(){
			ctx := GetContext()
			if ctx == nil {
				t.Fatal("ctx is nil")
			}
			if ctx.Id() != 1 {
				t.Fatal("ctx id wrong")
			}
			if ctx.Get("hello").(string) != "world" {
				t.Fatal("ctx value wrong")
			}
			wg.Done()
		})

	})
	wg.Wait()
	if len(contexts) > 0 {
		t.Error("memory leaks!")
	}
}

func TestStart(t *testing.T) {
	start(256, func() {
		id := getContextId()
		if id != 256 {
			t.Fatal("ctx id wrong!")
		}
	})
	start(65536, func() {
		id := getContextId()
		if id != 65536 {
			t.Fatal("ctx id wrong!")
		}
	})
	start(0xabcdef, func() {
		id := getContextId()
		if id != 0xabcdef {
			t.Fatal("ctx id wrong!")
		}
	})
}