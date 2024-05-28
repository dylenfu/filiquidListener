package dao

import (
	"fmt"

	"github.com/filiquid/listener/model"
)

type UnStake struct {
	StakeId   string `gorm:"index"`
	Staker    string `gorm:"index"`
	Amount    string
	Start     uint64
	End       uint64
	RealEnd   uint64
	Minted    string
	TimeStamp uint64
}

func (s *UnStake) Up(raw *model.Unstaked) *UnStake {
	s.StakeId = raw.Id.String()
	s.TimeStamp = raw.Timestamp
	s.Staker = raw.Staker.Hex()
	s.Amount = raw.Amount.String()
	s.Start = raw.Start.Uint64()
	s.End = raw.End.Uint64()
	s.RealEnd = raw.RealEnd.Uint64()
	s.Minted = raw.Minted.String()
	return s
}

func (s *UnStake) Down() *model.UnstakedData {
	d := new(model.UnstakedData)
	d.Id = s.StakeId
	d.Timestamp = fmt.Sprintf("%d", s.TimeStamp)
	d.Amount = s.Amount
	d.Start = fmt.Sprintf("%d", s.Start)
	d.End = fmt.Sprintf("%d", s.End)
	d.RealEnd = fmt.Sprintf("%d", s.RealEnd)
	d.Minted = s.Minted
	return d
}

func (d *Dao) InsertUnstake(data *UnStake) error {
	return d.db.Model(&UnStake{}).Create(data).Error
}

func (d *Dao) GetUnstake(staker string) ([]UnStake, error) {
	var list []UnStake
	err := d.db.Model(&UnStake{}).Where("staker=?", staker).Scan(&list).Error
	return list, err
}
