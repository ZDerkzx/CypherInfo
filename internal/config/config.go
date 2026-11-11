// Package config carga la configuración desde variables de entorno.
package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken string
	GuildID      string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}
	c := Config{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		GuildID:      os.Getenv("GUILD_ID"),
	}
	if c.DiscordToken == "" {
		return c, errors.New("The variable was not found DISCORD_TOKEN")
	}
	return c, nil
}
