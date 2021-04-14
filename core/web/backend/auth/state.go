package auth

import (
	"GoBot-Recode/core/logger"
	"GoBot-Recode/core/web/backend/auth/flow"
	"math/rand"
	"time"
)

func RandomState() {
	logger.LogModuleNoNewline(logger.TypeInfo, "GoBot/Web", "Generating random State...")
	length := 16
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	flow.State = string(b)
	logger.AppendDone()
}