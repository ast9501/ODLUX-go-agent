package internal

import (
	"time"
	. "agent/model"
)

func GenerateToken() *Token{
	token := &Token{}
	// Fixed token
	token.Token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJhZG1pbkBzZG4iLCJyb2xlcyI6WyJ1c2VyIiwiYWRtaW4iXSwiaXNzIjoiT05BUC1TRE5DIiwibmFtZSI6ImFkbWluQHNkbiIsImV4cCI6MTY0NjY1MjY1NywiZmFtaWx5X25hbWUiOiIifQ.9VBxip13i9k-zSg_BqFuhq5SR6_ANlk_hmLnjJLbUGQ"
	token.ExpiredTime = time.Now().Unix() + 2000
    token.Type = "Bearer"
    return token
}