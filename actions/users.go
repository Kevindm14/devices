package actions

import (
    "net/http"
	"devices/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

func UsersNew(c buffalo.Context) error {
	c.Set("user", &models.User{})
	return c.Render(http.StatusOK, r.HTML("users/new.html"))
}

func UsersIndex(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	users := &models.Users{}

	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	q := tx.PaginateFromParams(c.Params())
	if err := q.All(users); err != nil {
		return errors.WithStack(err)
	}

	c.Set("pagination", q.Paginator)
	c.Set("users", users)
	return c.Render(http.StatusOK, r.HTML("users/index.html"))
}

func UsersCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	managerEmail := c.Param("ManagerEmail")
	if managerEmail == "" {
		user.ManagerEmail = "No"
	}

	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("user", user)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("users/new.plush.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "User was created successfully")

	return c.Redirect(302, "/users/%s", user.ID)
}

func UsersDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	c.Set("user", user)
	return c.Render(200, r.HTML("users/show"))
}

func UsersDestroy(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(user); err != nil {
		return err
	}

	c.Flash().Add("success", "user was destroyed successfully")
	// Redirect to the devices page
	return c.Redirect(302, "/users")
}

