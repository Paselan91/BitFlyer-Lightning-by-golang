package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type ConfigList struct {
	ApiKey      string
	ApiSecret   string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
}

var Config ConfigList

func init() {
	EnvLoad()
	log.Printf("init!!!!")
	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	Config = ConfigList{
		ApiKey:        os.Getenv("api_key"),
		ApiSecret:     os.Getenv("api_secret"),
		ProductCode:   os.Getenv("product_code"),
		Durations:     durations,
		TradeDuration: durations[os.Getenv("trade_duration")],
	}
}

// TODO: main()と２重で呼び出しているので、どうにかしたい
// EnvLoad 環境変数の取得
func EnvLoad() {
	if os.Getenv("ENV_MODE") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
