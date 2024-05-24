package dao

import (
	"gorm.io/gorm"
)

type User struct {
	FamilyAddr string `gorm:"primaryKey"`
	MinerCount uint64
}

func (d *Dao) InsertUser(family string) error {
	user := User{FamilyAddr: family, MinerCount: 0}
	return d.db.Where(User{FamilyAddr: family}).FirstOrCreate(&user).Error
}

func (d *Dao) IncreaseMinerCount(familyAddr string) error {
	err := d.db.Model(&User{}).Where("family_addr=?", familyAddr).Update("miner_count", gorm.Expr("miner_count + ?", 1)).Error
	return err
}

func (d *Dao) DecreaseMinerCount(familyAddr string) error {
	err := d.db.Model(&User{}).Where("family_addr=?", familyAddr).Update("miner_count", gorm.Expr("miner_count - ?", 1)).Error
	return err
}

func (d *Dao) DeleteUser(familyAddr string) error {
	condition := &User{FamilyAddr: familyAddr, MinerCount: 0}
	return d.db.Delete(condition).Error
}

func (d *Dao) GetUserFamilyAddress() ([]string, error) {
	var list []string
	err := d.db.Model(&User{}).Select("family_addr").Scan(&list).Error
	return list, err
}
