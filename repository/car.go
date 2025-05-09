package repository

import (
	"database/sql"
	"interview/entity"
)

type CarRepository interface {
	GetAll()([]entity.Car, error)
	Create(car entity.Car) (*entity.Car, error)
	Delete(id int) error
	Update(id int, car entity.Car)(*entity.Car, error)
	GetCarById(id int)(*entity.Car, error)
}

type carRepositoryDatabase struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) CarRepository{
	return &carRepositoryDatabase{db: db}
}

func (r *carRepositoryDatabase) GetAll()([]entity.Car, error){
	query := `select id, car_name, day_rate, month_rate, image from cars`
	rows, err := r.db.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var cars []entity.Car
	for rows.Next(){
		var car entity.Car
		err := rows.Scan(
			&car.ID,
			&car.Name,
			&car.DayRate,
			&car.MonthRate,
			&car.Image,
		)
		if err != nil{
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *carRepositoryDatabase) GetCarById(id int)(*entity.Car, error){
	query := `select id, car_name, day_rate, month_rate, image from cars where id = $1`

	var car entity.Car
	err:=r.db.QueryRow(query, id).Scan(
		&car.ID,
		&car.Name,
		&car.DayRate,
		&car.MonthRate,
		&car.Image,
	)

	if err != nil{
		return nil, err
	}

	return &car, nil
}

func (r *carRepositoryDatabase) Create(car entity.Car) (*entity.Car, error){
	query := `
	insert into cars (car_name, day_rate, month_rate, image) 
	values($1, $2, $3, $4)
	returning id;`
	var id int
	err := r.db.QueryRow(
		query,
		car.Name,
		car.DayRate,
		car.MonthRate,
		car.Image,
	).Scan(&id)
	if err != nil{
		return nil, err
	}

	car.ID = id
	return &car, nil
}

func (r *carRepositoryDatabase) Delete(id int) error{
	query := `delete from cars where id = $1`
	_,err := r.db.Exec(query, id)
	return err
}

func (r *carRepositoryDatabase) Update(id int, car entity.Car)(*entity.Car, error){
	query := `
	update cars
	set car_name=$1, day_rate=$2, month_rate=$3, image=$4
	where id=$5
	`

	_, err := r.db.Exec(
		query,
		car.Name,
		car.DayRate,
		car.MonthRate,
		car.Image,
		id,
	)  
	if err != nil{
		return nil, err
	}

	car.ID = id
	return &car, nil
}