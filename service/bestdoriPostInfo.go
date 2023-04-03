package service

import (
	"github.com/6QHTSK/Bestdori-Proxy/cache"
	"github.com/6QHTSK/Bestdori-Proxy/config"
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"github.com/6QHTSK/Bestdori-Proxy/models"
	"github.com/6QHTSK/Bestdori-Proxy/models/bestdori"
	"strings"
)

func FetchBandName(bandID int) (bandName string, err error) {
	// If Cached
	bandName, err = cache.AllJson.GetBandName(bandID)
	if err == nil {
		return bandName, nil
	}
	// Else Fetch
	var replyBandList bestdori.BandAllJson
	err = httpGet(config.BandAllJsonUrl, &replyBandList)
	if err != nil {
		return bandName, err
	}
	bandList := replyBandList.Convert()
	err = cache.AllJson.SetBandList(bandList)
	if err != nil {
		return bandName, err
	}
	if bandName, ok := bandList[bandID]; !ok {
		return "", errors.BandNotFound
	} else {
		return bandName, nil
	}
}

func FetchBandoriPost(postID int, diff int, method int) (post models.Post, err error) {
	if method == models.MethodFull || method == models.MethodInfo {
		var bandoriPost bestdori.BandoriPost
		err = httpGet(config.BandoriPostListUrl(postID), &bandoriPost)
		if err != nil {
			if err == errors.RemoteReplyReject {
				return post, errors.PostNotFound
			}
			return post, err
		}
		assets := models.AssetsURL{
			Cover: bandoriPost.CoverUrl(postID),
			Audio: bandoriPost.AudioUrl(postID),
		}
		err = cache.Assets.SetCache(models.ServerBandori, postID, assets)
		if err != nil {
			return post, err
		}
		playLevel, ok := bandoriPost.Difficulty[diff]
		if !ok {
			return post, errors.PostNotFound
		}
		artists, err := FetchBandName(bandoriPost.BandID)
		if err != nil {
			return post, err
		}
		post.Info = &models.Info{
			PostID:   postID,
			Title:    bandoriPost.Title(),
			Artists:  artists,
			Username: "craftegg",
			Diff:     diff,
			Rating:   playLevel.PlayLevel,
			AudioUrl: assets.Audio,
			CoverUrl: assets.Cover,
			Time:     bandoriPost.Time(),
			Content:  bandoriPost.Content(),
		}
	}
	if method == models.MethodFull || method == models.MethodChart {
		var bandoriChart bestdori.V2Chart
		err = httpGet(config.BandoriPostUrl(postID, diff), &bandoriChart)
		if err != nil {
			return post, err
		}
		err = bandoriChart.ChartCheck()
		if err != nil {
			return post, err
		}
		post.Chart = &bandoriChart
	}
	return post, nil
}

func getBestdoriAssetsUrl(postID int, post bestdori.BestdoriPost) (assets models.AssetsURL, err error) {
	err = errors.AssetTypeErr
	if post.Post.Song.Type == "custom" {
		assets = models.AssetsURL{
			Cover: strings.ToValidUTF8(post.Post.Song.Cover, ""),
			Audio: strings.ToValidUTF8(post.Post.Song.Audio, ""),
		}
		err = nil
	} else if post.Post.Song.Type == "bandori" {
		assets, err = FetchBandoriAssetsUrl(post.Post.Song.ID)
	} else if post.Post.Song.Type == "llsif" {
		assets, err = FetchLLSifAssetsUrl(post.Post.Song.ID)
	}
	// Cache
	if err == nil {
		err = cache.Assets.SetCache(models.ServerBestdori, postID, assets)
	}
	return assets, err
}

func FetchBestdoriPost(postID int, method int) (post models.Post, err error) {
	var bestdoriPost bestdori.BestdoriPost
	err = httpGet(config.BestdoriPostUrl(postID), &bestdoriPost)
	if err != nil {
		return post, err
	}
	if !bestdoriPost.IsChart() {
		return post, errors.PostNotFound
	}
	assets, err := getBestdoriAssetsUrl(postID, bestdoriPost)
	if err != nil {
		return post, err
	}
	if method == models.MethodFull || method == models.MethodInfo {
		post.Info = &models.Info{
			PostID:   postID,
			Title:    bestdoriPost.GetTitle(),
			Artists:  bestdoriPost.GetArtists(),
			Username: bestdoriPost.GetUsername(),
			Nickname: bestdoriPost.GetNickname(),
			Diff:     bestdoriPost.Post.Diff,
			Rating:   bestdoriPost.Post.Level,
			AudioUrl: assets.Audio,
			CoverUrl: assets.Cover,
			Time:     bestdoriPost.Post.Time / 1000,
			Content:  bestdoriPost.GetContent(),
		}
	}
	if method == models.MethodFull || method == models.MethodChart {
		err = bestdoriPost.Post.Chart.ChartCheck()
		if err != nil {
			return post, err
		}
		post.Chart = &bestdoriPost.Post.Chart
	}
	return post, nil
}
