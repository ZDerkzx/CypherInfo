package models

import (
	"context"

	"github.com/google/uuid"
)

const (
	CategoryVPS       = "vps"
	CategoryVPSPlus   = "vps-plus"
	CategoryVDS       = "vds"
	CategoryStorage   = "storage"
	CategoryBareMetal = "bare-metal"
)

type Plan struct {
	Id           uuid.UUID `json:"id"`
	Category     string    `json:"category"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Specs        Specs     `json:"specs"`
	PriceMonthly Money     `json:"price_monthly"`
	BaseCost     Money     `json:"base_cost"`
	Available    bool      `json:"available"`
}

type Specs struct {
	Name        string  `json:"name"`
	Provider    string  `json:"provider"`
	VCpu        int     `json:"vcpu"`
	CpuModel    string  `json:"cpu_model"`
	RamGb       int     `json:"ram_gb"`
	RamNote     string  `json:"ram_note"`
	DiskGB      int     `json:"disk_gb"`
	Storage     string  `json:"storage"`
	Snapshots   int     `json:"snapshots"`
	Network     string  `json:"network"`
	Traffic     string  `json:"traffic"`
	Setup       string  `json:"setup"`
	BandwidthTB float64 `json:"bandwidth_tb"`
	Location    string  `json:"location"`
}

type PlanService interface {
	List(ctx context.Context) ([]Plan, error)
	ListByCategory(ctx context.Context, category string) ([]Plan, error)
}
