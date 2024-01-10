package model

const MediaTableName = "media"

type Media struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Md5     string `json:"md_5"`
	Type    string `json:"type"`
	Path    string `json:"path"`
	OssPath string `json:"oss_path"`
	Model
}

func NewMedia() (media *Media) {
	return &Media{}
}

func (m *Media) TableName() string {
	return MediaTableName
}
