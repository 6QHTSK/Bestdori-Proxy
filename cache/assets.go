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
	cache [models.TotalServerCnt]*bigcache.BigCache
}

var Assets assetsCache

func init() {
	var err error
	for i := 0; i < models.TotalServerCnt; i++ {
		Assets.cache[i], err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Hour * 24)) // Assets 过期时间为24小时
		if err != nil {
			panic(err)
		}
	}
}

func (c *assetsCache) SetCache(server int, postID int, assets models.AssetsURL) error {
	if server < 0 || server >= models.TotalServerCnt {
		return errors.UnknownServerErr
	}
	assetsRaw, err := json.Marshal(assets)
	if err != nil {
		return errors.JsonMarshalErr
	}
	err = c.cache[server].Set(strconv.Itoa(postID), assetsRaw)
	if err != nil {
		return errors.CacheSetErr
	}
	return nil
}

func (c *assetsCache) GetCache(server int, postID int) (assets models.AssetsURL, err error) {
	if server < 0 || server >= models.TotalServerCnt {
		return assets, errors.UnknownServerErr
	}
	assetsRaw, err := c.cache[server].Get(strconv.Itoa(postID))
	if err != nil {
		return assets, errors.CacheGetErr
	}
	err = json.Unmarshal(assetsRaw, &assets)
	if err != nil {
		return assets, errors.JsonUnMarshalError
	}
	return assets, nil
}
