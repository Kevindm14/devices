package actions

import (
	"devices/models"
	"fmt"
	"net/url"
)

func (as *ActionSuite) CreateUser() *models.User {
	user := &models.User{
		FirstName:    "Kevin",
		LastName:     "Diaz",
		Email:        "kevindiaz@gmail.com",
		Role:         "RegularUser",
		ManagerEmail: "admin@manager.com",
	}
	verrs, err := as.DB.ValidateAndCreate(user)
	as.NoError(err)
	as.False(verrs.HasAny())
	return user
}

func (as *ActionSuite) Test_Users_Index() {
	as.LoadFixture("lots of users")
	res := as.HTML("/users").Get()

	body := res.Body.String()
	as.Contains(body, "kevin")
	as.Contains(body, "kevin")
}

func (as *ActionSuite) Test_Users_show() {
	user := as.CreateUser()

	res := as.HTML("/users/%s", user.ID).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_users_New() {
	res := as.HTML("/users/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_users_Edit() {
	user := as.CreateUser()
	res := as.HTML("/users/%s/edit", user.ID).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
	user := models.User{
		FirstName:    "Kevin",
		LastName:     "Diaz",
		Email:        "kevindiaz@gmail.com",
		Role:         "RegularUser",
		ManagerEmail: "admin@manager.com",
	}

	res := as.HTML("/users").Post(&user)
	as.DB.Last(&user)
	as.Equal(302, res.Code)

	as.Equal(fmt.Sprintf("/users/%s", user.ID), res.Location())
	as.NotZero(user.ID)
}

func (as *ActionSuite) Test_Users_Update() {
	user := as.CreateUser()

	userNew := url.Values{
		"FirstName":    []string{"Kevin"},
		"LastName":     []string{"Diaz"},
		"Email":        []string{"kevin@gmail.com"},
		"Role":         []string{"RegularUser"},
		"ManagerEmail": []string{"admin@manager.com"},
	}

	res := as.HTML("/users/%s", user.ID).Put(userNew)
	as.Equal(302, res.Code)

	as.NoError(as.DB.Reload(user))
	as.Equal("kevin@gmail.com", user.Email)
}

func (as *ActionSuite) Test_Users_Destroy() {
	user := as.CreateUser()

	res := as.HTML("/users/%s", user.ID).Delete()
	as.Equal(302, res.Code)
	as.Equal("/users", res.Location())
}
