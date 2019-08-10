# golang-testing

## Result

### int

```bash
$ go test -v ./cmd/int_test.go -bench .
```

|method  |  |  |
|---|---|---|
|reflect.DeepEqual()  |20000000  |85.0 ns/op  |
|assert.EqualValues()  |1000000  |1841 ns/op  |
|cmp.Equal()  |5000000  |371 ns/op  |
|is.Equal()  |10000000  |189 ns/op  |

### int and interface

```bash
$ go test -v ./cmd/interface_int_test.go -bench .
```

|method  |  |  |
|---|---|---|
|reflect.DeepEqual()  |20000000  |72.4 ns/op  |
|assert.EqualValues()  |1000000  |1775 ns/op  |
|cmp.Equal()  |5000000  |359 ns/op  |
|is.Equal()  |10000000  |173 ns/op  |

### struct and interface

```bash
$ go test -v ./cmd/interface_struct_test.go -bench .
```

|method  |type  |  |  |
|---|---|---|---|
|reflect.DeepEqual()  |value |5000000  |257 ns/op  |
|reflect.DeepEqual()  |pointer |30000000  |46.6 ns/op  |
|assert.EqualValues()  |value |500000  |2013 ns/op  |
|assert.EqualValues()  |pointer |1000000  |1742 ns/op  |
|cmp.Equal()  |value |1000000  |1301 ns/op  |
|cmp.Equal()  |pointer |1000000  |1509 ns/op  |
|is.Equal()  |value |5000000  |364 ns/op  |
|is.Equal()  |pointer |10000000  |152 ns/op  |

## References

* [pkg/testing](https://golang.org/pkg/testing)
* [Testing_flags](https://golang.org/cmd/go/#hdr-Testing_flags)
