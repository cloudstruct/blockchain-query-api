package models

import (
	"time"
)

type Block struct {
	Id              int64      `gorm:"column:id"`
	Hash            []byte     `gorm:"column:hash"` // This is a "hash32type" column
	EpochNumber     uint16     `gorm:"column:epoch_no"`
	SlotNumber      uint32     `gorm:"column:slot_no"`
	EpochSlotNumber uint32     `gorm:"column:epoch_slot_no"`
	BlockNumber     uint32     `gorm:"column:block_no"`
	PreviousID      int64      `gorm:"column:previous_id"`    // block(id)
	SlotLeaderID    int64      `gorm:"column:slot_leader_id"` // slot_leader(id)
	Size            uint32     `gorm:"column:size"`
	Time            *time.Time `gorm:"column:time"`
	TxCount         int64      `gorm:"column:tx_count"`
	ProtoMajor      uint16     `gorm:"column:proto_major"`
	ProtoMinor      uint16     `gorm:"column:proto_minor"`
	VrfKey          string     `gorm:"column:vrf_key"`
	OpCert          []byte     `gorm:"column:op_cert"`         // This is a "hash32type" column
	OpCertCounter   uint32     `gorm:"column:op_cert_counter"` // This is a "word63type" column
}

// Override default pluralized table name
func (Block) TableName() string {
	return "block"
}
