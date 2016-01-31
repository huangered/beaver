package model

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Store save volumes & conf
type Store struct {
	volumeIndex *os.File
	Volumes     map[int32]*Volume
	config      *Config
}

func NewStore(c *Config) (s *Store) {
	s = &Store{}
	s.Volumes = make(map[int32]*Volume)
	s.config = c
	s.init()
	return
}

func (s *Store) init() {
	log.Println("Read store conf")
	log.Println("Read volume")
	s.volumeIndex, _ = os.Open("test")
	if err := s.parseVolumeIndex(); err != nil {
		return
	}

}

func (s *Store) parseVolumeIndex() (err error) {
	var (
		data  []byte
		lines []string
	)
	if data, err = ioutil.ReadAll(s.volumeIndex); err != nil {
		log.Panicf("Read volume index file err(%v)", err)
	}
	log.Println(string(data))
	lines = strings.Split(string(data), "\n")
	if err = s.parseIndex(lines); err != nil {
		return err
	}
	return
}

func (s *Store) parseIndex(lines []string) (err error) {
	var (
		id        int64
		vid       int32
		line      string
		blockFile string
		indexFile string
		seqs      []string
		v         *Volume
	)
	for _, line = range lines {
		log.Println(line)
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if seqs = strings.Split(line, ","); len(seqs) != 3 {
			return
		}
		if id, err = strconv.ParseInt(seqs[0], 10, 32); err != nil {
			log.Panicf("Volume index: \"%s\" format error", line[0])
			return
		}
		vid = int32(id)
		blockFile = seqs[1]
		indexFile = seqs[2]
		log.Printf("Parse volume index, id: %d, block: %s, index: %s", vid, blockFile, indexFile)
		if v, err = NewVolume(vid, blockFile, indexFile, s.config); err != nil {
			log.Panicf("New volume error(%v)", err)
			return
		}
		s.Volumes[vid] = v
	}
	log.Println("Parse volume index done.")
	return
}

func (s *Store) AddVolume() {

}

func (s *Store) DelVolume() {

}

func (s *Store) Close() {

}

func (s *Store) Clean() {

}
