package utils

import "golang.org/x/crypto/bcrypt"

// Hash
// This function is used to hash some word
func Hash(word string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(word), 14)
	return string(bytes), err
}

// CheckHash
// This function is used to check if the word matches the hashed word
func CheckHash(word, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(word))
	return err == nil
}
