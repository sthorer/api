package ipfs

import (
	"context"
	"fmt"
	"log"
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
	for attempt := 0; ; attempt++ {
		if attempt > 10 {
			return nil, fmt.Errorf("failed to ping IPFS node at: %s", ipfsNodeURL)
		}

		if _, err := sh.Request("ping").Send(context.Background()); err == nil {
			break
		}

		log.Println("failed to connect to IPFS node. Retrying...")
		time.Sleep(time.Second * 5)
	}

	return &IPFS{Shell: sh}, nil
}
