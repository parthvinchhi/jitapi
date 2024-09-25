package jitapi

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Config DbConfig
	db     *gorm.DB
}

func (p *Postgres) buildConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.Config.DBHost, p.Config.DBPort, p.Config.DBUser, p.Config.DBPassword, p.Config.DBName, p.Config.DBSslMode)
}

func (p *Postgres) Connect() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.buildConnectionString(),
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err == nil {
		err = sqlDB.Ping()
		if err != nil {
			return err
		}
	}
	p.db = db

	return nil
}

func (p *Postgres) Query(query string) ([]Data, error) {
	var data []Data

	if err := p.db.Raw(query).Scan(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
