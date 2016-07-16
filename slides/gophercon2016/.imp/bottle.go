package main

import (
	"github.com/goadesign/goa"
	"github.com/raphael/gophercon2016/cellar/app"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Create runs the create action.
func (c *BottleController) Create(ctx *app.CreateBottleContext) error {
	// TBD: implement
	return nil
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// TBD: implement
	res := &app.Bottle{
		ID:      ctx.ID,
		Name:    "gophercon",
		Vintage: 2016,
	}
	return ctx.OK(res)
}
