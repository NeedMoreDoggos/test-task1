package config

import (
	"errors"
	"fmt"
)

type Config struct {
	DBConfig DBConfig `yaml:"db" json:"db"`
}

type DBConfig struct {
	Type     string `yaml:"type" json:"type"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`

	//Для специфичных настроек бд. Потом можно создать
	//структуру аля pgMeta и забирать тайп асертом
	Meta interface{} `yaml:"meta" json:"meta"`
}

type PostgresMeta struct {
	SSLMode        string `yaml:"sslmode" json:"sslmode"`
	MaxConnections int    `yaml:"max_connections" json:"max_connections"`
}

var (
	ErrUnknownDatabase = errors.New("unknown database")
)

func (c *Config) GetDBConfig() (string, error) {
	switch c.DBConfig.Type {
	case "mysql":
		return c.getMySQLConfig()
	case "postgres":
		return c.getPostgresConfig()
	default:
		return "", ErrUnknownDatabase
	}
}

func (c *Config) GetMeta() *PostgresMeta {
	fn := "config.GetMeta"

	meta, ok := c.DBConfig.Meta.(map[string]interface{})
	if !ok {
		fmt.Println(c.DBConfig.Meta)
		panic("implent me if not ok" + fn)
	}

	maxConn, ok := meta["max_connections"].(int)
	if !ok {
		panic("implent me if not ok maxConn " + fn)
	}

	sslMode, ok := meta["sslmode"].(string)
	if !ok {
		panic("implent me if not ok sslMode " + fn)
	}

	return &PostgresMeta{
		SSLMode:        sslMode,
		MaxConnections: maxConn,
	}
}

func (c *Config) getPostgresConfig() (string, error) {
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		c.DBConfig.Host,
		c.DBConfig.Port,
		c.DBConfig.User,
		c.DBConfig.Password,
		c.DBConfig.Database,
	)

	return dns, nil
}

func (c *Config) getMySQLConfig() (string, error) {
	panic("not implemented")
}
