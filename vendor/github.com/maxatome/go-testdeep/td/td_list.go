// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"reflect"

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/maxatome/go-testdeep/internal/flat"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdList struct {
	baseOKNil
	items []reflect.Value
}

func newList(items ...any) tdList {
	return tdList{
		baseOKNil: newBaseOKNil(4),
		items:     flat.Values(items),
	}
}

func (l *tdList) String() string {
	return util.SliceToBuffer(bytes.NewBufferString(l.GetLocation().Func), l.items).
		String()
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	"github.com/maxatome/go-testdeep/internal/flat"
>>>>>>> 5ce8c7613 (update vendored files)
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdList struct {
	baseOKNil
	items []reflect.Value
}

func newList(items ...interface{}) tdList {
	return tdList{
		baseOKNil: newBaseOKNil(4),
		items:     flat.Values(items),
	}
}

func (l *tdList) String() string {
	return util.SliceToBuffer(bytes.NewBufferString(l.GetLocation().Func), l.items).
		String()
}
<<<<<<< HEAD

func (l *tdList) uniqTypeBehind() reflect.Type {
	var (
		lastIfType, lastType, curType reflect.Type
		severalIfTypes                bool
	)

	//
	for _, item := range l.items {
		if !item.IsValid() {
			return nil // no need to go further
		}

		if item.Type().Implements(testDeeper) {
			curType = item.Interface().(TestDeep).TypeBehind()

			// Ignore unknown TypeBehind
			if curType == nil {
				continue
			}

			// Ignore interfaces & interface pointers too (see Isa), but
			// keep them in mind in case we encounter always the same
			// interface pointer
			if curType.Kind() == reflect.Interface ||
				(curType.Kind() == reflect.Ptr &&
					curType.Elem().Kind() == reflect.Interface) {
				if lastIfType == nil {
					lastIfType = curType
				} else if lastIfType != curType {
					severalIfTypes = true
				}
				continue
			}
		} else {
			curType = item.Type()
		}

		if lastType != curType {
			if lastType != nil {
				return nil
			}
			lastType = curType
		}
	}

	// Only one type found
	if lastType != nil {
		return lastType
	}

	// Only one interface type found
	if lastIfType != nil && !severalIfTypes {
		return lastIfType
	}
	return nil
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
||||||| parent of 5ce8c7613 (update vendored files)

func (l *tdList) uniqTypeBehind() reflect.Type {
	var (
		lastIfType, lastType, curType reflect.Type
		severalIfTypes                bool
	)

	//
	for _, item := range l.items {
		if !item.IsValid() {
			return nil // no need to go further
		}

		if item.Type().Implements(testDeeper) {
			curType = item.Interface().(TestDeep).TypeBehind()

			// Ignore unknown TypeBehind
			if curType == nil {
				continue
			}

			// Ignore interfaces & interface pointers too (see Isa), but
			// keep them in mind in case we encounter always the same
			// interface pointer
			if curType.Kind() == reflect.Interface ||
				(curType.Kind() == reflect.Ptr &&
					curType.Elem().Kind() == reflect.Interface) {
				if lastIfType == nil {
					lastIfType = curType
				} else if lastIfType != curType {
					severalIfTypes = true
				}
				continue
			}
		} else {
			curType = item.Type()
		}

		if lastType != curType {
			if lastType != nil {
				return nil
			}
			lastType = curType
		}
	}

	// Only one type found
	if lastType != nil {
		return lastType
	}

	// Only one interface type found
	if lastIfType != nil && !severalIfTypes {
		return lastIfType
	}
	return nil
}
=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	"github.com/maxatome/go-testdeep/internal/flat"
>>>>>>> 6b7ce455e (update vendored files)
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdList struct {
	baseOKNil
	items []reflect.Value
}

func newList(items ...interface{}) tdList {
	return tdList{
		baseOKNil: newBaseOKNil(4),
		items:     flat.Values(items),
	}
}

func (l *tdList) String() string {
	return util.SliceToBuffer(bytes.NewBufferString(l.GetLocation().Func), l.items).
		String()
}
<<<<<<< HEAD

func (l *tdList) uniqTypeBehind() reflect.Type {
	var (
		lastIfType, lastType, curType reflect.Type
		severalIfTypes                bool
	)

	//
	for _, item := range l.items {
		if !item.IsValid() {
			return nil // no need to go further
		}

		if item.Type().Implements(testDeeper) {
			curType = item.Interface().(TestDeep).TypeBehind()

			// Ignore unknown TypeBehind
			if curType == nil {
				continue
			}

			// Ignore interfaces & interface pointers too (see Isa), but
			// keep them in mind in case we encounter always the same
			// interface pointer
			if curType.Kind() == reflect.Interface ||
				(curType.Kind() == reflect.Ptr &&
					curType.Elem().Kind() == reflect.Interface) {
				if lastIfType == nil {
					lastIfType = curType
				} else if lastIfType != curType {
					severalIfTypes = true
				}
				continue
			}
		} else {
			curType = item.Type()
		}

		if lastType != curType {
			if lastType != nil {
				return nil
			}
			lastType = curType
		}
	}

	// Only one type found
	if lastType != nil {
		return lastType
	}

	// Only one interface type found
	if lastIfType != nil && !severalIfTypes {
		return lastIfType
	}
	return nil
}
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

func (l *tdList) uniqTypeBehind() reflect.Type {
	var (
		lastIfType, lastType, curType reflect.Type
		severalIfTypes                bool
	)

	//
	for _, item := range l.items {
		if !item.IsValid() {
			return nil // no need to go further
		}

		if item.Type().Implements(testDeeper) {
			curType = item.Interface().(TestDeep).TypeBehind()

			// Ignore unknown TypeBehind
			if curType == nil {
				continue
			}

			// Ignore interfaces & interface pointers too (see Isa), but
			// keep them in mind in case we encounter always the same
			// interface pointer
			if curType.Kind() == reflect.Interface ||
				(curType.Kind() == reflect.Ptr &&
					curType.Elem().Kind() == reflect.Interface) {
				if lastIfType == nil {
					lastIfType = curType
				} else if lastIfType != curType {
					severalIfTypes = true
				}
				continue
			}
		} else {
			curType = item.Type()
		}

		if lastType != curType {
			if lastType != nil {
				return nil
			}
			lastType = curType
		}
	}

	// Only one type found
	if lastType != nil {
		return lastType
	}

	// Only one interface type found
	if lastIfType != nil && !severalIfTypes {
		return lastIfType
	}
	return nil
}
=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"github.com/maxatome/go-testdeep/internal/flat"
>>>>>>> 4d7e5ad26 (update vendored files)
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdList struct {
	baseOKNil
	items []reflect.Value
}

func newList(items ...interface{}) tdList {
	return tdList{
		baseOKNil: newBaseOKNil(4),
		items:     flat.Values(items),
	}
}

func (l *tdList) String() string {
	return util.SliceToBuffer(bytes.NewBufferString(l.GetLocation().Func), l.items).
		String()
}
<<<<<<< HEAD

func (l *tdList) uniqTypeBehind() reflect.Type {
	var (
		lastIfType, lastType, curType reflect.Type
		severalIfTypes                bool
	)

	//
	for _, item := range l.items {
		if !item.IsValid() {
			return nil // no need to go further
		}

		if item.Type().Implements(testDeeper) {
			curType = item.Interface().(TestDeep).TypeBehind()

			// Ignore unknown TypeBehind
			if curType == nil {
				continue
			}

			// Ignore interfaces & interface pointers too (see Isa), but
			// keep them in mind in case we encounter always the same
			// interface pointer
			if curType.Kind() == reflect.Interface ||
				(curType.Kind() == reflect.Ptr &&
					curType.Elem().Kind() == reflect.Interface) {
				if lastIfType == nil {
					lastIfType = curType
				} else if lastIfType != curType {
					severalIfTypes = true
				}
				continue
			}
		} else {
			curType = item.Type()
		}

		if lastType != curType {
			if lastType != nil {
				return nil
			}
			lastType = curType
		}
	}

	// Only one type found
	if lastType != nil {
		return lastType
	}

	// Only one interface type found
	if lastIfType != nil && !severalIfTypes {
		return lastIfType
	}
	return nil
}
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

func (l *tdList) uniqTypeBehind() reflect.Type {
	var (
		lastIfType, lastType, curType reflect.Type
		severalIfTypes                bool
	)

	//
	for _, item := range l.items {
		if !item.IsValid() {
			return nil // no need to go further
		}

		if item.Type().Implements(testDeeper) {
			curType = item.Interface().(TestDeep).TypeBehind()

			// Ignore unknown TypeBehind
			if curType == nil {
				continue
			}

			// Ignore interfaces & interface pointers too (see Isa), but
			// keep them in mind in case we encounter always the same
			// interface pointer
			if curType.Kind() == reflect.Interface ||
				(curType.Kind() == reflect.Ptr &&
					curType.Elem().Kind() == reflect.Interface) {
				if lastIfType == nil {
					lastIfType = curType
				} else if lastIfType != curType {
					severalIfTypes = true
				}
				continue
			}
		} else {
			curType = item.Type()
		}

		if lastType != curType {
			if lastType != nil {
				return nil
			}
			lastType = curType
		}
	}

	// Only one type found
	if lastType != nil {
		return lastType
	}

	// Only one interface type found
	if lastIfType != nil && !severalIfTypes {
		return lastIfType
	}
	return nil
}
=======
>>>>>>> 4d7e5ad26 (update vendored files)
