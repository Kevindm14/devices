package grifts

import (
	"devices/models"

	"github.com/markbates/grift/grift"
	"github.com/wawandco/fako"
)

var _ = grift.Namespace("db", func() {
	grift.Desc("seed", "Seeds a database")
	grift.Add("users", func(c *grift.Context) error {
		count := 0

		for count < 50 {
			type User struct {
				FirstName    string `fako:"first_name"`
				LastName     string `fako:"last_name"`
				Email        string `fako:"email_address"`
				ManagerEmail string `fako:"email_address"`
			}

			var user User
			fako.Fill(&user)

			if count%4 == 0 {
				userNew := &models.User{
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					Role:      "Manager",
				}
				models.DB.ValidateAndCreate(userNew)
			} else {
				userNew := &models.User{
					FirstName:    user.FirstName,
					LastName:     user.LastName,
					Email:        user.Email,
					Role:         "RegularUser",
					ManagerEmail: user.ManagerEmail,
				}
				models.DB.ValidateAndCreate(userNew)
			}

			count++
		}
		return nil
	})

	grift.Desc("seed", "Seeds a database")
	grift.Add("devices", func(c *grift.Context) error {
		images := []string{
			"https://i.pinimg.com/originals/67/54/78/675478c7dcc17f90ffa729387685615a.jpg",
			"https://falabella.scene7.com/is/image/FalabellaCO/4153277_1?q=i?wid=800&hei=800&qlt=70",
			"https://www.itechcolombia.co/wp-content/uploads/samsung-s20-plus-azul-600x600.jpg",
			"https://images-eu.ssl-images-amazon.com/images/I/418csJ322SL.jpg",
			"https://hsi.com.co/wp-content/uploads/2019/01/Mi-8-Lite-128gb-negro.jpg",
			"https://images.samsung.com/is/image/samsung/es-galaxy-s6-g920f-sm-g920fzkaphe-001-front-black?$PD_GALLERY_L_JPG$",
			"https://images-na.ssl-images-amazon.com/images/I/618cyFfBFkL._AC_SL1200_.jpg",
			"https://www.officedepot.com.pa/medias/22227-1200ftw?context=bWFzdGVyfHJvb3R8MjMyMTQ1fGltYWdlL2pwZWd8aGQzL2hmMS85NTM4OTQ5MDU0NDk0LmpwZ3w5ZDQ2ZjRiZjdhMTE1ZTRjODczZGE3ZDJmZjg4MTliNzAzM2U5ZDZlZTU3OTlmNDAyOGY5ZGNhMTg2Y2IyMDlm",
			"https://exitocol.vtexassets.com/arquivos/ids/1052619-800-auto?width=800&height=auto&aspect=true",
			"https://csmobiles.com/30515-large_default/samsung-galaxy-s10-lite-g770-8gb-ram-128gb-dual-sim-azul.jpg",
			"https://cnet4.cbsistatic.com/img/mMu445m35jov17fJgfNRBGQx8lM=/940x0/2020/01/03/d2b76a19-625e-429f-a041-bc2a4aad5753/galaxynote10litepr-mainff.jpg",
			"https://i.blogs.es/c3597a/fdd984e3-de6e-4936-af4c-ed93e402a654/1366_2000.jpeg",
			"https://www.yaphone.com/3947-large_default/xiaomi-mi-10-lite-5g.jpg",
			"https://exitocol.vtexassets.com/arquivos/ids/2133199-800-auto?width=800&height=auto&aspect=true",
			"https://images-na.ssl-images-amazon.com/images/I/71T0vySAgNL._AC_SX569_.jpg",
			"https://falabella.scene7.com/is/image/FalabellaCO/7802601_1?q=i?wid=800&hei=800&qlt=70",
			"https://exitocol.vtexassets.com/arquivos/ids/1052832/Celular-Xiaomi-Redmi-Note-8-64-GB.jpg?v=637130730362000000",
			"https://exitocol.vtexassets.com/arquivos/ids/1054794/Celular-Xiaomi-Redmi-Note-8-Blanco-128gb-4Ram---Forro-goma.jpg?v=637130795328800000",
			"https://d500.epimg.net/cincodias/imagenes/2020/03/05/smartphones/1583416907_266480_1583418090_sumario_normal.jpg",
			"https://cdn2.smart-gsm.com//2020/03/Realme-6-Pro-1.jpg",
		}

		for i := 0; i < 5; i++ {
			for j := 0; j < 20; j++ {
				type Device struct {
					Manufacture string `fako:"digits"`
					Make        string `fako:"product"`
					Model       string `fako:"model"`
					Storage     string `fako:"title"`
					Color       string `fako:"color"`
				}

				var device Device
				fako.Fill(&device)

				if j%2 == 0 {
					deviceNew := &models.Device{
						Manufacture:     device.Manufacture,
						Make:            device.Make,
						Model:           device.Model,
						Storage:         device.Storage,
						Cost:            100000,
						Color:           device.Color,
						OperatingSystem: "Android",
						Image:           images[j],
					}
					models.DB.ValidateAndCreate(deviceNew)

				} else if j%3 == 0 {
					deviceNew := &models.Device{
						Manufacture:     device.Manufacture,
						Make:            device.Make,
						Model:           device.Model,
						Storage:         device.Storage,
						Cost:            200000,
						Color:           device.Color,
						OperatingSystem: "IOS",
						Image:           images[j],
					}
					models.DB.ValidateAndCreate(deviceNew)

				} else {
					deviceNew := &models.Device{
						Manufacture:     device.Manufacture,
						Make:            device.Make,
						Model:           device.Model,
						Storage:         device.Storage,
						Cost:            300000,
						Color:           device.Color,
						OperatingSystem: "Windows",
						Image:           images[j],
					}
					models.DB.ValidateAndCreate(deviceNew)
				}
			}
		}
		return nil
	})

})
