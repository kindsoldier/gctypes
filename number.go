
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
// Number
//
type Number float64

func (this Number) MarshalJSON() ([]byte, error) {
    json, err := json.Marshal(float64(this))
    return json, err
}

func (this *Number) UnmarshalJSON(data []byte) error {
    var ref interface{}
    err := json.Unmarshal(data, &ref)
    if err != nil {
        *this = Number(0)
        return err
    }

    switch v := ref.(type) {
        case nil:
            *this = Number(0)
            return nil
        case string:
            var str string
            err := json.Unmarshal(data, &str)
            if err != nil {
                return err
            }
            var res float64
            err = json.Unmarshal([]byte(str), &res)
            *this = Number(res)
            return err
        case float64:
            var res float64
            err := json.Unmarshal(data, &res)
            *this = Number(res)
            return err
        case bool:
            *this = Number(0)
            //return nil
            return errors.New(fmt.Sprintf("unmarshall number: cannot convert bool value to number"))
        default:
            return errors.New(fmt.Sprintf("unmarshall number: i don't know about type %T", v))
    }
    return nil
}
