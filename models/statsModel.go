package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Stats struct {
	gorm.Model
	Grade         float64 `gorm:"size:11;not null;" json:"grade"`
	CorrectAnswer int     `gorm:"size:11;" json:"correctAnswer"`
	WrongAnswer   int     `gorm:"size:11;not null;" json:"wrongAnswer"`
	GroupId       int
	UserId        int
	Group         Group `gorm:"foreignKey:GroupId"`
	User          User  `gorm:"foreignKey:UserId"`
}

func GetStats(uid uint, gid uint) ([]Stats, error) {
	var n []Stats

	var query Group
	if err := DB.First(&query, gid).Error; err != nil {
		return []Stats{}, errors.New("GroupId not found")
	}

	err := DB.Model(Stats{}).Where("user_id = ? AND group_id = ?", uid, gid).Find(&n).Error
	if err != nil {
		return n, err
	}

	//hide user id
	for i := 0; i < len(n); i++ {
		n[i].UserId = -1
	}

	return n, nil
}

func CreateStats(s *Stats) error {
	DB.AutoMigrate(Stats{})
	var err error

	var query Group
	if err := DB.First(&query, s.GroupId).Error; err != nil {
		return errors.New("GroupId not found")
	}

	err = DB.Create(&s).Error
	if err != nil {
		return err
	}

	return nil
}
