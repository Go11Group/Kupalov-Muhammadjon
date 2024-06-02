package postgres

import (
	"module/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo{
	return &UserRepo{db}
}

// Create
func (u *UserRepo) Create(user model.User){
	u.db.Create(&user)
}
// Read
func (u *UserRepo) GetAllUsers() (*[]model.User, error) {
	users := []model.User{}
	u.db.Find(&users)
	return &users, nil
}
func (u *UserRepo) GetByFilter(fil model.Filter) (*[]model.User, error) {
	args := []interface{}{}
	cond := ""
	if fil.Age != nil {
		args = append(args, *fil.Age)
		cond += "age=? "
	}
	if fil.ID != 0 {
		args = append(args, fil.ID)
		cond += "and id=? "
	}
	if fil.Email != nil {
		args = append(args, *fil.Email)
		cond += "and email=? "
	}
	if fil.Field != nil {
		args = append(args, *fil.Field)
		cond += "and feild=? "
	}
	if fil.FirstName != nil {
		args = append(args, *fil.FirstName)
		cond += "and first_name=? "
	}
	if fil.LastName != nil {
		args = append(args, *fil.LastName)
		cond += "and last_name=? "
	}
	if fil.Password != nil {
		args = append(args, *fil.Password)
		cond += "and password=? "
	}

	users := []model.User{}
	tx := u.db.Where(cond, args...).Find(&users)
	return &users, tx.Error
}

func (u *UserRepo) GetByFirstName(firstName string) (*[]model.User, error) {
	users := []model.User{}
	u.db.Where("first_name = ?", firstName).Find(&users)
	return &users, nil
}

// // Update
func (u *UserRepo) UpdateUser(user model.User) error {
	tx := u.db.Save(&user)
	return tx.Error
}
// // Delete
func (u *UserRepo) DeleteUser(id uint) error {
	tx := u.db.Delete(&model.User{}, id)
	return tx.Error
}