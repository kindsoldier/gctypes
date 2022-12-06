# Custom types

Производные типы String, Bool, Number от string, bool, float64
и их трансляторы из/в JSON
для глючных синтаксически изуродованных одноэсниками структур

UPD Для сохранения семантики типов добавил Int & Float, оба подтипы float64

```
bool data: { "bar": true } result: {"bar":true}
bool data: { "bar": false } result: {"bar":false}
bool data: { "bar": "1" } result: {"bar":true}
bool data: { "bar": "0" } result: {"bar":false}
bool data: { "bar": "true" } result: {"bar":true}
bool data: { "bar": "false" } result: {"bar":false}
bool data: { "bar": null } result: {"bar":false}
bool data: { "bar": null } result: {"bar":false}
number data: { "bar": 123 } result: {"bar":123}
number data: { "bar": "123" } result: {"bar":123}
number data: { "bar": 123.45 } result: {"bar":123.45}
number data: { "bar": 123.457891012345 } result: {"bar":123.457891012345}
number data: { "bar": 123.45789101234567890123456789 } result: {"bar":123.45789101234568}
number data: { "bar": 12345789101234567890123456789 } result: {"bar":1.234578910123457e+28}
number data: { "bar": 123457891012345 } result: {"bar":123457891012345}
number data: { "bar": -123457891012345 } result: {"bar":-123457891012345}
number data: { "bar": -123457891012345 } result: {"bar":-123457891012345}
number data: { "bar": null } result: {"bar":0}
string data: { "bar": true } result: {"bar":"true"}
string data: { "bar": false } result: {"bar":"false"}
string data: { "bar": 12 } result: {"bar":"12"}
string data: { "bar": 12.1 } result: {"bar":"12.1"}
string data: { "bar": "true" } result: {"bar":"true"}
string data: { "bar": "false" } result: {"bar":"false"}
string data: { "bar": null } result: {"bar":""}
string data: { "bar": "foo bar" } result: {"bar":"foo bar"}
PASS
ok  	e2api/types	0.006s
```