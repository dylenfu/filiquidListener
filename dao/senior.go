package dao

import (
	"fmt"
	"math/big"

	"github.com/filiquid/listener/model"
)

type SeniorData struct {
	BlockHeight          uint64 `gorm:"primaryKey"`
	BlockTimeStamp       uint64
	TotalPendingInterest string
	TotalFIL             string
	Borrowing            string
	BorrowingAndPeriod   string
	AccumulatedPayback   string
	InterestExp          string
}

func (s *SeniorData) Up(d struct {
	BlockHeight          *big.Int
	BlockTimeStamp       *big.Int
	TotalPendingInterest *big.Int
	TotalFIL             *big.Int
	Borrowing            *big.Int
	BorrowingAndPeriod   *big.Int
	AccumulatedPayback   *big.Int
	InterestExp          *big.Int
}) *SeniorData {
	s.BlockHeight = d.BlockHeight.Uint64()
	s.BlockTimeStamp = d.BlockTimeStamp.Uint64()
	s.TotalPendingInterest = d.TotalPendingInterest.String()
	s.TotalFIL = d.TotalFIL.String()
	s.Borrowing = d.Borrowing.String()
	s.BorrowingAndPeriod = d.BorrowingAndPeriod.String()
	s.AccumulatedPayback = d.AccumulatedPayback.String()
	s.InterestExp = d.InterestExp.String()
	return s
}

func (s *SeniorData) Down() *model.SeniorDataStruct {
	d := new(model.SeniorDataStruct)

	d.BlockHeight = fmt.Sprintf("%d", s.BlockHeight)
	d.BlockTimeStamp = fmt.Sprintf("%d", s.BlockTimeStamp)
	d.TotalPendingInterest = s.TotalPendingInterest
	d.TotalFIL = s.TotalFIL
	d.Borrowing = s.Borrowing
	d.BorrowingAndPeriod = s.BorrowingAndPeriod
	d.AccumulatedPayback = s.AccumulatedPayback
	d.InterestExp = s.InterestExp

	return d
}

func (s *Dao) InsertSenior(data *SeniorData) error {
	return s.db.Model(&SeniorData{}).Create(data).Error
}

func (s *Dao) GetSeniorDataCount() (int64, error) {
	var num int64
	if err := s.db.Model(&SeniorData{}).Count(&num).Error; err != nil {
		return 0, err
	}
	return num, nil
}

func (s *Dao) GetLatestSeniorData() (*SeniorData, error) {
	data := new(SeniorData)
	if err := s.db.Model(&SeniorData{}).Order("block_height desc").Limit(1).Scan(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Dao) GetSeniorDataAll() ([]SeniorData, error) {
	var list []SeniorData
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Dao) GetSeniorDataInterval(interval int64) ([]SeniorData, error) {
	var list []SeniorData

	if err := s.db.Find(&list).Where("block_height % ? = 0", interval).Order("block_height asc").Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (s *Dao) GetSeniorDataLowerTimeCount(lowerBond int64) (int64, error) {
	var num int64
	if err := s.db.Model(&SeniorData{}).Select("count(*)").Where("block_time_stamp >= ?", lowerBond).Count(&num).Error; err != nil {
		return 0, err
	}
	return num, nil
}

func (s *Dao) GetSeniorDataLowerTimeList(lowerBond int64) ([]SeniorData, error) {
	var list []SeniorData
	if err := s.db.Model(&SeniorData{}).Select("*").Where("block_time_stamp >= ?", lowerBond).Order("block_height asc").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Dao) GetSeniorDataLowerTimeIntervalList(lowerBond int64, interval int64) ([]SeniorData, error) {
	var list []SeniorData
	if err := s.db.Model(&SeniorData{}).Select("*").Where("block_time_stamp >= ? && block_height % ? = 0", lowerBond, interval).Order("block_height asc").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
