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
// Float
//
type fooFloat struct {
    Bar Float `json:"bar"`
}

func convFloat(data string) (string, error) {
    var err error
    var j []byte

    var s fooFloat
    err = json.Unmarshal([]byte(data), &s)
    if err != nil {
        return string(j), err
    }

    j, err = json.Marshal(s)
    if err != nil {
        return string(j), err
    }
    fmt.Println("float data:", data, "result:", string(j))
    return string(j), err
}
// 123
func TestFloat11(t *testing.T) {
    data := `{ "bar": 123 }`
    p := `{"bar":123}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// "123"
func TestFloat15(t *testing.T) {
    data := `{ "bar": "123" }`
    p := `{"bar":123}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// 123.45
func TestFloat16(t *testing.T) {
    data := `{ "bar": 123.45 }`
    p := `{"bar":123.45}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}
// 123.457891012345
func TestFloat17(t *testing.T) {
    data := `{ "bar": 123.457891012345 }`
    p := `{"bar":123.457891012345}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// 123.45789101234567890123456789
func TestFloat18(t *testing.T) {
    data := `{ "bar": 123.45789101234567890123456789 }`
    p := `{"bar":123.45789101234568}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// 12345789101234567890123456789
func TestFloat19(t *testing.T) {
    data := `{ "bar": 12345789101234567890123456789 }`
    p := `{"bar":1.234578910123457e+28}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// 123457891012345
func TestFloat20(t *testing.T) {
    data := `{ "bar": 123457891012345 }`
    p := `{"bar":123457891012345}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// -123457891012345
func TestFloat21(t *testing.T) {
    data := `{ "bar": -123457891012345 }`
    p := `{"bar":-123457891012345}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// "-123457891012345"
func TestFloat22(t *testing.T) {
    data := `{ "bar": -123457891012345 }`
    p := `{"bar":-123457891012345}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}



// null
func TestFloat91(t *testing.T) {
    data := `{ "bar": null }`
    p := `{"bar":0}`
    j, err := convFloat(data)
    assert.Equal(t, p, j)
    assert.Equal(t, err, nil)
}

// true, false
func TestFloat92(t *testing.T) {
    data := `{ "bar": true }`
    p := `{"bar":0}`
    j, err := convFloat(data)
    assert.NotEqual(t, p, j)
    assert.NotEqual(t, err, nil)
}
