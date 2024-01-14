package infisicalclient

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

func DecryptProjectKey(serviceToken string, tokenDetail GetServiceTokenDetailsResponse) (string, error) {
	lastDot := strings.LastIndex(serviceToken, ".")
	if lastDot == -1 {
		return "", fmt.Errorf("invalid service token")
	}
	serviceTokenSecret := serviceToken[lastDot+1:]

	encryptedKey := tokenDetail.EncryptedKey
	iv := tokenDetail.Iv
	tag := tokenDetail.Tag

	projectKey, err := decrypt(encryptedKey, iv, tag, serviceTokenSecret)
	if err != nil {
		return "", fmt.Errorf("error while decrypting project key: %w", err)
	}

	return projectKey, nil
}

func DecryptSecret(secret EncryptedSecret, projectKey string) (key string, value string, err error) {
	keyCiphertext := secret.SecretKeyCiphertext
	keyIV := secret.SecretKeyIV
	keyTag := secret.SecretKeyTag

	secretKey, err := decrypt(keyCiphertext, keyIV, keyTag, projectKey)
	if err != nil {
		return "", "", fmt.Errorf("error while decrypting secret key: %w", err)
	}

	valueCiphertext := secret.SecretValueCiphertext
	valueIV := secret.SecretValueIV
	valueTag := secret.SecretValueTag
	secretValue, err := decrypt(valueCiphertext, valueIV, valueTag, projectKey)

	if err != nil {
		return "", "", fmt.Errorf("error while decrypting secret value: %w", err)
	}

	return secretKey, secretValue, nil
}

func decrypt(ciphertext string, iv string, authTag string, key string) (string, error) {
	nonceBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", fmt.Errorf("error while decoding iv: %w", err)
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error while creating new cipher: %w", err)
	}
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("error while decoding ciphertext: %w", err)
	}
	tagBytes, err := base64.StdEncoding.DecodeString(authTag)
	if err != nil {
		return "", fmt.Errorf("error while decoding authTag: %w", err)
	}

	gcm, err := cipher.NewGCMWithNonceSize(block, len(nonceBytes))
	if err != nil {
		return "", fmt.Errorf("error while creating new GCM: %w", err)
	}

	// We are using nonceBytes in Open function directly replacing previously used nonce variable
	plainBytes, err := gcm.Open(nil, nonceBytes, append(ciphertextBytes, tagBytes...), nil)
	if err != nil {
		panic(err)
	}

	return string(plainBytes), nil
}
