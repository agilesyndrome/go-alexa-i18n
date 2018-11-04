# Go i18n Support for Alexa Responses

First, create a `./i18.yml` file in the root of your SKILL folder (not your Intent subfolders)


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
