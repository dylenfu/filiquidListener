package dao

import (
	"fmt"

	"github.com/filiquid/listener/model"
)

type Stake struct {
	StakeId   string `gorm:"index"`
	Staker    string `gorm:"index"`
	Amount    string
	Start     uint64
	End       uint64
	Minted    string
	TimeStamp uint64
}

func (s *Stake) Up(raw *model.Staked) *Stake {
	s.StakeId = raw.Id.String()
	s.TimeStamp = raw.Timestamp
	s.Staker = raw.Staker.Hex()
	s.Amount = raw.Amount.String()
	s.Start = raw.Start.Uint64()
	s.End = raw.End.Uint64()
	s.Minted = raw.Minted.String()
	return s
}

func (s *Stake) Down() *model.StakedData {
	d := new(model.StakedData)
	d.Id = s.StakeId
	d.Amount = s.Amount
	d.Start = fmt.Sprintf("%d", s.Start)
	d.End = fmt.Sprintf("%d", s.End)
	d.Minted = s.Minted
	d.Timestamp = fmt.Sprintf("%d", s.TimeStamp)

	return d
}

func (d *Dao) InsertStake(data *Stake) error {
	return d.db.Model(&Stake{}).Create(data).Error
}

func (d *Dao) GetStakes(staker string) ([]Stake, error) {
	var list []Stake
	err := d.db.Model(&Stake{}).Where("staker=?", staker).Scan(list).Error
	return list, err
}

func (d *Dao) GetDistinctStakerNum() (int, error) {
	var num int
	err := d.db.Model(&Stake{}).Select("count(distinct staker)").Scan(&num).Error
	return num, err
}
