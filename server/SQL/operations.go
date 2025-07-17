//CRUD/operations.go
package SQL

import(
	"github.com/jasimpn/trace/models"
    "gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.User) error {
    result := db.Create(&user)

    return result.Error
    
}

func CreateStorageFolder(db *gorm.DB, storage models.Storage) error {
    result := db.Create(&storage)

    return result.Error
    
}

func CreateRepo(db *gorm.DB, repository models.Repo) error {
    result := db.Create(&repository)

    return result.Error
    
}

