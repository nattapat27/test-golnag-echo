package repository

import (
	"github.com/go-pg/pg"
	"github.com/nattapat27/test-golnag-echo/model"

	//"net/http"
)

type userRepository struct {
	db *pg.DB
}

func NewUserRepository(dbcon *pg.DB) UserRepositoryInf {
	return &userRepository{
		db: dbcon,
	}
}

func (u *userRepository) Create(user *model.User) error {
	return u.db.Insert(user)
}

func (u *userRepository) FetchOne(id int) (*model.User, error) {
	user := &model.User{}
	err := u.db.Model(user).Where("id = ?", id).First()
	if err != nil{
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Fetch() ([]*model.User, error) {
	users := []*model.User{}
	err := u.db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}