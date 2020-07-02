package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"github.com/setiadijoe/fullstack/app/auth"
	"github.com/setiadijoe/fullstack/app/helpers"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeSave ...
func (u *User) BeforeSave() error {
	hashedPassword, err := helpers.Hash(u.Password)
	if nil != err {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare ...
func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
}

// Validate ...
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("required_nickname")
		}
		if u.Password == "" {
			return errors.New("required_password")
		}
		if u.Email == "" {
			return errors.New("required_email")
		}
		if err := checkmail.ValidateFormat(u.Email); nil != err {
			return errors.New("invalid_email")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("required_password")
		}
		if u.Email == "" {
			return errors.New("required_email")
		}
		if err := checkmail.ValidateFormat(u.Email); nil != err {
			return errors.New("invalid_email")
		}
		return nil
	default:
		if u.Nickname == "" {
			return errors.New("required_nickname")
		}
		if u.Password == "" {
			return errors.New("required_password")
		}
		if u.Email == "" {
			return errors.New("required_email")
		}
		if err := checkmail.ValidateFormat(u.Email); nil != err {
			return errors.New("invalid_email")
		}
		return nil
	}
}

// Login ...
func (u *User) Login(db *gorm.DB, email, password string) (result interface{}, err error) {
	err = db.Debug().Model(u).Where("email = ?", email).Take(&u).Error
	if nil != err {
		return nil, err
	}
	err = helpers.VerifyPassword(u.Password, password)
	if nil != err && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	return auth.CreateToken(u.ID)
}

// Save ...
func (u *User) Save(db *gorm.DB) (user *User, err error) {
	err = db.Debug().Create(&u).Error
	if nil != err {
		return nil, err
	}
	return u, nil
}

// FindAll ...
func (u *User) FindAll(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if nil != err {
		return nil, err
	}

	return &users, nil
}

// FindByID ...
func (u *User) FindByID(db *gorm.DB, id uint32) (*User, error) {
	err := db.Debug().Model(User{}).Where("id = ?", id).Take(&u).Error
	if nil != err {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user_not_found")
	}
	return u, nil
}

// UpdateByID ...
func (u *User) UpdateByID(db *gorm.DB, id uint32, email, nickname, password string) (*User, error) {
	err := u.BeforeSave()
	if nil != err {
		return nil, err
	}

	err = db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Updates(
		User{
			Email:     email,
			Nickname:  nickname,
			Password:  password,
			UpdatedAt: time.Now().UTC(),
		},
	).Error

	if nil != err {
		return nil, err
	}

	err = db.Debug().Model(&u).Where("id = ?", id).Take(&u).Error
	if nil != err {
		return nil, err
	}

	return u, nil

}

// DeleteByID ...
func (u *User) DeleteByID(db *gorm.DB, id uint32) (int64, error) {
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).Delete(&User{}).Error

	if nil != err {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
