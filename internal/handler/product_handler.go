package handler

import (
	"encoding/json"
	"fmt"
	models "goshop/internal/model"
	"goshop/internal/repository"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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

func GenerateRandomString(length int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
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

	formImage, _ := c.MultipartForm()
	imageArray := []string{}

	for key, files := range formImage.File {
		if strings.HasPrefix(key, "image_product[") && strings.HasSuffix(key, "]") {
			for _, file := range files {
				src, err := file.Open()
				if err != nil {
					return c.JSON(400, models.ErrorJSON{
						Status:  400,
						Error:   true,
						Message: "Failed to open file: " + err.Error(),
					})
				}
				defer src.Close()
				year, month, day := time.Now().Date()
				fileNameByDate := fmt.Sprintf("product_%d%d%d_%s.%s", year, month, day, GenerateRandomString(5), strings.Split(file.Filename, ".")[1])
				dst, err := os.Create("storage/products/" + fileNameByDate)
				if err != nil {
					return c.JSON(400, models.ErrorJSON{
						Status:  400,
						Error:   true,
						Message: "Failed to create file: " + err.Error(),
					})
				}
				defer dst.Close()

				if _, err = io.Copy(dst, src); err != nil {
					return c.JSON(400, models.ErrorJSON{
						Status:  400,
						Error:   true,
						Message: "Failed to copy file: " + err.Error(),
					})
				}

				imageArray = append(imageArray, fileNameByDate)
			}
		}
	}
	jsonData, _ := json.Marshal(imageArray)
	form.Images = string(jsonData)

	productModel := models.Product{
		ProductID: 0,
		Name:      form.Name,
		Category:  form.Category,
		Price:     form.Price,
		Images:    form.Images,
		Quantity:  form.Quantity,
	}

	err := r.Repo.CreateProduct(productModel)
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

func (r *ProductRepository) UpdateDataProduct(c echo.Context) error {
	var form models.ProductForm
	if err := c.Bind(&form); err != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Failed to bind data product",
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	productModel := models.Product{
		Name:     form.Name,
		Category: form.Category,
		Price:    form.Price,
		Images:   form.Images,
		Quantity: form.Quantity,
	}
	err := r.Repo.UpdateProduct(productModel, id)
	if err != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Failed to update data product",
		})
	}
	return c.JSON(200, models.SuccessJSON{
		Status:  200,
		Error:   false,
		Message: "Success update data product",
	})
}
