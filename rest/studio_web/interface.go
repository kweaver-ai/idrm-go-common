package studio_web

import "context"

type Driven interface {
	InsertWebApps(ctx context.Context, req []*WebAppConfig) error
}

type App struct {
	TextZHCN      string `json:"textZHCN"`
	TextZHTW      string `json:"textZHTW"`
	TextENUS      string `json:"textENUS"`
	Icon          string `json:"icon,omitempty"`
	Type          string `json:"type"`
	IsDefaultOpen bool   `json:"isDefaultOpen,omitempty"`
}

type SubApp struct {
	Children   map[string]any `json:"children,omitempty"`
	Entry      string         `json:"entry,omitempty"`
	ActiveRule string         `json:"activeRule,omitempty"`
	BaseRouter string         `json:"baseRouter,omitempty"`
}

type WebAppConfig struct {
	Name       string `json:"name"`
	Parent     string `json:"parent"`
	OrderIndex int    `json:"orderIndex"`
	App        App    `json:"app"`
	SubApp     SubApp `json:"subapp"`
}
