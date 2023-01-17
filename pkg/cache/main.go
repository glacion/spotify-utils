package cache

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type cache struct {
	Path  string
	Token string
}

// Dependency injection?
// Never heard of 'er
var c cache

func GetToken() string {
	return c.Token
}

func SetToken(token string) error {
	c.Token = token
	return write()
}

func Init(path string) error {
	log.Debug().Msg("reading cache")
	c = cache{Path: path}

	err := ensureFile(path)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &c)
	if err != nil {
		return err
	}

	return nil
}

func write() error {
	log.Debug().Msg("writing cache")
	content, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.Path, content, os.ModePerm)
}

func ensureFile(path string) error {
	log.Debug().Str("path", path).Msg("ensuring cache storage")
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err == nil {
		// Write minimum applicable JSON
		_, err = file.WriteString("{}")
	}
	defer file.Close()
	return err
}
