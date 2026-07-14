package config

import (
	"github.com/beatenstratblues/maple-workbench/pkg/random"
)

type config struct {
	HostUrl           string
	DbUserName        string
	DbPassword        string
	AdminUsername     string
	AdminUserEmail    string
	AdminUserPassword string
}

func Load() (*config, error) {
	adminPassword, err := random.Generator(10)
	if err != nil {
		return nil, err
	}
	return &config{
		HostUrl:           "",
		DbUserName:        "",
		DbPassword:        "",
		AdminUsername:     "admin",
		AdminUserEmail:    "",
		AdminUserPassword: adminPassword,
	}, nil
}
