package model

const MediaTableName = "media"

type Media struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Md5  string `json:"md_5"`
	Type string `json:"type"`
	Path string `json:"path"`
	Model
}

func NewMedia() (media *Media) {
	return &Media{}
}

func (m *Media) TableName() string {
	return MediaTableName
}
