package attributevalue

import (
	"reflect"
	"strings"
	"sync"
)

var fieldCache = &fieldCacher{}

type fieldCacheKey struct {
	typ  reflect.Type
	opts structFieldOptions
}

type fieldCacher struct {
	cache sync.Map
}

func (c *fieldCacher) Load(key fieldCacheKey) (*cachedFields, bool) {
	if v, ok := c.cache.Load(key); ok {
		return v.(*cachedFields), true
	}
	return nil, false
}

func (c *fieldCacher) LoadOrStore(key fieldCacheKey, fs *cachedFields) (*cachedFields, bool) {
	v, ok := c.cache.LoadOrStore(key, fs)
	return v.(*cachedFields), ok
}

type cachedFields struct {
	fields       []field
	fieldsByName map[string]int
}

func (f *cachedFields) All() []field {
	return f.fields
}

func (f *cachedFields) FieldByName(name string) (field, bool) {
	if i, ok := f.fieldsByName[name]; ok {
		return f.fields[i], ok
	}
	for _, f := range f.fields {
		if strings.EqualFold(f.Name, name) {
			return f, true
		}
	}
	return field{}, false
}
