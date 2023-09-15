package cmd

import (
	"os"

	"github.com/dose-na-nuvem/vehicles/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configFile string
	cfg        = config.New()
)

// rootCmd representa o comando base quando chamado sem nenhum subcomando.
var rootCmd = &cobra.Command{
	Use: "vehicles",

	Short: "Controle de veículos e tags",

	Long: `Permite gerenciar os veículos e as tags de customers:
Será possível adicionar carros com multiplas tags e relacionar os carros com os donos atuais.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.yaml",
		"Define o arquivo de configuração a utilizar.")

	// Associa o Viper as flags
	if err := viper.BindPFlags(startCmd.Flags()); err != nil {
		cfg.Logger.Error("falha ao ligar as flags", zap.Error(err))
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	// Configura o nome padrão do arquivo de configuração, sem a extensão.
	viper.SetConfigFile(configFile)

	// Tenta ler o arquivo de configuração, ignorando erros caso o mesmo não seja encontrado
	// retorna um erro se não conseguirmos analisar o arquivo de configuração encontrado.
	if err := viper.ReadInConfig(); err != nil {
		// Não há problems se não existir um arquivo de configuração.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			cfg.Logger.Error("arquivo não encontrado",
				zap.String("arquivo", configFile),
				zap.Error(err),
			)

			return
		}

		cfg.Logger.Error("falha na leitura do arquivo de configuração", zap.Error(err))
	} else {
		cfg.Logger.Info("arquivo de configuração lido", zap.String("config", configFile))
	}

	// Converter o estado interno do Viper em nosso objeto de configuração
	if err := viper.Unmarshal(&cfg); err != nil {
		cfg.Logger.Error("falhou ao converter o arquivo de configuração", zap.Error(err))

		return
	}
}
