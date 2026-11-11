// Punto de entrada: aquí solo se cablean las dependencias (composition root).
package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"cypher/internal/config"
	"cypher/internal/handler"
	"cypher/internal/repository"
)

func main() {
	if err := repository.OpenDb(); err != nil {
		panic(err)
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		log.Error("Invalid configuration", "err", err)
		os.Exit(1)
	}

	repo := repository.NewPlanRepo(repository.Db)
	if err := repo.SeedPlans(); err != nil {
		log.Error("Unable to create plans", "err", err)
		os.Exit(1)
	}
	h := handler.New(repo, log)
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Error("The session could not be created.", "err", err)
		os.Exit(1)
	}
	dg.Identify.Intents = discordgo.IntentsGuilds
	h.Bind(dg, cfg.GuildID)
	if err := dg.Open(); err != nil {
		log.Error("Unable to connect to Discord", "err", err)
		os.Exit(1)
	}
	defer dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Info("turning off bot")
}
