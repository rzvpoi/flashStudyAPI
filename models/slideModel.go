package models

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Slide struct {
	gorm.Model
	Question string `gorm:"size:255;not null;" json:"question"`
	Answer   string `gorm:"size:255;" json:"answer"`
	Image    string `gorm:"size:255;" json:"image"`
	Tags     string `gorm:"size:255;" json:"tags"`
	GroupId  int
	Group    Group `gorm:"foreignKey:GroupId"`
}

func UpdateSlide(uid uint, in *Slide) (string, error) {
	var s Slide
	if err := DB.First(&s, uid).Error; err != nil {
		return "", errors.New("Slide not found")
	}

	s.Question = in.Question
	s.Answer = in.Answer
	s.Tags = in.Tags

	var newFileName string = ""
	if len(in.Image) > 0 {
		allowed := []string{".jpg", ".png", ".jpeg"}
		extension := filepath.Ext(in.Image)
		var isValid bool = false
		for _, ext := range allowed {
			if ext == extension {
				isValid = true
			}
		}
		if !isValid {
			return "", errors.New("Extension not supported")
		}

		result := DB.Last(&s, "group_id = ?", s.GroupId)

		newFileName = "slide-image_g" + strconv.Itoa(s.GroupId) + "n" + strconv.FormatInt(int64(result.RowsAffected), 10) + ".jpg"

		s.Image = "image/" + newFileName
	} else {
		err := os.Remove("./public/images-slide/" + s.Image[5:])
		if err != nil {
			return "", err
		}
		s.Image = ""
	}

	DB.Save(&s)
	return newFileName, nil
}

func DeleteSlide(s Slide) (string, error) {

	err := DB.First(&s).Error
	if err != nil {
		return "", err
	}

	err = DB.Unscoped().Delete(&s).Error
	if err != nil {
		return "", err
	}

	err = os.Remove("./public/images-slide/" + s.Image[5:])
	if err != nil {
		return "", err
	}

	return "Group Deleted", nil
}

func CreateSlide(s *Slide) (string, error) {
	allowed := []string{".jpg", ".png", ".jpeg"}
	extension := filepath.Ext(s.Image)
	var isValid bool = false
	for _, ext := range allowed {
		if ext == extension {
			isValid = true
		}
	}
	if !isValid {
		return "", errors.New("Extension not supported")
	}

	// Verify if groupId exists in the db
	var queryG Group
	if err := DB.First(&queryG, s.GroupId).Error; err != nil {
		return "", errors.New("GroupId not found")
	}

	//Get latest id from the slide within a specific group
	//for image nameing
	var query Slide
	err := DB.Last(&query, "group_id = ?", s.GroupId).Error
	if err != nil {
		query.ID = 0
	}

	newFileName := "slide-image_g" + strconv.Itoa(s.GroupId) + "n" + strconv.Itoa(int(query.ID)+1) + ".jpg"

	fmt.Println(newFileName)
	s.Image = "image/" + newFileName

	err = DB.Create(&s).Error
	if err != nil {
		return "", err
	}

	return newFileName, nil

}

func GetSlide(gid int) ([]Slide, error) {

	var s []Slide

	var query Group
	if err := DB.First(&query, gid).Error; err != nil {
		return []Slide{}, errors.New("GroupId not found")
	}

	err := DB.Model(Group{}).Where("group_id = ?", gid).Find(&s).Error
	if err != nil {
		return s, err
	}

	return s, nil

}
