package action

import (
	"github.com/geovani-moc/geac/entity"
)

//SendEmails this function sen an email using the smtp
func SendEmails(emails []entity.Email, user *entity.User) {
	for i := 0; i < len(emails); i++ {
		emails[i].Send()
	}
}
