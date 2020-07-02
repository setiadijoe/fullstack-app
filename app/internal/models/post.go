package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Post ...
type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare ...
func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = time.Now().UTC()
}

// Validate ...
func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("required_title")
	}
	if p.Content == "" {
		return errors.New("required_content")
	}
	if p.AuthorID < 1 {
		return errors.New("required_author")
	}
	return nil
}

// Save ...
func (p *Post) Save(db *gorm.DB) (post *Post, err error) {
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if nil != err {
		return nil, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if nil != err {
			return nil, err
		}
	}
	return p, nil
}

// FindAll ...
func (p *Post) FindAll(db *gorm.DB) (*[]Post, error) {
	posts := []Post{}
	err := db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error
	if nil != err {
		return nil, err
	}
	if len(posts) > 0 {
		for i := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return &posts, nil
}

// FindByID ...
func (p *Post) FindByID(db *gorm.DB, id uint64) (*Post, error) {
	err := db.Debug().Model(&Post{}).Where("id = ?", id).Take(&p).Error
	if nil != err {
		return nil, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

// UpdateByID ...
func (p *Post) UpdateByID(db *gorm.DB, id uint64, title, content string) (*Post, error) {
	err := db.Debug().Model(&Post{}).Where("id = ?", id).Updates(
		Post{
			Title:     title,
			Content:   content,
			UpdatedAt: time.Now().UTC(),
		},
	).Error
	if err != nil {
		return nil, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

// DeleteByID ...
func (p *Post) DeleteByID(db *gorm.DB, postID uint64, userID uint32) (int64, error) {
	err := db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", postID, userID).Take(&Post{}).Delete(&Post{}).Error

	if nil != err {
		if gorm.IsRecordNotFoundError(err) {
			return 0, errors.New("post_not_found")
		}
		return 0, err
	}
	return db.RowsAffected, nil
}
