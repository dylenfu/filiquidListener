package dao

import (
	"fmt"

	"github.com/filiquid/listener/model"
)

type Vote struct {
	Voter      string `gorm:"index"`
	ProposalId string `gorm:"index"`
	Category   string
	Amount     string
	TimeStamp  uint64
}

func (v *Vote) Up(raw *model.Voted) *Vote {
	v.TimeStamp = raw.Timestamp
	v.Voter = raw.Voter.Hex()
	v.ProposalId = raw.ProposalId.String()
	v.Category = fmt.Sprintf("%d", raw.Category)
	v.Amount = raw.Amount.String()
	return v
}

func (v *Vote) Down() *model.VotedData {
	d := new(model.VotedData)
	d.Voter = v.Voter
	d.Category = v.Category
	d.Amount = v.Amount
	return d
}

func (d *Dao) InsertVote(data *Vote) error {
	return d.db.Model(&Vote{}).Create(data).Error
}

func (d *Dao) GetVoteByProposalId(proposalId string) ([]Vote, error) {
	var list []Vote
	err := d.db.Model(&Vote{}).Where("proposal_id=?", proposalId).Order("time_stamp desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetVoteByProposalIdAndVote(proposalId, voter string) ([]Vote, error) {
	var list []Vote
	err := d.db.Model(&Vote{}).Where("proposal_id=? and voter=?", proposalId, voter).Order("time_stamp desc").Scan(&list).Error
	return list, err
}
