package register

import (
	"time"

	model "github.com/restuwahyu13/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	transaction := r.db.Begin()
	errorCode := make(chan string, 1)

	users := model.EntityUsers{
		Fullname:  input.Fullname,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now().Local(),
	}

	checkUserAccount := transaction.Where("email = ?", input.Email).First(&users).RowsAffected

	if checkUserAccount > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
	}

	addNewUser := transaction.Create(&users).Error
	transaction.Commit()

	if addNewUser != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}