package email

import "github.com/perfect-panel/server/internal/types"

type Platform int

const (
	SMTP Platform = iota
	unsupported
)

var platformNames = map[string]Platform{
	"smtp":        SMTP,
	"unsupported": unsupported,
}

func (p Platform) String() string {
	for k, v := range platformNames {
		if v == p {
			return k
		}
	}
	return "unsupported"
}

func parsePlatform(s string) Platform {
	if p, ok := platformNames[s]; ok {
		return p
	}
	return unsupported
}

func GetSupportedPlatforms() []types.PlatformInfo {
	return []types.PlatformInfo{
		{
			Platform:    SMTP.String(),
			PlatformUrl: "",
			PlatformFieldDescription: map[string]string{
				"host": "host",
				"port": "port",
				"user": "user",
				"pass": "pass",
				"from": "from",
				"ssl":  "ssl",
			},
		},
	}
}
