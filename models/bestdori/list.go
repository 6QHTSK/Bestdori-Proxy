package bestdori

type ListRequestBody struct {
	Following    bool   `json:"following"`
	CategoryName string `json:"categoryName,omitempty"`
	CategoryId   string `json:"categoryId,p,omitempty"`
	Order        string `json:"order"`
	Limit        uint   `json:"limit"`
	Offset       uint   `json:"offset"`
	Username     string `json:"username,omitempty"`
}

type ListReplyBody struct {
	Result bool `json:"result"`
	Count  uint `json:"count"`
	Posts  []struct {
		CategoryName string `json:"categoryName,omitempty"`
		CategoryId   string `json:"categoryId,omitempty"`
		Id           int    `json:"id"`
	} `json:"posts"`
}
