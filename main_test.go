package goodle

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
)

var (
	BASE  string
	TOKEN string
)

func TestMain(m *testing.M) {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	BASE = viper.GetString("BASE")
	TOKEN = viper.GetString("TOKEN")
	os.Exit(m.Run())
}
