
/*
 * Copyright 2020 Oleg Borodin  <borodin@unix7.org>
 */

package types

import (
    "fmt"
    "encoding/json"
    "errors"
    //"strconv"
)

//
// String
//
type String string

func (this String) MarshalJSON() ([]byte, error) {
    json, err := json.Marshal(string(this))
    return json, err
}

func (this *String) UnmarshalJSON(data []byte) error {
    var err error

    var ref interface{}
    err = json.Unmarshal(data, &ref)
    if err != nil {
        *this = String(string(data))
        return nil
    }

    switch v := ref.(type) {
        case nil:
            *this = String("")
            return nil
        case string:
            var res string
            err = json.Unmarshal(data, &res)
            *this = String(res)
            return err
        case int:
            *this = String(string(data))
            return nil
        case float64:
            *this = String(string(data))
            return nil
        case bool:
            *this = String(string(data))
            return nil
        default:
            return errors.New(fmt.Sprintf("unmarshall string: i don't know about type %T", v))
    }
    return nil
}
