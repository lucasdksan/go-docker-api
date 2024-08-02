package controller

import (
	"go-docker-api/internal/model"
	"go-docker-api/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{Message: "Id do produto não existe!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{Message: "Id do produto inválido!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	product, err := p.productUsecase.GetProductById(productId)

	if product == nil {
		response := model.Response{Message: "Produto não identificado!"}
		ctx.JSON(http.StatusNotFound, response)

		return
	}

	if err != nil {
		ctx.JSON(http.StatusNoContent, err)

		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := model.Response{Message: "Id do produto não existe!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{Message: "Id do produto inválido!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	if err := p.productUsecase.DeleteProduct(productId); err != nil {
		response := model.Response{Message: "Id do produto inválido!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product model.UpdateData

	if id == "" {
		response := model.Response{Message: "Id do produto não existe!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{Message: "Id do produto inválido!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := p.productUsecase.UpdateProduct(productId, product); err != nil {
		response := model.Response{Message: "Id do produto inválido!"}
		ctx.JSON(http.StatusNoContent, response)

		return
	}

	ctx.JSON(http.StatusOK, nil)
}
