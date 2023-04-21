package forms

import (
	"fmt"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}
func (f *Form) Valid() bool {
   return len(f.Errors)==0
}
 func (f *Form) Required(fields ...string){
    for _ , field:=range fields {
        value:=f.Get(field)
       if strings.TrimSpace(value)==""{
        f.Errors.Add(field,"This field cannot be empty")
       }
    }
 }
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string,) bool {
	x := f.Get(field)
	if x != "" {

		return true
	}
    f.Errors.Add(field,"This field has no data inputed")
	return false
}
 func (f *Form)MinLength(field string , length int) bool {
	x:=f.Get(field)
	if len(x) < length {
		f.Errors.Add(field,fmt.Sprintf("this field mus be at least %d character long",length))
		return false
	}
	return true
 }