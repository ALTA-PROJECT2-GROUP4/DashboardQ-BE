package services

import (
	"dashboardq-be/features/users"
	"dashboardq-be/helper"
	"dashboardq-be/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegis(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputData := users.Core{Name: "Alfian", Email: "alfian@gmail.com", Phone: "081234"}
	resData := users.Core{ID: uint(1), Name: "Alfian", Email: "alfian@gmail.com", Phone: "081234"}

	t.Run("success creating account", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Phone, res.Phone)
		repo.AssertExpectations(t)
	})
	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(users.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "email already registered")
		repo.AssertExpectations(t)
	})

	t.Run("access denied", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(users.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		repo.On("Register", uint(1), mock.Anything).Return(users.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Register(pToken, inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server error")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputEmail := "alfian@gmail.com"
	passwordHashed := helper.GeneratePassword("123")
	resData := users.Core{ID: uint(1), Name: "Adi Yuda", Email: "alfian@gmail.com", Password: passwordHashed}
	t.Run("login success", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.Name, res.Name)
		repo.AssertExpectations(t)
	})
	t.Run("account not found", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(users.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "123")
		assert.NotNil(t, token)
		assert.ErrorContains(t, err, "not")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "komuk@example.com"
		repo.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(repo)
		_, res, err := srv.Login(inputEmail, "342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password")
		assert.Empty(t, nil)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "komuk@example.com"
		repo.On("Login", inputEmail).Return(users.Core{}, errors.New("nip or password not allowed empty")).Once()
		srv := New(repo)
		_, res, err := srv.Login(inputEmail, "")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "empty")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := users.Core{ID: uint(1), Name: "adiyuda", Email: "adiyuda@example.com", Phone: "08123"}
	t.Run("success show profile", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken, uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData, res)
		repo.AssertExpectations(t)
	})
	t.Run("account not found", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(users.Core{}, errors.New("query error, problem with server")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("cannot access", func(t *testing.T) {
		repo.On("Profile", uint(1)).Return(users.Core{}, errors.New("access denied, cannot access")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "denied")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputData := users.Core{ID: 1, Name: "Adi Yuda", Phone: "08123"}
	resData := users.Core{ID: 1, Name: "Adi Yuda", Phone: "08123"}

	t.Run("success updating account", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("fail updating account", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(users.Core{}, errors.New("users not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not registered")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("email duplicated", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(users.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email duplicated")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("Update", uint(1), mock.Anything).Return(users.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateAdm(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputData := users.Core{ID: 1, Name: "Adi Yuda", Phone: "08123"}
	resData := users.Core{ID: 1, Name: "Adi Yuda", Phone: "08123"}
	t.Run("success updating account", func(t *testing.T) {
		repo.On("UpdateAdm", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateAdm(pToken, uint(1), inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("access denied", func(t *testing.T) {
		repo.On("UpdateAdm", uint(1), uint(1), inputData).Return(users.Core{}, errors.New("access denied")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateAdm(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})
	t.Run("fail updating account", func(t *testing.T) {
		repo.On("UpdateAdm", uint(1), uint(1), inputData).Return(users.Core{}, errors.New("users not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateAdm(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "account not registered")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("email duplicated", func(t *testing.T) {
		repo.On("UpdateAdm", uint(1), uint(1), inputData).Return(users.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateAdm(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email duplicated")
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("no data updated", func(t *testing.T) {
		repo.On("UpdateAdm", uint(1), uint(1), inputData).Return(users.Core{}, errors.New("no data updated")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.UpdateAdm(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestDeactive(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("deleting account successful", func(t *testing.T) {
		repo.On("Deactive", uint(1), uint(2)).Return(nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Deactive(pToken, uint(2))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	// internal server error, account fail to delete
	t.Run("internal server error, account fail to delete", func(t *testing.T) {
		repo.On("Deactive", uint(1), uint(2)).Return(errors.New("no user has delete")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Deactive(pToken, uint(2))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		repo.AssertExpectations(t)
	})
}

func TestShowAll(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := []users.Core{
		{
			ID:    1,
			Name:  "Adi Yuda",
			Email: "adiyuda@gmail.com",
			Phone: "081234",
		},
	}
	t.Run("get all employee successful", func(t *testing.T) {
		repo.On("ShowAll").Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAll()
		assert.Equal(t, res[0].ID, resData[0].ID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	// internal server error, account fail to ShowAll
	t.Run("internal server error, account fail to ShowAll", func(t *testing.T) {
		repo.On("ShowAll").Return([]users.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		res, err := srv.ShowAll()
		assert.Equal(t, res, []users.Core{})
		assert.ErrorContains(t, err, "data not found")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
