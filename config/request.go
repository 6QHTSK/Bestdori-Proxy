package config

import (
	"fmt"
)

var BandAllJsonUrl = "https://bestdori.com/api/bands/all.1.json"
var LLSifAllJsonUrl = "https://bestdori.com/api/misc/llsif.10.json"
var LLSifAssetPrefix = "https://card.niconi.co.ni/asset/"
var OfficialPostList = "https://bestdori.com/api/songs/all.5.json"
var BestdoriListUrl = "https://bestdori.com/api/post/list"

func BandoriPostListUrl(postID int) string {
	return fmt.Sprintf("https://bestdori.com/api/songs/%d.json", postID)
}

func BandoriPostUrl(postID int, diff int) string {
	str := "expert"
	if diff >= 0 && diff <= 4 {
		str = []string{"easy", "normal", "hard", "expert", "special"}[diff]
	}
	return fmt.Sprintf("https://bestdori.com/api/charts/%d/%s.json", postID, str)
}

func BestdoriPostUrl(postID int) string {
	return fmt.Sprintf("https://bestdori.com/api/post/details?id=%d", postID)
}

func BestdoriCoverUrl(region string, bundle int, jacket string) string {
	return fmt.Sprintf("https://bestdori.com/assets/%s/musicjacket/musicjacket%d_rip/"+
		"assets-star-forassetbundle-startapp-musicjacket-musicjacket%d-%s-jacket.png", region, bundle, bundle, jacket)
}

func BestdoriAudioUrl(region string, audioID int) string {
	return fmt.Sprintf("https://bestdori.com/assets/%s/sound/bgm%d_rip/bgm%d.mp3", region, audioID, audioID)
}
