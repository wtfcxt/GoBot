package web

import "html/template"

var tpl = template.Must(template.ParseFiles("html/index.html"))