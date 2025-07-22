package migrate

import (
	"time"

	"github.com/Js41313/Futuer-2/internal/model/user"
	"github.com/Js41313/Futuer-2/pkg/logger"
	"github.com/Js41313/Futuer-2/pkg/tool"
	"github.com/Js41313/Futuer-2/pkg/uuidx"
	"gorm.io/gorm"
)

// CreateAdminUser create admin user
func CreateAdminUser(email, password string, tx *gorm.DB) error {
	enable := true
	return tx.Transaction(func(tx *gorm.DB) error {
		// Prevent duplicate creation
		if tx.Model(&user.User{}).Find(&user.User{}).RowsAffected != 0 {
			logger.Info("User already exists, skip creating administrator account")
			return nil
		}

		u := user.User{
			Password:  tool.EncodePassWord(password),
			IsAdmin:   &enable,
			ReferCode: uuidx.UserInviteCode(time.Now().Unix()),
		}
		if err := tx.Model(&user.User{}).Save(&u).Error; err != nil {
			return err
		}
		method := user.AuthMethods{
			UserId:         u.Id,
			AuthType:       "email",
			AuthIdentifier: email,
			Verified:       true,
		}
		if err := tx.Model(&user.AuthMethods{}).Save(&method).Error; err != nil {
			return err
		}
		return nil
	})
}
