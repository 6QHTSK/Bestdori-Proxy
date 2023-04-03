package service

import (
	"github.com/6QHTSK/Bestdori-Proxy/cache"
	"github.com/6QHTSK/Bestdori-Proxy/config"
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"github.com/6QHTSK/Bestdori-Proxy/models"
	"github.com/6QHTSK/Bestdori-Proxy/models/bestdori"
)

func FetchAssetsUrl(postID int, server int) (assets models.AssetsURL, err error) {
	assets, err = cache.Assets.GetCache(server, postID)
	if err == nil {
		return assets, nil
	}
	switch server {
	case models.ServerBandori:
		return FetchBandoriAssetsUrl(postID)
	case models.ServerLLSif:
		return FetchLLSifAssetsUrl(postID)
	case models.ServerBestdori:
		post, err := FetchBestdoriPost(postID, models.MethodInfo)
		if err != nil {
			return assets, err
		}
		return models.AssetsURL{Cover: post.CoverUrl, Audio: post.AudioUrl}, nil
	}
	return assets, errors.UnknownServerErr
}

func FetchBandoriAssetsUrl(postID int) (assets models.AssetsURL, err error) {
	// If Cache
	assets, err = cache.Assets.GetCache(models.ServerBandori, postID)
	if err == nil {
		return assets, nil
	}
	// Else Fetch
	var officialPost bestdori.BandoriPost
	err = httpGet(config.BandoriPostListUrl(postID), &officialPost)
	if err != nil {
		if err == errors.RemoteReplyReject {
			return assets, errors.PostNotFound
		}
		return assets, err
	}
	assets = models.AssetsURL{
		Cover: officialPost.CoverUrl(postID),
		Audio: officialPost.AudioUrl(postID),
	}
	// Cache
	err = cache.Assets.SetCache(models.ServerBandori, postID, assets)
	return assets, nil
}

func FetchLLSifAssetsUrl(postID int) (assets models.AssetsURL, err error) {
	// If Cached
	assets, err = cache.Assets.GetCache(models.ServerLLSif, postID)
	if err == nil {
		return assets, nil
	}
	// Else Fetch
	var LLSifJson bestdori.LLSifAllJson
	err = httpGet(config.LLSifAllJsonUrl, &LLSifJson)
	if err != nil {
		return assets, err
	}
	// Cache
	for key, item := range LLSifJson {
		err = cache.Assets.SetCache(models.ServerLLSif, key, models.AssetsURL{
			Cover: item.CoverUrl(),
			Audio: item.AudioUrl(),
		})
		if err != nil {
			return assets, err
		}
	}
	// 找到对应Asset返回
	if item, ok := LLSifJson[postID]; !ok {
		return assets, errors.PostNotFound
	} else {
		return models.AssetsURL{
			Cover: item.CoverUrl(),
			Audio: item.AudioUrl(),
		}, nil
	}
}
