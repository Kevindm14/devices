package actions

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gobuffalo/plush"
	"github.com/stretchr/testify/require"
)

func Test_isPlural(t *testing.T) {
	var res string
	var err error
	a := require.New(t)

	devices := `<%= isPlural(number, "device") %>`
	users := `<%= isPlural(number, "user") %>`
	tcases := []struct {
		Number         int
		ExpectedResult string
	}{
		{Number: 2, ExpectedResult: "devices"},
		{Number: 1, ExpectedResult: "device"},
		{Number: 0, ExpectedResult: "devices"},
		{Number: 2, ExpectedResult: "users"},
		{Number: 1, ExpectedResult: "user"},
		{Number: 0, ExpectedResult: "users"},
	}

	i := 0
	for _, tcase := range tcases {
		contextData := map[string]interface{}{
			"isPlural": isPlural,
			"number":   tcase.Number,
		}

		ctx := plush.NewContextWith(contextData)

		if strings.Contains(tcase.ExpectedResult, "device") {
			res, err = plush.Render(devices, ctx)
		} else if strings.Contains(tcase.ExpectedResult, "user") {
			res, err = plush.Render(users, ctx)
		}

		a.NoError(err)
		a.Equal(tcase.ExpectedResult, res)
		i++
	}
}

func Test_isActive(t *testing.T) {
	var res string
	var err error
	assert := require.New(t)

	linkUser := `<%= isActive("user") %>`
	linkDevice := `<%= isActive("device") %>`
	tcases := []struct {
		Route          string
		View           string
		ExpectedResult string
	}{
		{Route: "/users/new", ExpectedResult: "active"},
		{Route: "/devices/edit", ExpectedResult: "active"},
	}

	for _, tcase := range tcases {
		request, er := http.NewRequest("GET", tcase.Route, nil)
		assert.NoError(er)

		ctx := plush.NewContextWith(map[string]interface{}{
			"request":  request,
			"isActive": isActive,
		})

		if strings.Contains(tcase.Route, "users") {
			res, err = plush.Render(linkUser, ctx)

		} else if strings.Contains(tcase.Route, "devices") {
			res, err = plush.Render(linkDevice, ctx)
		}

		assert.NoError(err)
		assert.Equal(res, tcase.ExpectedResult)
	}
}
