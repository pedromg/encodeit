# encodeit
small cli program to encode a value from an io.Reader source into an encoding (default Base64)


### Usage

```
$ ./encodeit [-source=/path/to/file] [-encoding=base64] text
```

Examples:

```
$ ./encodeit hash_me
$ ./encodeit "hash me too"
$ ./encodeit -source=path/to/file
$ ./encodeit -encoding=base32 me@example.com
```


Params:

-	__-source__: (string) the name of a file to read the first line from;
-	__-encoding__: the encoding to encode to, examples: base32, base64 (default)


### Tests

`$ go test`


### Debug

[Delve](https://github.com/derekparker/delve) can be used to debug, add `runtime.Breakpoint()` on the code, or add breakpoints manually on the Delve CLI


### Cross Compile

If you are building on OSX for Linux usage, make sure your Go is prepared to generate binaries for other architectures. To enable it for Linux:

```
$ cd  $GOROOT/src
$ GOOS=linux GOARCH=386 ./make.bash
```
Then to generate a linux specific binary:
```
$ GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o encodeit.linux encodeit.go
```

