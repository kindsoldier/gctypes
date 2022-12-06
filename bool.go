
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
// Bool
//
type Bool bool

func (this Bool) MarshalJSON() ([]byte, error) {
    json, err := json.Marshal(bool(this))
    return json, err
}

func (this *Bool) UnmarshalJSON(data []byte) error {
    var err error

   var ref interface{}
    err = json.Unmarshal(data, &ref)

    if err != nil {
        *this = false
        return nil
    }

    switch v := ref.(type) {
        case nil:
            *this = false
            return nil
        case string:
            var res string
            err := json.Unmarshal(data, &res)
            if err != nil {
                *this = false
                return err
            }
            switch res {
                case "true", "1":
                    *this = true
                case "false", "0", "-1":
                    *this = false
                default:
                    *this = false
                    return errors.New(fmt.Sprintf("unmarshall bool: cannot convert string value '%s' to bool", res))
            }
            return nil
        case float64:
            var res float64
            err := json.Unmarshal(data, &res)
            if err != nil {
                *this = false
                return err
            }
            switch res {
                case 1:
                    *this = true
                case 0, -1:
                    *this = false
                default:
                    *this = false
                    return errors.New(fmt.Sprintf("unmarshall bool: cannot convert number value %f to bool", res))
            }
            return nil
        case bool:
            var res bool
            err := json.Unmarshal(data, &res)
            *this = Bool(res)
            return err
        default:
            return errors.New(fmt.Sprintf("unmarshall bool: i don't know about type %T", v))
    }
    return nil
}
