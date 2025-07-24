package pkg

import (
	"encoding/base32"
	"strings"

	"github.com/google/uuid"
)

func GenerateCalendarID() string {
	id := uuid.New()

	encoder := base32.HexEncoding.WithPadding(base32.NoPadding)

	encoded := encoder.EncodeToString(id[:]) 
	return strings.ToLower(encoded)   
}
