package bestdori

import (
	"strings"
)

type BestdoriPost struct {
	Result bool `json:"result"`
	Post   struct {
		CategoryName string `json:"categoryName"`
		CategoryId   string `json:"categoryId"`
		Title        string `json:"title"`
		Artists      string `json:"artists"`
		Diff         int    `json:"diff"`
		Level        int    `json:"level"`
		Time         int64  `json:"time"`
		Author       struct {
			Username string `json:"username"`
			Nickname string `json:"nickname"`
		} `json:"author"`
		Song struct {
			Type  string `json:"type"`
			Audio string `json:"audio"`
			Cover string `json:"cover"`
			ID    int    `json:"id"`
		}
		Chart   V2Chart `json:"chart"`
		Content []struct {
			Data string `json:"data"`
			Type string `json:"type"`
		} `json:"content"`
	} `json:"post"`
}

func (p *BestdoriPost) IsChart() bool {
	return p.Result && p.Post.CategoryName == "SELF_POST" && p.Post.CategoryId == "chart"
}

func (p *BestdoriPost) GetTitle() string {
	return strings.ToValidUTF8(p.Post.Title, "")
}

func (p *BestdoriPost) GetArtists() string {
	return strings.ToValidUTF8(p.Post.Artists, "")
}

func (p *BestdoriPost) GetUsername() string {
	return strings.ToValidUTF8(p.Post.Author.Username, "")
}

func (p *BestdoriPost) GetNickname() string {
	return strings.ToValidUTF8(p.Post.Author.Nickname, "")
}

func (p *BestdoriPost) GetContent() (content string) {
	for _, lane := range p.Post.Content {
		if lane.Type == "text" {
			content += lane.Data + "\n"
		}
	}
	return content
}
