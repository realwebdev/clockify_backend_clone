package handlers

import "github.com/realwebdev/Bilal/clockify3/datastore"

type Handler struct {
	DB datastore.DBController
}

func New(dbHandler datastore.DBController) *Handler {
	return &Handler{DB: dbHandler}
}
