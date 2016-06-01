package config

type Config struct {
	Telegram struct {
		BotToken      string
		AllowedGroups []int64
	}
}
