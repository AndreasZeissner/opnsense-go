// Code generated by internal/generate/api/main.go; DO NOT EDIT.

package ipsec

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var VTIOpts = api.ReqOpts{
	AddEndpoint:         "/api/ipsec/vti/add",
	GetEndpoint:         "/api/ipsec/vti/get",
	UpdateEndpoint:      "/api/ipsec/vti/set",
	DeleteEndpoint:      "/api/ipsec/vti/del",
	ReconfigureEndpoint: ipsecReconfigureEndpoint,
	Monad:               "VTI",
}

// Data structs

type VTI struct {
	UUID      string          `json:"uuid"`
	Name      string          `json:"name"`
	Interface api.SelectedMap `json:"interface"`
	Local     string          `json:"local"`
	Remote    string          `json:"remote"`
	Enabled   bool            `json:"enabled"`
}

// CRUD operations

func (c *Controller) AddVTI(ctx context.Context, resource *VTI) (string, error) {
	return api.Add(c.Client(), ctx, VTIOpts, resource)
}

func (c *Controller) GetVTI(ctx context.Context, id string) (*VTI, error) {
	return api.GetFilter(c.Client(), ctx, VTIOpts, &VTI{}, id)
}

func (c *Controller) GetVTIAll(ctx context.Context) ([]VTI, error) {
	return api.GetAll(c.Client(), ctx, VTIOpts, []VTI{})
}

func (c *Controller) UpdateVTI(ctx context.Context, id string, resource *VTI) error {
	return api.Update(c.Client(), ctx, VTIOpts, resource, id)
}

func (c *Controller) DeleteVTI(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, VTIOpts, id)
}
