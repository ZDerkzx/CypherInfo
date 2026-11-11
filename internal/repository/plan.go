package repository

import (
	"context"
	"cypher/internal/lib/srz"
	"cypher/internal/models"
	"sort"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

type PlanRepo struct {
	db *bbolt.DB
}

func NewPlanRepo(db *bbolt.DB) *PlanRepo {
	return &PlanRepo{db: db}
}

func (r *PlanRepo) List(ctx context.Context) ([]models.Plan, error) {
	plans := []models.Plan{}
	err := r.db.View(func(tx *bbolt.Tx) error {
		c := tx.Bucket([]byte("plans")).Cursor()
		for k, data := c.First(); k != nil; k, data = c.Next() {
			p, err := srz.Deserialize[models.Plan](data)
			if err != nil {
				return err
			}
			plans = append(plans, p)
		}
		return nil
	})
	return plans, err
}

// type PlanFilter struct {
// 	MinRamGb int
// 	MaxPrice float64
// }

func (r *PlanRepo) ListByCategory(ctx context.Context, category string) ([]models.Plan, error) {
	all, err := r.List(ctx)
	if err != nil {
		return nil, err
	}
	plans := []models.Plan{}
	for _, p := range all {
		if p.Category == category {
			plans = append(plans, p)
		}
	}
	sort.SliceStable(plans, func(i, j int) bool {
		a := plans[i].PriceMonthly.Dollars*100 + int64(plans[i].PriceMonthly.Cents)
		b := plans[j].PriceMonthly.Dollars*100 + int64(plans[j].PriceMonthly.Cents)
		return a < b
	})
	return plans, nil
}

func (r *PlanRepo) CreatePlan(p models.Plan) error {
	if p.Id == uuid.Nil {
		p.Id = uuid.NewSHA1(uuid.NameSpaceOID, []byte(p.Category+"/"+p.Name))
	}
	data, err := srz.Serialize(p)
	if err != nil {
		return err
	}
	return r.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("plans")).Put([]byte(p.Id.String()), data)
	})
}

// func DeletePlan(db *bbolt.DB, key uuid.UUID) error {
// 	err := db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("plans"))
// 		return b.Delete([]byte(key.String()))
// 	})
// 	return err
// }

// /*
// Mete en una lista todos los planes deserializandolo con la libreria estandar.
// */
// func ViewPlans(db *bbolt.DB) ([]Plan, error) {
// 	plans := []Plan{}
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("plans"))
// 		c := b.Cursor()
// 		for k, data := c.First(); k != nil; k, data = c.Next() {
// 			p, err := srz.Deserialize[Plan](data)
// 			if err != nil {
// 				return err
// 			}
// 			plans = append(plans, p)
// 		}
// 		return nil
// 	})
// 	return plans, err
// }
