package model

import (
	"model/block"
	"model/index"
	"stat"
)

type Volume struct {
	Id      int32             `json:"id"`
	Status  *stat.Stats       `json:"stats"`
	Block   *block.Superblock `json:"block"`
	Indexer *index.Indexer    `json:"index"`
	// data
	conf    *Config
	needles map[int64]int64
	ch      chan uint32
}

func NewVolume(id int32, blockFile string, indexFile string, c *Config) (v *Volume, err error) {
	v = &Volume{}
	v.Id = id
	v.conf = c
	v.Status = &stat.Stats{}
	v.needles = make(map[int64]int64)
	v.ch = make(chan uint32)
	if v.Block, err = block.NewSuperBlock(blockFile); err != nil {
		return nil, err
	}
	if v.Indexer, err = index.NewIndexer(indexFile); err != nil {
		return nil, err
	}
	if err = v.init(); err != nil {
		return nil, err
	}
	return
}

func (v *Volume) init() (err error) {
	return
}
