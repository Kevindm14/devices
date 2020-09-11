package actions

import (
	"devices/models"
	"fmt"
)

func (as *ActionSuite) Test_Devices_index() {
	as.LoadFixture("lots of devices")
	res := as.HTML("/devices").Get()

	body := res.Body.String()
	as.Contains(body, "iphone")
	as.Contains(body, "samsung")
}

func (as *ActionSuite) Test_Devices_Create() {
	device := models.Device{
		Manufacture:     "123546",
		Make:            "Iphone",
		Model:           "xr",
		Storage:         "64gb",
		Cost:            100000,
		OperatingSystem: "Android",
		Color:           "Yellow",
		Image:           "https://images-na.ssl-images-amazon.com/images/I/61LA-THTwWL._AC_SL1500_.jpg",
		IsNew:           false,
	}

	res := as.HTML("/devices").Post(&device)
	as.DB.Last(&device)
	as.Equal(302, res.Code)

	as.Equal(fmt.Sprintf("/devices/%s", device.ID), res.Location())
	as.NotZero(device.ID)
	as.NotZero(device.CreatedAt)
}

func (as *ActionSuite) Test_Devices_destroy() {
	device := &models.Device{
		Manufacture:     "123546",
		Make:            "Iphone",
		Model:           "xr",
		Storage:         "64gb",
		Cost:            100000,
		OperatingSystem: "Android",
		Color:           "Yellow",
		Image:           "https://images-na.ssl-images-amazon.com/images/I/61LA-THTwWL._AC_SL1500_.jpg",
		IsNew:           false,
	}

	as.DB.Create(device)
	res := as.HTML("/devices/%s", device.ID).Delete()
	as.Equal(302, res.Code)
	as.Equal("/devices", res.Location())
}
