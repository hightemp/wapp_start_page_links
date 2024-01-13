package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/django/v3"
	"github.com/hightemp/wapp_start_page_links/lib/config"
	"github.com/hightemp/wapp_start_page_links/lib/session"

	"github.com/joho/godotenv"
)

//go:embed assets
var embedAsssets embed.FS

const (
	DEFAULT_HOST   = "0.0.0.0"
	DEFAULT_PORT   = "8585"
	DEFAULT_CONFIG = "./config.yaml"
)

func GetEnv(name string, defaultValue string) string {
	value, exists := os.LookupEnv(name)

	if !exists {
		value = defaultValue
	}

	return value
}

func getElementAtIndex(slice interface{}, index int) (interface{}, bool) {
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Array {
		return nil, false
	}

	if index >= 0 && index < sliceValue.Len() {
		return sliceValue.Index(index).Interface(), true
	}

	return nil, false
}

func createDirectoryIfNotExists(dirPath string) error {
	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		//
	}

	return nil
}

func removeItemFromArray(slice interface{}, index int) interface{} {
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Slice {
		panic("Input is not a slice")
	}

	length := sliceValue.Len()

	if index < 0 || index >= length {
		panic("Index out of range")
	}

	return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	serverPort := GetEnv("PORT", DEFAULT_PORT)
	serverHost := GetEnv("HOST", DEFAULT_HOST)
	configFile := GetEnv("CONFIG_FILE", DEFAULT_CONFIG)

	serveStr := serverHost + ":" + serverPort

	configObj := config.New(configFile)

	engine := django.NewPathForwardingFileSystem(http.FS(embedAsssets), "/assets/views", ".tpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	store := session.NewSessionStorage()

	_ = createDirectoryIfNotExists("./images")

	app.Use("images", filesystem.New(filesystem.Config{
		Root: http.Dir("./images"),
	}))

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(embedAsssets),
		PathPrefix: "assets",
		Browse:     true,
	}))

	//fmt.Printf("%v", spew.Sdump(configObj))
	//spew.Printf("List: %#+v", configObj.Data.List)

	app.Get("/", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			configObj.Load()

			return c.Render("main_page", fiber.Map{
				"config": configObj,
			})
		})
	})

	app.Get("/edit_list", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			configObj.Load()

			return c.Render("edit_list_page", fiber.Map{
				"config": configObj,
			})
		})
	})

	app.Get("/edit_list/edit/:id", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			id, err := strconv.Atoi(c.Params("id"))
			if err != nil {
				return c.Render("error_page", fiber.Map{})
			}
			item, exists := getElementAtIndex(configObj.Data.List, id)
			if !exists {
				return c.Render("error_page", fiber.Map{})
			}

			return c.Render("edit_site_item_page", fiber.Map{
				"item": item,
				"id":   id,
			})
		})
	})

	app.Get("/edit_list/add", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			return c.Render("edit_site_item_page", fiber.Map{
				"item": config.ConfigSite{},
				"id":   -1,
			})
		})
	})

	app.Get("/edit_list/delete/:id", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			id, err := strconv.Atoi(c.Params("id"))
			if err != nil {
				return c.Render("error_page", fiber.Map{})
			}
			_, exists := getElementAtIndex(configObj.Data.List, id)
			if !exists {
				return c.Render("error_page", fiber.Map{})
			}

			configObj.Data.List = removeItemFromArray(configObj.Data.List, id).([]config.ConfigSite)

			configObj.Save()

			return c.Redirect("/edit_list")
		})
	})

	app.Post("/edit_list/update/:id", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			id, err := strconv.Atoi(c.Params("id"))
			if err != nil {
				return c.Render("error_page", fiber.Map{})
			}
			if id != -1 {
				_, exists := getElementAtIndex(configObj.Data.List, id)
				if !exists {
					return c.Render("error_page", fiber.Map{})
				}
			}
			site := config.ConfigSite{
				Name:        c.FormValue("name"),
				Description: c.FormValue("description"),
				Image:       c.FormValue("image"),
				Url:         c.FormValue("url"),
			}

			if id == -1 {
				configObj.Data.List = append(configObj.Data.List, site)
			} else {
				configObj.Data.List[id] = site
			}
			configObj.Save()

			if c.FormValue("save_and_edit") == "1" {
				referer := c.Get("Referer")
				return c.Redirect(referer)
			}

			return c.Redirect("/edit_list")
		})
	})

	app.Get("/settings", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			configObj.Load()

			return c.Render("settings", fiber.Map{
				"config": configObj,
			})
		})
	})

	app.Post("/settings/update", func(c *fiber.Ctx) error {
		return store.Wrap(c, func(c *fiber.Ctx, s *session.Session) error {
			configObj.Data.Settings.Theme = c.FormValue("theme")

			if c.FormValue("open_links_in_new_window") == "1" {
				configObj.Data.Settings.OpenLinksInNewWindow = true
			} else {
				configObj.Data.Settings.OpenLinksInNewWindow = false
			}

			configObj.Save()

			return c.Redirect("/settings")
		})
	})

	log.Fatal(app.Listen(serveStr))
}
