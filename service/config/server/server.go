package server

import (
	pb "github.com/micro/micro/v3/proto/config"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config/handler"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
	"github.com/urfave/cli/v2"
)

const (
	name    = "config"
	address = ":8001"
)

var (
	// Flags specific to the config service
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config_secret_key",
			EnvVars: []string{"MICRO_CONFIG_SECRET_KEY"},
			Usage:   "set config secret key",
		},
	}
)

// Run micro config
func Run(c *cli.Context) error {
	srv := service.New(
		service.Name(name),
		service.Address(address),
	)

	store.DefaultStore.Init(store.Table("config"))

	// register the handler
	pb.RegisterConfigHandler(srv.Server(), handler.NewConfig(c.String("config_secret_key")))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
	return nil
}
