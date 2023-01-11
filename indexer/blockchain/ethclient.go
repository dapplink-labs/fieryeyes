package chain

import (
	"context"
	"crypto/tls"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultTimeout = 5 * time.Second
)

type EvmChainClient struct {
	context   context.Context
	ethClient *ethclient.Client
}

func NewEvmChainClient(ctx context.Context, url string, disableHTTP2 bool) (*EvmChainClient, error) {
	ctxt, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	var ethClient *ethclient.Client
	if strings.HasPrefix(url, "http") {
		httpClient := new(http.Client)
		if disableHTTP2 {
			log.Info("Disabled HTTP/2 support in L1 eth client")
			httpClient.Transport = &http.Transport{
				TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			}
		}
		rpcClient, err := rpc.DialHTTPWithClient(url, httpClient)
		if err != nil {
			return nil, err
		}
		ethClient, _ = ethclient.NewClient(rpcClient), nil
	}
	ethClient, _ = ethclient.DialContext(ctxt, url)
	return &EvmChainClient{
		context:   ctx,
		ethClient: ethClient,
	}, nil
}
