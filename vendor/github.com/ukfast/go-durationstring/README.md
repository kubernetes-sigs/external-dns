# durationstring

A simple Go package for working with string format durations, e.g. `1d4h`

## Usage

#### Parsing

```go
d, err :=  durationstring.Parse("1d4h5ns")

assert.Equal(t, 1, d.Days)
assert.Equal(t, 4, d.Hours)
assert.Equal(t, 5, d.Nanoseconds)
```

#### String formatting

```go
d := durationstring.NewDuration(1, 0, 0, 4, 0, 0, 0, 0, 0)
s := d.String()

assert.Equal(t, "1y4h", s)
```
