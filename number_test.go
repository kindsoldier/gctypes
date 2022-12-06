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
// Number
//
type fooNumber struct {
    Bar Number `json:"bar"`
}

func convNumber(data string) (string, error) {
    var err error
    var j []byte

    var s fooNumber
    err = json.Unmarshal([]byte(data), &s)
    if err != nil {
        return string(j), err
    }

    j, err = json.Marshal(s)
    if err != nil {
        return string(j), err
    }
    fmt.Println("number data:", data, "result:", string(j))
    return string(j), err
}
// 123
func TestNumber11(t *testing.T) {
    data := `{ "bar": 123 }`
    p := `{"bar":123}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// "123"
func TestNumber15(t *testing.T) {
    data := `{ "bar": "123" }`
    p := `{"bar":123}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// 123.45
func TestNumber16(t *testing.T) {
    data := `{ "bar": 123.45 }`
    p := `{"bar":123.45}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// 123.457891012345
func TestNumber17(t *testing.T) {
    data := `{ "bar": 123.457891012345 }`
    p := `{"bar":123.457891012345}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// 123.45789101234567890123456789
func TestNumber18(t *testing.T) {
    data := `{ "bar": 123.45789101234567890123456789 }`
    p := `{"bar":123.45789101234568}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// 12345789101234567890123456789
func TestNumber19(t *testing.T) {
    data := `{ "bar": 12345789101234567890123456789 }`
    p := `{"bar":1.234578910123457e+28}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// 123457891012345
func TestNumber20(t *testing.T) {
    data := `{ "bar": 123457891012345 }`
    p := `{"bar":123457891012345}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// -123457891012345
func TestNumber21(t *testing.T) {
    data := `{ "bar": -123457891012345 }`
    p := `{"bar":-123457891012345}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// "-123457891012345"
func TestNumber22(t *testing.T) {
    data := `{ "bar": -123457891012345 }`
    p := `{"bar":-123457891012345}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}



// null
func TestNumber91(t *testing.T) {
    data := `{ "bar": null }`
    p := `{"bar":0}`
    j, err := convNumber(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// true, false
func TestNumber92(t *testing.T) {
    data := `{ "bar": true }`
    p := `{"bar":0}`
    j, err := convNumber(data)
    assert.NotEqual(t, p, j)
    assert.NotEqual(t, err, nil)
}
