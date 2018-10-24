package alexai18n

import (
  "github.com/arienmalec/alexa-go" 
  "github.com/qor/i18n"
  "github.com/qor/i18n/backends/yaml"
)

var (
  I18n *i18n.I18n
)

func WorldString(req alexa.Request,  stringID string) (string) {
  cultureId := req.Body.Locale
  templ := I18n.T(cultureId, stringID)
  return (string(templ))
}

func init() {
  I18n = i18n.New(
     yaml.New("./i18n.yml"),
  )
}

