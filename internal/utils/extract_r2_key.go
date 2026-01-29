package utils

import "strings"

// Extrai chave da URL de um objeto salvo no bucket R2.
//
// Parameters:
// 	- stored ProfilePicture salva atualmente. Ex: "https://bucketname.r2.cloudflarestorage.com/users/abc.jpg"
//
// Returns:
// 	- string: chave extraída. Ex: "users/abc.jpg". Retorna string vazia se não for possível extrair.
func ExtractR2Key(stored string) string {
	stored = strings.TrimSpace(stored)
	if stored == "" {
		return ""
	}
	
	if strings.HasPrefix(stored, "users/") {
		return stored
	}
	
	if strings.HasPrefix(stored, "http://") || strings.HasPrefix(stored, "https://") {
		i := strings.Index(stored, "://")
		if i == -1 {
			return ""
		}
		rest := stored[i+3:]               // dominio/path
		slash := strings.Index(rest, "/")  // começo do path
		if slash == -1 {
			return ""
		}
		path := rest[slash+1:] // sem a "/" inicial
		return path
	}
	return ""
}