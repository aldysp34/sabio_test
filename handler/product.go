package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"test.sabio/dto"
	"test.sabio/usecase"
)

type ProductHandler struct{
	productUsecase usecase.ProductUsecase
}



func NewProductHandler(pu usecase.ProductUsecase) *ProductHandler{
	return &ProductHandler{productUsecase: pu}
}

func (p *ProductHandler) Create(ctx *gin.Context){
	newRequest := dto.ProductRequest{}
	err := ctx.ShouldBindJSON(&newRequest)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	resp, err := p.productUsecase.Create(ctx, newRequest)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{Data: resp})
}

func (p *ProductHandler) ReadByID(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	resp, err := p.productUsecase.ReadByID(ctx, id)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: resp})
}

func (p *ProductHandler) Read(ctx *gin.Context){
	resp, err := p.productUsecase.Read(ctx)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: resp})
}

func (p *ProductHandler) Update (ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	newRequest := dto.ProductRequest{}
	err = ctx.ShouldBindJSON(&newRequest)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	resp, err := p.productUsecase.Update(ctx, id, newRequest)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: resp})
}

func (p *ProductHandler) Delete(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = p.productUsecase.Delete(ctx, id)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}