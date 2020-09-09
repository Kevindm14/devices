package actions

import (
	"devices/models"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

func DevicesIndex(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	devices := &models.Devices{}

	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	q := tx.PaginateFromParams(c.Params())
	if err := q.All(devices); err != nil {
		return errors.WithStack(err)
	}

	c.Set("pagination", q.Paginator)
	c.Set("devices", devices)
	return c.Render(http.StatusOK, r.HTML("devices/index.html"))
}

func DevicesCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	device := &models.Device{}

	if err := c.Bind(device); err != nil {
		return err
	}

	image := c.Param("Image")
	device.Image = base64.StdEncoding.EncodeToString([]byte(image))

	verrs, err := tx.ValidateAndCreate(device)

	fmt.Println("-----------------")
	fmt.Println("device:", device)
	fmt.Println("------------------")

	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("device", device)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("devices/newDevice.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Device was created successfully")

	return c.Redirect(302, "/devices/")
}

func DevicesNew(c buffalo.Context) error {
	c.Set("device", models.Device{})
	return c.Render(http.StatusOK, r.HTML("devices/newDevice.html"))
}
