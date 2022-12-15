package models

import (
	"errors"
	"fmt"

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

	err := DB.Model(&l).Where("user_id = ? AND group_id = ?", uid, gid).Find(&l).Error
	if err != nil {
		return err
	}

	return nil
}

func LikeGroup(uid uint, gid uint) error {
	var query Group
	if err := DB.First(&query, gid).Error; err != nil {
		return errors.New("GroupId not found")
	}

	l := Likes{}
	l.GroupId = gid
	l.UserId = uid

	g := Group{}

	err := DB.Model(&l).Where("user_id = ? AND group_id = ?", uid, gid).Find(&l).Error
	fmt.Println(l)
	if err != nil {
		// group is not liked by the user
		err = DB.Create(&l).Error
		if err != nil {
			return err
		}

		err = DB.Model(Group{}).Where("id = ?", gid).Find(&g).Error
		if err != nil {
			return err
		}

		err = DB.Model(Group{}).Where("id = ?", gid).Update("likes", g.Likes+1).Error
		if err != nil {
			return err
		}

	} else {
		//group is liked by user
		err = DB.Unscoped().Delete(&l).Error
		if err != nil {
			return err
		}

		err = DB.Model(Group{}).Where("id = ?", gid).Find(&g).Error
		if err != nil {
			return err
		}

		err = DB.Model(Group{}).Where("id = ?", gid).Update("likes", g.Likes-1).Error
		if err != nil {
			return err
		}
	}

	return nil
}
