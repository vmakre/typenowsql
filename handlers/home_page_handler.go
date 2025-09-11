package handlers

import (
	"net/http"
	"os"
	"time"

	"typenowsql/html"
	"typenowsql/service"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	. "maragu.dev/gomponents"
	//hx "maragu.dev/gomponents-htmx/http"
	// . "maragu.dev/gomponents/html"
	ghttp "maragu.dev/gomponents/http"
)

type HomePageHandler struct {
	homepageService service.HomePageService
}

func NewHomePageHandler(homepageService service.HomePageService) *HomePageHandler {
	return &HomePageHandler{homepageService: homepageService}
}

func (h *HomePageHandler) GetHomePage() http.HandlerFunc {
	return ghttp.Adapt(
		func(w http.ResponseWriter, r *http.Request) (Node, error) {
			return html.HomePage(html.PageProps{}), nil
		})
	// return ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
	//db := database.NewDB()
	//sessionManager := service.NewSessionManager()

	//sessionManager.Put(r.Context(), "flash", "This is a flash message")
	//things, err := db.GetThings(r.Context())
	// if err != nil {
	// 	return nil, err
	// }
	//things := []string{"thing1", "thing2", "thing3"}

}
func (h *HomePageHandler) GetGueryPage() http.HandlerFunc {
	return ghttp.Adapt(
		func(w http.ResponseWriter, r *http.Request) (Node, error) {
			return html.QueryPage2(html.PageProps{}, time.Now()), nil
		})
}

// if application is in cmd/app public folder is not accesicle so check where is main.go first
func Checkexists(path1 string) string {
	if _, err := os.Stat(path1); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return "../../" + path1
		} else {
			// other error
			return path1
		}
	} else {
		return path1
	}
}

// Static assets handler, which serves files from the root that have an extension, and everything from
// the images, scripts, and styles directories.
func Static(r chi.Router) {
	staticHandler := http.FileServer(http.Dir(Checkexists("public")))
	r.Get(`/{:[^.]+\.[^.]+}`, staticHandler.ServeHTTP)
	r.Get(`/{:images|scripts|styles}/*`, staticHandler.ServeHTTP)
}
