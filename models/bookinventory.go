package models

import (
	// _"net/html"
	// "net/http"
	_"strings"
	_"errors"
	// "github.com/gin-gonic/gin"

	_"github.com/jinzhu/gorm"
	_"golang.org/x/crypto/bcrypt"
	_"libraryBackend/utils/token"
)

type  Bookinventory struct {
	ISBN string 	`json:"isbn"`
	Libid int `json:"libid"`
	Title string `json:"title`
	Authors string `json:"authors"`
	Publisher string `json:"publisher"`
	Version string `json:"version"`
	Totalcopies int `json:"totalcopies"`
	Avaiblecopies int `json:"avaiblecopies"`
}

func (u *Bookinventory) SaveBook() (*Bookinventory, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Bookinventory{}, err
	}
	return u, nil
}

