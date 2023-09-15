package vehicles

import (
	"context"
	"fmt"

	"github.com/dose-na-nuvem/vehicles/config"
	"github.com/dose-na-nuvem/vehicles/pkg/server"
	"go.uber.org/zap"
)

func Start(ctx context.Context) error {
	cfg := config.New()

	cfg.Logger.Info("service started")

	// http server start
	http := server.NewHTTP(cfg)
	go func() {
		err := http.Start(ctx)
		if err != nil {
			cfg.Logger.Error("falha ao iniciar servidor http", zap.Error(err))
		}
	}()

	return nil
}

func Stop(_ context.Context) error {
	fmt.Printf("service stoped")

	return nil
}
