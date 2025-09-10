package html

import (
	"fmt"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// HomePage is the front page of the app.
func QueryPage2(props PageProps, now time.Time) Node {
	props.Title = "Home"

	return page(props,
		Div(
			ID("tab-container"),
			Div(
				Class("tabs tabs-lift"),
				Div(
					Class("tab tab-buttons"),
					Button(
						Class("tab-button active"),
						Data("tab", "1"),
						Text("Tab 1"),
						Script(Raw(`
            me().on("click", ev => {  let numb = me(ev).attribute("data-tab"); any(".tab-button").classRemove("active"); me(ev).classAdd("active"); any(".tab-content2").classRemove("active"); any("#tabcon" + numb).classAdd("active") ; })
          `)),
					),
					Button(
						Class("tab-exit active"),
						Data("tab", "1"),
						Text("x"),
						Script(Raw(`
            me().on("click", ev => {  let numb = me(ev).attribute("data-tab"); event.preventDefault(); any("#tabcon" + numb).remove(); me(".tab-buttons").remove();  let alltabs = any(".tab-button"); if (alltabs.length >0) { let lasttab = alltabs[alltabs.length -1]; lasttab.classAdd("active"); let lasttabnum = lasttab.attribute("data-tab"); any("#tabcon" + lasttabnum).classAdd("active") }; })
          `)),
					),
				),
				Div(
					Class("tab tab-buttons"),
					Button(
						Class("tab-button"),
						Data("tab", "2"),
						Text("Tab 2"),
						Script(Raw(`
            me().on("click", ev => {  let numb = me(ev).attribute("data-tab"); any(".tab-button").classRemove("active"); me(ev).classAdd("active"); any(".tab-content2").classRemove("active"); any("#tabcon" + numb).classAdd("active") ;})
          `)),
					),
					Button(
						Class("tab-exit active"),
						Data("tab", "2"),
						Text("x"),
						Script(Raw(`
            me().on("click", ev => {  let numb = me(ev).attribute("data-tab"); console.log(numb ) ;event.preventDefault(); any("#tabcon" + numb).remove();  me(".tab-buttons").remove();  let alltabs = any(".tab-button"); if (alltabs.length >0) { let lasttab = alltabs[alltabs.length -1]; lasttab.classAdd("active"); let lasttabnum = lasttab.attribute("data-tab"); any("#tabcon" + lasttabnum).classAdd("active") }; })
          `)),
					),
				),
				Div(
					Class("tab tab-buttons"),
					Button(
						Class("tab-button"),
						Data("tab", "3"),
						Text("Tab 3"),
						Script(Raw(`
            me().on("click", ev => {  let numb = me(ev).attribute("data-tab"); any(".tab-button").classRemove("active"); me(ev).classAdd("active"); any(".tab-content2").classRemove("active"); any("#tabcon" + numb).classAdd("active") ;})
          `)),
					),
					Button(
						Class("tab-exit active"),
						Data("tab", "3"),
						Text("x"),
						Script(Raw(`
            me().on("click", ev => {  let numb = me(ev).attribute("data-tab"); event.preventDefault(); any("#tabcon" + numb).remove();  me(".tab-buttons").remove();  let alltabs = any(".tab-button"); if (alltabs.length >0) { let lasttab = alltabs[alltabs.length -1]; lasttab.classAdd("active"); let lasttabnum = lasttab.attribute("data-tab"); any("#tabcon" + lasttabnum).classAdd("active") }; })
          `)),
					),
				),
			),
			Div(
				Class("tab-contents"),
				Div(
					Class("tab-content-container"),
					Div(
						ID("tabcon1"),
						Class("tab-content2 active"),
						H3(
							Text("Content for Tab 1"),
						),
						Div(
							Data("tablename", "db1"),
							Textarea(
								ID("keyeventelement1"),
								Class("textarea textarea-info textarea-md resize"),
								Text("select * from city"),
							),
							Script(Raw(`
                        me().on("keydown", async event => {
                        let e = me(event)
                            if ((event.ctrlKey || event.metaKey) && event.key === 'e' ) {
                                event.preventDefault();
								let tablename = me(event).attribute("data-tablename");
                                clickedCtrLE(e.children[0], tablename ,"#responseQuery1")
                            }
                        })
                        `),
							),
						),
						Div(
							ID("responseQuery1"),
							Class("m-2 p-2 border border-base-300 rounded"),
							Text("Response"),
						),
					),
					Div(
						ID("tabcon2"),
						Class("tab-content2"),
						H3(
							Text("Content for Tab 2"),
						),
						Div(
							Data("tablename", "main"),
							Textarea(
								ID("keyeventelement2"),
								Class("textarea textarea-info textarea-md resize"),
								Text("Press Ctrl+Q"),
							),
							Script(Raw(`
                        me().on("keydown", async event => {
                        let e = me(event)
                            if ((event.ctrlKey || event.metaKey) && event.key === 'e' ) {
                                event.preventDefault();
								let tablename = me(event).attribute("data-tablename");
                                clickedCtrLE(e.children[0], tablename ,"#responseQuery2")
                            }
                        })
                        `),
							),
						),
						Div(
							ID("responseQuery2"),
							Class("m-2 p-2 border border-base-300 rounded"),
							Text("Response"),
						),
					),
					Div(
						ID("tabcon3"),
						Class("tab-content2"),
						H3(
							Text("Content for Tab 3"),
						),
						Div(
							Textarea(
								ID("keyeventelement3"),
								Class("textarea textarea-info textarea-md resize"),
								Text("Press Ctrl+Q"),
							),
							Script(Raw(`
                        me().on("keydown", async event => {
                        let e = me(event)
                            if ((event.ctrlKey || event.metaKey) && event.key === 'e' ) {
                                event.preventDefault();
								let tablename = me(event).attribute("data-tablename");
                                clickedCtrLE(e.children[0], tablename ,"#responseQuery3")
                            }
                        })
                        `),
							),
						),
						Div(
							ID("responseQuery3"),
							Class("m-2 p-2 border border-base-300 rounded"),
							Text("Response"),
						),
					),
				),
			),
		),

		Script(Raw(`
        //onloadAdd(()=>{
                function clickedCtrLE(myInput  , tablename ,target) {
                    // if ((event.ctrlKey || event.metaKey) && event.key === 'e') {
                    // event.preventDefault(); // Prevent the default browser save action
                    var start = myInput.selectionStart;
                    var finish = myInput.selectionEnd;
                    // Obtain the selected text
                    var selectedText = myInput.value.substring(start, finish);
                    if (selectedText == "") {
                        selectedText = myInput.value;
                    }
                    // console.log("Ctrl+E pressed, selectedText is:"+selectedText);
                    return sendQuery(selectedText, tablename, target );
                //}
                };

                // Autocomplete
                //  myInput.addEventListener('keydown', function(event) {
                //  // Check if Ctrl (or Cmd on Mac) and 'Space' key are pressed
                //  if ((event.ctrlKey || event.metaKey) && event.key === ' ' ) {
                //    event.preventDefault(); // Prevent the default browser save action
                //    // If there's a selection, get its position
                //    start = myInput.selectionStart;
                //    end = myInput.selectionEnd;
                //    // alert (myInput.top);
                //    const cursorRow = myInput.value.substr(0, start).split("\n").length;
                //    // const textBeforeCursor = myInput.value.substring(0, start);
                //    console.log("Ctrl+SPACE  pressed, custom save action triggered!" + "cursor row position "+cursorRow ) ;
                //    console.log(JSON.stringify(getElementCoordinates(myInput)));
                //    // Example: call a function to save data, submit a form, etc.
                //    // Add to document
                //     // document.body.appendChild(divWindow);
                //    sendAutocmplete(myInput.value,cursorRow);
                //  }
                // });
                // function sendAutocmplete(elm,cursorRow) {
                //  htmx.ajax('GET', '/query/qq', {
                //  values: { query: elm , cursorRow: cursorRow},
                //  target: '#response',
                //  swap: 'innerHTML'
                //  });
                // }
                async function sendQuery(text , tablename , target ) {
                    try {
                        const response = await fetch('/query/'+tablename, {
                        method: 'POST', // Specify the HTTP method as POST
                      //  headers: {
                     //       'Content-Type': 'application/json' // Indicate that the body contains JSON data
                      //  },
                        body: JSON.stringify({ query: text }) // Convert the JavaScript object to a JSON string
                    });
                        const result = await response.text();
                        any(target)[0].innerHTML = result;
                    } catch (error) {
                        console.error("Error:", error);
                    }

                }
                    //})`),
		),
		dbQueryModal(props),
	)
}

func AddTabFunctionality(lasttab int) Node {
	tabnumber := lasttab + 1
	return Div(
		ID("tabcon"+fmt.Sprintf("%d", tabnumber)),
		Class("tab-content2"),
		H3(
			Text("Content for Tab "+fmt.Sprintf("%d", tabnumber)),
		),
		Div(
			Textarea(
				ID("keyeventelement"+fmt.Sprintf("%d", tabnumber)),
				Class("textarea textarea-info textarea-md resize"),
				Text("Press Ctrl+QQQQQ"),
			),
			Script(Raw(`
                  me().on("keydown", async event => {
                  let e = me(event)
                  if ((event.ctrlKey || event.metaKey) && event.key === 'e' ) {
                    event.preventDefault();
                    clickedCtrLE(e.children[0], "#responseQuery`+fmt.Sprintf("%d", tabnumber)+`)
                  }
                 })
              `),
			),
		),
		Div(
			ID("responseQuery"+fmt.Sprintf("%d", tabnumber)),
			Class("m-2 p-2 border border-base-300 rounded"),
			Text("Response"),
		),
	)
}

func searchBar(props PageProps) Node {
	return Dialog(
		ID("my_search_modal"),
		Class("modal"),
		Div(
			Class("modal-box"),
			H3(
				Class("font-bold text-lg"),
				Text("Search"),
			),
			Input(
				Type("text"),
				Placeholder("Type here"),
				Class("input input-bordered w-full max-w-xs"),
			),
			Div(
				Class("modal-action"),
				Form(
					Method("dialog"),
					Button(
						Class("btn"),
						Text("Close"),
					),
				),
			),
		),
	)
}

func dbQueryModal(props PageProps) Node {
	//dbmanager.DBConfig
	return Dialog(
		ID("my_query_modal_1"),
		Class("modal"),
		Div(
			Class("modal-box"),
			H3(
				Class("text-lg font-bold"),
				Text("Update/Insert query"),
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
