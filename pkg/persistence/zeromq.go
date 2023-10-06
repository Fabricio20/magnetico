package persistence

import (
	"encoding/hex"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	zmq "github.com/zeromq/goczmq"
	"go.uber.org/zap"
	"time"
)

type zeromq struct {
	context *zmq.Sock
	cache   *cache.Cache
}

func makeZeroMQ() (Database, error) {
	instance := new(zeromq)
	context, err := zmq.NewPub("tcp://0.0.0.0:2222")
	if err != nil {
		return nil, err
	}
	instance.context = context
	instance.cache = cache.New(5*time.Minute, 10*time.Minute)
	return instance, nil
}

func (instance *zeromq) Engine() databaseEngine {
	return ZeroMQ
}

func (instance *zeromq) DoesTorrentExist(infoHash []byte) (bool, error) {
	_, found := instance.cache.Get(string(infoHash))
	return found, nil
}

func (instance *zeromq) AddNewTorrent(infoHash []byte, name string, files []File) error {
	data, err := json.Marshal(SimpleTorrentSummary{
		InfoHash: hex.EncodeToString(infoHash),
		Name:     name,
		Files:    files,
	})
	if err != nil {
		return errors.Wrap(err, "Failed to encode metadata")
	}
	err = instance.context.SendMessage([][]byte{data})
	if err != nil {
		return errors.Wrap(err, "Failed to transmit")
	}
	instance.cache.Set(string(infoHash), data, cache.DefaultExpiration)
	zap.L().Debug(string(data))
	return nil
}

func (instance *zeromq) Close() {
	instance.context.Destroy()
}
