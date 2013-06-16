package memcache

import (
	memcache_impl "github.com/bradfitz/gomemcache/memcache"
	"github.com/pierrre/imageserver"
)

type MemcacheCache struct {
	Memcache *memcache_impl.Client
}

func (cache *MemcacheCache) Get(key string, parameters imageserver.Parameters) (image *imageserver.Image, err error) {
	item, err := cache.Memcache.Get(key)
	if err != nil {
		return
	}
	image = &imageserver.Image{}
	err = image.Unserialize(item.Value)
	if err != nil {
		image = nil
	}
	return
}

func (cache *MemcacheCache) Set(key string, image *imageserver.Image, parameters imageserver.Parameters) (err error) {
	serialized, err := image.Serialize()
	if err != nil {
		return
	}
	item := &memcache_impl.Item{
		Key:   key,
		Value: serialized,
	}
	err = cache.Memcache.Set(item)
	return
}
