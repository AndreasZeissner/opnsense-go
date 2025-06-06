// Code generated by internal/generate/api/main.go; DO NOT EDIT.

package ipsec

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var ConnectionOpts = api.ReqOpts{
	AddEndpoint:         "/api/ipsec/connections/add",
	GetEndpoint:         "/api/ipsec/connections/get",
	UpdateEndpoint:      "/api/ipsec/connections/set",
	DeleteEndpoint:      "/api/ipsec/connections/del",
	ReconfigureEndpoint: ipsecReconfigureEndpoint,
	Monad:               "Connection",
}

// Data structs

type LocalConfig struct {
	Address        string          `json:"address"`
	Port           string          `json:"port"`
	Authentication api.SelectedMap `json:"authentication"`
	Certificate    api.SelectedMap `json:"certificate"`
}

type RemoteConfig struct {
	Address        string          `json:"address"`
	Port           string          `json:"port"`
	Authentication api.SelectedMap `json:"authentication"`
	Certificate    api.SelectedMap `json:"certificate"`
}

type Connection struct {
	UUID     string       `json:"uuid"`
	Name     string       `json:"name"`
	Enabled  bool         `json:"enabled"`
	Local    LocalConfig  `json:"local"`
	Remote   RemoteConfig `json:"remote"`
	Children []interface{} `json:"children"`
}

// CRUD operations

func (c *Controller) AddConnection(ctx context.Context, resource *Connection) (string, error) {
	return api.Add(c.Client(), ctx, ConnectionOpts, resource)
}

func (c *Controller) GetConnection(ctx context.Context, id string) (*Connection, error) {
	return api.GetFilter(c.Client(), ctx, ConnectionOpts, &Connection{}, id)
}

func (c *Controller) GetConnectionAll(ctx context.Context) ([]Connection, error) {
	return api.GetAll(c.Client(), ctx, ConnectionOpts, []Connection{})
}

func (c *Controller) UpdateConnection(ctx context.Context, id string, resource *Connection) error {
	return api.Update(c.Client(), ctx, ConnectionOpts, resource, id)
}

func (c *Controller) DeleteConnection(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, ConnectionOpts, id)
}
