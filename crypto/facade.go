package crypto

// Crypto provides object-style access to the cryptography functions in this package.
var Crypto = facade{}

type facade struct{}

func (facade) MD5(value string) string                  { return MD5(value) }
func (facade) Md5(value string) string                  { return Md5(value) }
func (facade) MD5V(value []byte, prefix ...byte) string { return MD5V(value, prefix...) }
func (facade) HashSHA256(value string) string           { return HashSHA256(value) }
func (facade) BcryptHash(password string) string        { return BcryptHash(password) }
func (facade) BcryptHashWithError(password string) (string, error) {
	return BcryptHashWithError(password)
}
func (facade) BcryptCheck(password, hash string) bool { return BcryptCheck(password, hash) }
func (facade) GenerateRSAKeys() (string, string, error) {
	return GenerateRSAKeys()
}
func (facade) EncryptRSA(publicKey string, message []byte) ([]byte, error) {
	return EncryptRSA(publicKey, message)
}
func (facade) DecryptRSA(privateKey string, ciphertext []byte) ([]byte, error) {
	return DecryptRSA(privateKey, ciphertext)
}
func (facade) Encrypt(key, message string) ([]byte, error) {
	return (&AESEncryptor{}).Encrypt(key, message)
}
func (facade) Decrypt(key string, ciphertext []byte) ([]byte, error) {
	return (&AESEncryptor{}).Decrypt(key, ciphertext)
}
