package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
	Likes       int    `gorm:"size:11;" json:"likes"`
	IsPublic    bool   `gorm:"not null;" json:"isPublic"`
	Color       string `gorm:"size:255;" json:"color"`
	UserId      int
	User        User `gorm:"foreignKey:UserId"`
}

func DeleteGroup(uid uint, gid uint) (string, error) {
	var g Group
	g.ID = gid

	err := DB.Unscoped().Delete(&g).Error

	if err != nil {
		return "", err
	}

	return "Group Deleted", nil

}

func UpdateGroup(uid uint, in *Group) (string, error) {
	var g Group

	if err := DB.First(&g, uid).Error; err != nil {
		return "", errors.New("Group not found")
	}

	g.Name = in.Name
	g.Description = in.Description
	g.IsPublic = in.IsPublic
	g.Color = in.Color

	DB.Save(&g)
	return "Group Updated!", nil
}

func SaveGroup(g *Group) (*Group, error) {
	var err error
	err = DB.Create(&g).Error
	if err != nil {
		return &Group{}, err
	}

	return g, nil
}

func GetGroups(uid uint) ([]Group, error) {

	var g []Group

	err := DB.Model(Group{}).Where("user_id = ?", uid).Find(&g).Error
	if err != nil {
		fmt.Print(err)
		return g, err

	}

	//hide user id
	for i := 0; i < len(g); i++ {
		g[i].UserId = -1
	}

	return g, nil
}
