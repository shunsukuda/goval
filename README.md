# goval
Go wrapper value

## Code Generation

```bash
go run ./_internal/template/gen.go
```

## Convert

```Numeric``` to ```Numeric```, ```ToXXXEq()``` returns:

|FromType-ToType      |Int8|Int16|Int32|Int64|Uint8|Uint16|Uint32|Uint64|Float32|Float64|Complex64|Complex128|
|:-:    |   :-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|**Int8**|(true)|true|true|true|0 <= e|0 <= e|0 <= e|0 <= e|true|true|true|true|
|**Int16**|minI8 <= e && e <= maxI8|(true)|true|true|0 <= e && e <= maxU8|0 <= e|0 <= e|0 <= e|true|true|true|true
|**Int32**|minI8 <= e && e <= maxI8|min16 <= e && e <= max16|(true)|true|0 <= e && e <= maxU8|0 <= e && e <= maxU16|0 <= e|0 <= e|minI24 <= e && e <= maxI24|true|minI24 <= e && e <= maxI24|true|
|**Int64**|minI8 <= e && e <= maxI8|minI16 <= e && e <= maxI16|minI32 <= e && e <= maxI32|(true)|0 <= e && e <= maxU8|0 <= e && e <= maxU16|0 <= e && e <= maxU32|0 <= e|minI24 <= e && e <= maxI24|minI53 <= e && e <= maxI53|minI24 <= e && e <= maxI24|minI53 <= e && e <= maxI53|
|**Uint8**|e <= maxI8|true|true|true|(true)|true|true|true|true|true|true|true
|**Uint16**|e <= maxI8|e <= maxI16|true|true|e <= maxU8|(true)|true|true|true|true|true|true
|**Uint32**|e <= maxI8|e <= maxI16|e <= maxI32|true|e <= maxU8|e <= maxU16|(ture)|true|e <= maxI24|true|e <= maxI24|true|
|**Uint64**|e <=  maxI8|e <=  maxI16|e <= maxI32|e <= maxI64|e <= maxU8|e <= maxU16|e <= maxU32|(true)|e <= maxI24|e <= maxI53|e <=maxI24|e <= maxI53|
|**Float32**|minI8 <= e && e <= maxI8 && IsInt|minI16 <= e && e <= maxI16 && IsInt|minI32 <= e && e <= maxI32 && IsInt|minI64 <= e && e <= maxI64 && IsInt|minU8 <= e && e <= maxU8 && IsInt|minU16 <= e && e <= maxU16 && IsInt|minU32 <= e && e <= maxU32 && IsInt|minU64 <= e && e <= maxU64 && IsInt|(true)|false|true|false|
|**Float64**|minI8 <= e && e <= maxI8 && IsInt|minI16 <= e && e <= maxI16 && IsInt|minI32 <= e && e <= maxI32 && IsInt|minI64 <= e && e <= maxI64 && IsInt|minU8 <= e && e <= maxU8 && IsInt|minU16 <= e && e <= maxU16 && IsInt|minU32 <= e && e <= maxU32 && IsInt|minU64 <= e && e <= maxU64 && IsInt|false|(true)|false|true|
|**Complex64**|minI8 <= e && e <= maxI8 && IsInt|minI16 <= e && e <= maxI16 && IsInt|minI32 <= e && e <= maxI32 && IsInt|minI64 <= e && e <= maxI64 && IsInt|minU8 <= e && e <= maxU8 && IsInt|minU16 <= e && e <= maxU16 && IsInt|minU32 <= e && e <= maxU32 && IsInt|minU64 <= e && e <= maxU64 && IsInt|imag(e) == 0|false|(true)|false|
|**Complex128**|minI8 <= e && e <= maxI8 && IsInt|minI16 <= e && e <= maxI16 && IsInt|minI32 <= e && e <= maxI32 && IsInt|minI64 <= e && e <= maxI64 && IsInt|minU8 <= e && e <= maxU8 && IsInt|minU16 <= e && e <= maxU16 && IsInt|minU32 <= e && e <= maxU32 && IsInt|minU64 <= e && e <= maxU64 && IsInt|false|imag(e) == 0|false|(true)|

```IsInt: e == e.ToInt().ToFloat()```