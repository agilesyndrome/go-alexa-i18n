# Go i18n Support for Alexa Responses

First, create a `./i18n.yml` file in the root of your SKILL folder (not your Intent subfolders)


```
en-US:
  myintent:
    title: "Card Title"
    text: "Words words words read by Alexa"
```


In your intent Handler `.go` file...
```
import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-i18n/alexai18n"
)

func Handler(request alexa.Request) (alexa.Response, error) {

	title := alexai18n.WorldString(request, "myintent.title")
        text := alexai18n.WorldString(request, "myintent.text")
        return alexa.NewSimpleResponse(title, text), nil
}

```


## Multiple Cultures

In the i18n.yml file
```
en-US:
  myintent:
    title: "Card Title"
    text: "Hello World!"
es-ES:
  myintent:
    title: "Card Title"
    text: "Hola"
```

If a request comes into your skill, and you do not yet have the culture in your i18n.yml file, it will return the FIRST culture listed.
