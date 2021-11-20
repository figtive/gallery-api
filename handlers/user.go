package handlers

import (
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/models"
)

func (m *module) UserGetOneByEmail(email string) (userInfo dtos.User, err error) {
	var user models.User
	if user, err = m.db.userOrmer.GetOneByEmail(email); err != nil {
		return
	}
	userInfo = dtos.User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
	return
}

func (m *module) UserInsert(userInfo dtos.User) (id string, err error) {
	if id, err = m.db.userOrmer.Insert(models.User{
		Name:    userInfo.Name,
		Email:   userInfo.Email,
		IsAdmin: false,
	}); err != nil {
		return
	}
	return
}

func (m *module) UserUpdate(userInfo dtos.User) (err error) {
	if err = m.db.userOrmer.Update(models.User{
		Email:   userInfo.Email,
		IsAdmin: userInfo.IsAdmin,
	}); err != nil {
		return
	}
	return
}
