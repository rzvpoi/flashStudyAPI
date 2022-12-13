package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Exam struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;" json:"name"`
	Description string `gorm:"size:max;" json:"description"`
	ExamDate    string `gorm:"size:255;" json:"examDate"`
	UserId      int
	User        User `gorm:"foreignKey:UserId"`
}

func DeleteExam(id uint) error {
	var n Exam
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

func UpdateExam(in *Exam, id uint) error {
	var n Exam

	if err := DB.First(&n, id).Error; err != nil {
		return errors.New("Exam not found")
	}

	n.Name = in.Name
	n.Description = in.Description
	n.ExamDate = in.ExamDate

	DB.Save(&n)
	return nil
}

func CreateExam(n *Exam) error {
	var err error

	err = DB.Create(&n).Error
	if err != nil {
		return err
	}

	return nil
}

func GetExam(uid uint) ([]Exam, error) {
	var e []Exam
	DB.AutoMigrate(Exam{})
	err := DB.Model(Exam{}).Where("user_id = ?", uid).Find(&e).Error
	if err != nil {
		return e, err

	}

	//hide user id
	for i := 0; i < len(e); i++ {
		e[i].UserId = -1
	}

	return e, nil
}
