package dao

type Family struct {
	MinerId    uint64 `gorm:"primaryKey"`
	FamilyAddr string
}

func (d *Dao) GetFamilyAll() ([]Family, error) {
	var list []Family
	err := d.db.Model(&Family{}).Scan(&list).Error
	return list, err
}

func (d *Dao) InsertFamily(data *Family) error {
	return d.db.Model(&Family{}).Create(data).Error
}

func (d *Dao) DeleteFamily(minerId uint64) error {
	return d.db.Model(&Family{}).Delete("miner_id = ?", minerId).Error
}

func (d *Dao) GetFamilyCount() (int64, error) {
	var num int64
	err := d.db.Model(&Family{}).Count(&num).Error
	return num, err
}
