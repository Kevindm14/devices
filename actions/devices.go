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

	verrs, err := tx.ValidateAndCreate(device)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("device", device)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("devices/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Device was created successfully")

	return c.Redirect(302, "/devices/%s", device.ID)
}

func DevicesNew(c buffalo.Context) error {
	c.Set("device", models.Device{})
	return c.Render(http.StatusOK, r.HTML("devices/new.plush.html"))
}

func DevicesDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	device := &models.Device{}

	if err := tx.Find(device, c.Param("device_id")); err != nil {
		return c.Error(404, err)
	}

	str, err := base64.StdEncoding.DecodeString(device.Image)
	if err != nil {
		fmt.Println(err)
	}

	device.Image = string(str)

	c.Set("device", device)
	return c.Render(200, r.HTML("devices/show"))
}

func DevicesEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	device := &models.Device{}

	if err := tx.Find(device, c.Param("device_id")); err != nil {
		return c.Error(404, err)
	}

	c.Set("device", device)
	return c.Render(http.StatusOK, r.HTML("devices/edit.plush.html"))
}

func DevicesUpdate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	device := &models.Device{}

	if err := tx.Find(device, c.Param("device_id")); err != nil {
		return c.Error(404, err)
	}

	if err := c.Bind(device); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(device)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("device", device)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("devices/edit.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Device was updated successfully")
	return c.Redirect(302, "/devices/%s", device.ID)
}

func DevicesDestroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	device := &models.Device{}

	if err := tx.Find(device, c.Param("device_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(device); err != nil {
		return err
	}

	c.Flash().Add("success", "device was destroyed successfully")
	// Redirect to the devices page
	return c.Redirect(302, "/devices")
}
