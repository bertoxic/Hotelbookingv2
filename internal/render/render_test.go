package render

import (
	"net/http"
	"testing"

	"github.com/bertoxic/bert/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error("failed")
	}
	sessions.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not in session adddefaultdata test ")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/GODDID", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = sessions.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}

func TestRenderTemplate(t *testing.T) {
	//pathToTemplate = "C:/Users/HP/Desktop/Bert/templates"
	pathToTemplate = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = Template(&ww, r, "home.page.gtpl", &models.TemplateData{})

	if err != nil {
		t.Error("error writing template to browser")
	}
	err = Template(&ww,r,"non-existing.page.gtpl",&models.TemplateData{})

		if err == nil {
			t.Error("rendered template that doest exist")
		}

	
}

func TestNewTemplate(t *testing.T){
		NewRenderer(app)	
}

func TestCreateTemplate(t *testing.T){
	_,err := CreateTemplateCache()
	if err != nil {
		t.Error("cannot test for template")
	}
}
