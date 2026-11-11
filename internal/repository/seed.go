package repository

import "cypher/internal/models"

const (
	noSetup   = "Sin cargo de configuración"
	unlimited = "Ilimitado*"
)

func usd(d int64, c int32) models.Money {
	return models.Money{Dollars: d, Cents: c}
}

func mk(category, name string, price, base models.Money, available bool, s models.Specs) models.Plan {
	return models.Plan{Category: category, Name: name, PriceMonthly: price, BaseCost: base, Available: available, Specs: s}
}

func (r *PlanRepo) SeedPlans() error {
	for _, p := range seedPlans() {
		if err := r.CreatePlan(p); err != nil {
			return err
		}
	}
	return nil
}

func seedPlans() []models.Plan {
	v, vp, vd, st, bm := models.CategoryVPS, models.CategoryVPSPlus, models.CategoryVDS, models.CategoryStorage, models.CategoryBareMetal
	return []models.Plan{
		mk(v, "Cloud VPS 4", usd(6, 90), usd(5, 28), true, models.Specs{VCpu: 4, RamGb: 8, DiskGB: 100, Storage: "SSD 100 GB", Snapshots: 1, Network: "200 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(v, "Cloud VPS 6", usd(9, 40), usd(7, 20), true, models.Specs{VCpu: 6, RamGb: 12, DiskGB: 200, Storage: "SSD 200 GB / NVMe 100 GB", Snapshots: 2, Network: "300 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(v, "Cloud VPS 8", usd(17, 50), usd(13, 44), true, models.Specs{VCpu: 8, RamGb: 24, DiskGB: 300, Storage: "SSD 300 GB / NVMe 150 GB", Snapshots: 3, Network: "600 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(v, "Cloud VPS 12", usd(31, 20), usd(24, 0), true, models.Specs{VCpu: 12, RamGb: 48, DiskGB: 400, Storage: "SSD 400 GB / NVMe 200 GB", Snapshots: 3, Network: "800 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(v, "Cloud VPS 16", usd(46, 30), usd(35, 60), true, models.Specs{VCpu: 16, RamGb: 64, DiskGB: 500, Storage: "SSD 500 GB / NVMe 250 GB", Snapshots: 3, Network: "1 Gbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(v, "Cloud VPS 18", usd(61, 20), usd(47, 4), true, models.Specs{VCpu: 18, RamGb: 96, DiskGB: 600, Storage: "SSD 600 GB / NVMe 300 GB", Snapshots: 3, Network: "1 Gbit/s", Traffic: unlimited, Setup: noSetup}),

		mk(vp, "Cloud VPS Plus 4", usd(16, 90), usd(13, 0), true, models.Specs{VCpu: 4, CpuModel: "AMD EPYC", RamGb: 8, DiskGB: 150, Storage: "NVMe 150 GB (más espacio disponible)", Snapshots: 5, Network: "500 Mbit/s", Traffic: unlimited}),
		mk(vp, "Cloud VPS Plus 6", usd(23, 90), usd(18, 40), true, models.Specs{VCpu: 6, CpuModel: "AMD EPYC", RamGb: 12, DiskGB: 300, Storage: "NVMe 300 GB (más espacio disponible)", Snapshots: 5, Network: "500 Mbit/s", Traffic: unlimited}),
		mk(vp, "Cloud VPS Plus 8", usd(43, 70), usd(33, 60), true, models.Specs{VCpu: 8, CpuModel: "AMD EPYC", RamGb: 24, DiskGB: 450, Storage: "NVMe 450 GB (más espacio disponible)", Snapshots: 5, Network: "1 Gbit/s", Traffic: unlimited}),
		mk(vp, "Cloud VPS Plus 12", usd(73, 80), usd(56, 80), true, models.Specs{VCpu: 12, CpuModel: "AMD EPYC", RamGb: 48, DiskGB: 600, Storage: "NVMe 600 GB (más espacio disponible)", Snapshots: 5, Network: "1 Gbit/s", Traffic: unlimited}),
		mk(vp, "Cloud VPS Plus 16", usd(98, 80), usd(76, 0), true, models.Specs{VCpu: 16, CpuModel: "AMD EPYC", RamGb: 64, DiskGB: 750, Storage: "NVMe 750 GB (más espacio disponible)", Snapshots: 5, Network: "1 Gbit/s", Traffic: unlimited}),
		mk(vp, "Cloud VPS Plus 18", usd(123, 80), usd(95, 20), true, models.Specs{VCpu: 18, CpuModel: "AMD EPYC", RamGb: 96, DiskGB: 900, Storage: "NVMe 900 GB (más espacio disponible)", Snapshots: 5, Network: "1 Gbit/s", Traffic: unlimited}),

		mk(vd, "Cloud VDS S", usd(67, 60), usd(52, 0), true, models.Specs{VCpu: 6, CpuModel: "AMD EPYC 7282 2.8 GHz", RamGb: 24, DiskGB: 180, Storage: "NVMe 180 GB (más espacio disponible)", Network: "250 Mbit/s", Traffic: unlimited}),
		mk(vd, "Cloud VDS M", usd(84, 80), usd(65, 20), true, models.Specs{VCpu: 8, CpuModel: "AMD EPYC 7282 2.8 GHz", RamGb: 32, DiskGB: 240, Storage: "NVMe 240 GB (más espacio disponible)", Network: "500 Mbit/s", Traffic: unlimited}),
		mk(vd, "Cloud VDS L", usd(119, 60), usd(92, 0), true, models.Specs{VCpu: 12, CpuModel: "AMD EPYC 7282 2.8 GHz", RamGb: 48, DiskGB: 360, Storage: "NVMe 360 GB (más espacio disponible)", Network: "750 Mbit/s", Traffic: unlimited}),
		mk(vd, "Cloud VDS XL", usd(134, 20), usd(103, 20), true, models.Specs{VCpu: 16, CpuModel: "AMD EPYC 7282 2.8 GHz", RamGb: 64, DiskGB: 480, Storage: "NVMe 480 GB (más espacio disponible)", Network: "1 Gbit/s", Traffic: unlimited}),
		mk(vd, "Cloud VDS XXL", usd(236, 10), usd(181, 60), true, models.Specs{VCpu: 24, CpuModel: "AMD EPYC 7282 2.8 GHz", RamGb: 96, DiskGB: 720, Storage: "NVMe 720 GB (más espacio disponible)", Network: "1 Gbit/s", Traffic: unlimited}),

		mk(st, "Almacenamiento VPS 10", usd(6, 90), usd(5, 28), true, models.Specs{VCpu: 2, RamGb: 4, DiskGB: 300, Storage: "SSD 300 GB", Network: "200 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(st, "Almacenamiento VPS 20", usd(9, 40), usd(7, 20), true, models.Specs{VCpu: 3, RamGb: 8, DiskGB: 400, Storage: "SSD 400 GB", Network: "300 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(st, "Almacenamiento VPS 30", usd(17, 50), usd(13, 44), true, models.Specs{VCpu: 6, RamGb: 18, DiskGB: 1000, Storage: "SSD 1 TB", Network: "600 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(st, "Almacenamiento VPS 40", usd(31, 20), usd(24, 0), true, models.Specs{VCpu: 8, RamGb: 30, DiskGB: 1200, Storage: "SSD 1,2 TB", Network: "800 Mbit/s", Traffic: unlimited, Setup: noSetup}),
		mk(st, "Almacenamiento VPS 50", usd(46, 30), usd(35, 60), true, models.Specs{VCpu: 14, RamGb: 50, DiskGB: 1400, Storage: "SSD 1,4 TB", Network: "1 Gbit/s", Traffic: unlimited, Setup: noSetup}),

		mk(bm, "Núcleos AMD Ryzen 12", usd(171, 30), usd(131, 75), false, models.Specs{VCpu: 12, CpuModel: "3,70 GHz · AMD Ryzen 9 7900", RamGb: 64, RamNote: "hasta 128 GB", DiskGB: 1000, Storage: "1 TB NVMe (más espacio disponible)", Network: "1 Gbit/s", Traffic: unlimited}),
		mk(bm, "AMD Genoa de 24 núcleos", usd(279, 60), usd(215, 10), false, models.Specs{VCpu: 24, CpuModel: "2,50 GHz (3,70 GHz máx.) · AMD EPYC 9224", RamGb: 128, RamNote: "REG ECC · hasta 768 GB", DiskGB: 2000, Storage: "2 discos SSD de 1 TB (más espacio disponible)", Network: "1 Gbit/s", Traffic: unlimited}),
		mk(bm, "AMD Turín de 32 núcleos", usd(560, 40), usd(431, 10), true, models.Specs{VCpu: 32, CpuModel: "3,55 GHz (4,20 GHz máx.) · AMD EPYC 9355P", RamGb: 128, RamNote: "hasta 768 GB", DiskGB: 2000, Storage: "2 × 1 TB NVMe (más espacio disponible)", Network: "1 Gbit/s (puerto de 10 Gbit/s disponible)", Traffic: unlimited}),
		mk(bm, "AMD Turín de 64 núcleos", usd(981, 60), usd(755, 10), false, models.Specs{VCpu: 64, CpuModel: "3,20 GHz (4,20 GHz máx.) · AMD EPYC 9555P", RamGb: 192, RamNote: "hasta 1152 GB", DiskGB: 2000, Storage: "2 × 1 TB NVMe (más espacio disponible)", Network: "1 Gbit/s (puerto de 10 Gbit/s disponible)", Traffic: unlimited}),
	}
}
