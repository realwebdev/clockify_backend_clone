package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Client struct {
	Db *gorm.DB
}

func New(connStr string) (h *Dbhandler, err error) {
	db, err := GetDbConnection(connStr)
	if err != nil {
		return nil, err
	}

	return &Client{Db: db}, nil
}

func (c *Client) Close() error {
	db, err := c.Db.DB().Begin()
	if err != nil {
		fmt.Println("error creating db object")
	}

	return db.Close()
}

func GetDbConnection(dbURI string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to Database!")

	return db, nil
}
