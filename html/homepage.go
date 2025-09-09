package html

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// HomePage is the front page of the app.
func HomePage(props PageProps) Node {
	props.Title = "Home"

	return page(props,
		Div(
			Class("hero bg-base-200 min-h-screen"),
			Div(
				Class("hero-content text-center"),
				Div(
					Class("max-w-md"),
					H1(
						Class("text-5xl font-bold"),
						Text("Hello there"),
					),
					P(
						Class("py-6"),
						Text("Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem\n        quasi. In deleniti eaque aut repudiandae et a id nisi."),
					),
					Button(
						Class("btn btn-primary"),
						Text("Get Started"),
					),
				),
			),
		),
		Div(Class("prose prose-indigo prose-lg md:prose-xl"),
			H1(Text("Welcome to the gomponents starter kit")),

			P(Text("It uses gomponents, HTMX, and Tailwind CSS, and you can use it as a template for your new app. 😎")),

			P(A(Href("https://github.com/maragudk/gomponents-starter-kit"), Text("See gomponents-starter-kit on GitHub"))),

			H2(Text("Try HTMX")),

			Button(
				Class("rounded-md bg-indigo-600 px-2.5 py-1.5 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"),
			//	Textf("Get things with HTMX"), hx.Get("/"), hx.Target("#things")),
			),
			Div(ID("things")), //ThingsPartial(things, now),

		),
	)
}

// ThingsPartial is a partial for showing a list of things, returned directly if the request is an HTMX request,
// and used in [HomePage].
// func ThingsPartial(things []model.Thing, now time.Time) Node {
// 	return Group{
// 		P(Textf("Here are %v things from the mock database (updated %v):", len(things), now.Format(time.TimeOnly))),
// 		Ul(
// 			Map(things, func(t model.Thing) Node {
// 				return Li(Text(t.Name))
// 			}),
// 		),
// 	}
// }

// TablePartial is a partial for showing a table of results from a query
func TablePartial(things []map[string]interface{}, columns []string) Node {

	return Group{
		Div(
			Class("overflow-x-auto"),
			Table(Class("table table-xs"),
				THead(
					Tr(
						Group(
							Map(columns, func(t string) Node {
								return Th(Text(t))
							}),
						),
					),
				),
				TBody(
					Group(
						Map(things, func(row map[string]interface{}) Node {
							return Tr(Class("bg-base-300"),
								Group(
									Map(columns, func(col string) Node {
										value := row[col]
										return Td(Text(fmt.Sprintf("%v", value)))
									}),
								),
							)
						}),
					),
				),
			),
		),
	}
}
