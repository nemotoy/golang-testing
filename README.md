# golang-testing

## Result

```bash
$ go test -v ./... -bench .
```

### int

|method  |  |  |
|---|---|---|
|reflect.DeepEqual()  |20000000  |85.0 ns/op  |
|assert.EqualValues()  |1000000  |1841 ns/op  |
|cmp.Equal()  |5000000  |371 ns/op  |
|is.Equal()  |10000000  |189 ns/op  |

### int and interface

|method  |  |  |
|---|---|---|
|reflect.DeepEqual()  |20000000  |72.4 ns/op  |
|assert.EqualValues()  |1000000  |1775 ns/op  |
|cmp.Equal()  |5000000  |359 ns/op  |
|is.Equal()  |10000000  |173 ns/op  |

## References

* [pkg/testing](https://golang.org/pkg/testing)
* [Testing_flags](https://golang.org/cmd/go/#hdr-Testing_flags)
