package utilitys

import (
	"net/mail"
	"strconv"

	"github.com/matthewhartstonge/argon2"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashPassword(password string) string {
	argon := argon2.DefaultConfig()
	encoded, _ := argon.HashEncoded([]byte(password))
	return string(encoded)
}

func StingToInt64(str string) int64 {
	result, _ := strconv.ParseInt(str, 10, 64)
	return result
}

func Paginate(pindex int64, plimit int64, totalrows int64) Pagination {
	if pindex < 1 {
		pindex = 1
	}
	pagination := Pagination{
		PageSkip:   pindex*plimit - plimit,
		TotalRows:  totalrows,
		PagesTotal: totalrows / plimit,
		PageLimit:  plimit,
		PageIndex:  pindex,
	}
	return pagination
}
