package upload

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lrth06/go-chat/lib/utils/config"
)
func HandleUpload(c *fiber.Ctx) error {
	env, err := config.GetConfig()
	if err != nil {
		return c.SendStatus(400)
	}
	appEnv := env.AppEnv
	//empty array for file urls
	var urls []string

	if appEnv == "development" {
		id := c.GetRespHeader("id")
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["uploads"]
		os.MkdirAll("./tmp/uploads/users/"+id, 0777)
		for _, file := range files {
			err := c.SaveFile(file, fmt.Sprintf("./tmp/uploads/users/%s/%s", id, file.Filename))
			if err != nil {
				return err
			}
			urls = append(urls, fmt.Sprintf("http://%s/images/users/%s/%s", c.Hostname(), id, file.Filename))

		}
		return c.Status(201).JSON(fiber.Map{
			"urls": urls,
		})
	}
	//return array of file urls as "data" in response
	return c.Status(201).JSON(fiber.Map{
		"data": urls,
	})

}
