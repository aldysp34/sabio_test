package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"test.sabio/handler"
	"test.sabio/repository"
	"test.sabio/usecase"
)

type Router struct{
	productHandler *handler.ProductHandler
}
func main(){
	db, err := NewPG(context.Background())
	if err != nil{
		log.Println(err)
	}

	pr := repository.NewProductRepository(db)
	pu := usecase.NewProductUsecase(pr)
	ph := handler.NewProductHandler(pu)

	opts := Router{
		productHandler: ph,
	}
	
	r := NewRouter(opts)
	srv := http.Server{
		Addr: ":8080",
		Handler: r,
	}

	err = srv.ListenAndServe()
	if err != nil{
		log.Fatalln(err)
	}

}

func NewRouter(opts Router) *gin.Engine{
	router := gin.New()

	products := router.Group("/product")
	products.GET("", opts.productHandler.Read)
	products.GET("/:id", opts.productHandler.ReadByID)
	products.POST("", opts.productHandler.Create)
	products.PUT("/:id", opts.productHandler.Update)
	products.DELETE("/:id", opts.productHandler.Delete)

	return router
}

func NewPG(ctx context.Context) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(`postgresql://%s:%s@%s:%s/%s?sslmode=disable`,
		"postgres", "postgres", "localhost", "5432", "sabio_test")
	db, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}