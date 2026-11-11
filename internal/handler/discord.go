// Package handler es la capa de presentación: traduce interacciones de Discord
// a llamadas al servicio y formatea las respuestas. No contiene lógica de negocio.
package handler

import (
	"context"
	"cypher/internal/models"
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
)

type serie struct {
	command     string
	description string
	title       string
	category    string
	color       int
}

var seriesList = []serie{
	{"vps", "Planes de la Serie Cloud VPS", "☁️ **Serie Cloud VPS**", models.CategoryVPS, 0x3498DB},
	{"vpsplus", "Planes de la Serie Cloud VPS Plus (AMD EPYC High-Performance)", "⚡ **Serie Cloud VPS Plus** (AMD EPYC™ High-Performance)", models.CategoryVPSPlus, 0xE67E22},
	{"vds", "Planes de la Serie Cloud VDS (Virtual Dedicated Servers)", "🖥️ **Serie Cloud VDS** (Virtual Dedicated Servers)", models.CategoryVDS, 0x9B59B6},
	{"almacenamiento", "Planes de la Serie Almacenamiento VPS", "💾 **Serie Almacenamiento VPS** (High-Capacity Storage)", models.CategoryStorage, 0x2ECC71},
	{"baremetal", "Servidores Dedicados Bare-Metal", "🏗️ **Serie Servidores Dedicados Bare-Metal**", models.CategoryBareMetal, 0xE74C3C},
}

var tierEmojis = []string{"🌱", "🌿", "🌳", "🚀", "🔥", "💎"}

type Handler struct {
	plans models.PlanService
	log   *slog.Logger
}

func New(plans models.PlanService, log *slog.Logger) *Handler {
	return &Handler{plans: plans, log: log}
}

func commands() []*discordgo.ApplicationCommand {
	cmds := make([]*discordgo.ApplicationCommand, 0, len(seriesList))
	for _, sr := range seriesList {
		cmds = append(cmds, &discordgo.ApplicationCommand{Name: sr.command, Description: sr.description})
	}
	return cmds
}

// Bind registra los handlers y publica los slash commands cuando el bot está listo.
func (h *Handler) Bind(s *discordgo.Session, guildID string) {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, guildID, commands())
		if err != nil {
			h.log.Error("no se pudieron registrar los comandos", "err", err)
			return
		}
		h.log.Info("bot listo", "user", s.State.User.Username)
	})
	s.AddHandler(h.onInteraction)
}

func (h *Handler) onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//itera sobre la lista y si el comando es igual entonces jala wei
	for _, sr := range seriesList {
		if sr.command == i.ApplicationCommandData().Name {
			h.listSeries(ctx, s, i, sr)
			return
		}
	}
}

func (h *Handler) listSeries(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate, sr serie) {
	plans, err := h.plans.ListByCategory(ctx, sr.category)
	if err != nil {
		h.fail(s, i, err)
		return
	}
	if len(plans) == 0 {
		h.reply(s, i, "No hay planes disponibles en esta serie.", nil)
		return
	}
	if len(plans) > 10 {
		plans = plans[:10]
	}
	embeds := make([]*discordgo.MessageEmbed, 0, len(plans))
	for idx, p := range plans {
		embeds = append(embeds, planEmbed(p, tierEmojis[idx%len(tierEmojis)], sr.color))
	}
	h.reply(s, i, sr.title, embeds)
}

func planEmbed(p models.Plan, tier string, color int) *discordgo.MessageEmbed {
	desc := fmt.Sprintf("💵 **$%d.%02d USD / mes**", p.PriceMonthly.Dollars, p.PriceMonthly.Cents)
	if !p.Available {
		desc += "\n❌ **Agotado**"
	} else if p.Category == models.CategoryBareMetal {
		desc += "\n✅ **Disponible**"
	}
	unit := "vCPU"
	if p.Category == models.CategoryBareMetal {
		unit = "núcleos"
	}
	cpu := fmt.Sprintf("%d %s", p.Specs.VCpu, unit)
	if p.Specs.CpuModel != "" {
		cpu += "\n" + p.Specs.CpuModel
	}
	ram := fmt.Sprintf("%d GB", p.Specs.RamGb)
	if p.Specs.RamNote != "" {
		ram += "\n" + p.Specs.RamNote
	}
	fields := []*discordgo.MessageEmbedField{
		{Name: "🧠 CPU", Value: cpu, Inline: true},
		{Name: "🧬 RAM", Value: ram, Inline: true},
		{Name: "💽 Almacenamiento", Value: p.Specs.Storage, Inline: true},
	}
	if p.Specs.Snapshots > 0 {
		word := "instantáneas"
		if p.Specs.Snapshots == 1 {
			word = "instantánea"
		}
		fields = append(fields, &discordgo.MessageEmbedField{Name: "📸 Copias de seguridad", Value: fmt.Sprintf("%d %s", p.Specs.Snapshots, word), Inline: true})
	}
	fields = append(fields,
		&discordgo.MessageEmbedField{Name: "🌐 Red", Value: p.Specs.Network, Inline: true},
		&discordgo.MessageEmbedField{Name: "📶 Tráfico", Value: p.Specs.Traffic, Inline: true},
	)
	embed := &discordgo.MessageEmbed{
		Title:       tier + " " + p.Name,
		Description: desc,
		Color:       color,
		Fields:      fields,
	}
	if p.Specs.Setup != "" {
		embed.Footer = &discordgo.MessageEmbedFooter{Text: "🛠️ " + p.Specs.Setup}
	}
	return embed
}

func (h *Handler) reply(s *discordgo.Session, i *discordgo.InteractionCreate, content string, embeds []*discordgo.MessageEmbed) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content, Embeds: embeds},
	})
	if err != nil {
		h.log.Error("Error in answering.", "err", err)
	}
}

func (h *Handler) fail(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	h.log.Error("mistake in command", "err", err)
	h.reply(s, i, "An error occurred, please check again - Cypher Hosting.", nil)
}
