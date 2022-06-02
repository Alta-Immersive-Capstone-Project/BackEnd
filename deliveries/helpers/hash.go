package helpers

import (
	"crypto/sha512"
	"encoding/hex"
	"kost/configs"
	"kost/entities"
)

func Hash512(req entities.Callback) bool {
	key := req.OrderID + req.StatusCode + req.GrossAmount + configs.Get().Payment.MidtransServerKey
	hash := sha512.New()
	hash.Write([]byte(key))
	hash_key := hex.EncodeToString(hash.Sum(nil))
	hash_value := req.SignatureKey

	if hash_key == hash_value {
		return !false
	}

	return false
}
