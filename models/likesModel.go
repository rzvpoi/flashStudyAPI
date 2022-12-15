package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Likes struct {
	gorm.Model
	GroupId uint
	UserId  uint
	Group   Group `gorm:"foreignKey:GroupId"`
	User    User  `gorm:"foreignKey:UserId"`
}

func IsLiked(uid uint, gid uint) error {
	l := Likes{}
	l.GroupId = gid
	l.UserId = uid

	err := DB.Model(&l).Where("user_id = ? && group_id = ?", uid, gid).Find(&l).Error
	if err != nil {
		return err
	}

	return nil
}

func LikeGroup(uid uint, gid uint) error {
	DB.AutoMigrate(Likes{})
	var query Group
	if err := DB.First(&query, gid).Error; err != nil {
		return errors.New("GroupId not found")
	}

	l := Likes{}
	l.GroupId = gid
	l.UserId = uid

	err := DB.Model(&l).Where("user_id = ? && group_id = ?", uid, gid).Find(&l).Error
	if err != nil {
		// group is not liked by the user
		err = DB.Create(&l).Error
		if err != nil {
			return err
		}
	} else {
		//group is liked by user
		err = DB.Unscoped().Delete(&l).Error
		if err != nil {
			return err
		}
	}

	return nil
}
