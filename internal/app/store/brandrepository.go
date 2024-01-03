package store

import "web-server/internal/app/model"

type BrandRepository struct {
	store *Store
}

func (r *BrandRepository) Create(b *model.Brand) (*model.Brand, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO (name) VALUES ($1) RETURNING id",
		b.Name,
	).Scan(&b.Id); err != nil {
		return nil, err
	}
	return b, nil
}

func (r *BrandRepository) FindByName(name string) (*model.Brand, error) {
	return nil, nil
}
