package config

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

func Http2Config() *http.Client {
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true, // penting untuk HTTP/2 tanpa TLS
			DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
				// gunakan koneksi TCP biasa tanpa TLS
				return (&net.Dialer{}).DialContext(ctx, network, addr)
			},
		},
	}
	return client
}
