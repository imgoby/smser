package api

import (
	"cn.sockstack/smser/api/variable"
	"cn.sockstack/smser/pkg"
)

func V1(app *pkg.App) {
	app.R.Route(
		"GET",
		"/test",
		pkg.RegisterService(
			variable.AddCategoryEndpoint(variable.CategoryService{}),
			variable.CreateAddParams(),
			variable.CreateAddResult(),
		),
	)
}
