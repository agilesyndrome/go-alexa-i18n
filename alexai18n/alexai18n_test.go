package alexai18n

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestENUSJsonText(t *testing.T) {
     request := CultureRequest("en-US")
     response := WorldString(request, "demo.hello")
     response2 := WorldString(request, "demo.three")
     assert.Equal(t, "Hello, world", response)
     assert.Equal(t, "Three!", response2)
}

func TestESESJsonText(t *testing.T) {
     request := CultureRequest("es-ES")
     response := WorldString(request, "demo.hello")
     assert.Equal(t, "Hola!", response)
}

func TestNonExistingCulture(t *testing.T) {
     request := CultureRequest("xx-XX")
     response := WorldString(request, "demo.hello")
     assert.Equal(t, "Hello, world", response)
}

func TestKeyNotExist(t *testing.T) {
     request := CultureRequest("en-US")
     response := WorldString(request, "key-not-exist")
     assert.Equal(t, "key-not-exist", response)
}

