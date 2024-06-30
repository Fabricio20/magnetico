package tracker

import (
	"context"
	"net"
	"net/url"

	trHttp "github.com/anacrolix/torrent/tracker/http"
	"github.com/anacrolix/torrent/tracker/udp"
	"github.com/anacrolix/torrent/types/infohash"
)

type Client interface {
	Announce(context.Context, AnnounceRequest, AnnounceOpt) (AnnounceResponse, error)
	Scrape(ctx context.Context, ihs []infohash.T) (out udp.ScrapeResponse, err error)
	Close() error
}

type AnnounceOpt = trHttp.AnnounceOpt

type NewClientOpts struct {
	Http trHttp.NewClientOpts
	// Overrides the network in the scheme. Probably a legacy thing.
	UdpNetwork   string
	ListenPacket func(network, addr string) (net.PacketConn, error)
}

func NewClient(urlStr string, opts NewClientOpts) (Client, error) {
	_url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	switch _url.Scheme {
	case "http", "https":
		return trHttp.NewClient(_url, opts.Http), nil
	case "udp", "udp4", "udp6":
		network := _url.Scheme
		if opts.UdpNetwork != "" {
			network = opts.UdpNetwork
		}
		cc, err := udp.NewConnClient(udp.NewConnClientOpts{
			Network:      network,
			Host:         _url.Host,
			ListenPacket: opts.ListenPacket,
		})
		if err != nil {
			return nil, err
		}
		return &udpClient{
			cl:         cc,
			requestUri: _url.RequestURI(),
		}, nil
	default:
		return nil, ErrBadScheme
	}
}
