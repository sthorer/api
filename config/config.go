package config

import (
	"log"
	"os"
	"strconv"

	"github.com/sthorer/api/utils"

	"github.com/go-playground/validator/v10"

	"github.com/sthorer/api/ipfs"

	"github.com/sthorer/api/database"
)

type Config struct {
	// The IPFS node URL
	IPFSNodeURL string

	// IPFS shell connection
	Shell *ipfs.IPFS

	// Database client
	Client *database.Database

	// Listening host
	Host string

	// Listening port
	Port uint16

	// A secret used to encrypt and decrypt JWT tokens
	Secret string

	// Validator instance
	Validator *validator.Validate
}

const (
	defaultHost = "127.0.0.1"
	defaultPort = 1234
)

func Initialize() (conf *Config, err error) {
	host := os.Getenv("STHORER_HOST")
	if host == "" {
		host = defaultHost
	}

	port := uint16(defaultPort)
	if rawPort := os.Getenv("STHORER_PORT"); rawPort != "" {
		p, err := strconv.ParseUint("", 10, 16)
		if err != nil {
			return nil, err
		}

		port = uint16(p)
	}

	secret := os.Getenv("STHORER_SECRET")
	if secret == "" {
		log.Println("STHORER_SECRET is not set. A temporary random secret will be generated")
		newSecret, err := utils.GenerateSecret(40)
		if err != nil {
			return nil, err
		}

		secret = newSecret
	}
	conf = &Config{
		Host:      host,
		Port:      port,
		Secret:    secret,
		Validator: validator.New(),
	}

	if conf.Shell, err = ipfs.Initialize(); err != nil {
		return nil, err
	}

	if conf.Client, err = database.Initialize(); err != nil {
		return nil, err
	}

	return conf, nil
}
