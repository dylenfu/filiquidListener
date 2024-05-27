package cache

import (
	"fmt"
	"math/big"

	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
	"github.com/filiquid/listener/model"
	"github.com/filiquid/listener/utils"
)

func (c *CacheData) SetFamilies(raw []byte) {
	c.Families = raw
}
func (c *CacheData) GetFamilies() []byte {
	return c.Families
}

// todo(xk): delete pannel data
func (c *CacheData) SetBasicSeniorPanel(basic, senior, panel []byte) {
	c.Basic = basic
	c.Senior = senior
	c.Panel = panel
}
func (c *CacheData) GetBasicSeniorPanel() (basic, senior, panel []byte) {
	basic = c.Basic
	senior = c.Senior
	panel = c.Panel
	return
}

func (c *CacheData) FetchAndSaveBasicSeniorData() error {
	var (
		front = new(model.DataStructFront)
	)

	nBasic, err := c.dao.GetBasicDataCount()
	if err != nil {
		return err
	}
	nSenior, err := c.dao.GetSeniorDataCount()
	if err != nil {
		return err
	}

	rawLatestBasic, err := c.dao.GetLatestBasicData()
	if err != nil {
		return err
	}

	latestBasic := rawLatestBasic.Down()
	panel, err := basicData2Front(latestBasic)
	if err != nil {
		return err
	}
	front.Panel = *panel

	rawLatestSenior, err := c.dao.GetLatestSeniorData()
	if err != nil {
		return err
	}
	latestSenior := rawLatestSenior.Down()
	apy, err := utils.CalculateAPY(latestSenior.InterestExp, latestSenior.TotalFIL)
	if err != nil {
		return err
	}
	front.Panel.APY = apy

	// 取样basic data，如果数据总量小于MAXIMUMRESULTS(100), 直接获取整张表；
	// 否则，取样间隔为 count/MAXIMUMRESULTS
	maxIntervalNum := int64(config.MAXIMUMRESULTS)
	var basicIntervalList []dao.BasicData
	if nBasic < maxIntervalNum {
		if basicIntervalList, err = c.dao.GetBasicDataAll(); err != nil {
			return err
		}
	} else {
		if basicIntervalList, err = c.dao.GetBasicDataInterval(nBasic / maxIntervalNum); err != nil {
			return err
		}
	}
	front.Basic = make([]model.BasicDataStructFront, len(basicIntervalList))
	for i := 0; i < len(basicIntervalList); i++ {
		d := basicIntervalList[i].Down()
		rBasicI, err := utils.GetBasicDataStructFront(d)
		if err != nil {
			return err
		}
		front.Basic[i] = *rBasicI
	}

	//4 APY = InterestExp / TotalFIL
	// senior数据抽样
	var seniorIntervalList []dao.SeniorData
	if nSenior < maxIntervalNum {
		if seniorIntervalList, err = c.dao.GetSeniorDataAll(); err != nil {
			return err
		}
	} else {
		if seniorIntervalList, err = c.dao.GetSeniorDataInterval(nSenior / maxIntervalNum); err != nil {
			return err
		}
	}
	front.Senior = make([]model.SeniorDataStructFront, 0)
	for i := 0; i < len(seniorIntervalList); i++ {
		d := seniorIntervalList[i].Down()
		if temp, err := utils.GetSeniorDataStructFront3(d); err == nil {
			front.Senior = append(front.Senior, *temp)
		}
	}

	//1 DAY
	basic1DayBond := int64(rawLatestBasic.BlockTimeStamp - config.DAY1)
	basic1DayList, err := c.getBasicSamplingData(basic1DayBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Basic1Day = basic1DayList

	senior1DayBond := int64(rawLatestSenior.BlockTimeStamp - config.DAY1)
	senior1DayList, err := c.getSeniorSamplingData(senior1DayBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Senior1Day = senior1DayList

	// 7 day
	basic7DayBond := int64(rawLatestBasic.BlockTimeStamp - config.DAY7)
	basic7DayList, err := c.getBasicSamplingData(basic7DayBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Basic7Day = basic7DayList

	senior7DayBond := int64(rawLatestSenior.BlockTimeStamp - config.DAY7)
	senior7DayList, err := c.getSeniorSamplingData(senior7DayBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Senior7Day = senior7DayList

	// 1 month
	basic1MonthBond := int64(rawLatestBasic.BlockTimeStamp - config.MONTH)
	basic1MonthList, err := c.getBasicSamplingData(basic1MonthBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Basic1Month = basic1MonthList

	senior1MonthBond := int64(rawLatestSenior.BlockTimeStamp - config.MONTH)
	senior1MonthList, err := c.getSeniorSamplingData(senior1MonthBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Senior1Month = senior1MonthList

	// 3 month
	basic3MonthBond := int64(rawLatestBasic.BlockTimeStamp - config.MONTH3)
	basic3MonthList, err := c.getBasicSamplingData(basic3MonthBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Basic3Month = basic3MonthList

	senior3MonthBond := int64(rawLatestSenior.BlockTimeStamp - config.MONTH3)
	senior3MonthList, err := c.getSeniorSamplingData(senior3MonthBond, maxIntervalNum)
	if err != nil {
		return err
	}
	front.Senior3Month = senior3MonthList

	bs, sn, pn := utils.DataStructFrontConvert(front)
	c.SetBasicSeniorPanel(
		[]byte(utils.ToString(bs)),
		[]byte(utils.ToString(sn)),
		[]byte(utils.ToString(pn)))
	return nil
}

func (c *CacheData) getBasicSamplingData(lowerBond, maxIntervalNum int64) ([]model.BasicDataStructFront, error) {

	nBasic1Day, err := c.dao.GetBasicDataLowerTimeCount(lowerBond)
	if err != nil {
		return nil, err
	}
	var basic1DayIntervalList []dao.BasicData
	if nBasic1Day < maxIntervalNum {
		if basic1DayIntervalList, err = c.dao.GetBasicDataLowerTimeList(lowerBond); err != nil {
			return nil, err
		}
	} else {
		if basic1DayIntervalList, err = c.dao.GetBasicDataLowerTimeIntervalList(lowerBond, nBasic1Day/maxIntervalNum); err != nil {
			return nil, err
		}
	}
	list := make([]model.BasicDataStructFront, len(basic1DayIntervalList))
	for i := 0; i < len(basic1DayIntervalList); i++ {
		d := basic1DayIntervalList[i].Down()
		rb, err := utils.GetBasicDataStructFront(d)
		if err != nil {
			return nil, err
		}
		list[i] = *rb
	}
	return list, nil
}

func (c *CacheData) getSeniorSamplingData(lowerBond, maxIntervalNum int64) ([]model.SeniorDataStructFront, error) {
	nSenior1Day, err := c.dao.GetSeniorDataLowerTimeCount(lowerBond)
	if err != nil {
		return nil, err
	}
	var senior1DayIntervalList []dao.SeniorData
	if nSenior1Day < maxIntervalNum {
		if senior1DayIntervalList, err = c.dao.GetSeniorDataLowerTimeList(nSenior1Day); err != nil {
			return nil, err
		} else {
			if senior1DayIntervalList, err = c.dao.GetSeniorDataLowerTimeIntervalList(lowerBond, nSenior1Day/maxIntervalNum); err != nil {
				return nil, err
			}
		}
	}
	list := make([]model.SeniorDataStructFront, 0)
	for i := 0; i < len(senior1DayIntervalList); i++ {
		d := senior1DayIntervalList[i].Down()
		if b, err := utils.GetSeniorDataStructFront3(d); err == nil {
			list = append(list, *b)
		}
	}
	return list, nil
}

func basicData2Front(d *model.BasicDataStruct) (*model.PanelStructFront, error) {
	panel := new(model.PanelStructFront)
	rRateBase, err := utils.Str2Uint64(d.RateBase)
	if err != nil {
		return nil, err
	}
	rateBase := float64(rRateBase)
	panel.TotalFIL = d.TotalFIL
	panel.FitTotalSupply = d.FitTotalSupply
	rUtilizationRate, err := utils.Str2Uint64(d.UtilizationRate)
	if err != nil {
		return nil, err
	}
	panel.UtilizationRate = float64(rUtilizationRate) / rateBase
	rExchange, err := utils.Str2Uint64(d.ExchangeRate)
	if err != nil {
		return nil, err
	}
	panel.FIL_FIT = float64(rExchange) / rateBase
	panel.AvailableFILRedeem = d.AvailableFIL

	pp, ok := new(big.Int).SetString(d.TotalFIL, 10)
	if !ok {
		return nil, fmt.Errorf("TotalFIL value not legal")
	}

	qq, ok := new(big.Int).SetString(d.UtilizedLiquidity, 10)
	if !ok {
		return nil, fmt.Errorf("UtilizedLiquidity value not legal")
	}

	pp.Mul(pp, big.NewInt(9)).Div(pp, big.NewInt(10))
	r := pp.Cmp(qq)
	if r <= 0 {
		panel.AvailableFIL = "0"
	} else {
		panel.AvailableFIL = pp.Sub(pp, qq).String()
	}

	panel.UtilizedLiquidity = d.UtilizedLiquidity
	if panel.UtilizationRate >= config.U_M {
		rRM, err := utils.Str2Uint64(d.RM)
		if err != nil {
			return nil, err
		}
		panel.InterestRate = float64(rRM) / rateBase
	} else {
		rInterestRate, err := utils.Str2Uint64(d.InterestRate)
		if err != nil {
			return nil, err
		}
		panel.InterestRate = float64(rInterestRate) / rateBase
	}
	panel.FigTotalSupply = d.FigTotalSupply
	panel.FigReleased = d.ReleasedFigStake
	panel.AccumulatedStakeMint = d.AccumulatedStakeMint
	panel.AccumulatedInterestMint = d.AccumulatedInterestMint

	return panel, nil
}
