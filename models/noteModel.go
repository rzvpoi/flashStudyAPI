package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Title   string `gorm:"size:255;not null;" json:"title"`
	Text    string `gorm:"size:max;" json:"text"`
	GroupId int
	UserId  int
	Group   Group `gorm:"foreignKey:GroupId"`
	User    User  `gorm:"foreignKey:UserId"`
}

func DeleteNote(id uint) error {
	var n Note
	n.ID = id

	err := DB.First(&n).Error
	if err != nil {
		return err
	}

	err = DB.Unscoped().Delete(&n).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateNote(in *Note, id uint) error {
	var n Note

	if err := DB.First(&n, id).Error; err != nil {
		return errors.New("Note not found")
	}

	n.Text = in.Text
	n.Title = in.Title

	DB.Save(&n)
	return nil
}

func CreateNote(n *Note) error {
	var err error

	var query Group
	if err := DB.First(&query, n.GroupId).Error; err != nil {
		return errors.New("GroupId not found")
	}

	err = DB.Create(&n).Error
	if err != nil {
		return err
	}

	return nil
}

func GetNote(uid uint, gid uint) ([]Note, error) {
	var n []Note

	var query Group
	if err := DB.First(&query, gid).Error; err != nil {
		return []Note{}, errors.New("GroupId not found")
	}

	err := DB.Model(Note{}).Where("user_id = ? && group_id = ?", uid, gid).Find(&n).Error
	if err != nil {
		fmt.Print(err)
		return n, err

	}

	//hide user id
	for i := 0; i < len(n); i++ {
		n[i].UserId = -1
	}

	return n, nil
}
