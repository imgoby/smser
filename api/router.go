package api

import (
	"cn.sockstack/smser/api/variable"
	"cn.sockstack/smser/src"
)

func V1(app *src.App) {
	app.R.Route(
		"GET",
		"/test",
		src.RegisterService(
			variable.AddCategoryEndpoint(variable.CategoryService{}),
			variable.CreateAddParams(),
			variable.CreateAddResult(),
		),
	)
}
