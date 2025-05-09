package usecase

import (
	"errors"
	"interview/dto"
	"interview/entity"
	"interview/repository"
)

type CarUsecase interface {
	GetAllCars()([]entity.Car, error)
	CreateOneCar(req dto.CreateCarReq) (*entity.Car, error)
	DeleteOneCar(id int) (error)
	GetCarById(id int) (*entity.Car, error)
	UpdateOneCar(id int, req dto.UpdateCarReq) (*entity.Car, error)
}

type carUsecase struct {
	carRepository repository.CarRepository
}

func NewCarUsecase(carRepository repository.CarRepository) CarUsecase{
	return &carUsecase{
		carRepository: carRepository,
	}
}

func (u *carUsecase) GetAllCars()([]entity.Car, error){
	return u.carRepository.GetAll()
}

func (u *carUsecase) CreateOneCar(req dto.CreateCarReq) (*entity.Car, error){
	car := entity.Car{
		Name: req.Name,
		DayRate: req.DayRate,
		MonthRate: req.MonthRate,
		Image: req.Image,
	}

	return u.carRepository.Create(car)
}

func (u *carUsecase) DeleteOneCar(id int) (error){
	carExist, err := u.carRepository.GetCarById(id)
	if err != nil{
		return err
	}
	if carExist == nil {
		return errors.New("car not found")
	}

	return u.carRepository.Delete(id)
}

func (u *carUsecase) UpdateOneCar(id int, req dto.UpdateCarReq) (*entity.Car, error){
	carExist, err := u.carRepository.GetCarById(id)
	if err != nil{
		return nil, err
	}
	if carExist == nil {
		return nil, errors.New("car not found")
	}

	if req.Name != ""{
		carExist.Name = req.Name
	}

	if req.DayRate > 0{
		carExist.DayRate = req.DayRate
	}

	if req.MonthRate > 0{
		carExist.MonthRate = req.MonthRate
	}

	if req.Image != ""{
		carExist.Image = req.Image
	}

	return u.carRepository.Update(id, *carExist)
}

func (u *carUsecase) GetCarById(id int) (*entity.Car, error){
	car, err := u.carRepository.GetCarById(id)
	if err != nil || car == nil{
		return nil, errors.New("car not found")
	}
	return car, nil
}