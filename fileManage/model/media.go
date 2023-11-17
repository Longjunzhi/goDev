package model

type Media struct {
	Name string `json:"name"`
	Size string `json:"size"`
	Md5  string `json:"md_5"`
	Type string `json:"type"`
	Model
}
