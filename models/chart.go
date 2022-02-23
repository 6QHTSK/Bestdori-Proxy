package models

type DiffType int

const (
	Diff_Easy DiffType = iota
	Diff_Normal
	Diff_Hard
	Diff_Expert
	Diff_Special
)

type Author struct {
	Username string `json:"username"` // 谱面作者用户名
	Nickname string `json:"nickname"` // 谱面作者昵称
}

type AssetsURL struct {
	Cover string `json:"cover"`
	Audio string `json:"audio"`
}

type ProcessedChartItem struct {
	ChartID  int         `json:"id"`              // Bestdori的谱面ID
	Title    string      `json:"title"`           // 谱面的标题
	Artists  string      `json:"artists"`         // 谱面的艺术家
	Author   Author      `json:"author"`          // 谱面作者
	Diff     DiffType    `json:"diff"`            // 谱面大难度
	Level    int         `json:"level"`           // 谱面小难度
	SongUrl  AssetsURL   `json:"song_url"`        // 谱面资源
	Official bool        `json:"official"`        //是否为官谱
	Likes    int         `json:"likes,omitempty"` // 喜爱数
	Chart    interface{} `json:"chart,omitempty"` // 谱面
	Time     int64       `json:"time,omitempty"`  // ms时间戳
	Content  string      `json:"content"`         // 内容
}

type SonolusChartItem struct {
	Description string `json:"description"`
	Item        struct {
		Name    string `json:"name"`
		Rating  int    `json:"rating"`
		Title   string `json:"title"`
		Artists string `json:"artists"`
		Cover   struct {
			Url string `json:"url"`
		} `json:"cover"`
		Bgm struct {
			Url string `json:"url"`
		} `json:"bgm"`
	} `json:"item"`
}

type BestdoriChartItem struct {
	Result bool `json:"result"`
	Post   struct {
		CategoryName string `json:"categoryName"`
		CategoryId   string `json:"categoryId"`
		Diff         int    `json:"diff"`
		Time         int64  `json:"time"`
		Author       struct {
			Username string `json:"username"`
			Nickname string `json:"nickname"`
		} `json:"author"`
		Likes int             `json:"likes"`
		Chart BestdoriV2Chart `json:"chart"`
	} `json:"post"`
}
type BestdoriMapItem struct {
	Result bool `json:"result"`
	Post   struct {
		Chart BestdoriV2Chart `json:"chart"`
	}
}

type ListQueryResultItem struct {
	ChartID int    `json:"id"`
	Author  Author `json:"author"`
	Diff    int    `json:"diff"`
	Level   int    `json:"level"`
	Likes   int    `json:"likes"`
}
