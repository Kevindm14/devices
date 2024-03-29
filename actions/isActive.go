package actions

import (
	"fmt"
	"html/template"
)

func isActive(route string) template.HTML {
	if route == "/users/" {
		return template.HTML(fmt.Sprint(
			`<li class='item'><a class='active' href='<%= usersPath() %>'>Users</a></li>
			 <li class='item'><a href='<%= devicesPath() %>'>Devices</a></li>`))
	}

	return template.HTML(fmt.Sprint(
		`<li class='item'><a href='<%= usersPath() %>'>Users</a></li>
		 <li class='item'><a class='active' href='<%= devicesPath() %>'>Devices</a></li>`))
}
