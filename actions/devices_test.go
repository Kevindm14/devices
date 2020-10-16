package actions

import (
	"devices/models"
	"net/url"

	"github.com/gobuffalo/httptest"
)

func (as *ActionSuite) CreateItem() *models.Device {
	device := &models.Device{
		Manufacture:     "123546",
		Make:            "Iphone",
		Model:           "xr",
		Storage:         "64gb",
		Cost:            100.000,
		OperatingSystem: "Android",
		Color:           "Yellow",
		Image:           "https://images-na.ssl-images-amazon.com/images/I/61LA-THTwWL._AC_SL1500_.jpg",
	}
	verrs, err := as.DB.ValidateAndCreate(device)
	as.NoError(err)
	as.False(verrs.HasAny())
	return device
}

// func (as *ActionSuite) TableChange(table string, count int, function func()) {
// 	beforeCount, err := as.DB.Count(table)
// 	as.NoError(err)

// 	function()

// 	afterCount, err := as.DB.Count(table)
// 	as.NoError(err)
// 	as.Equal(count, afterCount-beforeCount)
// }

func (as *ActionSuite) Test_Devices_index() {
	as.LoadFixture("lots of devices")
	res := as.HTML("/devices").Get()

	as.Equal(200, res.Code)

	body := res.Body.String()
	as.Contains(body, "iphone", "samsung")
	as.NotContains(body, "xiaomi")
}

func (as *ActionSuite) Test_Devices_show() {
	device := as.CreateItem()

	res := as.HTML("/devices/%s", device.ID).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Devices_New() {
	res := as.HTML("/devices/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Devices_Edit() {
	device := as.CreateItem()
	res := as.HTML("/devices/%s/edit", device.ID).Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Devices_Create() {
	device := models.Device{
		Manufacture:     "123546",
		Make:            "Iphone",
		Model:           "xr",
		Storage:         "64gb",
		OperatingSystem: "Android",
	}

	deviceVar := models.DeviceVariations{
		Storage: []string{"32GB", "64GB"},
		Cost:    []float64{10000, 20000},
		Color:   []string{"Blue", "Red"},
		Image: []string{
			"https://encrypted-tbn2.gstatic.com/shopping?q=tbn:ANd9GcQ_F2DSXXspdC8ytBT2K6xmCwJX5A5bMR80AkvpRK-W_8XK9c2JCQM&usqp=CAc",
			"https://m.media-amazon.com/images/I/51n24DedexL.jpg",
		},
	}

	var res *httptest.Response

	for i := 0; i < len(deviceVar.Storage); i++ {
		devices := models.Device{
			Manufacture:     device.Manufacture,
			Make:            device.Make,
			Model:           device.Model,
			Storage:         deviceVar.Storage[i],
			Cost:            deviceVar.Cost[i],
			OperatingSystem: device.OperatingSystem,
			Image:           deviceVar.Image[i],
		}

		res = as.HTML("/devices").Post(&devices)

		as.DB.Last(&devices)
		as.Equal(302, res.Code)
		as.NotZero(devices.ID)
		as.NotZero(devices.CreatedAt)
	}

	as.Equal("/devices", res.Location())

}

func (as *ActionSuite) Test_Devices_Update() {
	device := as.CreateItem()

	deviceNew := url.Values{
		"Manufacture":     []string{"123546"},
		"Make":            []string{"Iphone"},
		"Model":           []string{"11"},
		"Storage":         []string{"64gb"},
		"Cost":            []string{"100000"},
		"OperatingSystem": []string{"Android"},
		"Color":           []string{"Yellow"},
		"Image":           []string{"https://images-na.ssl-images-amazon.com/images/I/61LA-THTwWL._AC_SL1500_.jpg"},
	}

	res := as.HTML("/devices/%s", device.ID).Put(deviceNew)
	as.Equal(302, res.Code)

	as.NoError(as.DB.Reload(device))
	as.Equal("11", device.Model)
}

func (as *ActionSuite) Test_Devices_destroy() {
	device := as.CreateItem()

	res := as.HTML("/devices/%s", device.ID).Delete()
	as.Equal(302, res.Code)
	as.Equal("/devices", res.Location())
}
