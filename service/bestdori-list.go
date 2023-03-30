package service

import (
	"Bestdori-Proxy/cache"
	"Bestdori-Proxy/models/bestdori"
	"Bestdori-Proxy/utils"
	"sort"
)

// FetchChartList Bestdori 社区-谱面页 总自制谱面列表 建议Limit设为20 Username 为空时为全谱面
func FetchChartList(offset uint, limit uint, username string) (count uint, chartList []int, err error) {
	if username == "craftegg" {
		return FetchOfficialChartList(offset, limit)
	}
	query := bestdori.ListRequestBody{
		Following:    false,
		CategoryName: "SELF_POST",
		CategoryId:   "chart",
		Order:        "TIME_DESC",
		Limit:        limit,
		Offset:       offset,
		Username:     "",
	}
	var list bestdori.ListReplyBody
	URL := "https://bestdori.com/api/post/list"
	err = utils.HttpPost(URL, query, &list)
	if err != nil {
		return count, chartList, err
	}
	for _, item := range list.Posts {
		chartList = append(chartList, item.Id)
	}
	return list.Count, chartList, nil
}

func FetchOfficialChartList(offset uint, limit uint) (count uint, chartList []int, err error) {
	officialChartList, err := cache.AllJson.GetOfficialChartList()
	if err == nil {
		return uint(len(officialChartList)), utils.GetListSliced(officialChartList, offset, limit), nil
	}
	// Fetch
	var replyChartList bestdori.ChartAllJson
	URL := "https://bestdori.com/api/bands/all.5.json"
	err = utils.HttpGet(URL, &replyChartList)
	if err != nil {
		return count, chartList, err
	}
	for key := range replyChartList {
		officialChartList = append(chartList, key)
	}
	sort.Slice(officialChartList, func(i, j int) bool {
		return officialChartList[i] > officialChartList[j]
	})
	err = cache.AllJson.SetOfficialChartList(officialChartList)
	return uint(len(officialChartList)), utils.GetListSliced(officialChartList, offset, limit), nil
}
