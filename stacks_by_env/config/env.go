package config

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Env interface {
	Identifier() EnvName
	Config() *Config
}

type EnvNameFormatter interface {
	ToString() string
	ToPascalCase() string
}

type EnvName string
type EnvShortName string

type EnvDevelop struct{}
type EnvProduction struct{}

const (
	EnvNameDevelop         EnvName      = "develop"
	EnvNameProduction      EnvName      = "production"
	EnvShortNameDevelop    EnvShortName = "dev"
	EnvShortNameProduction EnvShortName = "prod"
)

func (e EnvDevelop) Identifier() EnvName    { return EnvNameDevelop }
func (e EnvProduction) Identifier() EnvName { return EnvNameProduction }

func (e EnvName) ToString() string {
	return string(e)
}

func (e EnvName) ToPascalCase() string {
	return cases.Title(language.English).String(string(e))
}

func (e EnvShortName) ToString() string {
	return string(e)
}

func (e EnvShortName) ToPascalCase() string {
	return cases.Title(language.English).String(string(e))
}

func (e EnvName) ShortName() EnvNameFormatter {
	switch e {
	case EnvDevelop{}.Identifier():
		return EnvShortNameDevelop
	case EnvProduction{}.Identifier():
		return EnvShortNameProduction
	default:
		return nil
	}
}

func NewEnv(identifier string) Env {
	switch EnvName(identifier) {
	case EnvDevelop{}.Identifier():
		return EnvDevelop{}
	case EnvProduction{}.Identifier():
		return EnvProduction{}
	default:
		return nil
	}
}
