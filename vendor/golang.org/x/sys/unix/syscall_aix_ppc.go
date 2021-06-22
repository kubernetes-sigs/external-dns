// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build aix && ppc
// +build aix,ppc
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix
// +build ppc
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build aix
// +build ppc
=======
//go:build aix && ppc
// +build aix,ppc
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix
// +build ppc
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

package unix

//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error) = getrlimit64
//sysnb	Setrlimit(resource int, rlim *Rlimit) (err error) = setrlimit64
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = lseek64

//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: int32(sec), Usec: int32(usec)}
}

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint32(length)
}

func (msghdr *Msghdr) SetControllen(length int) {
	msghdr.Controllen = uint32(length)
}

func (msghdr *Msghdr) SetIovlen(length int) {
	msghdr.Iovlen = int32(length)
}

func (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint32(length)
}

func Fstat(fd int, stat *Stat_t) error {
	return fstat(fd, stat)
}

func Fstatat(dirfd int, path string, stat *Stat_t, flags int) error {
	return fstatat(dirfd, path, stat, flags)
}

func Lstat(path string, stat *Stat_t) error {
	return lstat(path, stat)
}

func Stat(path string, statptr *Stat_t) error {
	return stat(path, statptr)
}
