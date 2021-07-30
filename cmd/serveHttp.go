package cmd

import (
	"context"
	"encoding/json"
	"net/http"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrdhira/arcana/config"
	"github.com/mrdhira/arcana/pkg/logger"
	"github.com/mrdhira/arcana/src/adapters"
	deliveriesHttp "github.com/mrdhira/arcana/src/deliveries/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// serveHttpCmd command
var serveHttpCmd = &cobra.Command{
	Use:   "serveHttp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initiate Config
		var Config *config.AppConfig
		err := viper.Unmarshal(&Config)
		if err != nil {
			logger.Error("error when unmarshal config", zap.Error(err))
			os.Exit(1)
		}

		// Debug Only
		ConfigDebug, _ := json.Marshal(Config)
		logger.Infow("Debug Config", "Config", string(ConfigDebug))

		// Initiate Adapter
		postgreSQLClient := adapters.NewPostgreSQLAdapter(context.Background(), &Config.Postgresql)

		// Initiate Routes
		Routes := deliveriesHttp.NewRoutes(postgreSQLClient)

		HttpServer := &http.Server{
			Addr:         ":3000",
			Handler:      Routes,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 60 * time.Second,
		}

		var GracefulStop = make(chan os.Signal)
		signal.Notify(GracefulStop, syscall.SIGTERM)
		signal.Notify(GracefulStop, syscall.SIGINT)

		go func() {
			if err := HttpServer.ListenAndServe(); err != http.ErrServerClosed {
				logger.Fatal("error on listen and serve", zap.Error(err))
			}
		}()

		<-GracefulStop

		// Closed adapter connection
		postgreSQLClient.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := HttpServer.Shutdown(ctx); err != nil {
			panic(err)
		}
		logger.Info("Http server closed")
	},
}

func init() {
	rootCmd.AddCommand(serveHttpCmd)
}
