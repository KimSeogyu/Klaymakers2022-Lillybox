package http

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"lillybox-backend/internal/database"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// VerifySignature ...
func VerifySignature(h *Handlers, bytes []byte) (bool, error) {
	body := ParseBytesToMapObject(bytes)
	if body == nil {
		return false, fiber.ErrBadRequest
	}
	if _, ok := body[ADDR]; !ok {
		return false, fiber.ErrBadRequest
	}
	if _, ok := body[SIGNATURE]; !ok {
		return false, fiber.ErrBadRequest
	}
	user, err := h.Database.ReadUser(body[ADDR])
	if err != nil {
		return false, fiber.ErrBadRequest
	}
	message := lillySignature + user.LoginRequestID
	str := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(str)
	sigDecoded := hexutil.MustDecode(body[SIGNATURE])
	if sigDecoded[64] == 27 || sigDecoded[64] == 28 {
		sigDecoded[64] -= 27
	}
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), sigDecoded)
	if sigPublicKeyECDSA == nil || err != nil {
		return false, fiber.ErrBadRequest
	}
	result := crypto.PubkeyToAddress(*sigPublicKeyECDSA).String()
	if strings.ToLower(string(result)) != body[ADDR] {
		return false, fiber.ErrForbidden
	}
	return true, nil
}

// GetOrInsertUser ...
func GetOrInsertUser(h *Handlers, bytes []byte) (*database.Users, error) {
	body := ParseBytesToMapObject(bytes)
	if body == nil {
		return nil, fiber.ErrBadRequest
	}
	if ok := KeyCheck(body, ADDR); !ok {
		return nil, fiber.ErrBadRequest
	}
	resp, err := h.Database.ReadUser(body[ADDR])
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, gorm.ErrRecordNotFound
	case errors.Is(err, nil):
		if _, err := h.Database.UpdateUser(resp.Address); err != nil {
			return nil, err
		}
		if resp, err := h.Database.ReadUser(resp.Address); err == nil {
			return resp, nil
		}
		return nil, err
	default:
		return nil, err
	}
}

// ParseBytesToMapObject ...
func ParseBytesToMapObject(bytes []byte) map[string]string {
	var object map[string]string
	if err := json.Unmarshal(bytes, &object); err != nil {
		return nil
	}
	return object
}
