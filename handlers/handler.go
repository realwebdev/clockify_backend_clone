package handlers

import "github.com/realwebdev/Bilal/clockify3/datastore"

type Handler struct {
	DB datastore.Dbhandler
}

func New(dbHandler datastore.Dbhandler) *Handler {
	return &Handler{
		DB: dbHandler,
	}
}
