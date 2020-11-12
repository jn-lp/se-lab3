package plants

import (
	"database/sql"
	"time"
)

type Plant struct {
	ID                int64   `json:"id"`
	SoilMoistureLevel float64 `json:"soilMoistureLevel"`
	SoilDataTimestamp string  `json:"soilDataTimestamp"`
}

type Store struct {
	Db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// ListPlants returns a list of plants that need to be moisturised
func (s *Store) ListPlants() ([]*Plant, error) {
	rows, err := s.Db.Query("SELECT * FROM plants WHERE soilMoistureLevel <= 0.2 ORDER BY soilDataTimestamp DESC LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Plant
	for rows.Next() {
		var plant Plant
		if err := rows.Scan(&plant.ID, &plant.SoilMoistureLevel, &plant.SoilDataTimestamp); err != nil {
			return nil, err
		}
		soilDataTimestamp, _ := time.Parse(time.RFC3339, plant.SoilDataTimestamp)
		plant.SoilDataTimestamp = soilDataTimestamp.Format(time.RFC3339)
		res = append(res, &plant)
	}
	if res == nil {
		res = make([]*Plant, 0)
	}
	return res, nil
}

// CreatePlant creates a new plant in DB
func (s *Store) CreatePlant(level float64) (Plant, error) {
	rows, err := s.Db.Query("INSERT INTO plants (soilMoistureLevel) VALUES ($1) RETURNING *", level)
	var plant Plant
	rows.Next()
	if err := rows.Scan(&plant.ID, &plant.SoilMoistureLevel, &plant.SoilDataTimestamp); err != nil {
		return plant, err
	}
	soilDataTimestamp, _ := time.Parse(time.RFC3339, plant.SoilDataTimestamp)
	plant.SoilDataTimestamp = soilDataTimestamp.Format(time.RFC3339)
	return plant, err
}

// UpdatePlant updates a plant in DB
func (s *Store) UpdatePlant(id int64, newLevel float64) error {
	_, err := s.Db.Exec("UPDATE plants SET soilMoistureLevel=$1 WHERE id=$2", newLevel, id)
	return err
}
