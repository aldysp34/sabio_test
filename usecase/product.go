package usecase

import (
	"context"

	"test.sabio/dto"
	"test.sabio/repository"
)

type ProductUsecase interface{
	Create(context.Context,dto.ProductRequest) (int, error)
	Read(context.Context) ([]dto.ProductRequest, error)
	ReadByID(context.Context, int)(dto.ProductRequest, error)
	Update(context.Context, int, dto.ProductRequest) (int, error)
	Delete(context.Context, int) error
}

type productUsecase struct{
	productRepo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase{
	return &productUsecase{
		productRepo: repo,
	}
}

func (u *productUsecase) Create(ctx context.Context, req dto.ProductRequest)(int, error){
	return u.productRepo.Create(ctx, req)
}

func (u *productUsecase) Read(ctx context.Context)([]dto.ProductRequest, error){
	return u.productRepo.Read(ctx)
}

func (u *productUsecase) ReadByID(ctx context.Context, id int)(dto.ProductRequest, error){
	return u.productRepo.ReadByID(ctx, id)
}

func (u *productUsecase) Update(ctx context.Context, id int, req dto.ProductRequest)(int, error){
	return u.productRepo.Update(ctx, id, req)
}

func (u *productUsecase) Delete(ctx context.Context, id int)error{
	return u.productRepo.Delete(ctx, id)
}