package main

import (
	"encoding/json"

	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func getMigrations(data amqp.Delivery) error {
	var migrationID string
	if err := json.Unmarshal(data.Body, &migrationID); err != nil {
		return errors.Errorf("[getMigrations] Unmarshal message body error: %s", err)
	}

	migration := models.Migration{ID: migrationID}
	if err := ctx.ES.GetByID(&migration); err != nil {
		return errors.Errorf("[getMigrations] Find migration error: %s", err)
	}

	if err := parseMigration(migration); err != nil {
		return errors.Errorf("[getMigrations] Compute error message: %s", err)
	}

	return nil
}

func parseMigration(migration models.Migration) error {
	if err := ctx.ES.UpdateContractMigrationsCount(migration.Address, migration.Network); err != nil {
		return err
	}

	logger.Info("Migration %s processed", migration.ID)
	return nil
}
