package gateway

import (
	"github.com/allrivenjs/course-phones-review/gadgets/smartphones/models"
	"github.com/allrivenjs/course-phones-review/internal/database"
	"github.com/allrivenjs/course-phones-review/internal/logs"
)

type SmartphoneStorageGateway interface {
	Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

type SmartphoneStorage struct {
	*database.MySqlClient
}

func (s *SmartphoneStorage) Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	tx, err := s.MySqlClient.Begin()
	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into smartphones (name, price, country_origin, os) values (?, ?, ?, ?)`, cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.Os)
	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fetch last  id")
		_ = tx.Rollback()
		return nil, err
	}
	_ = tx.Commit()
	return &models.Smartphone{
		Id:            id,
		Name:          cmd.Name,
		Price:         cmd.Price,
		CountryOrigin: cmd.CountryOrigin,
		Os:            cmd.Os,
	}, nil
}
