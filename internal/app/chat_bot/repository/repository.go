package repository

import (
	"context"
	"errors"

	"github.com/ahargunyllib/thera-be/domain/contracts"
	"github.com/jmoiron/sqlx"
)

type chatBotRepository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

// Begin implements contracts.ChatBotRepository.
func (c *chatBotRepository) Begin(ctx context.Context) error {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	c.tx = tx

	return nil
}

// Commit implements contracts.ChatBotRepository.
func (c *chatBotRepository) Commit() error {
	if c.tx == nil {
		return errors.New("transaction is nil")
	}

	err := c.tx.Commit()
	if err != nil {
		return err
	}

	c.tx = nil

	return nil
}

// Rollback implements contracts.ChatBotRepository.
func (c *chatBotRepository) Rollback() error {
	if c.tx == nil {
		return errors.New("transaction is nil")
	}

	err := c.tx.Rollback()
	if err != nil {
		return err
	}

	c.tx = nil

	return nil
}

func NewChatBotRepository(db *sqlx.DB) contracts.ChatBotRepository {
	return &chatBotRepository{
		db: db,
	}
}
