package useCase

import "github.com/nattapat27/test-golnag-echo/model"

type UserUseCaseInf interface {
	Create(user *model.User) error
	FetchOne(id int) (*model.User, error)
	FetchAll() ([]*model.User, error)
}

type RelationUseCaseInf interface {
	Create(relation *model.Relation) error
	FetchByUserId(useId int) ([]*model.Relation, error)
}
