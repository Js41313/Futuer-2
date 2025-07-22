package constant

import (
	"encoding/json"
)

// Used for type cloning conversion
const (
	Int64   int64  = 0
	Uint32  uint32 = 0
	DevMode        = "dev"
)

// VerifyType is the type of verification code
type VerifyType uint8

const (
	Register VerifyType = iota + 1
	Security
)

func ParseVerifyType(i uint8) VerifyType {
	return VerifyType(i)
}

func (v VerifyType) String() string {
	switch v {
	case Register:
		return "register"
	case Security:
		return "security"
	default:
		return "unknown"
	}
}

// TempOrderCacheKey Cache to Redis Key
// eg: temp_order:order_no
const TempOrderCacheKey = "temp_order:%s"

type TemporaryOrderInfo struct {
	OrderNo    string `json:"order_no"`
	Identifier string `json:"identifier"`
	AuthType   string `json:"auth_type"`
	Password   string `json:"password"`
	InviteCode string `json:"invite_code,omitempty"`
}

func (t TemporaryOrderInfo) Marshal() string {
	value, _ := json.Marshal(t)
	return string(value)
}
