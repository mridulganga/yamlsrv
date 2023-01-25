package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

func MapPath(m map[string]interface{}, path string) interface{} {
	var obj any = m
	var val any = nil

	parts := strings.Split(path, "/")
	for _, p := range parts {
		if v, ok := obj.(map[string]any); ok {
			obj = v[p]
			val = obj
		} else if v, ok := obj.(map[any]any); ok {
			obj = v[p]
			val = obj
		} else if v, ok := obj.([]any); ok {
			pInt, _ := strconv.Atoi(p)
			if pInt >= 0 && pInt < len(v) {
				obj = v[pInt]
				val = obj
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
	return val
}

func main() {
	yamlFile, err := os.ReadFile("main.yaml")
	if err != nil {
		panic(err)
	}

	data := map[string]any{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Get("*", func(c *fiber.Ctx) error {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		dataBytes, err := json.Marshal(MapPath(data, c.Path()[1:]))
		if err != nil || string(dataBytes) == "null" {
			return c.Status(404).SendString("error 404")
		}
		return c.SendString(string(dataBytes))
	})
	app.Listen(":3000")

}
