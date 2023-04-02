package bestdori

type ListRequestBody struct {
	Following    bool   `json:"following"`
	CategoryName string `json:"categoryName,omitempty"`
	CategoryId   string `json:"categoryId,p,omitempty"`
	Order        string `json:"order"`
	Limit        uint64 `json:"limit"`
	Offset       uint64 `json:"offset"`
	Username     string `json:"username,omitempty"`
}

type ListResponseBody struct {
	Result bool   `json:"result"`
	Count  uint64 `json:"count"`
	Posts  []struct {
		CategoryName string `json:"categoryName,omitempty"`
		CategoryId   string `json:"categoryId,omitempty"`
		Id           int    `json:"id"`
	} `json:"posts"`
}
