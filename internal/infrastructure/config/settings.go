package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fluxa-cli/internal/domain"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Settings struct {
	v    *viper.Viper
	path string
}

func New() *Settings {
	v := viper.New()
	v.SetEnvPrefix("FLUXA")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()
	v.SetDefault("base-url", "https://fluxa.nuculabs.dev")
	v.SetDefault("output", "table")
	return &Settings{v: v}
}

func (s *Settings) BindFlags(flags *pflag.FlagSet) error {
	for _, key := range []string{"base-url", "entity", "output"} {
		if err := s.v.BindPFlag(key, flags.Lookup(key)); err != nil {
			return err
		}
	}
	return nil
}

func (s *Settings) Load(override string) error {
	path := override
	if path == "" {
		dir, err := os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("find config directory: %w", err)
		}
		path = filepath.Join(dir, "fluxa", "config.yaml")
	}
	s.path = path
	s.v.SetConfigFile(path)
	if err := s.v.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("read config: %w", err)
	}
	if s.Output() != "table" && s.Output() != "json" {
		return errors.New("output must be table or json")
	}
	return nil
}

func (s *Settings) BaseURL() string     { return s.v.GetString("base-url") }
func (s *Settings) EntityID() domain.ID { return domain.ID(s.v.GetInt("entity")) }
func (s *Settings) Output() string      { return s.v.GetString("output") }
func (s *Settings) Path() string        { return s.path }

func (s *Settings) SaveEntity(id domain.ID) error  { return s.save("entity", id) }
func (s *Settings) SaveBaseURL(value string) error { return s.save("base-url", value) }

func (s *Settings) save(key string, value any) error {
	if s.path == "" {
		return errors.New("configuration is not loaded")
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0700); err != nil {
		return err
	}
	w := viper.New()
	w.SetConfigFile(s.path)
	if err := w.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		return err
	}
	w.Set(key, value)
	if err := w.WriteConfigAs(s.path); err != nil {
		return err
	}
	if err := os.Chmod(s.path, 0600); err != nil {
		return err
	}
	s.v.Set(key, value)
	return nil
}
