package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	apiKey  string
	port    int
	logLevel string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "yt-stats",
	Short: "YouTube Stats - Fetch and display YouTube video statistics",
	Long: `A CLI tool and API server for fetching YouTube video statistics.

Examples:
  # Start API server
  yt-stats serve

  # Get stats for a video
  yt-stats get dQw4w9WgXcQ

  # Use with custom API key
  yt-stats get VIDEO_ID --api-key YOUR_KEY`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yt-stats.yaml)")
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "YouTube Data API key (env: YTSTATS_API_KEY)")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8998, "API server port")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level (debug, info, warn, error)")

	// Bind flags to viper
	viper.BindPFlag("apiKey", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("logLevel", rootCmd.PersistentFlags().Lookup("log-level"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".yt-stats" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".yt-stats")
	}

	// Environment variables
	viper.SetEnvPrefix("YTSTATS")
	viper.AutomaticEnv()

	// Read config file if it exists
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
