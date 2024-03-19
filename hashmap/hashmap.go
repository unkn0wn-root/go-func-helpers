package hashmap

import (
	"sync"
)

type Hash struct {
	sync.RWMutex
	Records map[string]interface{}
	path    string
}

func NewHash(records map[string]interface{}) *Hash {
	return &Hash{Records: records}
}

func (i *Hash) Add(id string, v interface{}) {
	i.Lock()
	defer i.Unlock()
	i.Records[id] = v

    return
}

func (i *Hash) Remove(id string) {
	i.Lock()
	defer i.Unlock()
	delete(i.Records, id)
}

func (i *Hash) Exists(id string) bool {
	i.Lock()
	defer i.Unlock()
	_, exists := i.Records[id]

    return exists
}

func (i *Hash) Get(id string) interface{} {
	i.Lock()
	defer i.Unlock()
	r := i.Records[id]

    return r
}

func (i *Hash) AddMultiple(m map[string]interface{}) {
	i.Lock()
	defer i.Unlock()
	for k, r := range m {
		i.Records[k] = r
	}
}

func (i *Hash) Map(fn func(record interface{}) bool) {
	i.Lock()
	defer i.Unlock()
    for _, r := range i.Records {
		if fn(r) {
			return
		}
	}
}

func (i *Hash) Length() int {
	i.Lock()
	defer i.Unlock()

    return len(i.Records)
}

func (i *Hash) Copy() map[string]interface{} {
	i.Lock()
	defer i.Unlock()
	c := make(map[string]interface{})
    for k, v := range i.Records {
		c[k] = v
	}

    return c
}
