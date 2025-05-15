package course

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
)

type (
	Repository interface {
		Create(course *Course) error
		GetById(id string) (*Course, error)
		GetAll(filters Filters, offset, limit int) ([]Course, error)
		Update(id string, updateRequest UpdateRequest) error
		Delete(id string) error
		Count(filters Filters) (int, error)
	}

	repo struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (r *repo) Create(course *Course) error {
	if err := r.db.Create(course).Error; err != nil {
		return err
	}
	r.log.Println("Course created:", course)
	return nil
}
func (r *repo) GetById(id string) (*Course, error) {
	var course Course
	if err := r.db.First(&course, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *repo) GetAll(filters Filters, offset, limit int) ([]Course, error) {
	var courses []Course
	query := r.db.Model(&Course{})
	query = applyFilters(query, filters)
	if err := query.Order("created_at desc").Limit(limit).Offset(offset).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *repo) Count(filters Filters) (int, error) {
	var count int64
	db := r.db.Model(&Course{})
	db = applyFilters(db, filters)
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *repo) Update(id string, updateRequest UpdateRequest) error {
	res := r.db.
		Model(&Course{}).
		Where("id = ?", id).
		Updates(updateRequest)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *repo) Delete(id string) error {
	var course Course
	if err := r.db.First(&course, "id = ?", id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&course).Error; err != nil {
		return err
	}
	return nil
}

func applyFilters(db *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		db = db.Where("lower(name) LIKE ?", filters.Name)
	}
	return db
}
