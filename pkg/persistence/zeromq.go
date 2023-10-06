package persistence

import (
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	zmq "github.com/zeromq/goczmq"
)

func makeZeroMQ() (Database, error) {
	instance := new(zeromq)
	context, err := zmq.NewPub("udp://*:2222")
	if err != nil {
		return nil, err
	}
	instance.context = context
	return instance, nil
}

type zeromq struct {
	context *zmq.Sock
}

func (s *zeromq) Engine() databaseEngine {
	return ZeroMQ
}

func (s *zeromq) DoesTorrentExist(infoHash []byte) (bool, error) {
	// Always say that "No the torrent does not exist" because we do not have
	// a way to know if we have seen it before or not.
	// TODO:
	// A possible improvement would be using bloom filters (with low false positive
	// probabilities) to apply some reasonable filtering.
	return false, nil
}

func (s *zeromq) AddNewTorrent(infoHash []byte, name string, files []File) error {
	data, err := json.Marshal(SimpleTorrentSummary{
		InfoHash: hex.EncodeToString(infoHash),
		Name:     name,
		Files:    files,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to encode metadata")
	}
	err = s.context.SendMessage([][]byte{data})
	if err != nil {
		return errors.Wrap(err, "Failed to transmit")
	}
	return nil
}

func (s *zeromq) Close() {
	s.context.Destroy()
}
