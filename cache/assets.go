package cache

import (
	"Bestdori-Proxy/errors"
	"Bestdori-Proxy/models"
	"encoding/json"
	"github.com/allegro/bigcache"
	"strconv"
	"time"
)

type assetsCache struct {
	cache *bigcache.BigCache
}

var Assets assetsCache

func init() {
	var err error
	Assets.cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Hour * 24)) // Assets 过期时间为24小时
	defer func(cache *bigcache.BigCache) {
		err := cache.Close()
		if err != nil {
			panic(err)
		}
	}(Assets.cache)
	if err != nil {
		panic(err)
	}
}

func (c *assetsCache) SetCache(chartID int, assets models.AssetsURL) error {
	assetsRaw, err := json.Marshal(assets)
	if err != nil {
		return errors.JsonMarshalErr
	}
	err = c.cache.Set(strconv.Itoa(chartID), assetsRaw)
	if err != nil {
		return errors.CacheSetErr
	}
	return nil
}

func (c *assetsCache) GetCache(chartID int) (assets models.AssetsURL, err error) {
	assetsRaw, err := c.cache.Get(strconv.Itoa(chartID))
	if err != nil {
		return assets, errors.CacheGetErr
	}
	err = json.Unmarshal(assetsRaw, &assets)
	if err != nil {
		return assets, errors.JsonUnMarshalError
	}
	return assets, nil
}
