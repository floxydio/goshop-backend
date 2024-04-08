package handler

import (
	models "goshop/internal/model"
	"goshop/internal/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Repo repository.ProductRepo
}

func ProductControllerInit(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Repo: repository.ProductRepoIntf(db),
	}
}

func (r *ProductRepository) FindDataProduct(c echo.Context) error {
	// Query Param
	name := c.QueryParam("name")
	category := c.QueryParam("category")
	products, err := r.Repo.FindProduct(name, category)
	if err != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Failed to get data product",
		})
	}
	return c.JSON(200, models.SuccessJSON{
		Status:  200,
		Error:   false,
		Data:    products,
		Message: "Success get data product",
	})
}

func (r *ProductRepository) CreateDataProduct(c echo.Context) error {
	var form models.ProductForm
	if err := c.Bind(&form); err != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Failed to bind data product",
		})
	}
	err := r.Repo.CreateProduct(form)
	if err != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Failed to create data product",
		})
	}
	return c.JSON(201, models.SuccessJSON{
		Status:  201,
		Error:   false,
		Message: "Success create data product",
	})
}
