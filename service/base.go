package service

import (
	"baileys/interfaces"
	"baileys/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type BaseServiceInterface[T interfaces.Entity] interface {
	GetById(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, payload T) (T, error)
	Update(ctx context.Context, id string, payload T) (T, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, payload T) ([]T, error)
}

type baseService[T interfaces.Entity] struct {
	baseRepository repository.BaseRepository[T]
}

func NewBaseService[T interfaces.Entity]() BaseServiceInterface[T] {
	bc := new(baseService[T])
	bc.baseRepository = repository.NewBaseRepository[T]()

	return bc
}

func (bc *baseService[T]) GetById(ctx context.Context, id string) (res T, err error) {
	return bc.baseRepository.GetById(ctx, id)
}

func (bc *baseService[T]) Create(ctx context.Context, payload T) (res T, err error) {
	err = payload.SetCol("Id", uuid.NewString())
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	err = bc.baseRepository.Create(ctx, &payload)

	return res, err
}

func (bc *baseService[T]) Update(ctx context.Context, id string, payload T) (res T, err error) {
	err = bc.baseRepository.Update(ctx, id, &payload)

	return payload, err
}

func (bc *baseService[T]) Delete(ctx context.Context, id string) (err error) {
	err = bc.baseRepository.Delete(ctx, id)

	return err
}

func (bc *baseService[T]) Get(ctx context.Context, payload T) (res []T, err error) {
	fieldMap, err := payload.GetStructData()
	if err != nil {
		return res, err
	}

	return bc.baseRepository.Get(ctx, fieldMap)
}
