package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBRepository interface {
	Select(dest interface{}, query interface{}, args ...interface{}) error
	SelectMany(dest interface{}, query interface{}, args ...interface{}) error

	InsertOrUpdate(dest interface{}) error
	Insert(dest interface{}) error
	Update(dest interface{}) error
	Delete(dest interface{}) error
}

type DBService struct {
	DataContext *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
	return &DBService{
		DataContext: db,
	}
}

func (u *DBService) Select(dest interface{}, query interface{}, args ...interface{}) error {
	if err := u.DataContext.Where(query, args...).First(dest).Error; err != nil {
		return err
	}
	return nil
}

func (u *DBService) SelectMany(dest interface{}, query interface{}, args ...interface{}) error {
	db := u.DataContext
	if query != nil {
		db = db.Where(query, args...)
	}

	if err := db.Find(dest).Error; err != nil {
		return err
	}
	return nil
}

func (u *DBService) InsertOrUpdate(dest interface{}) error {
	if err := u.DataContext.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(dest).Error; err != nil {
		return err
	}

	return nil
}

func (u *DBService) Insert(user interface{}) error {
	if err := u.DataContext.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *DBService) Update(user interface{}) error {
	r := u.DataContext.Model(user).Updates(user)
	if r.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	if err := r.Error; err != nil {
		return err
	}
	return nil
}

func (u *DBService) Delete(dest interface{}) error {
	r := u.DataContext.Delete(dest)
	if r.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	if err := r.Error; err != nil {
		return err
	}

	return nil
}
