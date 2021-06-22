// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

// Package td (aka go-testdeep) allows extremely flexible deep
// comparison, it is built for testing.
//
// It is a go rewrite and adaptation of wonderful Test::Deep perl
// module (see https://metacpan.org/pod/Test::Deep).
//
// In golang, comparing data structure is usually done using
// reflect.DeepEqual or using a package that uses this function behind
// the scene.
//
// This function works very well, but it is not flexible. Both
// compared structures must match exactly.
//
// The purpose of td package is to do its best to introduce this
// missing flexibility using "operators" when the expected value (or
// one of its component) cannot be matched exactly.
//
// See https://go-testdeep.zetta.rocks/ for details, and for easy HTTP
// API testing, see the tdhttp helper
// https://pkg.go.dev/github.com/maxatome/go-testdeep/helpers/tdhttp
//
// Example of use
//
// Imagine a function returning a struct containing a newly created
// database record. The Id and the CreatedAt fields are set by the
// database layer:
//
//   type Record struct {
//     Id        uint64
//     Name      string
//     Age       int
//     CreatedAt time.Time
//   }
//
//   func CreateRecord(name string, age int) (*Record, error) {
//     // Do INSERT INTO … and return newly created record or error if it failed
//   }
//
// Using standard testing package
//
// To check the freshly created record contents using standard testing
// package, we have to do something like that:
//
//   import (
//     "testing"
//     "time"
//   )
//
//   func TestCreateRecord(t *testing.T) {
//     before := time.Now().Truncate(time.Second)
//     record, err := CreateRecord()
//
//     if err != nil {
//       t.Errorf("An error occurred: %s", err)
//     } else {
//       expected := Record{Name: "Bob", Age: 23}
//
//       if record.Id == 0 {
//         t.Error("Id probably not initialized")
//       }
//       if before.After(record.CreatedAt) ||
//         time.Now().Before(record.CreatedAt) {
//         t.Errorf("CreatedAt field not expected: %s", record.CreatedAt)
//       }
//       if record.Name != expected.Name {
//         t.Errorf("Name field differs, got=%s, expected=%s",
//           record.Name, expected.Name)
//       }
//       if record.Age != expected.Age {
//         t.Errorf("Age field differs, got=%s, expected=%s",
//           record.Age, expected.Age)
//       }
//     }
//   }
//
// Using basic go-testdeep approach
//
// td package, via its Cmp* functions, handles the tests and all the
// error message boiler plate. Let's do it:
//
//   import (
//     "testing"
//     "time"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestCreateRecord(t *testing.T) {
//     before := time.Now().Truncate(time.Second)
//     record, err := CreateRecord()
//
//     if td.CmpNoError(t, err) {
//       td.Cmp(t, record.Id, td.NotZero(), "Id initialized")
//       td.Cmp(t, record.Name, "Bob")
//       td.Cmp(t, record.Age, 23)
//       td.Cmp(t, record.CreatedAt, td.Between(before, time.Now()))
//     }
//   }
//
// As we cannot guess the Id field value before its creation, we use
// the NotZero operator to check it is set by CreateRecord() call. The
// same it true for the creation date field CreatedAt. Thanks to the
// Between operator we can check it is set with a value included
// between the date before CreateRecord() call and the date just
// after.
//
// Note that if `Id` and `CreateAt` could be known in advance, we could
// simply do:
//
//   import (
//     "testing"
//     "time"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestCreateRecord(t *testing.T) {
//     before := time.Now().Truncate(time.Second)
//     record, err := CreateRecord()
//
//     if td.CmpNoError(t, err) {
//       td.Cmp(t, record, &Record{
//         Id:        1234,
//         Name:      "Bob",
//         Age:       23,
//         CreatedAt: time.Date(2019, time.May, 1, 12, 13, 14, 0, time.UTC),
//      })
//     }
//   }
//
// But unfortunately, it is common to not know exactly the value of some
// fields…
//
// Using advanced go-testdeep technique
//
// Of course we can test struct fields one by one, but with go-testdeep,
// the whole struct can be compared with one Cmp call.
//
//   import (
//     "testing"
//     "time"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestCreateRecord(t *testing.T) {
//     before := time.Now().Truncate(time.Second)
//     record, err := CreateRecord()
//
//     if td.CmpNoError(t, err) {
//       td.Cmp(t, record,
//         td.Struct(
//           &Record{
//             Name: "Bob",
//             Age:  23,
//           },
//           td.StructFields{
//             "Id":        td.NotZero(),
//             "CreatedAt": td.Between(before, time.Now()),
//           }),
//         "Newly created record")
//     }
//   }
//
// See the use of the Struct operator. It is needed here to overcome
// the go static typing system and so use other go-testdeep operators
// for some fields, here NotZero and Between.
//
// Not only structs can be compared. A lot of operators can
// be found below to cover most (all?) needed tests. See
// https://pkg.go.dev/github.com/maxatome/go-testdeep/td#TestDeep
//
// Using go-testdeep Cmp shortcuts
//
// The Cmp function is the keystone of this package, but to make
// the writing of tests even easier, the family of Cmp* functions are
// provided and act as shortcuts. Using CmpStruct function, the
// previous example can be written as:
//
//   import (
//     "testing"
//     "time"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestCreateRecord(t *testing.T) {
//     before := time.Now().Truncate(time.Second)
//     record, err := CreateRecord()
//
//     if td.CmpNoError(t, err) {
//       td.CmpStruct(t, record,
//         &Record{
//           Name: "Bob",
//           Age:  23,
//         },
//         td.StructFields{
//           "Id":        td.NotZero(),
//           "CreatedAt": td.Between(before, time.Now()),
//         },
//         "Newly created record")
//     }
//   }
//
// Using T type
//
// testing.T can be encapsulated in td.T type, simplifying again the
// test:
//
//   import (
//     "testing"
//     "time"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestCreateRecord(tt *testing.T) {
//     t := td.NewT(tt)
//
//     before := time.Now().Truncate(time.Second)
//     record, err := CreateRecord()
//
//     if t.CmpNoError(err) {
//       t.RootName("RECORD").Struct(record,
//         &Record{
//           Name: "Bob",
//           Age:  23,
//         },
//         td.StructFields{
//           "Id":        td.NotZero(),
//           "CreatedAt": td.Between(before, time.Now()),
//         },
//         "Newly created record")
//     }
//   }
//
// Note the use of RootName method, it allows to name what we are
// going to test, instead of the default "DATA".
//
// A step further with operator anchoring
//
// Overcome the go static typing system using the Struct operator is
// sometimes heavy. Especially when structs are nested, as the Struct
// operator needs to be used for each level surrounding the level in
// which an operator is involved. Operator anchoring feature has been
// designed to avoid this heaviness:
//
//   import (
//     "testing"
//     "time"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestCreateRecord(tt *testing.T) {
//     before := time.Now().Truncate(time.Second)
//     record, err := CreateRecord()
//
//     t := td.NewT(tt) // operator anchoring needs a *td.T instance
//
//     if t.CmpNoError(err) {
//       t.Cmp(record,
//         &Record{
//           Name:      "Bob",
//           Age:       23,
//           ID:        t.A(td.NotZero(), uint64(0)).(uint64),
//           CreatedAt: t.A(td.Between(before, time.Now())).(time.Time),
//         },
//         "Newly created record")
//     }
//   }
//
// See the A method (or its full name alias Anchor) documentation for
// details.
package td // import "github.com/maxatome/go-testdeep/td"
