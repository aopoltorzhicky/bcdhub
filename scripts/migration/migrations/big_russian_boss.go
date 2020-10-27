package migrations

import (
	"github.com/baking-bad/bcdhub/internal/config"
)

const yes = "yes"

// BigRussianBoss -
type BigRussianBoss struct{}

// Key -
func (m *BigRussianBoss) Key() string {
	return "big_russian_boss"
}

// Description -
func (m *BigRussianBoss) Description() string {
	return "Script for filling missing or manual metadata"
}

// Do - migrate function
func (m *BigRussianBoss) Do(ctx *config.Context) error {
	if err := m.fillTZIP(ctx); err != nil {
		return err
	}
	if err := m.createTZIP(ctx); err != nil {
		return err
	}
	if err := m.fillAliases(ctx); err != nil {
		return err
	}
	return nil
}

func (m *BigRussianBoss) fillTZIP(ctx *config.Context) error {
	answer, err := ask("Do you want to fill TZIP data from repository? (yes/no)")
	if err != nil {
		return err
	}
	if answer == yes {
		migration := FillTZIP{}
		if err := migration.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (m *BigRussianBoss) createTZIP(ctx *config.Context) error {
	answer, err := ask("Do you want to create missing TZIP data? (yes/no)")
	if err != nil {
		return err
	}
	if answer == yes {
		migration := CreateTZIP{}
		if err := migration.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (m *BigRussianBoss) fillAliases(ctx *config.Context) error {
	answer, err := ask("Do you want to fill aliases? (yes/no)")
	if err != nil {
		return err
	}
	if answer == yes {
		migration := Aliases{}
		if err := migration.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}
