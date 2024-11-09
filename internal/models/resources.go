package models

type Resources struct {
	Data []struct {
		Name  string `json:"name"`
		Code  string `json:"code"`
		Skill string `json:"skill"`
		Level int    `json:"level"`
		Drops []struct {
			Code        string `json:"code"`
			Rate        int    `json:"rate"`
			MinQuantity int    `json:"min_quantity"`
			MaxQuantity int    `json:"max_quantity"`
		} `json:"drops"`
	} `json:"data"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
	Pages int `json:"pages"`
}
