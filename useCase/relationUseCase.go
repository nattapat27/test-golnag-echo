package useCase

import (
	"github.com/nattapat27/test-golnag-echo/model"
	"github.com/nattapat27/test-golnag-echo/repository"
)

type relationUseCase struct {
	relationRepo repository.RelationRepositoryInf
}

func NewRelationUseCase(relationRepo repository.RelationRepositoryInf) RelationUseCaseInf{
	return &relationUseCase{
		relationRepo:relationRepo,
	}
}

func (r *relationUseCase) Create(relation *model.Relation) error {
	return r.relationRepo.Create(relation)
}

func (r *relationUseCase) FetchByUserId(useId int) ([]*model.Relation, error) {
	return r.relationRepo.FetchByUserId(useId)
}