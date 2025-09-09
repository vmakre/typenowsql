// Package HTML holds all the common HTML components and utilities.
package html

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

var hashOnce sync.Once
var appCSSPath, appJSPath, htmxJSPath, hyperJSPath, surrealJSPath string

// PageProps are properties for the [page] component.
type PageProps struct {
	Title       string
	Description string
}

// page layout with header, footer, and container to restrict width and set base padding.
func page(props PageProps, children ...Node) Node {

	// JavaScript code to be injected
	// jsCode := `
	// 	document.addEventListener('DOMContentLoaded', function() {
	// 		document.getElementById('demo').textContent = 'Hello from JavaScript!';
	// 	});
	// `

	// Hash the paths for easy cache busting on changes
	hashOnce.Do(func() {
		appCSSPath = getHashedPath("public/styles/app.css")
		htmxJSPath = getHashedPath("public/scripts/htmx.js")
		appJSPath = getHashedPath("public/scripts/app.js")
		hyperJSPath = getHashedPath("public/scripts/hyperscript.js")
		surrealJSPath = getHashedPath("public/scripts/surreal.js")
	})

	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    "en",
		Head: []Node{
			// HTML(Attr("data-theme", "light")),
			Link(Rel("stylesheet"), Href(appCSSPath)),
			Script(Src(htmxJSPath), Defer()),
			Script(Src(hyperJSPath), Defer()),
			Script(Src(surrealJSPath)),
			Script(Src(appJSPath), Defer()),
		},
		Body: []Node{
			header(),
			Div(
				Class("grow"),
				container(true, true,
					Group(children),
				),
			),
			dbconnectionsModal(),
			footer(),
			// add javascript
			// Script(
			// 	Type("text/javascript"), // Optional, but good practice
			// 	Raw(jsCode),
			// ),
		},
	})
}

// header bar with logo and navigation.
func header() Node {
	return Div(
		Class("navbar bg-base-100 shadow-sm"),
		Div(
			Class("navbar-start"),
			Div(
				Class("flex-1"),
				Button(
					Class("tab-button"),
					Text("sqL"),
					Script(Raw(`
			//<button class="btn" onclick="my_modal_1.showModal()">open modal</button>
            me().on("click", ev => {  my_modal_1.showModal() ;})
          `)),
				),
				// Button(
				// 	Class("btn btn-ghost text-xl"),
				// 	Text("sqL"),
				// ),

				// Script(Raw(`
				//   me().on("click", async event => {
				//   if ((event.EnterKey || event.key === 'Enter') && !event.shiftKey) {
				//     event.preventDefault();
				// 	let e = me(event).children[0];
				// 	alert ("Search for: " + e.value);
				//     //clickedCtrLE(e.children[0], "#responseQuery)
				//   }
				// });
				// me().on("click" , event => { let e = me(event).children[0]; e.style.backgroundColor = "#353434"; } );
				// `)),
			),
		),
		Div(
			Class("navbar-center"),
			Input(
				Type("text"),
				Placeholder("Search"),
				Class("input  w-84 md:w-full"),
			),
			Script(Raw(`
                  me().on("keydown", async event => {
				  if ((event.EnterKey || event.key === 'Enter') && !event.shiftKey) {
                    event.preventDefault();
					let e = me(event).children[0];
					alert ("Search for: " + e.value);
                    //clickedCtrLE(e.children[0], "#responseQuery)
                  }
				});
				me().on("click" , event => { let e = me(event).children[0]; e.style.backgroundColor = "#353434"; } );  
				`)),
		),
		Div(
			Class("navbar-end"),
			Div(
				Class("dropdown"),
				Div(
					TabIndex("0"),
					Role("button"),
					Class("btn m-1"),
					Text("-*-"),
				),
				Ul(
					TabIndex("0"),
					Class("dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"),
					Li(
						A(
							Text("Theme"),
						),
					),
					Li(
						A(
							Text("Settings"),
						),
					),
				),
			),
		),
	)

}

// container restricts the width and sets padding.
func container(padX, padY bool, children ...Node) Node {
	return Div(
		Classes{
			"max-w-7xl mx-auto":     true,
			"px-4 md:px-8 lg:px-16": padX,
			"py-4 md:py-8":          padY,
		},
		Group(children),
	)
}

func getProjectRoot() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get caller information")
	}
	// Assuming the project root is one level up from the directory containing this file
	return filepath.Abs(filepath.Join(filepath.Dir(filename), ".."))
}

// footer with a link to the gomponents website.
func footer() Node {
	return Div(Class("bg-indigo-600 text-white"),
		container(true, false,
			Div(Class("h-16 flex items-center justify-center"),
				A(Href("https://www.gomponents.com"), Text("www.gomponents.com")),
			),
		),
	)
}

func dbconnectionsModal(dbhandler *DBHandler) Node {
	//dbmanager.DBConfig
	return Dialog(
		ID("my_modal_1"),
		Class("modal"),
		Div(
			Class("modal-box"),
			H3(
				Class("text-lg font-bold"),
				Text("Hello!"),
			),
			P(
				Class("py-4"),
				Text("Press ESC key or click the button below to close"),
			),
			Div(
				Class("modal-action"),
				Form(
					Method("dialog"),
					// if there is a button in form, it will close the modal
					Button(
						Class("btn"),
						Text("Close"),
					),
				),
			),
		),
	)
}

func getHashedPath(path1 string) string {
	root, err := getProjectRoot()
	if err != nil {
		fmt.Println("Error getting project root:", err)
		return ""
	}
	fmt.Println("Project Root (estimated):", root)
	externalPath := strings.TrimPrefix(path1, "public")
	ext := path.Ext(path1)
	data, err := os.ReadFile(path.Join(root, path1))
	if err != nil {
		return fmt.Sprintf("%v.x%v", strings.TrimSuffix(externalPath, ext), ext)
	}

	return fmt.Sprintf("%v.%x%v", strings.TrimSuffix(externalPath, ext), sha256.Sum256(data), ext)
}
