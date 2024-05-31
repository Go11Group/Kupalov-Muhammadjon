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
func (u *UserRepo) GetById(id uint) (*model.User, error) {
	user := model.User{}
	tx := u.db.First(&user, id)
	return &user, tx.Error
}
func (u *UserRepo) GetByFirstName(firstName string) (*[]model.User, error) {
	users := []model.User{}
	u.db.Where("first_name = ?", firstName).Find(&users)
	return &users, nil
}
func (u *UserRepo) GetByLastName(lastName string) (*[]model.User, error) {
	users := []model.User{}
	u.db.Where("last_name = ?", lastName).Find(&users)
	return &users, nil
}
func (u *UserRepo) GetByGender(gender string) (*[]model.User, error) {
	users := []model.User{}
	u.db.Where("gender = ?", gender).Find(&users)
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