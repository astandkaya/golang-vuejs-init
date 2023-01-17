package services

import(
    "crypto/sha256"
    "encoding/hex"
)

type HashService struct {
}

func Hash() *HashService {
    return &HashService{
    }
}

func (r *HashService) Make(str string) string { 
    b := []byte(str)

    sha256 := sha256.Sum256(b)
    hashed := hex.EncodeToString(sha256[:])
 
    return hashed
}