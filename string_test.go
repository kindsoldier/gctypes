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
// String
//
type fooString struct {
    Bar String `json:"bar"`
}

func convString(data string) (string, error) {
    var err error
    var j []byte

    var s fooString
    err = json.Unmarshal([]byte(data), &s)
    if err != nil {
        return string(j), err
    }

    j, err = json.Marshal(s)
    if err != nil {
        return string(j), err
    }
    fmt.Println("string data:", data, "result:", string(j))
    return string(j), err
}

// true, false
func TestString11(t *testing.T) {
    data := `{ "bar": true }`
    p := `{"bar":"true"}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

func TestString12(t *testing.T) {
    data := `{ "bar": false }`
    p := `{"bar":"false"}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// "12", "02"
func TestString21(t *testing.T) {
    data := `{ "bar": 12 }`
    p := `{"bar":"12"}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

//func TestString22(t *testing.T) {
//    data := `{ "bar": 02 }`
//    p := `{"bar":"02"}`
//    j, err := convString(data)
//    assert.Equal(t, p, j)
//    assert.Equal(t, err, nil)
//}

// "12.1", "02.1"
func TestString24(t *testing.T) {
    data := `{ "bar": 12.1 }`
    p := `{"bar":"12.1"}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// msg:"invalid character '2' after object key:value pair"
//func TestString25(t *testing.T) {
//    data := `{ "bar": 02.1 }`
//    p := `{"bar":"02.1"}`
//    j, err := convString(data)
//    assert.Equal(t, p, j)
//    assert.Equal(t, err, nil)
//}


// "true", "false"
func TestString31(t *testing.T) {
    data := `{ "bar": "true" }`
    p := `{"bar":"true"}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

func TestString32(t *testing.T) {
    data := `{ "bar": "false" }`
    p := `{"bar":"false"}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// null
func TestString91(t *testing.T) {
    data := `{ "bar": null }`
    p := `{"bar":""}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// NaN
//func TestString92(t *testing.T) {
//    data := `{ "bar": NaN }`
//    p := `{"bar":"null"}`
//    j, err := convString(data)
//    assert.Equal(t, p, j)
//    assert.Equal(t, err, nil)
//}

// "foo bar"
func TestString99(t *testing.T) {
    data := `{ "bar": "foo bar" }`
    p := `{"bar":"foo bar"}`
    j, err := convString(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
