package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"test.sabio/dto"
)

type ProductRepository interface{
	Create(context.Context,dto.ProductRequest) (int, error)
	Read(context.Context) ([]dto.ProductRequest, error)
	ReadByID(context.Context, int)(dto.ProductRequest, error)
	Update(context.Context, int, dto.ProductRequest) (int, error)
	Delete(context.Context, int) error
}

type productRepository struct{
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository{
	return &productRepository{
		db: db,
	}
}

func (u *productRepository) Create(ctx context.Context, req dto.ProductRequest)(int, error){
	var id int
	err := u.db.QueryRow(ctx, "INSERT INTO product (name, price, stock) VALUES ($1, $2, $3) RETURNING id", req.Name, req.Price, req.Stock).Scan(&id)
	if err != nil{
		return 0, err
	}
	return id, nil
}

func (u *productRepository) Read(ctx context.Context)([]dto.ProductRequest, error){
	rows, _ := u.db.Query(ctx, "SELECT p.id, p.name, p.price, p.stock FROM product p")
	defer rows.Close()

	products := []dto.ProductRequest{}
	for rows.Next(){
		product := dto.ProductRequest{}
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
		products = append(products, product)
	}

	return products, nil
}

func (u *productRepository) ReadByID(ctx context.Context, id int)(dto.ProductRequest, error){
	var resp dto.ProductRequest
	err := u.db.QueryRow(ctx, "SELECT p.id, p.name, p.price, p.stock FROM product p WHERE p.id = $1", id).Scan(&resp.Id, &resp.Name, &resp.Price, &resp.Stock)
	if err != nil{
		return resp, err
	}
	return resp, nil

}

func (u *productRepository) Update(ctx context.Context, id int, req dto.ProductRequest)(int, error){
	var respId int
	err := u.db.QueryRow(ctx, "UPDATE product SET name = $1, price = $2, stock = $3 WHERE id = $4", req.Name, req.Price, req.Stock, id).Scan(&respId)
	if err != nil{
		return 0, err
	}
	return respId, nil
}

func (u *productRepository) Delete(ctx context.Context, id int) error{
	_, err := u.db.Exec(ctx, "DELETE FROM product WHERE id = $1", id)
	if err != nil{
		return err
	}
	return nil
}