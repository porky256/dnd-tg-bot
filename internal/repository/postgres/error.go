package postgres

import (
	"github.com/lib/pq"
	"github.com/porky256/dnd-tg-bot/internal/database"
	"log"
)

func ParsePqError(err error) error {
	if pqerr, ok := err.(*pq.Error); ok {
		log.Println("errorcode: ", pqerr.Code)
		switch pqerr.Code {
		case "23505":
			return database.ErrViolatesUnique
		default:
			return err
		}
	} else {
		return err
	}
}
