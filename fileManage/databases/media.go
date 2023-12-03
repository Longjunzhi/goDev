package databases

import (
	"Img/model"
)

func CreateMedia(media *model.Media) (err error) {
	err = Db.Debug().Unscoped().Where(model.Media{Md5: media.Md5}).FirstOrCreate(&media).Error
	return err
}
