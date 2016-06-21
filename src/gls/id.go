package gls
import "sync"

type IdPool struct  {
	lock     sync.Mutex
	maxId    uint32
	released []uint32
}

var idPool = IdPool{ maxId: 1 }

func (self *IdPool) get() uint32 {
	self.lock.Lock()
	var ret uint32
	size := len(self.released)
	if size > 0 {
		ret = self.released[size-1]
		self.released = self.released[0:size-1]
		return ret
	} else {
		ret = self.maxId
		self.maxId++
	}
	self.lock.Unlock()
	return ret
}

func (self *IdPool) put(i uint32) {
	self.lock.Lock()
	self.released = append(self.released, i)
	self.lock.Unlock()
}


