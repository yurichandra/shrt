package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yurichandra/shrt/model"
)

// URLRepository represent repository of URL model.
type URLRepository struct {
	db *gorm.DB
}

// Get return all available urls.
func (r *URLRepository) Get() []model.URL {
	urls := make([]model.URL, 0)
	r.db.Find(&urls)

	return urls
}

// Find return single url.
func (r *URLRepository) Find(id uint) model.URL {
	url := model.URL{}
	r.db.First(&url)

	return url
}

// Create create new data and will return error if occured.
func (r *URLRepository) Create(data *model.URL) error {
	return r.db.Create(&data).Error
}

// Update saving existing data and will return error if occured.
func (r *URLRepository) Update(data *model.URL) error {
	return r.db.Save(&data).Error
}

// Delete soft deletes existing data and will return error if occured.
func (r *URLRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&model.URL{}).Error
}
