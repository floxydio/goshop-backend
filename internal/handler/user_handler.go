package handler

import (
	models "goshop/internal/model"
	"goshop/internal/repository"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repo repository.UserRepo
}

func UserControllerInit(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repo: repository.UserRepoIntf(db),
	}
}

func (r *UserRepository) SignUp(c echo.Context) error {
	var form models.UserForm

	if err := c.Bind(&form); err != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Failed to bind data",
		})
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 10)

	formModel := models.User{
		Nama:     form.Nama,
		Username: form.Username,
		Password: form.Password,
		NoTelp:   form.NoTelp,
	}
	formModel.Password = string(bytes)

	signUpResult := r.Repo.SignUp(formModel)

	if signUpResult != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: signUpResult.Error(),
		})
	}
	return c.JSON(201, models.SuccessJSON{
		Status:  201,
		Error:   false,
		Message: "Success sign up",
	})
}

func (r *UserRepository) SignIn(c echo.Context) error {
	var form models.UserForm

	if err := c.Bind(&form); err != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Failed to bind data",
		})
	}

	formModel := models.User{
		Username: form.Username,
		Password: form.Password,
	}
	user, err := r.Repo.FindUser(formModel)

	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(400, models.ErrorJSON{
				Status:  400,
				Error:   true,
				Message: "User not found",
			})

		}
	}
	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if errHash != nil {
		return c.JSON(400, models.ErrorJSON{
			Status:  400,
			Error:   true,
			Message: "Password wrong!",
		})
	}
	return c.JSON(200, models.SuccessJSON{
		Status:  200,
		Error:   false,
		Message: "Success sign in",
	})

}
