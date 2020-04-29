package api

import "cn.sockstack/smser/pkg/service"

func V1(app *service.App) {
	app.R.Route("GET", "/test", service.RegisterService(Endpoint(Service{}), CreateRequest(), CreateResponse()))
}
