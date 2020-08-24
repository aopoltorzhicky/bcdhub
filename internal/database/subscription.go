package database

import "github.com/jinzhu/gorm"

// Subscription model
type Subscription struct {
	gorm.Model
	UserID    uint   `gorm:"primary_key;not null"`
	Address   string `gorm:"primary_key;not null"`
	Network   string `gorm:"primary_key;not null"`
	Alias     string
	WatchMask uint
	SentryDSN string
}

func (d *db) GetSubscription(userID uint, address, network string) (s Subscription, err error) {
	err = d.ORM.
		Where("user_id = ? AND address = ? AND network = ?", userID, address, network).
		First(&s).Error
	return
}

func (d *db) GetSubscriptions(address, network string) ([]Subscription, error) {
	var subs []Subscription

	err := d.ORM.
		Where("address = ? AND network = ?", address, network).
		Find(&subs).Error

	return subs, err
}

func (d *db) ListSubscriptions(userID uint) ([]Subscription, error) {
	var subs []Subscription

	err := d.ORM.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&subs).Error

	return subs, err
}

func (d *db) UpsertSubscription(s *Subscription) error {
	return d.ORM.
		Where("user_id = ? AND address = ? AND network = ?", s.UserID, s.Address, s.Network).
		Assign(Subscription{Alias: s.Alias, WatchMask: s.WatchMask, SentryDSN: s.SentryDSN}).
		FirstOrCreate(s).Error
}

func (d *db) DeleteSubscription(s *Subscription) error {
	return d.ORM.Unscoped().
		Where("user_id = ? AND address = ? AND network = ?", s.UserID, s.Address, s.Network).
		Delete(Subscription{}).Error
}

func (d *db) GetSubscriptionsCount(address, network string) (count int, err error) {
	err = d.ORM.
		Model(&Subscription{}).
		Where("address = ? AND network = ?", address, network).
		Count(&count).Error
	return
}
