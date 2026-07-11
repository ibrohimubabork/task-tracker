package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash = errors.New("invalid password hash")
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var defaultParams = Argon2Params{
	Memory:      19 * 1024, // 64 * 1024
	Iterations:  2,         // 3
	Parallelism: 1,         // 4
	SaltLength:  16,
	KeyLength:   32,
}

func HashPassword(password string) (string, error) {
	salt := make([]byte, defaultParams.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate password salt: %w", err)
	}
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultParams.Iterations,
		defaultParams.Memory,
		defaultParams.Parallelism,
		defaultParams.KeyLength,
	)

	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultParams.Memory,
		defaultParams.Iterations,
		defaultParams.Parallelism,
		saltEncoded,
		hashEncoded,
	)

	return encodedHash, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	params, salt, expectedHash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	actualHash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	if subtle.ConstantTimeCompare(actualHash, expectedHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedHash string) (Argon2Params, []byte, []byte, error) {
	var params Argon2Params

	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return params, nil, nil, ErrInvalidHash
	}

	if parts[1] != "argon2id" {
		return params, nil, nil, ErrInvalidHash
	}

	var version int

	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return params, nil, nil, ErrInvalidHash
	}

	if version != argon2.Version {
		return params, nil, nil, ErrInvalidHash
	}

	if _, err := fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&params.Memory,
		&params.Iterations,
		&params.Parallelism,
	); err != nil {
		return params, nil, nil, ErrInvalidHash
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return params, nil, nil, ErrInvalidHash
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return params, nil, nil, ErrInvalidHash
	}

	params.SaltLength = uint32(len(salt))
	params.KeyLength = uint32(len(expectedHash))

	if params.Memory == 0 ||
		params.Iterations == 0 ||
		params.Parallelism == 0 ||
		params.KeyLength == 0 {
		return params, nil, nil, ErrInvalidHash
	}

	return params, salt, expectedHash, nil
}
