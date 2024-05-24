package dao

import (
	"fmt"

	"github.com/filiquid/listener/model"
)

type WithdrawFig struct {
	Id        string `gorm:"index"`
	Staker    string `gorm:"index"`
	Amount    string
	TimeStamp uint64
}

func (w *WithdrawFig) Up(raw *model.WithdrawnFig) *WithdrawFig {
	w.Id = raw.Id.String()
	w.Amount = raw.Amount.String()
	w.Staker = raw.Staker.Hex()
	w.TimeStamp = raw.Timestamp
	return w
}

func (w *WithdrawFig) Down() *model.WithdrawnFigData {
	d := new(model.WithdrawnFigData)
	d.Amount = w.Amount
	d.Id = w.Id
	d.Timestamp = fmt.Sprintf("%d", w.TimeStamp)
	return d
}

func (d *Dao) InsertWithdrawFig(data *WithdrawFig) error {
	return d.db.Model(&WithdrawFig{}).Create(data).Error
}

func (d *Dao) GetWithdrawFig(staker string) ([]WithdrawFig, error) {
	var list []WithdrawFig
	err := d.db.Model(&WithdrawFig{}).Where("staker=?", staker).Scan(&list).Error
	return list, err
}
