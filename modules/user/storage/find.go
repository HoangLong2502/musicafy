package userstorage

import (
	usermodels "example.com/musicafy_be/modules/user/models"
	"gorm.io/gorm"
)

func (s *store) FindUser(arg usermodels.AccountQueries) (usermodels.User, error) {
	var user usermodels.User

	results := s.db.Table(usermodels.User{}.TableName()).Where(
		func(db *gorm.DB) *gorm.DB {
			if arg.ID != nil {
				db = db.Where("id = ?", *arg.ID)
			}
			if arg.Username != nil {
				db = db.Where("username = ?", *arg.Username)
			}
			if arg.Email != nil {
				db = db.Where("email = ?", *arg.Email)
			}
			return db
		}(s.db),
	).First(&user)

	if results.Error != nil {
		return user, results.Error
	}

	return user, nil
}
