package database

import "github.com/jinzhu/gorm"

func (d *db) GetOrCreateUser(u *User) error {
	return d.ORM.Where("login = ?", u.Login).FirstOrCreate(u).Error
}

func (d *db) GetUser(userID uint) (*User, error) {
	var user User

	if err := d.ORM.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *db) ListSubscriptions(userID uint) ([]Subscription, error) {
	var subs []Subscription

	if err := d.ORM.Where("user_id = ?", userID).Order("created_at DESC").Find(&subs).Error; err != nil {
		return nil, err
	}

	return subs, nil
}

func (d *db) ListSubscriptionsWithLimit(userID uint, limit int) ([]Subscription, error) {
	var subs []Subscription

	if err := d.ORM.Order("created_at desc").Limit(limit).Where("user_id = ?", userID).Find(&subs).Error; err != nil {
		return nil, err
	}

	return subs, nil
}

func (d *db) CreateSubscription(s *Subscription) error {
	return d.ORM.Create(s).Error
}

func (d *db) DeleteSubscription(s *Subscription) error {
	return d.ORM.Unscoped().Where("entity_id = ? AND user_id = ? and entity_type = ?", s.EntityID, s.UserID, s.EntityType).Delete(Subscription{}).Error
}

func (d *db) GetSubscription(sID, typ string) (s Subscription, err error) {
	err = d.ORM.Where("entity_id = ? AND entity_type = ?", sID, typ).Find(&s).Error
	return
}

func (d *db) GetSubscriptionRating(entityID string) (SubRating, error) {
	var s SubRating
	if err := d.ORM.Model(&Subscription{}).Where("entity_id = ?", entityID).Count(&s.Count).Error; err != nil {
		return s, err
	}

	if err := d.ORM.Raw(`
		SELECT users.login, users.avatar_url
		FROM subscriptions
		INNER JOIN users ON subscriptions.user_id=users.id
		WHERE entity_id = ? AND subscriptions.deleted_at IS NULL
		LIMIT 5;`, entityID).Scan(&s.Users).Error; err != nil {
		return s, err
	}

	return s, nil
}

func (d *db) GetAliases(network string) ([]Alias, error) {
	var aliases []Alias

	if err := d.ORM.Where("network = ?", network).Find(&aliases).Error; err != nil {
		return nil, err
	}

	return aliases, nil
}

func (d *db) GetAlias(address, network string) (string, error) {
	var alias Alias
	if err := d.ORM.Select("alias").Where("address = ? AND network = ?", address, network).First(&alias).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return "", err
		}
		return "", nil
	}
	return alias.Alias, nil
}

func (d *db) CreateAlias(alias, address, network string) error {
	return d.ORM.Create(&Alias{
		Alias:   alias,
		Address: address,
		Network: network,
	}).Error
}

func (d *db) Close() {
	d.ORM.Close()
}
