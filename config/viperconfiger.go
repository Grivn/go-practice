package config

import (
	"io"
	"log"
	"math"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// ViperConfiger implements Configer through viper
type ViperConfiger struct {
	config *viper.Viper
}

// New creates a new instance of viper configer
func New() *ViperConfiger {
	return &ViperConfiger{}
}

// Peer represents an entry in the list of peers in configuration
type Peer struct {
	ID   int
	Addr string
}

// LoadConfig parses the config file
func (c *ViperConfiger) LoadConfig(filePath string) {
	confPath, confName := filepath.Split(filePath)
	ext := filepath.Ext(confName)
	name := confName[:len(confName)-len(ext)]

	if len(ext) < 2 || len(name) == 0 {
		panic("Invalid config file name (should specify file extension for parsing).")
	}

	config := viper.New()
	config.SetConfigName(name)
	config.SetConfigType(ext[1:]) // exclude the char '.'
	config.AddConfigPath(confPath)
	config.AddConfigPath("./") // support relative path
	err := config.ReadInConfig()
	if err != nil {
		panic("Failed to parse config file: " + err.Error())
	}

	c.config = config
}

// ReadConfig reads configuration from the specified io.Reader parsing
// it according to the specified type supported by
// github.com/spf13/viper package, e.g. "yaml"
func (c *ViperConfiger) ReadConfig(in io.Reader, cfgType string) error {
	config := viper.New()
	config.SetConfigType(cfgType)
	err := config.ReadConfig(in)
	if err != nil {
		return err
	}
	c.config = config
	return nil
}

// IsInitialized returns true if the config file is correctly parsed
func (c *ViperConfiger) IsInitialized() bool {
	return nil != c.config
}

//============ Helpers ============

// getUint32 checks the boundary of the config value and returns uint32
func (c *ViperConfiger) getUint32(key string) uint32 {
	n := c.config.GetInt64(key)
	if n < int64(0) || n > math.MaxUint32 {
		panic("interger overflow: " + key)
	}
	return uint32(n)
}

// getTimeDuration returns time duration of the config value
func (c *ViperConfiger) getTimeDuration(key string) time.Duration {
	return c.config.GetDuration(key)
}

//============ Configer Interface =============

// N returns the total number of replicas
func (c *ViperConfiger) N() uint32 {
	return c.getUint32("protocol.n")
}

// Peers returns a list peers
func (c *ViperConfiger) Peers() []Peer {
	peers := []Peer{}
	err := c.config.UnmarshalKey("peers", &peers)
	if err != nil {
		log.Print("Failed to unmarshal peers:", err)
		return nil
	}
	return peers
}
