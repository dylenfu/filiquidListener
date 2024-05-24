package dao

import (
	"fmt"
	"math/big"

	"github.com/filiquid/listener/model"
)

type Interest struct {
	Minter    string `gorm:"index"`
	Principal string
	Interest  string
	Minted    string
	TimeStamp uint64
}

func (i *Interest) Up(raw *model.Interest) *Interest {
	i.Minter = raw.Minter.Hex()
	i.Principal = raw.Principal.String()
	i.Interest = raw.Interest.String()
	i.Minted = raw.Minted.String()
	i.TimeStamp = raw.Timestamp
	return i
}

func (i *Interest) Down() *model.InterestData {
	data := new(model.InterestData)
	data.Minted = i.Minted
	interest, ok := new(big.Int).SetString(i.Interest, 10)
	if !ok {
		return nil
	}
	principal, ok := new(big.Int).SetString(i.Principal, 10)
	if !ok {
		return nil
	}
	data.Amount = new(big.Int).Add(interest, principal).String()
	data.Timestamp = fmt.Sprintf("%d", i.TimeStamp)
	return data
}

func (d *Dao) InsertInterest(data *Interest) error {
	return d.db.Model(&Interest{}).Create(data).Error
}

func (d *Dao) GetInterest(minter string) ([]Interest, error) {
	var list []Interest
	err := d.db.Model(&Interest{}).Where("minter = ?", minter).Scan(list).Error
	return list, err
}
