package repository

import (
	"github.com/go-pg/pg"
	"github.com/nattapat27/test-golnag-echo/model"
	"log"
)

type relationRepository struct {
	db *pg.DB
}

func NewRelationRepository(dbcon *pg.DB) RelationRepositoryInf {
	return &relationRepository{
		db:dbcon,
	}
}

func (r *relationRepository) Create(relation *model.Relation) error {
	return r.db.Insert(relation)
}

func (r *relationRepository) FetchByUserId(userId int) ([]*model.Relation, error) {
	relation := []*model.Relation{}
	err := r.db.Model(&relation).Where("user_id = ?", userId).Select()
	log.Println(relation)
	if err != nil{
		return nil, err
	}
	return  relation, nil
}