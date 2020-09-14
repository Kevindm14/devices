package actions

import (
	"devices/models"
	"fmt"
)

func (as *ActionSuite) Test_Users_Index() {
	as.LoadFixture("lots of users")
	res := as.HTML("/users").Get()

	body := res.Body.String()
	as.Contains(body, "kevin")
	as.Contains(body, "kevin")
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

func (as *ActionSuite) Test_Users_Destroy() {
	user := &models.User{
		FirstName:    "Kevin",
		LastName:     "Diaz",
		Email:        "kevindiaz@gmail.com",
		Role:         "RegularUser",
		ManagerEmail: "admin@manager.com",
	}

	as.DB.Create(user)
	res := as.HTML("/users/%s", user.ID).Delete()
	as.Equal(302, res.Code)
	as.Equal("/users", res.Location())
}

