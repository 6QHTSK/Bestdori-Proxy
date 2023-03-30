package models

import "Bestdori-Proxy/models/bestdori"

type WorkMethod int

const (
	FULL  = 0
	INFO  = 1
	CHART = 2
) // 工作模式

type AssetsURL struct {
	Cover string `json:"cover"`
	Audio string `json:"audio"`
}

type Post struct {
	*Info
	Chart *bestdori.V2Chart `json:"chart,omitempty"` // 谱面
}

type Info struct {
	ChartID  int    `json:"id"`                 // Bestdori的谱面ID
	Title    string `json:"title"`              // 谱面的标题
	Artists  string `json:"artists"`            // 谱面的艺术家
	Username string `json:"username"`           // 谱面作者 官谱返回craftegg
	Nickname string `json:"nickname,omitempty"` // 谱面作者昵称	可为空
	Diff     int    `json:"diff"`               // 谱面难度
	Rating   int    `json:"rating"`             // 谱面等级
	MusicUrl string `json:"music_url"`          // 谱面资源
	CoverUrl string `json:"cover_url"`          // 封面资源
	Time     int64  `json:"time,omitempty"`     // ms时间戳
	Content  string `json:"content"`            // 内容
}
