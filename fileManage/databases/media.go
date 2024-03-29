package databases

import (
	"context"
	"fileManage/model"
)

func CreateMedia(media *model.Media) (err error) {
	err = Db.Debug().Unscoped().Where(model.Media{Md5: media.Md5}).FirstOrCreate(&media).Error
	return err
}

func UpdateMediaByMedia(media *model.Media) (err error) {
	err = Db.Debug().Unscoped().Table(media.TableName()).Where(model.Media{ID: media.ID}).Update(media).Error
	return err
}

func GetMediaByMd5(media *model.Media) (err error) {
	err = Db.Debug().Unscoped().Where(model.Media{Md5: media.Md5}).First(media).Error
	if IsErrorRecordNotFound(err) {
		return nil
	}
	return err
}

func GetMediaByMd5s(md5 string) (media *model.Media, err error) {
	err = Db.Debug().Unscoped().Where(model.Media{Md5: md5}).First(media).Error
	if IsErrorRecordNotFound(err) {
		return nil, nil
	}
	return media, err
}

func MediaGetByOffset(ctx context.Context, offset, count int) (media []model.Media, page int, err error) {
	media = make([]model.Media, 0)
	err = Db.Debug().Table(model.MediaTableName).Offset(offset).Limit(count).Order("id DESC").Find(&media).Error
	if IsErrorRecordNotFound(err) {
		return media, count, nil
	}
	return media, count, nil
}
func GetMediaById(id uint) (media *model.Media, err error) {
	media = &model.Media{}
	err = Db.Debug().Unscoped().Where("id = ?", id).First(&media).Error
	if IsErrorRecordNotFound(err) {
		return nil, nil
	}
	return media, err
}

type MediaCount struct {
	Total int64 `json:"total"`
}

func NewMediaCount() *MediaCount {
	return &MediaCount{}
}

//
//func GetMediaTotal(ctx context.Context) (boxCount *BoxCount, err error) {
//	boxCount = NewBoxCount()
//	err = HermesServer.Debug().Table(BoxTableName).Select("count(*) as total").Find(&boxCount).Error
//	if err != nil && !IsErrorRecordNotFound(err) {
//		return nil, err
//	}
//	return boxCount, nil
//}
