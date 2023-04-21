package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/bertoxic/bert/internal/config"
	"github.com/bertoxic/bert/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{
	"humanDate":HumanDate ,
	"formatDate":FormatDate ,
	"iterate":Iterate ,
	"add":Add ,
}
var app *config.AppConfig

//var pathToTemplate = "C:/Users/HP/Desktop/Bert/templates"
var pathToTemplate = "templates"

//returns formatted time
func HumanDate (t time.Time) string{
	return t.Format("2006-01-02")
}
func FormatDate (t time.Time,f string) string{
	return t.Format(f)
}
func NewRenderer(a *config.AppConfig) {
	app = a
}
func Add (a,b int) int{
	return a + b
}
//iterate returns a slice of ints starting at 1 and going to count
func Iterate(count int) []int {
	var i int
	var items []int

	for i = 0; i < count; i++ {
		items = append(items, i)
	}
return items
}
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.Sessions.PopString(r.Context(), "flash")
	td.Error = app.Sessions.PopString(r.Context(), "error")
	td.Warning = app.Sessions.PopString(r.Context(), "warning")
	if app.Sessions.Exists(r.Context(),"user_id") {
	td.IsAuthenticated = 1
	}

	return td
}
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	// _ , err := RenderTemplateTest(w)
	// if err != nil {
	// 	return err
	// }
	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// parsedTemplate.Execute(w, nil)
	// return nil
	var tc map[string]*template.Template

	if app.UserCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, OK := tc[tmpl]
	if !OK {
		//log.Fatal("could1 not get template from template cache")
		return errors.New("couldn't get template from cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myChache := map[string]*template.Template{}
	//fmt.Println("page is currently xxxx ")
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.gtpl", pathToTemplate))
	//fmt.Printf(fmt.Sprintf("zzzzzzzzzzzzzzz %d ", len(pages)))

	if err != nil {
		return myChache, err

	}
	//fmt.Println("page is currently vvvvv ")
	for _, page := range pages {
		name := filepath.Base(page)
		//	fmt.Println("page is currently kkkk")
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myChache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*layout.gtpl", pathToTemplate))
		//	fmt.Printf(fmt.Sprintf("aaaaaaaaaaa %d ", len(matches)))

		if err != nil {
			return myChache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*layout.gtpl", pathToTemplate))
			if err != nil {
				return myChache, err
			}
			//println(name)
		}

		myChache[name] = ts
	}
	return myChache, nil
}
