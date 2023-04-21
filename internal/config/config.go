package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"

	"github.com/bertoxic/bert/models"
)

type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Sessions      *scs.SessionManager
	MailChan      chan models.MailData
}
