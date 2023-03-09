package services

// import (
// 	"dashboardq-be/features/class"
// 	"dashboardq-be/helper"
// 	"dashboardq-be/mocks"
// 	"errors"
// 	"testing"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestCreate(t *testing.T) {
// 	repo := mocks.NewClassData(t)
// 	inputData := class.Core{
// 		Name:       "Back End Batch 15",
// 		StartClass: "2022-09-14",
// 		EndClass:   "2022-11-10",
// 		IdUser:     1,
// 		User:       users.Core{
// 			Name: "adul",
// 		},
// 	}
// 	resData := class.Core{
// 		ID:         uint(1),
// 		Name:       "Back End Batch 15",
// 		StartClass: "2022-09-14",
// 		EndClass:   "2022-11-10",
// 		IdUser:     1,
// 		UserName:   "Adul Ganteng",
// 	}

// 	t.Run("success creating class", func(t *testing.T) {
// 		repo.On("Create", uint(1), mock.Anything).Return(resData, nil).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Create(pToken, inputData)
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData.ID, res.ID)
// 		assert.Equal(t, resData.UserName, res.UserName)
// 		repo.AssertExpectations(t)
// 	})
// 	// t.Run("Duplicated", func(t *testing.T) {
// 	// 	repo.On("Create", uint(1), mock.Anything).Return(class.Core{}, errors.New("name duplicated")).Once()
// 	// 	srv := New(repo)
// 	// 	_, token := helper.GenerateToken(1)
// 	// 	pToken := token.(*jwt.Token)
// 	// 	pToken.Valid = true
// 	// 	res, err := srv.Create(pToken, inputData)
// 	// 	assert.NotNil(t, err)
// 	// 	assert.Equal(t, uint(0), res.ID)
// 	// 	assert.ErrorContains(t, err, "name already Createed")
// 	// 	repo.AssertExpectations(t)
// 	// })

// 	t.Run("access denied", func(t *testing.T) {
// 		repo.On("Create", uint(1), mock.Anything).Return(class.Core{}, errors.New("access denied")).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Create(pToken, inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "access denied")
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("internal server error", func(t *testing.T) {
// 		repo.On("Create", uint(1), mock.Anything).Return(class.Core{}, errors.New("server error")).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Create(pToken, inputData)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, uint(0), res.ID)
// 		assert.ErrorContains(t, err, "server error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestShow(t *testing.T) {
// 	repo := mocks.NewClassData(t)
// 	resData := class.Core{
// 		ID:         uint(1),
// 		Name:       "Back End Batch 16",
// 		StartClass: "2023-09-14",
// 		EndClass:   "2023-11-10",
// 		IdUser:     1,
// 		UserName:   "Cici Cantik",
// 	}
// 	t.Run("success show class", func(t *testing.T) {
// 		repo.On("Show", uint(1)).Return(resData, nil).Once()
// 		srv := New(repo)
// 		res, err := srv.Show(uint(1))
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData, res)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("class not found", func(t *testing.T) {
// 		repo.On("Show", uint(1)).Return(class.Core{}, errors.New("query error, problem with server")).Once()
// 		srv := New(repo)
// 		res, err := srv.Show(uint(1))
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "server")
// 		assert.Equal(t, class.Core{}, res)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("cannot access", func(t *testing.T) {
// 		repo.On("Show", uint(1)).Return(class.Core{}, errors.New("access denied, cannot access")).Once()
// 		srv := New(repo)
// 		res, err := srv.Show(uint(1))
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "denied")
// 		assert.Equal(t, class.Core{}, res)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	repo := mocks.NewClassData(t)
// 	inputData := class.Core{
// 		ID:         1,
// 		Name:       "Back End Batch 16",
// 		StartClass: "2023-09-14",
// 		EndClass:   "2023-11-10",
// 		IdUser:     1,
// 		UserName:   "Cici Cantik",
// 	}
// 	resData := class.Core{
// 		ID:         1,
// 		Name:       "Back End Batch 16",
// 		StartClass: "2023-09-14",
// 		EndClass:   "2023-11-10",
// 		IdUser:     1,
// 		UserName:   "Cici Cantik",
// 	}

// 	t.Run("success updating class", func(t *testing.T) {
// 		repo.On("Update", uint(1), uint(1), inputData).Return(resData, nil).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Update(pToken, uint(1), inputData)
// 		assert.Nil(t, err)
// 		assert.Equal(t, resData.ID, res.ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("fail updating class", func(t *testing.T) {
// 		repo.On("Update", uint(1), uint(1), inputData).Return(class.Core{}, errors.New("class not found")).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Update(pToken, uint(1), inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "not Createed")
// 		assert.Equal(t, class.Core{}, res)
// 		repo.AssertExpectations(t)
// 	})
// 	// t.Run("name duplicated", func(t *testing.T) {
// 	// 	repo.On("Update", uint(2), mock.Anything).Return(class.Core{}, errors.New("name duplicated")).Once()
// 	// 	srv := New(repo)
// 	// 	_, token := helper.GenerateToken(2)
// 	// 	pToken := token.(*jwt.Token)
// 	// 	pToken.Valid = true
// 	// 	res, err := srv.Update(pToken, uint(2), inputData)
// 	// 	assert.NotNil(t, err)
// 	// 	assert.ErrorContains(t, err, "name duplicated")
// 	// 	assert.Equal(t, class.Core{}, res)
// 	// 	repo.AssertExpectations(t)
// 	// })
// 	t.Run("access denied", func(t *testing.T) {
// 		repo.On("Update", uint(1), uint(1), inputData).Return(class.Core{}, errors.New("access denied")).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		res, err := srv.Update(pToken, uint(1), inputData)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "access denied")
// 		assert.Equal(t, class.Core{}, res)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	repo := mocks.NewClassData(t)
// 	t.Run("delete class successful", func(t *testing.T) {
// 		repo.On("Delete", uint(1)).Return(nil).Once()
// 		srv := New(repo)
// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true

// 		err := srv.Delete(pToken, uint(1))
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// 	// internal server error, class fail to delete
// 	t.Run("internal server error, class fail to delete", func(t *testing.T) {
// 		repo.On("Deactive", uint(1)).Return(errors.New("no class has delete")).Once()
// 		srv := New(repo)

// 		_, token := helper.GenerateToken(1)
// 		pToken := token.(*jwt.Token)
// 		pToken.Valid = true
// 		err := srv.Delete(pToken, uint(1))
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "error")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestShowAll(t *testing.T) {
// 	repo := mocks.NewClassData(t)
// 	resData := []class.Core{
// 		{
// 			ID:         1,
// 			Name:       "Back End Batch 16",
// 			StartClass: "2023-09-14",
// 			EndClass:   "2023-11-10",
// 			IdUser:     1,
// 			UserName:   "Cici Cantik",
// 		},
// 	}
// 	t.Run("show all class successful", func(t *testing.T) {
// 		repo.On("ShowAll").Return(resData, nil).Once()
// 		srv := New(repo)
// 		res, err := srv.ShowAll()
// 		assert.Equal(t, res[0].ID, resData[0].ID)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// 	// internal server error, class fail to ShowAll
// 	t.Run("internal server error, class fail to ShowAll", func(t *testing.T) {
// 		repo.On("ShowAll").Return([]class.Core{}, errors.New("data not found")).Once()
// 		srv := New(repo)
// 		res, err := srv.ShowAll()
// 		assert.Equal(t, res, []class.Core{})
// 		assert.ErrorContains(t, err, "data not found")
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// }
