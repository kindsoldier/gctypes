/*
 * Copyright 2020 Oleg Borodin  <borodin@unix7.org>
 */

package types

import (
    "testing"
    "fmt"
    "encoding/json"
    "github.com/stretchr/testify/assert"
)

//
// Bool
//
type fooBool struct {
    Bar Bool `json:"bar"`
}

func convBool(data string) (string, error) {
    var err error
    var j []byte

    var s fooBool
    err = json.Unmarshal([]byte(data), &s)
    if err != nil {
        return string(j), err
    }

    j, err = json.Marshal(s)
    if err != nil {
        return string(j), err
    }
    fmt.Println("bool data:", data, "result:", string(j))
    return string(j), err
}
// true, false
func TestBool11(t *testing.T) {
    data := `{ "bar": true }`
    p := `{"bar":true}`
    j, err := convBool(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

func TestBool12(t *testing.T) {
    data := `{ "bar": false }`
    p := `{"bar":false}`
    j, err := convBool(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// "1", "0"
func TestBool21(t *testing.T) {
    data := `{ "bar": "1" }`
    p := `{"bar":true}`
    j, err := convBool(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

func TestBool22(t *testing.T) {
    data := `{ "bar": "0" }`
    p := `{"bar":false}`
    j, err := convBool(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// "true", "false"
func TestBool31(t *testing.T) {
    data := `{ "bar": "true" }`
    p := `{"bar":true}`
    j, err := convBool(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

func TestBool32(t *testing.T) {
    data := `{ "bar": "false" }`
    p := `{"bar":false}`
    j, err := convBool(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// null
func TestBool91(t *testing.T) {
    data := `{ "bar": null }`
    p := `{"bar":false}`
    j, err := convBool(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

func TestBool92(t *testing.T) {
    data := `{ "bar": null }`
    p := `{"bar":true}`
    j, err := convBool(data)
    assert.NotEqual(t, p, j)
    assert.Equal(t, err, nil)
}
