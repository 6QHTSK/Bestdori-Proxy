package service

import (
	"fmt"
	"github.com/6QHTSK/ayachan-bestdoriAPI/models"
	"github.com/6QHTSK/ayachan-bestdoriAPI/utils"
	"net/http"
	"strings"
)

func getDiffStr(diff models.DiffType) string {
	switch diff {
	case models.Diff_Easy:
		return "easy"
	case models.Diff_Normal:
		return "normal"
	case models.Diff_Hard:
		return "hard"
	case models.Diff_Expert:
		return "expert"
	case models.Diff_Special:
		return "special"
	default:
		return "expert"
	}
}

func FetchChartInfo(chartID int, diff models.DiffType) (chart models.ProcessedChartItem, errorCode int, err error) {
	var bestdoriChart models.BestdoriChartItem
	var sonolusChart models.SonolusChartItem
	if chartID <= 1010 {
		var officialMap models.BestdoriV2Chart
		officialURL := fmt.Sprintf("https://bestdori.com/api/charts/%d/%s.json", chartID, getDiffStr(diff))
		_, err := utils.HttpGet(officialURL, officialMap)
		// 如果是官谱
		if err == nil {
			officialURL := fmt.Sprintf("https://servers.sonolus.com/bandori/sonolus/levels/bandori-%d-%s", chartID, getDiffStr(diff))
			_, err := utils.HttpGet(officialURL, &sonolusChart)
			if err != nil {
				return chart, http.StatusBadGateway, err
			}
			chart = models.ProcessedChartItem{
				ChartID: chartID,
				Title:   sonolusChart.Item.Title,
				Artists: sonolusChart.Item.Artists,
				Author:  models.Author{Username: "__CraftEgg__"},
				Diff:    diff,
				Level:   sonolusChart.Item.Rating,
				SongUrl: models.AssetsURL{
					Cover: sonolusChart.Item.Cover.Url,
					Audio: sonolusChart.Item.Bgm.Url,
				},
				Official: true,
				Chart:    officialMap,
			}
			return chart, http.StatusOK, nil
		}
	}
	//如果是自制
	customURL := fmt.Sprintf("https://bestdori.com/api/post/details?id=%d", chartID)
	errorCode, err = utils.HttpGet(customURL, &bestdoriChart)
	if err != nil {
		return chart, errorCode, err
	}
	if !bestdoriChart.Result || bestdoriChart.Post.CategoryName != "SELF_POST" || bestdoriChart.Post.CategoryId != "chart" {
		return chart, http.StatusNotFound, fmt.Errorf("谱面未找到")
	}
	customURL = fmt.Sprintf("https://servers.sonolus.com/bestdori/sonolus/levels/bestdori-%d", chartID)
	errorCode, err = utils.HttpGet(customURL, &sonolusChart)
	if err != nil {
		return chart, errorCode, err
	}
	bestdoriChart.Post.Author.Username = strings.ToValidUTF8(bestdoriChart.Post.Author.Username, "")
	bestdoriChart.Post.Author.Nickname = strings.ToValidUTF8(bestdoriChart.Post.Author.Nickname, "")
	chart = models.ProcessedChartItem{
		ChartID: chartID,
		Title:   strings.ToValidUTF8(sonolusChart.Item.Title, ""),
		Artists: strings.ToValidUTF8(sonolusChart.Item.Artists, ""),
		Author:  bestdoriChart.Post.Author,
		Diff:    models.DiffType(bestdoriChart.Post.Diff),
		Level:   sonolusChart.Item.Rating,
		SongUrl: models.AssetsURL{
			Cover: strings.ToValidUTF8(sonolusChart.Item.Cover.Url, ""),
			Audio: strings.ToValidUTF8(sonolusChart.Item.Bgm.Url, ""),
		},
		Official: false,
		Likes:    bestdoriChart.Post.Likes,
		Time:     bestdoriChart.Post.Time / 1000,
		Chart:    bestdoriChart.Post.Chart,
		Content:  strings.ToValidUTF8(sonolusChart.Description, ""),
	}
	return chart, http.StatusOK, nil
}

func FetchMap(chartID int, diff models.DiffType) (Map models.BestdoriV2Chart, errorCode int, err error) {
	if chartID < 1010 {
		officialURL := fmt.Sprintf("https://bestdori.com/api/charts/%d/%s.json", chartID, getDiffStr(diff))
		_, err := utils.HttpGet(officialURL, &Map)
		if err == nil {
			return Map, http.StatusOK, nil
		}
	}
	var bestdoriChart models.BestdoriCustomMap
	customURL := fmt.Sprintf("https://bestdori.com/api/post/details?id=%d", chartID)
	errCode, err := utils.HttpGet(customURL, &bestdoriChart)
	if err != nil {
		return Map, errCode, err
	}
	if !bestdoriChart.Result || bestdoriChart.Post.CategoryName != "SELF_POST" || bestdoriChart.Post.CategoryId != "chart" {
		return Map, http.StatusNotFound, fmt.Errorf("谱面未找到")
	}
	//result,err := bestdoriChart.Post.Chart.MapCheck()
	//if !result{
	//	return Map, http.StatusBadGateway, err
	//}
	return bestdoriChart.Post.Chart, http.StatusOK, nil
}

type listQueryObject struct {
	Following    bool   `json:"following"`
	CategoryName string `json:"categoryName"`
	CategoryId   string `json:"categoryId"`
	Order        string `json:"order"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}
type listQueryResult struct {
	Result bool                         `json:"result"`
	Count  int                          `json:"count"`
	Posts  []models.ListQueryResultItem `json:"posts"`
}

func FetchChartList(page int, limit int) (count int, chartList []models.ListQueryResultItem, errorCode int, err error) {
	query := listQueryObject{
		Following:    false,
		CategoryName: "SELF_POST",
		CategoryId:   "chart",
		Order:        "TIME_DESC",
		Limit:        limit,
		Offset:       page * limit,
	}
	var list listQueryResult
	URL := "https://bestdori.com/api/post/list"
	errorCode, err = utils.HttpPost(URL, query, &list)
	if err != nil {
		return count, chartList, errorCode, err
	}
	for _, item := range list.Posts {
		chartList = append(chartList, item)
	}
	return list.Count, chartList, http.StatusOK, nil
}

type engineQueryResult struct {
	Item map[string]interface{} `json:"item"`
}

func FetchEngine() (engine models.Engine, errCode int, err error) {
	var queryResult engineQueryResult
	URL := "https://servers.sonolus.com/bestdori/sonolus/engines/bandori"
	errCode, err = utils.HttpGet(URL, &queryResult)
	if err != nil {
		return engine, errCode, err
	}
	engine.Engine = queryResult.Item
	engine.SetTrue()
	return engine, http.StatusOK, nil
}
