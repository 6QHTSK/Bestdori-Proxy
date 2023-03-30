package config

import (
	"fmt"
)

var BandAllJsonUrl = "https://bestdori.com/api/bands/all.1.json"
var LLSifAllJsonUrl = "https://bestdori.com/api/misc/llsif.10.json"
var LLSifAssetPrefix = "https://card.niconi.co.ni/asset/"

func OfficialPostUrl(chartID int) string {
	return fmt.Sprintf("https://bestdori.com/api/songs/%d.json", chartID)
}

func OfficialChartUrl(chartID int, diff int) string {
	str := "expert"
	if diff >= 0 && diff <= 4 {
		str = []string{"easy", "normal", "hard", "expert", "special"}[diff]
	}
	return fmt.Sprintf("https://bestdori.com/api/charts/%d/%s.json", chartID, str)
}

func FanmadeUrl(chartID int) string {
	return fmt.Sprintf("https://bestdori.com/api/post/details?id=%d", chartID)
}

func BestdoriCoverUrl(region string, bundle int, jacket string) string {
	return fmt.Sprintf("https://bestdori.com/assets/%s/musicjacket/musicjacket%d_rip/"+
		"assets-star-forassetbundle-startapp-musicjacket-musicjacket%d-%s-jacket.png", region, bundle, bundle, jacket)
}

func BestdoriBGMUrl(region string, musicID int) string {
	return fmt.Sprintf("https://bestdori.com/assets/%s/sound/bgm%d_rip/bgm%d.mp3", region, musicID, musicID)
}
