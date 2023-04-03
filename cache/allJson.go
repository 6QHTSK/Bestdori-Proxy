package cache

import (
	"encoding/json"
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"github.com/6QHTSK/Bestdori-Proxy/models/bestdori"
	"github.com/allegro/bigcache"
	"time"
)

type allJsonCache struct {
	cache *bigcache.BigCache
}

var AllJson allJsonCache

func init() {
	var err error
	AllJson.cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Hour * 24)) // AllJson 过期时间为24小时
	defer func(cache *bigcache.BigCache) {
		err := cache.Close()
		if err != nil {
			panic(err)
		}
	}(AllJson.cache)
	if err != nil {
		panic(err)
	}
}

func (c *allJsonCache) GetBandName(bandID int) (bandName string, err error) {
	BandListRaw, err := AllJson.cache.Get("Band")
	if err != nil {
		return bandName, errors.CacheGetErr
	}
	var bandList bestdori.BandNameMap
	err = json.Unmarshal(BandListRaw, &bandList)
	if err != nil {
		return bandName, errors.JsonUnMarshalError
	}
	bandName, ok := bandList[bandID]
	if !ok {
		return bandName, errors.CacheNotFound
	}
	return bandName, nil
}

func (c *allJsonCache) SetBandList(bandList bestdori.BandNameMap) error {
	bandListRaw, err := json.Marshal(bandList)
	if err != nil {
		return errors.JsonMarshalErr
	}
	err = AllJson.cache.Set("Band", bandListRaw)
	if err != nil {
		return errors.CacheSetErr
	}
	return nil
}

func (c *allJsonCache) GetOfficialPostList() (officialList []int, err error) {
	postListRaw, err := AllJson.cache.Get("officialPost")
	if err != nil {
		return officialList, errors.CacheGetErr
	}
	err = json.Unmarshal(postListRaw, &officialList)
	if err != nil {
		return officialList, errors.JsonUnMarshalError
	}
	return officialList, nil
}

func (c *allJsonCache) SetOfficialPostList(officialPostList []int) error {
	postListRaw, err := json.Marshal(officialPostList)
	if err != nil {
		return errors.JsonMarshalErr
	}
	err = c.cache.Set("officialPost", postListRaw)
	if err != nil {
		return errors.CacheSetErr
	}
	return nil
}
