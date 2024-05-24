package dao

import "gorm.io/gorm/clause"

type Height struct {
	Id          uint64 `gorm:"primaryKey"`
	BlockHeight uint64 `gorm:"unique"`
}

func (d *Dao) GetLastHeight() (uint64, error) {
	var height uint64
	err := d.db.Model(&Height{}).Select("block_height").Order("block_height desc").Limit(1).Scan(&height).Error
	return height, err
}

func (d *Dao) UpdateLastHeight(height uint64) error {
	data := &Height{Id: 0, BlockHeight: height}
	return d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"block_height"}),
	}).Create(data).Error
}
