package email

import (
	"encoding/json"
	"fmt"

	"github.com/Js41313/Futuer-2/pkg/email/smtp"
	"github.com/Js41313/Futuer-2/pkg/logger"
)

type Sender interface {
	Send(to []string, subject, body string) error
}

func NewSender(platform, config, siteName string) (Sender, error) {
	switch parsePlatform(platform) {
	case SMTP:
		cfg := smtp.Config{}
		if err := json.Unmarshal([]byte(config), &cfg); err != nil {
			logger.Error("unmarshal email config failed", logger.Field("error", err.Error()), logger.Field("config", config))
			return nil, err
		}
		cfg.SiteName = siteName
		return smtp.NewClient(&cfg), nil
	default:
		return nil, fmt.Errorf("unsupported platform: %s", platform)
	}
}
