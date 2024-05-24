package dao

import (
	"fmt"

	"github.com/filiquid/listener/model"
)

type Proposal struct {
	ProposalId      uint64 `gorm:"uniqueIndex"`
	Proposer        string `gorm:"index"`
	DiscussionIndex string
	Category        string
	Deadline        uint64 `gorm:"index"`
	Deposited       string
	Text            string
	Values          string
	Params          string
	Executed        bool
	Result          string
	TimeStamp       uint64
}

func (data *Proposal) Up(raw *model.Proposed) *Proposal {
	data.Proposer = raw.Proposer.Hex()
	data.ProposalId = raw.ProposalId.Uint64()
	data.DiscussionIndex = raw.DiscussionIndex.String()
	data.Category = fmt.Sprintf("%d", raw.Category)
	data.Deadline = raw.Deadline.Uint64()
	data.Deposited = raw.Deposited.String()
	data.Text = raw.Text
	data.TimeStamp = raw.Timestamp
	data.Executed = raw.Executed
	data.Values = raw.ValueString()
	data.Result = fmt.Sprintf("%d", raw.Result)
	return data
}

func (s *Dao) InsertProposal(data *Proposal) error {
	return s.db.Model(&Proposal{}).Create(data).Error
}

func (s *Dao) UpdateProposal(result string, proposalId uint64) error {
	condition := &Proposal{Result: result, Executed: true}
	return s.db.Model(&Proposal{}).Where("proposal_id=?", proposalId).UpdateColumns(condition).Error
}

func (d *Dao) GetAllProposalIds() ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Order("proposal_id desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetProposalIdByProposer(proposer string) ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Where("proposer=?", proposer).Order("proposal_id desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetProposalIdByDeadline(deadline uint64) ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Where("deadline>=?", deadline).Order("proposal_id desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetProposalIdByProposerAndDeadline(proposer string, deadline uint64) ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Where("proposer=? and deadline>=?", proposer, deadline).Order("proposal_id desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetFinishedProposal(deadline uint64) ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Where("deadline<? and executed=0", deadline).Order("proposal_id desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetFinishedProposalByProposer(proposer string, deadline uint64) ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Where("proposer=? and deadline<? and executed=0", proposer, deadline).Order("proposal_id desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetExcutedProposal() ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Where("executed=1").Order("proposal_id desc").Scan(&list).Error
	return list, err
}

func (d *Dao) GetExcutedProposalByProposer(proposer string) ([]string, error) {
	var list []string
	err := d.db.Model(&Proposal{}).Select("proposal_id").Where("proposer=? and executed=1", proposer).Order("proposal_id desc").Scan(&list).Error
	return list, err
}
