package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config          *Config
	db              *sql.DB
	brandRepository *BrandRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	if err := s.db.Close(); err != nil {
		return
	}
}

func (s *Store) Brand() *BrandRepository {
	if s.brandRepository != nil {
		return s.brandRepository
	}

	s.brandRepository = &BrandRepository{
		store: s,
	}

	return s.brandRepository
}
