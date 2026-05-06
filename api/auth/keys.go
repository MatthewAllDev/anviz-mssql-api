package auth

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type APIKeyStore struct {
	entries []apiKeyEntry
}

type apiKeyEntry struct {
	scope Scope
	value string
	hash  bool
}

func LoadAPIKeys(envValue, filename string) (*APIKeyStore, error) {
	if strings.TrimSpace(filename) != "" {
		content, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		return parseAPIKeys(string(content), true)
	}
	if strings.TrimSpace(envValue) == "" {
		return nil, errors.New("API_KEYS_FILE or API_KEYS is required")
	}
	return parseAPIKeys(envValue, false)
}

func (s *APIKeyStore) ScopeFor(token string) (Scope, bool) {
	for _, entry := range s.entries {
		if entry.matches(token) {
			return entry.scope, true
		}
	}
	return "", false
}

func (e apiKeyEntry) matches(token string) bool {
	if e.hash {
		return bcrypt.CompareHashAndPassword([]byte(e.value), []byte(token)) == nil
	}
	return subtle.ConstantTimeCompare([]byte(e.value), []byte(token)) == 1
}

func parseAPIKeys(value string, hashed bool) (*APIKeyStore, error) {
	store := &APIKeyStore{}
	for _, item := range strings.FieldsFunc(value, func(r rune) bool { return r == ',' || r == ';' || r == '\n' }) {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		entry, err := parseAPIKeyEntry(item, hashed)
		if err != nil {
			return nil, err
		}
		store.entries = append(store.entries, entry)
	}
	if len(store.entries) == 0 {
		return nil, errors.New("no API keys configured")
	}
	return store, nil
}

func parseAPIKeyEntry(item string, hashed bool) (apiKeyEntry, error) {
	parts := strings.SplitN(item, ":", 2)
	if len(parts) != 2 {
		return apiKeyEntry{}, fmt.Errorf("invalid API key entry %q", item)
	}

	secret := strings.TrimSpace(parts[0])
	scope := Scope(strings.TrimSpace(parts[1]))
	if secret == "" {
		return apiKeyEntry{}, errors.New("API key cannot be empty")
	}
	if scope != ReadScope && scope != CrudScope {
		return apiKeyEntry{}, fmt.Errorf("invalid API key scope %q", scope)
	}
	if hashed {
		if _, err := bcrypt.Cost([]byte(secret)); err != nil {
			return apiKeyEntry{}, fmt.Errorf("invalid API key hash: %w", err)
		}
	}

	return apiKeyEntry{
		scope: scope,
		value: secret,
		hash:  hashed,
	}, nil
}
