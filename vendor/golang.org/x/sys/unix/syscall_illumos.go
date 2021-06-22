<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// illumos system calls not present on Solaris.

//go:build amd64 && illumos
// +build amd64,illumos

package unix

import (
	"unsafe"
)

func bytes2iovec(bs [][]byte) []Iovec {
	iovecs := make([]Iovec, len(bs))
	for i, b := range bs {
		iovecs[i].SetLen(len(b))
		if len(b) > 0 {
			iovecs[i].Base = &b[0]
		} else {
			iovecs[i].Base = (*byte)(unsafe.Pointer(&_zero))
		}
	}
	return iovecs
}

//sys	readv(fd int, iovs []Iovec) (n int, err error)

func Readv(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = readv(fd, iovecs)
	return n, err
}

//sys	preadv(fd int, iovs []Iovec, off int64) (n int, err error)

func Preadv(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = preadv(fd, iovecs, off)
	return n, err
}

//sys	writev(fd int, iovs []Iovec) (n int, err error)

func Writev(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = writev(fd, iovecs)
	return n, err
}

//sys	pwritev(fd int, iovs []Iovec, off int64) (n int, err error)

func Pwritev(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = pwritev(fd, iovecs, off)
	return n, err
}

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4

func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	nfd, err = accept4(fd, &rsa, &len, flags)
	if err != nil {
		return
	}
	if len > SizeofSockaddrAny {
		panic("RawSockaddrAny too small")
	}
	sa, err = anyToSockaddr(fd, &rsa)
	if err != nil {
		Close(nfd)
		nfd = 0
	}
	return
}
<<<<<<< HEAD

//sys	putmsg(fd int, clptr *strbuf, dataptr *strbuf, flags int) (err error)

func Putmsg(fd int, cl []byte, data []byte, flags int) (err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Len: int32(len(cl)),
			Buf: (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Len: int32(len(data)),
			Buf: (*int8)(unsafe.Pointer(&data[0])),
		}
	}
	return putmsg(fd, clp, datap, flags)
}

//sys	getmsg(fd int, clptr *strbuf, dataptr *strbuf, flags *int) (err error)

func Getmsg(fd int, cl []byte, data []byte) (retCl []byte, retData []byte, flags int, err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Maxlen: int32(len(cl)),
			Buf:    (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Maxlen: int32(len(data)),
			Buf:    (*int8)(unsafe.Pointer(&data[0])),
		}
	}

	if err = getmsg(fd, clp, datap, &flags); err != nil {
		return nil, nil, 0, err
	}

	if len(cl) > 0 {
		retCl = cl[:clp.Len]
	}
	if len(data) > 0 {
		retData = data[:datap.Len]
	}
	return retCl, retData, flags, nil
}

func IoctlSetIntRetInt(fd int, req uint, arg int) (int, error) {
	return ioctlRet(fd, req, uintptr(arg))
}

func IoctlSetString(fd int, req uint, val string) error {
	bs := make([]byte, len(val)+1)
	copy(bs[:len(bs)-1], val)
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&bs[0])))
	runtime.KeepAlive(&bs[0])
	return err
}

// Lifreq Helpers

func (l *Lifreq) SetName(name string) error {
	if len(name) >= len(l.Name) {
		return fmt.Errorf("name cannot be more than %d characters", len(l.Name)-1)
	}
	for i := range name {
		l.Name[i] = int8(name[i])
	}
	return nil
}

func (l *Lifreq) SetLifruInt(d int) {
	*(*int)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruInt() int {
	return *(*int)(unsafe.Pointer(&l.Lifru[0]))
}

func (l *Lifreq) SetLifruUint(d uint) {
	*(*uint)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruUint() uint {
	return *(*uint)(unsafe.Pointer(&l.Lifru[0]))
}

func IoctlLifreq(fd int, req uint, l *Lifreq) error {
	return ioctl(fd, req, uintptr(unsafe.Pointer(l)))
}

// Strioctl Helpers

func (s *Strioctl) SetInt(i int) {
	s.Len = int32(unsafe.Sizeof(i))
	s.Dp = (*int8)(unsafe.Pointer(&i))
}

func IoctlSetStrioctlRetInt(fd int, req uint, s *Strioctl) (int, error) {
	return ioctlRet(fd, req, uintptr(unsafe.Pointer(s)))
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2009 The Go Authors. All rights reserved.
||||||| parent of 5ce8c7613 (update vendored files)
// Copyright 2009 The Go Authors. All rights reserved.
=======
// Copyright 2021 The Go Authors. All rights reserved.
>>>>>>> 5ce8c7613 (update vendored files)
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// illumos system calls not present on Solaris.

//go:build amd64 && illumos
// +build amd64,illumos

package unix

import (
	"fmt"
	"runtime"
	"unsafe"
)

func bytes2iovec(bs [][]byte) []Iovec {
	iovecs := make([]Iovec, len(bs))
	for i, b := range bs {
		iovecs[i].SetLen(len(b))
		if len(b) > 0 {
			// somehow Iovec.Base on illumos is (*int8), not (*byte)
			iovecs[i].Base = (*int8)(unsafe.Pointer(&b[0]))
		} else {
			iovecs[i].Base = (*int8)(unsafe.Pointer(&_zero))
		}
	}
	return iovecs
}

//sys	readv(fd int, iovs []Iovec) (n int, err error)

func Readv(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = readv(fd, iovecs)
	return n, err
}

//sys	preadv(fd int, iovs []Iovec, off int64) (n int, err error)

func Preadv(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = preadv(fd, iovecs, off)
	return n, err
}

//sys	writev(fd int, iovs []Iovec) (n int, err error)

func Writev(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = writev(fd, iovecs)
	return n, err
}

//sys	pwritev(fd int, iovs []Iovec, off int64) (n int, err error)

func Pwritev(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = pwritev(fd, iovecs, off)
	return n, err
}

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4

func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	nfd, err = accept4(fd, &rsa, &len, flags)
	if err != nil {
		return
	}
	if len > SizeofSockaddrAny {
		panic("RawSockaddrAny too small")
	}
	sa, err = anyToSockaddr(fd, &rsa)
	if err != nil {
		Close(nfd)
		nfd = 0
	}
	return
}

//sys	putmsg(fd int, clptr *strbuf, dataptr *strbuf, flags int) (err error)

func Putmsg(fd int, cl []byte, data []byte, flags int) (err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Len: int32(len(cl)),
			Buf: (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Len: int32(len(data)),
			Buf: (*int8)(unsafe.Pointer(&data[0])),
		}
	}
	return putmsg(fd, clp, datap, flags)
}

//sys	getmsg(fd int, clptr *strbuf, dataptr *strbuf, flags *int) (err error)

func Getmsg(fd int, cl []byte, data []byte) (retCl []byte, retData []byte, flags int, err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Maxlen: int32(len(cl)),
			Buf:    (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Maxlen: int32(len(data)),
			Buf:    (*int8)(unsafe.Pointer(&data[0])),
		}
	}

	if err = getmsg(fd, clp, datap, &flags); err != nil {
		return nil, nil, 0, err
	}

	if len(cl) > 0 {
		retCl = cl[:clp.Len]
	}
	if len(data) > 0 {
		retData = data[:datap.Len]
	}
	return retCl, retData, flags, nil
}

func IoctlSetIntRetInt(fd int, req uint, arg int) (int, error) {
	return ioctlRet(fd, req, uintptr(arg))
}

func IoctlSetString(fd int, req uint, val string) error {
	bs := make([]byte, len(val)+1)
	copy(bs[:len(bs)-1], val)
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&bs[0])))
	runtime.KeepAlive(&bs[0])
	return err
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// Lifreq Helpers

func (l *Lifreq) SetName(name string) error {
	if len(name) >= len(l.Name) {
		return fmt.Errorf("name cannot be more than %d characters", len(l.Name)-1)
	}
	for i := range name {
		l.Name[i] = int8(name[i])
	}
	return nil
}

func (l *Lifreq) SetLifruInt(d int) {
	*(*int)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruInt() int {
	return *(*int)(unsafe.Pointer(&l.Lifru[0]))
}

func (l *Lifreq) SetLifruUint(d uint) {
	*(*uint)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruUint() uint {
	return *(*uint)(unsafe.Pointer(&l.Lifru[0]))
}

func IoctlLifreq(fd int, req uint, l *Lifreq) error {
	return ioctl(fd, req, uintptr(unsafe.Pointer(l)))
}

// Strioctl Helpers

func (s *Strioctl) SetInt(i int) {
	s.Len = int32(unsafe.Sizeof(i))
	s.Dp = (*int8)(unsafe.Pointer(&i))
}

func IoctlSetStrioctlRetInt(fd int, req uint, s *Strioctl) (int, error) {
	return ioctlRet(fd, req, uintptr(unsafe.Pointer(s)))
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2009 The Go Authors. All rights reserved.
||||||| parent of 6b7ce455e (update vendored files)
// Copyright 2009 The Go Authors. All rights reserved.
=======
// Copyright 2021 The Go Authors. All rights reserved.
>>>>>>> 6b7ce455e (update vendored files)
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// illumos system calls not present on Solaris.

//go:build amd64 && illumos
// +build amd64,illumos

package unix

import (
	"fmt"
	"runtime"
	"unsafe"
)

func bytes2iovec(bs [][]byte) []Iovec {
	iovecs := make([]Iovec, len(bs))
	for i, b := range bs {
		iovecs[i].SetLen(len(b))
		if len(b) > 0 {
			// somehow Iovec.Base on illumos is (*int8), not (*byte)
			iovecs[i].Base = (*int8)(unsafe.Pointer(&b[0]))
		} else {
			iovecs[i].Base = (*int8)(unsafe.Pointer(&_zero))
		}
	}
	return iovecs
}

//sys	readv(fd int, iovs []Iovec) (n int, err error)

func Readv(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = readv(fd, iovecs)
	return n, err
}

//sys	preadv(fd int, iovs []Iovec, off int64) (n int, err error)

func Preadv(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = preadv(fd, iovecs, off)
	return n, err
}

//sys	writev(fd int, iovs []Iovec) (n int, err error)

func Writev(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = writev(fd, iovecs)
	return n, err
}

//sys	pwritev(fd int, iovs []Iovec, off int64) (n int, err error)

func Pwritev(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = pwritev(fd, iovecs, off)
	return n, err
}

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4

func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	nfd, err = accept4(fd, &rsa, &len, flags)
	if err != nil {
		return
	}
	if len > SizeofSockaddrAny {
		panic("RawSockaddrAny too small")
	}
	sa, err = anyToSockaddr(fd, &rsa)
	if err != nil {
		Close(nfd)
		nfd = 0
	}
	return
}

//sys	putmsg(fd int, clptr *strbuf, dataptr *strbuf, flags int) (err error)

func Putmsg(fd int, cl []byte, data []byte, flags int) (err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Len: int32(len(cl)),
			Buf: (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Len: int32(len(data)),
			Buf: (*int8)(unsafe.Pointer(&data[0])),
		}
	}
	return putmsg(fd, clp, datap, flags)
}

//sys	getmsg(fd int, clptr *strbuf, dataptr *strbuf, flags *int) (err error)

func Getmsg(fd int, cl []byte, data []byte) (retCl []byte, retData []byte, flags int, err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Maxlen: int32(len(cl)),
			Buf:    (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Maxlen: int32(len(data)),
			Buf:    (*int8)(unsafe.Pointer(&data[0])),
		}
	}

	if err = getmsg(fd, clp, datap, &flags); err != nil {
		return nil, nil, 0, err
	}

	if len(cl) > 0 {
		retCl = cl[:clp.Len]
	}
	if len(data) > 0 {
		retData = data[:datap.Len]
	}
	return retCl, retData, flags, nil
}

func IoctlSetIntRetInt(fd int, req uint, arg int) (int, error) {
	return ioctlRet(fd, req, uintptr(arg))
}

func IoctlSetString(fd int, req uint, val string) error {
	bs := make([]byte, len(val)+1)
	copy(bs[:len(bs)-1], val)
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&bs[0])))
	runtime.KeepAlive(&bs[0])
	return err
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// Lifreq Helpers

func (l *Lifreq) SetName(name string) error {
	if len(name) >= len(l.Name) {
		return fmt.Errorf("name cannot be more than %d characters", len(l.Name)-1)
	}
	for i := range name {
		l.Name[i] = int8(name[i])
	}
	return nil
}

func (l *Lifreq) SetLifruInt(d int) {
	*(*int)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruInt() int {
	return *(*int)(unsafe.Pointer(&l.Lifru[0]))
}

func (l *Lifreq) SetLifruUint(d uint) {
	*(*uint)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruUint() uint {
	return *(*uint)(unsafe.Pointer(&l.Lifru[0]))
}

func IoctlLifreq(fd int, req uint, l *Lifreq) error {
	return ioctl(fd, req, uintptr(unsafe.Pointer(l)))
}

// Strioctl Helpers

func (s *Strioctl) SetInt(i int) {
	s.Len = int32(unsafe.Sizeof(i))
	s.Dp = (*int8)(unsafe.Pointer(&i))
}

func IoctlSetStrioctlRetInt(fd int, req uint, s *Strioctl) (int, error) {
	return ioctlRet(fd, req, uintptr(unsafe.Pointer(s)))
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2009 The Go Authors. All rights reserved.
||||||| parent of 4d7e5ad26 (update vendored files)
// Copyright 2009 The Go Authors. All rights reserved.
=======
// Copyright 2021 The Go Authors. All rights reserved.
>>>>>>> 4d7e5ad26 (update vendored files)
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// illumos system calls not present on Solaris.

//go:build amd64 && illumos
// +build amd64,illumos

package unix

import (
	"fmt"
	"runtime"
	"unsafe"
)

func bytes2iovec(bs [][]byte) []Iovec {
	iovecs := make([]Iovec, len(bs))
	for i, b := range bs {
		iovecs[i].SetLen(len(b))
		if len(b) > 0 {
			// somehow Iovec.Base on illumos is (*int8), not (*byte)
			iovecs[i].Base = (*int8)(unsafe.Pointer(&b[0]))
		} else {
			iovecs[i].Base = (*int8)(unsafe.Pointer(&_zero))
		}
	}
	return iovecs
}

//sys	readv(fd int, iovs []Iovec) (n int, err error)

func Readv(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = readv(fd, iovecs)
	return n, err
}

//sys	preadv(fd int, iovs []Iovec, off int64) (n int, err error)

func Preadv(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = preadv(fd, iovecs, off)
	return n, err
}

//sys	writev(fd int, iovs []Iovec) (n int, err error)

func Writev(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = writev(fd, iovecs)
	return n, err
}

//sys	pwritev(fd int, iovs []Iovec, off int64) (n int, err error)

func Pwritev(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = pwritev(fd, iovecs, off)
	return n, err
}

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4

func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	nfd, err = accept4(fd, &rsa, &len, flags)
	if err != nil {
		return
	}
	if len > SizeofSockaddrAny {
		panic("RawSockaddrAny too small")
	}
	sa, err = anyToSockaddr(fd, &rsa)
	if err != nil {
		Close(nfd)
		nfd = 0
	}
	return
}

//sys	putmsg(fd int, clptr *strbuf, dataptr *strbuf, flags int) (err error)

func Putmsg(fd int, cl []byte, data []byte, flags int) (err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Len: int32(len(cl)),
			Buf: (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Len: int32(len(data)),
			Buf: (*int8)(unsafe.Pointer(&data[0])),
		}
	}
	return putmsg(fd, clp, datap, flags)
}

//sys	getmsg(fd int, clptr *strbuf, dataptr *strbuf, flags *int) (err error)

func Getmsg(fd int, cl []byte, data []byte) (retCl []byte, retData []byte, flags int, err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Maxlen: int32(len(cl)),
			Buf:    (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Maxlen: int32(len(data)),
			Buf:    (*int8)(unsafe.Pointer(&data[0])),
		}
	}

	if err = getmsg(fd, clp, datap, &flags); err != nil {
		return nil, nil, 0, err
	}

	if len(cl) > 0 {
		retCl = cl[:clp.Len]
	}
	if len(data) > 0 {
		retData = data[:datap.Len]
	}
	return retCl, retData, flags, nil
}

func IoctlSetIntRetInt(fd int, req uint, arg int) (int, error) {
	return ioctlRet(fd, req, uintptr(arg))
}

func IoctlSetString(fd int, req uint, val string) error {
	bs := make([]byte, len(val)+1)
	copy(bs[:len(bs)-1], val)
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&bs[0])))
	runtime.KeepAlive(&bs[0])
	return err
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// Lifreq Helpers

func (l *Lifreq) SetName(name string) error {
	if len(name) >= len(l.Name) {
		return fmt.Errorf("name cannot be more than %d characters", len(l.Name)-1)
	}
	for i := range name {
		l.Name[i] = int8(name[i])
	}
	return nil
}

func (l *Lifreq) SetLifruInt(d int) {
	*(*int)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruInt() int {
	return *(*int)(unsafe.Pointer(&l.Lifru[0]))
}

func (l *Lifreq) SetLifruUint(d uint) {
	*(*uint)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruUint() uint {
	return *(*uint)(unsafe.Pointer(&l.Lifru[0]))
}

func IoctlLifreq(fd int, req uint, l *Lifreq) error {
	return ioctl(fd, req, uintptr(unsafe.Pointer(l)))
}

// Strioctl Helpers

func (s *Strioctl) SetInt(i int) {
	s.Len = int32(unsafe.Sizeof(i))
	s.Dp = (*int8)(unsafe.Pointer(&i))
}

func IoctlSetStrioctlRetInt(fd int, req uint, s *Strioctl) (int, error) {
	return ioctlRet(fd, req, uintptr(unsafe.Pointer(s)))
}
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)

//sys	putmsg(fd int, clptr *strbuf, dataptr *strbuf, flags int) (err error)

func Putmsg(fd int, cl []byte, data []byte, flags int) (err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Len: int32(len(cl)),
			Buf: (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Len: int32(len(data)),
			Buf: (*int8)(unsafe.Pointer(&data[0])),
		}
	}
	return putmsg(fd, clp, datap, flags)
}

//sys	getmsg(fd int, clptr *strbuf, dataptr *strbuf, flags *int) (err error)

func Getmsg(fd int, cl []byte, data []byte) (retCl []byte, retData []byte, flags int, err error) {
	var clp, datap *strbuf
	if len(cl) > 0 {
		clp = &strbuf{
			Maxlen: int32(len(cl)),
			Buf:    (*int8)(unsafe.Pointer(&cl[0])),
		}
	}
	if len(data) > 0 {
		datap = &strbuf{
			Maxlen: int32(len(data)),
			Buf:    (*int8)(unsafe.Pointer(&data[0])),
		}
	}

	if err = getmsg(fd, clp, datap, &flags); err != nil {
		return nil, nil, 0, err
	}

	if len(cl) > 0 {
		retCl = cl[:clp.Len]
	}
	if len(data) > 0 {
		retData = data[:datap.Len]
	}
	return retCl, retData, flags, nil
}

func IoctlSetIntRetInt(fd int, req uint, arg int) (int, error) {
	return ioctlRet(fd, req, uintptr(arg))
}

func IoctlSetString(fd int, req uint, val string) error {
	bs := make([]byte, len(val)+1)
	copy(bs[:len(bs)-1], val)
	err := ioctl(fd, req, uintptr(unsafe.Pointer(&bs[0])))
	runtime.KeepAlive(&bs[0])
	return err
}

// Lifreq Helpers

func (l *Lifreq) SetName(name string) error {
	if len(name) >= len(l.Name) {
		return fmt.Errorf("name cannot be more than %d characters", len(l.Name)-1)
	}
	for i := range name {
		l.Name[i] = int8(name[i])
	}
	return nil
}

func (l *Lifreq) SetLifruInt(d int) {
	*(*int)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruInt() int {
	return *(*int)(unsafe.Pointer(&l.Lifru[0]))
}

func (l *Lifreq) SetLifruUint(d uint) {
	*(*uint)(unsafe.Pointer(&l.Lifru[0])) = d
}

func (l *Lifreq) GetLifruUint() uint {
	return *(*uint)(unsafe.Pointer(&l.Lifru[0]))
}

func IoctlLifreq(fd int, req uint, l *Lifreq) error {
	return ioctl(fd, req, uintptr(unsafe.Pointer(l)))
}

// Strioctl Helpers

func (s *Strioctl) SetInt(i int) {
	s.Len = int32(unsafe.Sizeof(i))
	s.Dp = (*int8)(unsafe.Pointer(&i))
}

func IoctlSetStrioctlRetInt(fd int, req uint, s *Strioctl) (int, error) {
	return ioctlRet(fd, req, uintptr(unsafe.Pointer(s)))
}
=======
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// illumos system calls not present on Solaris.

// +build amd64,illumos

package unix

import "unsafe"

func bytes2iovec(bs [][]byte) []Iovec {
	iovecs := make([]Iovec, len(bs))
	for i, b := range bs {
		iovecs[i].SetLen(len(b))
		if len(b) > 0 {
			// somehow Iovec.Base on illumos is (*int8), not (*byte)
			iovecs[i].Base = (*int8)(unsafe.Pointer(&b[0]))
		} else {
			iovecs[i].Base = (*int8)(unsafe.Pointer(&_zero))
		}
	}
	return iovecs
}

//sys	readv(fd int, iovs []Iovec) (n int, err error)

func Readv(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = readv(fd, iovecs)
	return n, err
}

//sys	preadv(fd int, iovs []Iovec, off int64) (n int, err error)

func Preadv(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = preadv(fd, iovecs, off)
	return n, err
}

//sys	writev(fd int, iovs []Iovec) (n int, err error)

func Writev(fd int, iovs [][]byte) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = writev(fd, iovecs)
	return n, err
}

//sys	pwritev(fd int, iovs []Iovec, off int64) (n int, err error)

func Pwritev(fd int, iovs [][]byte, off int64) (n int, err error) {
	iovecs := bytes2iovec(iovs)
	n, err = pwritev(fd, iovecs, off)
	return n, err
}

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4

func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error) {
	var rsa RawSockaddrAny
	var len _Socklen = SizeofSockaddrAny
	nfd, err = accept4(fd, &rsa, &len, flags)
	if err != nil {
		return
	}
	if len > SizeofSockaddrAny {
		panic("RawSockaddrAny too small")
	}
	sa, err = anyToSockaddr(fd, &rsa)
	if err != nil {
		Close(nfd)
		nfd = 0
	}
	return
}

//sysnb	pipe2(p *[2]_C_int, flags int) (err error)

func Pipe2(p []int, flags int) error {
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]_C_int
	err := pipe2(&pp, flags)
	p[0] = int(pp[0])
	p[1] = int(pp[1])
	return err
}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
