# Generation tool

Inspired by [go-cloudstack], which is entirely generated, this tool aims at
finding errors in the implementation of the interfaces. Because naming things
are hard, going full _code generation_ seems far fetched at the current times.

## Setup

Download the API definition from your CloudStack instance using the famous [cs cli tool][cs].

```console
$ cs listApis > listApis.json
```

## Find all errors

Reading the Go code and the JSON description (`-apis`) it lists the errors per struct.

```console
$ go run generate/main.go -apis listApis.json
...
```

## Find specific error

Then, inspect the errors of a particular command using `-cmd`

```
$ go run generate/main.go -apis listApis.json -cmd deleteSnapshot
tag:id: missing `doc:"The ID of the snapshot"`

snapshots.go:97.6: DeleteSnapshot has 1 error(s)
```

## Find errors in response

All commands response may also be checked, but you have to give the expected output type.

```
go run generate/main.go -apis listApis.json -cmd listApis -type API
```

Nothing more... nothing less...

## TODO

- Check that the `APIName()` (Go) matches the command name (JSON).
- Find the corresponding type for the responses.


[go-cloudstack]: https://github.com/xanzy/go-cloudstack
[cs]: https://pypi.org/project/cs/
