// Code generated by internal/generate/api/main.go; DO NOT EDIT.

package ipsec

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var SettingsOpts = api.ReqOpts{
	AddEndpoint:         "",
	GetEndpoint:         "/api/ipsec/settings/get",
	UpdateEndpoint:      "/api/ipsec/settings/set",
	DeleteEndpoint:      "",
	ReconfigureEndpoint: ipsecReconfigureEndpoint,
	Monad:               "Settings",
}

// Data structs

type Settings struct {
	Enabled             bool `json:"enabled"`
	NATTraversal        bool `json:"nat_traversal"`
	NATTraversalPort    int  `json:"nat_traversal_port"`
	EnableCompression   bool `json:"enable_compression"`
	EnableRedirect      bool `json:"enable_redirect"`
	EnableRedirectRoute bool `json:"enable_redirect_route"`
}

// CRUD operations

func (c *Controller) AddSettings(ctx context.Context, resource *Settings) (string, error) {
	return api.Add(c.Client(), ctx, SettingsOpts, resource)
}

func (c *Controller) GetSettings(ctx context.Context, id string) (*Settings, error) {
	return api.Get(c.Client(), ctx, SettingsOpts, &Settings{}, id)
}

func (c *Controller) GetSettingsAll(ctx context.Context) ([]Settings, error) {
	return api.GetAll(c.Client(), ctx, SettingsOpts, []Settings{})
}

func (c *Controller) UpdateSettings(ctx context.Context, id string, resource *Settings) error {
	return api.Update(c.Client(), ctx, SettingsOpts, resource, id)
}

func (c *Controller) DeleteSettings(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, SettingsOpts, id)
}
