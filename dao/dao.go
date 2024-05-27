package dao

import (
	"fmt"
	"log"
	"time"

	"github.com/filiquid/listener/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(c *config.DbConfig) (*Dao, error) {
	styp := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", c.User, c.Password, c.Url, c.DbName, styp)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("DB opened successfully")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	log.Println("DB ping successfully")
	return &Dao{db: db}, nil
}

func (d *Dao) Close() {
	sqlDB, err := d.db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func (d *Dao) Migrate() {
	d.db.AutoMigrate(&Height{})
	d.db.AutoMigrate(&User{})
	d.db.AutoMigrate(&Family{})
	d.db.AutoMigrate(&Interest{})
	d.db.AutoMigrate(&Stake{})
	d.db.AutoMigrate(&UnStake{})
	d.db.AutoMigrate(&WithdrawFig{})
	d.db.AutoMigrate(&Proposal{})
	d.db.AutoMigrate(&Vote{})
	d.db.AutoMigrate(&BasicData{})
	d.db.AutoMigrate(&SeniorData{})
}
