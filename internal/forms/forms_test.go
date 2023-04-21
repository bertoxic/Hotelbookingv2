package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_valid(t *testing.T){

    r := httptest.NewRequest("POST","/into",nil)
    form:= New(r.PostForm)

        isValid := form.Valid()
        if !isValid {
            t.Error("failed validation when it ought not to")
        }
}

func TestForm_Required(t *testing.T){

    postDat := url.Values{}
    postDat.Add("ven","masarati")
    formx:= New(postDat)
    formx.Required("bus")
   if formx.Valid(){
            t.Error("required check failed ")
    }


    postData := url.Values{}
     postData.Add("a","a")
     postData.Add("b","b")
     
     forms:= New(postData)
      forms.Required("a","b")
      if !forms.Valid() {
            t.Error("somthing wong")
      }

}

func TestForm_MinLength(t *testing.T){
    postData:= url.Values{}
     forms:= New(postData)
     postData.Add("lens","beautiful")
     forms.MinLength("lens",5)
     if !forms.Valid() {
            t.Error("length check should be valid but it is not")
     }

     postDat:= url.Values{}
     formsx:= New(postDat)
     formsx.MinLength("conn",7)
     if formsx.Valid(){
        t.Error("length check should be invalid cuz field is non-existent")
     }

}

func TestForm_Has(t *testing.T){
    postData := url.Values{}
    postData.Add("tips","")
    form:= New(postData)
   
    form.Has("tips")

    if form.Valid() {
        t.Error("tips has no value hence should fail the test")
    }
}