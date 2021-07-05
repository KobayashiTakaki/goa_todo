package design

import . "goa.design/goa/v3/dsl"

var _ = API("todo", func() {
	Title("Todo Service")
	Description("Service that manage todo.")
	Server("todo", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var Todo = ResultType("Todo", func() {
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("title", String, "Title")
		Attribute("is_done", Boolean, "IsDone")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("is_done")
	})
})

var _ = Service("todo", func() {
	Description("Service that manage todo.")
	Method("hello", func() {
		Payload(func() {
			Attribute("name", String, "Name")
			Required("name")
		})
		Result(String)
		HTTP(func() {
			GET("/hello/{name}")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Payload(func() {
			Attribute("id", Int, "ID")
			Required("id")
		})
		Result(Todo)
		HTTP(func() {
			GET("/todo/{id}")
			Response(StatusOK)
		})
	})

	Method("create", func() {
		Payload(func() {
			Attribute("title", String, "Title")
			Required("title")
		})
		Result(String)
		HTTP(func() {
			POST("/todo")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Payload(func() {
			Attribute("id", Int, "ID")
			Attribute("is_done", Boolean, "IsDone")
			Required("id")
			Required("is_done")
		})
		Result(String)
		HTTP(func() {
			POST("/todo/{id}")
			Response(StatusOK)
		})
	})
})
