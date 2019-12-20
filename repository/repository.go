package repository

import "github.com/nattapat27/test-golnag-echo/model"

type UserRepositoryInf interface {
	Create(user *model.User) error
	FetchOne(id int) (*model.User, error)
	Fetch() ([]*model.User, error)
}

type RelationRepositoryInf interface {
	Create(relation *model.Relation) error
	FetchByUserId(userId int) ([]*model.Relation, error)
}
