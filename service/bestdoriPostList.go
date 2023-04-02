package service

import (
	"Bestdori-Proxy/cache"
	"Bestdori-Proxy/config"
	"Bestdori-Proxy/models/bestdori"
	"sort"
)

func getListSliced(List []int, offset uint64, limit uint64) (SlicedList []int) {
	if offset >= uint64(len(List)) {
		return SlicedList
	}
	end := offset + limit
	if end > uint64(len(List)) {
		end = uint64(len(List))
	}
	return List[offset:end]
}

// FetchBestdoriPostList Bestdori 社区-谱面页 总自制谱面列表 建议Limit设为20 Username 为空时为全谱面
func FetchBestdoriPostList(offset uint64, limit uint64, username string) (count uint64, postList []int, err error) {
	query := bestdori.ListRequestBody{
		Following:    false,
		CategoryName: "SELF_POST",
		CategoryId:   "chart",
		Order:        "TIME_DESC",
		Limit:        limit,
		Offset:       offset,
		Username:     username,
	}
	var list bestdori.ListResponseBody
	err = httpPost(config.BestdoriListUrl, query, &list)
	if err != nil {
		return count, postList, err
	}
	for _, item := range list.Posts {
		postList = append(postList, item.Id)
	}
	return list.Count, postList, nil
}

// FetchBandoriPostList Bandori 官方谱面列表 建议Limit设为20
func FetchBandoriPostList(offset uint64, limit uint64) (count uint64, postList []int, err error) {
	officialPostList, err := cache.AllJson.GetOfficialPostList()
	if err == nil {
		return uint64(len(officialPostList)), getListSliced(officialPostList, offset, limit), nil
	}
	// Fetch
	var responsePostList bestdori.OfficialPostAllJson
	err = httpGet(config.OfficialPostList, &responsePostList)
	if err != nil {
		return count, postList, err
	}
	for key := range responsePostList {
		officialPostList = append(officialPostList, key)
	}
	sort.Slice(officialPostList, func(i, j int) bool {
		return officialPostList[i] > officialPostList[j]
	})
	err = cache.AllJson.SetOfficialPostList(officialPostList)
	return uint64(len(officialPostList)), getListSliced(officialPostList, offset, limit), nil
}
