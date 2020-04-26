package ipfs

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFS struct {
	*shell.Shell
}

const defaultIPFSNodeURL = "127.0.0.1:5001"

func Initialize() (*IPFS, error) {
	ipfsNodeURL := os.Getenv("STHORER_IPFS_NODE_URL")
	if ipfsNodeURL == "" {
		ipfsNodeURL = defaultIPFSNodeURL
	}

	sh := shell.NewShellWithClient(ipfsNodeURL, &http.Client{Timeout: time.Second * 10})
	if _, err := sh.Request("ping").Send(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping IPFS node at: %s", ipfsNodeURL)
	}

	return &IPFS{Shell: sh}, nil
}
