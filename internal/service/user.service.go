package service

import "example.com/ecommerce/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user repo
func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
