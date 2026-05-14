# spew

Originally copied from
[github.com/davecgh/go-spew](https://github.com/davecgh/go-spew), it
has been patched to fit go-testdeep needs as the original repository
seems to have no activity anymore.

spew implements a deep pretty printer for Go data structures to aid
in debugging.  A comprehensive suite of tests with 100% test coverage
is provided to ensure proper functionality. This directory and its
sub-directories are licensed under the [copyfree](http://copyfree.org)
[ISC License](internal/spew/LICENSE), so it may be used in open
source or commercial projects.

If you're interested in reading about how the original go-spew package
came to life and some of the challenges involved in providing a deep
pretty printer, there is a blog post about it
[here](https://web.archive.org/web/20160304013555/https://blog.cyphertite.com/go-spew-a-journey-into-dumping-go-data-structures/).

## Documentation

https://pkg.go.dev/github.com/maxatome/go-testdeep/internal/spew

## Quick Start

Add this import line to the file you're working in:

```go
import "github.com/maxatome/go-testdeep/internal/spew"
```

To dump a variable with full newlines, indentation, type and pointer
information use Sdump:

```go
str := spew.Sdump(myVar1)
```

## Sample Sdump Output

```
(main.Foo) {
 unexportedField: (*main.Bar)(0xf84002e210)({
  flag: (main.Flag) flagTwo,
  data: (uintptr) <nil>
 }),
 ExportedField: (map[interface {}]interface {}) {
  (string) "one": (bool) true
 }
}
([]uint8) {
 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20  |............... |
 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30  |!"#$%&'()*+,-./0|
 00000020  31 32                                             |12|
}
```

## Configuration Options

Configuration of spew is handled by fields in the ConfigState type. For
convenience, all of the top-level functions use a global state available via the
spew.Config global.

It is also possible to create a ConfigState instance that provides methods
equivalent to the top-level functions. This allows concurrent configuration
options. See the ConfigState documentation for more details.

```
* Indent
	String to use for each indentation level for Sdump function.
	It is a single space by default.  A popular alternative is "\t".

* MaxDepth
	Maximum number of levels to descend into nested data structures.
	There is no limit by default.

* DisableMethods
	Disables invocation of error and fmt.Stringer interface methods.
	Method invocation is enabled by default.

* DisablePointerMethods
	Disables invocation of error and fmt.Stringer interface methods on types
	which only accept pointer receivers from non-pointer variables.  This option
	relies on access to the unsafe package, so it will not have any effect when
	running in environments without access to the unsafe package such as Google
	App Engine or with the "safe" build tag specified.
	Pointer method invocation is enabled by default.

* DisablePointerAddresses
	DisablePointerAddresses specifies whether to disable the printing of
	pointer addresses. This is useful when diffing data structures in tests.

* EnableCapacities
	EnableCapacities specifies whether to enable the printing of capacities
	for arrays, slices and channels.
```

## Unsafe Package Dependency

This package relies on the unsafe package to perform some of the more advanced
features, however it also supports a "limited" mode which allows it to work in
environments where the unsafe package is not available.  By default, it will
operate in this mode on Google App Engine and when compiled with GopherJS.  The
"safe" build tag may also be specified to force the package to build without
using the unsafe package.

## License

This directory and its sub-directories are licensed under the
[copyfree](http://copyfree.org) [ISC License](internal/spew/LICENSE),
so it may be used in open source or commercial projects.
