package cache

import (
	"Bestdori-Proxy/errors"
	"Bestdori-Proxy/models/bestdori"
	"encoding/json"
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

func (c *allJsonCache) GetLLSifItem(musicID int) (item bestdori.LLSifAllJsonItem, err error) {
	LLSifListRaw, err := AllJson.cache.Get("LLSif")
	if err != nil {
		return item, err
	}
	var LLSifList bestdori.LLSifAllJson
	err = json.Unmarshal(LLSifListRaw, &LLSifList)
	if err != nil {
		return item, errors.JsonUnMarshalError
	}
	item, ok := LLSifList[musicID]
	if !ok {
		return item, errors.CacheNotFound
	}
	return item, nil
}

func (c *allJsonCache) SetLLSifList(LLSifList bestdori.LLSifAllJson) error {
	LLSifListRaw, err := json.Marshal(LLSifList)
	if err != nil {
		return errors.JsonMarshalErr
	}
	err = AllJson.cache.Set("LLSif", LLSifListRaw)
	if err != nil {
		return errors.CacheSetErr
	}
	return nil
}

func (c *allJsonCache) GetOfficialChartList() (officialList []int, err error) {
	chartListRaw, err := AllJson.cache.Get("Chart")
	if err != nil {
		return officialList, errors.CacheGetErr
	}
	err = json.Unmarshal(chartListRaw, &officialList)
	if err != nil {
		return officialList, errors.JsonUnMarshalError
	}
	return officialList, nil
}

func (c *allJsonCache) SetOfficialChartList(officialChartList []int) error {
	chartListRaw, err := json.Marshal(officialChartList)
	if err != nil {
		return errors.JsonMarshalErr
	}
	err = c.cache.Set("Chart", chartListRaw)
	if err != nil {
		return errors.CacheSetErr
	}
	return nil
}
