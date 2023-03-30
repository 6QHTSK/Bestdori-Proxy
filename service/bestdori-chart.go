package service

import (
	"Bestdori-Proxy/cache"
	"Bestdori-Proxy/config"
	"Bestdori-Proxy/errors"
	"Bestdori-Proxy/models"
	"Bestdori-Proxy/models/bestdori"
	"Bestdori-Proxy/utils"
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
	err = utils.HttpGet(config.BandAllJsonUrl, &replyBandList)
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

func FetchOfficialChart(chartID int, diff int, method models.WorkMethod) (post models.Post, err error) {
	if method == models.FULL || method == models.INFO {
		var officialPost bestdori.OfficialPost
		err = utils.HttpGet(config.OfficialPostUrl(chartID), &officialPost)
		if err != nil {
			return post, err
		}
		assets := models.AssetsURL{
			Cover: officialPost.CoverUrl(chartID),
			Audio: officialPost.SongUrl(chartID),
		}
		err = cache.Assets.SetCache(chartID, assets)
		if err != nil {
			return post, err
		}
		playLevel, ok := officialPost.Difficulty[diff]
		if !ok {
			return post, errors.ChartNotFound
		}
		artists, err := FetchBandName(officialPost.BandID)
		if err != nil {
			return post, err
		}
		post.Info = &models.Info{
			ChartID:  chartID,
			Title:    officialPost.Title(),
			Artists:  artists,
			Username: "craftegg",
			Diff:     diff,
			Rating:   playLevel.PlayLevel,
			MusicUrl: assets.Audio,
			CoverUrl: assets.Cover,
			Time:     officialPost.Time(),
			Content:  officialPost.Content(),
		}
	}
	if method == models.FULL || method == models.CHART {
		var officialMap bestdori.V2Chart
		err = utils.HttpGet(config.OfficialChartUrl(chartID, diff), &officialMap)
		if err != nil {
			return post, err
		}
		err = officialMap.MapCheck()
		if err != nil {
			return post, err
		}
		post.Chart = &officialMap
	}
	return post, nil
}

func FetchOfficialAssetsUrl(chartID int) (assets models.AssetsURL, err error) {
	// If Cache
	assets, err = cache.Assets.GetCache(chartID)
	if err == nil {
		return assets, nil
	}
	// Else Fetch
	var officialPost bestdori.OfficialPost
	err = utils.HttpGet(config.OfficialPostUrl(chartID), &officialPost)
	if err != nil {
		return assets, err
	}
	return models.AssetsURL{
		Cover: officialPost.CoverUrl(chartID),
		Audio: officialPost.SongUrl(chartID),
	}, nil
}

func FetchLLSifAssetsUrl(musicID int) (assets models.AssetsURL, err error) {
	// If Cached
	LLSifItem, err := cache.AllJson.GetLLSifItem(musicID)
	if err == nil {
		return models.AssetsURL{
			Cover: LLSifItem.CoverUrl(),
			Audio: LLSifItem.SongUrl(),
		}, nil
	}
	// Else Fetch
	var LLSifList bestdori.LLSifAllJson
	err = utils.HttpGet(config.LLSifAllJsonUrl, &LLSifList)
	if err != nil {
		return assets, err
	}
	err = cache.AllJson.SetLLSifList(LLSifList)
	if err != nil {
		return assets, err
	}
	if item, ok := LLSifList[musicID]; !ok {
		return assets, errors.ChartNotFound
	} else {
		return models.AssetsURL{
			Cover: item.CoverUrl(),
			Audio: item.SongUrl(),
		}, nil
	}
}

func FetchFanmadeAssetsUrl(chartID int, post bestdori.FanmadePost) (assets models.AssetsURL, err error) {
	err = errors.AssetTypeErr
	if post.Post.Song.Type == "custom" {
		assets = models.AssetsURL{
			Cover: strings.ToValidUTF8(post.Post.Song.Cover, ""),
			Audio: strings.ToValidUTF8(post.Post.Song.Audio, ""),
		}
		err = nil
	} else if post.Post.Song.Type == "bandori" {
		assets, err = FetchOfficialAssetsUrl(post.Post.Song.ID)
	} else if post.Post.Song.Type == "llsif" {
		assets, err = FetchLLSifAssetsUrl(post.Post.Song.ID)
	}
	// 统一Cache 刷新Cache时间
	if err == nil {
		err = cache.Assets.SetCache(chartID, assets)
	}
	return assets, err
}

func FetchFanmadeChart(chartID int, method models.WorkMethod) (post models.Post, err error) {
	var bestdoriChart bestdori.FanmadePost
	err = utils.HttpGet(config.FanmadeUrl(chartID), &bestdoriChart)
	if err != nil {
		return post, err
	}
	if !bestdoriChart.IsChart() {
		return post, errors.ChartNotFound
	}
	assets, err := FetchFanmadeAssetsUrl(chartID, bestdoriChart)
	if err != nil {
		return post, err
	}
	if method == models.FULL || method == models.INFO {
		post.Info = &models.Info{
			ChartID:  chartID,
			Title:    bestdoriChart.GetTitle(),
			Artists:  bestdoriChart.GetArtists(),
			Username: bestdoriChart.GetUsername(),
			Nickname: bestdoriChart.GetNickname(),
			Diff:     bestdoriChart.Post.Diff,
			Rating:   bestdoriChart.Post.Level,
			MusicUrl: assets.Audio,
			CoverUrl: assets.Cover,
			Time:     bestdoriChart.Post.Time / 1000,
			Content:  bestdoriChart.GetContent(),
		}
	}
	if method == models.FULL || method == models.CHART {
		err = bestdoriChart.Post.Chart.MapCheck()
		if err != nil {
			return post, err
		}
		post.Chart = &bestdoriChart.Post.Chart
	}
	return post, nil
}

func FetchChartInfo(chartID int, diff int, official bool, method models.WorkMethod) (chart models.Post, err error) {
	if chartID < 921 || official {
		chart, err = FetchOfficialChart(chartID, diff, method)
	} else {
		chart, err = FetchFanmadeChart(chartID, method)
	}
	return chart, err
}

func FetchAssetsUrl(chartID int) (assets models.AssetsURL, err error) {
	assets, err = cache.Assets.GetCache(chartID)
	if err == nil {
		return assets, nil
	}
	// 如果是官谱
	if chartID <= 1010 {
		assets, err = FetchOfficialAssetsUrl(chartID)
		if err == nil {
			// 此处要cache
			err = cache.Assets.SetCache(chartID, assets)
			return assets, nil
		}
	}
	chart, err := FetchFanmadeChart(chartID, models.INFO)
	if err != nil {
		return assets, err
	}
	// 已经Cache过了
	return models.AssetsURL{
		Cover: chart.CoverUrl,
		Audio: chart.MusicUrl,
	}, nil
}
