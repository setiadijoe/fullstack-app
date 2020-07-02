package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/setiadijoe/fullstack/app/internal/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

// Load ...
func Load(db *gorm.DB) {
	ok := db.Debug().HasTable(&models.User{})
	ok = db.Debug().HasTable(&models.Post{})

	err := db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		db.Debug().Rollback()
		log.Fatalf("cannot migrate table: %v", err)
	}

	if !ok {
		for i := range users {
			err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)
			}
			posts[i].AuthorID = users[i].ID

			err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
			if err != nil {
				log.Fatalf("cannot seed posts table: %v", err)
			}
		}
	}

}
