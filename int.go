
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
// Int
//
type Int float64

func (this Int) MarshalJSON() ([]byte, error) {
    json, err := json.Marshal(float64(this))
    return json, err
}

func (this *Int) UnmarshalJSON(data []byte) error {
    var ref interface{}
    err := json.Unmarshal(data, &ref)
    if err != nil {
        *this = Int(0)
        return err
    }

    switch v := ref.(type) {
        case nil:
            *this = Int(0)
            return nil
        case string:
            var str string
            err := json.Unmarshal(data, &str)
            if err != nil {
                return err
            }
            var res float64
            err = json.Unmarshal([]byte(str), &res)
            *this = Int(res)
            return err
        case float64:
            var res float64
            err := json.Unmarshal(data, &res)
            *this = Int(res)
            return err
        case bool:
            *this = Int(0)
            //return nil
            return errors.New(fmt.Sprintf("unmarshall int: cannot convert bool value to int"))
        default:
            return errors.New(fmt.Sprintf("unmarshall int: i don't know about type %T", v))
    }
    return nil
}
