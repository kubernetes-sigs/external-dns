// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build riscv64 && gc
// +build riscv64
// +build gc
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build riscv64,gc
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build riscv64,gc
=======
//go:build riscv64 && gc
// +build riscv64
// +build gc
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build riscv64,gc
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// +build riscv64,gc
=======
//go:build riscv64 && gc
// +build riscv64
// +build gc
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build riscv64,gc
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// +build riscv64,gc
=======
//go:build riscv64 && gc
// +build riscv64
// +build gc
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build riscv64,gc
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

#include "textflag.h"

//
// System calls for linux/riscv64.
//
// Where available, just jump to package syscall's implementation of
// these functions.

TEXT ·Syscall(SB),NOSPLIT,$0-56
	JMP	syscall·Syscall(SB)

TEXT ·Syscall6(SB),NOSPLIT,$0-80
	JMP	syscall·Syscall6(SB)

TEXT ·SyscallNoError(SB),NOSPLIT,$0-48
	CALL	runtime·entersyscall(SB)
	MOV	a1+8(FP), A0
	MOV	a2+16(FP), A1
	MOV	a3+24(FP), A2
	MOV	trap+0(FP), A7	// syscall entry
	ECALL
	MOV	A0, r1+32(FP)	// r1
	MOV	A1, r2+40(FP)	// r2
	CALL	runtime·exitsyscall(SB)
	RET

TEXT ·RawSyscall(SB),NOSPLIT,$0-56
	JMP	syscall·RawSyscall(SB)

TEXT ·RawSyscall6(SB),NOSPLIT,$0-80
	JMP	syscall·RawSyscall6(SB)

TEXT ·RawSyscallNoError(SB),NOSPLIT,$0-48
	MOV	a1+8(FP), A0
	MOV	a2+16(FP), A1
	MOV	a3+24(FP), A2
	MOV	trap+0(FP), A7	// syscall entry
	ECALL
	MOV	A0, r1+32(FP)
	MOV	A1, r2+40(FP)
	RET
