// Code generated by ccgo. DO NOT EDIT.

// threadtest2
//  /*
//  ** 2004 January 13
//  **
//  ** The author disclaims copyright to this source code.  In place of
//  ** a legal notice, here is a blessing:
//  **
//  **    May you do good and not evil.
//  **    May you find forgiveness for yourself and forgive others.
//  **    May you share freely, never taking more than you give.
//  **
//  *************************************************************************
//  ** This file implements a simple standalone program used to test whether
//  ** or not the SQLite library is threadsafe.
//  **
//  ** This file is NOT part of the standard SQLite library.  It is used for
//  ** testing only.
//  */
package main

import (
	"math"
	"os"
	"unsafe"

	"github.com/cznic/crt"
	"github.com/cznic/sqlite/internal/bin"
)

var argv []*int8

func main() {
	for _, v := range os.Args {
		argv = append(argv, (*int8)(crt.CString(v)))
	}
	argv = append(argv, nil)
	X_start(crt.NewTLS(), int32(len(os.Args)), &argv[0])
}

func X_start(tls *crt.TLS, _argc int32, _argv **int8) {
	crt.X__register_stdfiles(tls, Xstdin, Xstdout, Xstderr)
	crt.X__builtin_exit(tls, Xmain(tls, _argc, _argv))
}

var Xstdin unsafe.Pointer

func init() {
	Xstdin = unsafe.Pointer(&X__stdfiles)
}

var X__stdfiles [3]unsafe.Pointer

var Xstdout unsafe.Pointer

func init() {
	Xstdout = unsafe.Pointer(uintptr(unsafe.Pointer(&X__stdfiles)) + 4)
}

var Xstderr unsafe.Pointer

func init() {
	Xstderr = unsafe.Pointer(uintptr(unsafe.Pointer(&X__stdfiles)) + 8)
}

// C comment
//  /*
//  ** Initialize the database and start the threads
//  */
func Xmain(tls *crt.TLS, _argc int32, _argv **int8) (r0 int32) {
	var _i, _rc int32
	var _1_zJournal *int8
	var _db unsafe.Pointer
	var _aThread [5]uint32
	r0 = int32(0)
	if crt.Xstrcmp(tls, str(0), str(8)) != 0 {
		_1_zJournal = bin.Xsqlite3_mprintf(tls, str(17), unsafe.Pointer(str(0)))
		crt.Xunlink(tls, str(0))
		crt.Xunlink(tls, _1_zJournal)
		bin.Xsqlite3_free(tls, unsafe.Pointer(_1_zJournal))
	}
	bin.Xsqlite3_open(tls, str(0), (**bin.Xsqlite3)(unsafe.Pointer(&_db)))
	if _db == nil {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(28))
		crt.Xexit(tls, int32(1))
	}
	_rc = bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), str(59), nil, nil, nil)
	if _rc != 0 {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(79), _rc)
		crt.Xexit(tls, int32(1))
	}
	bin.Xsqlite3_close(tls, (*bin.Xsqlite3)(_db))
	_i = int32(0)
_3:
	if uint32(_i) >= uint32(5) {
		goto _6
	}
	crt.Xpthread_create(tls, elem0((*uint32)(unsafe.Pointer(&_aThread)), uintptr(_i)), nil, Xworker, crt.U2P(uintptr(_i)))
	_i += 1
	goto _3
_6:
	_i = int32(0)
_7:
	if uint32(_i) >= uint32(5) {
		goto _10
	}
	crt.Xpthread_join(tls, *elem0((*uint32)(unsafe.Pointer(&_aThread)), uintptr(_i)), nil)
	_i += 1
	goto _7
_10:
	if Xall_stop == 0 {
		crt.Xprintf(tls, str(107))
		return int32(0)
	}
	crt.Xprintf(tls, str(129))
	return int32(1)

	_ = _aThread
	panic(0)
}

// C comment
//  /*
//  ** This is the worker thread
//  */
func Xworker(tls *crt.TLS, _workerArg unsafe.Pointer) (r0 unsafe.Pointer) {
	var _id, _rc, _cnt int32
	var _db unsafe.Pointer
	_id = int32(crt.P2U(_workerArg))
	_cnt = int32(0)
	crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(147), _id)
_0:
	if Xall_stop != 0 || postInc1(&_cnt, 1) >= int32(10000) {
		goto _1
	}
	if (_cnt % int32(100)) == int32(0) {
		crt.Xprintf(tls, str(167), _id, _cnt)
	}
_3:
	if bin.Xsqlite3_open(tls, str(0), (**bin.Xsqlite3)(unsafe.Pointer(&_db))) != int32(0) {
		crt.Xsched_yield(tls)
		goto _3
	}
	bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), str(175), nil, nil, nil)
	if Xall_stop != 0 {
		bin.Xsqlite3_close(tls, (*bin.Xsqlite3)(_db))
		goto _1
	}
	_rc = bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), str(198), nil, nil, nil)
	bin.Xsqlite3_close(tls, (*bin.Xsqlite3)(_db))
	goto _0
_1:
	crt.Xfprintf(tls, (*crt.XFILE)(Xstderr), str(234), _id)
	return nil

	_ = _rc
	panic(0)
}

// C comment
//  /*
//  ** When this variable becomes non-zero, all threads stop
//  ** what they are doing.
//  */
var Xall_stop int32

func bool2int(b bool) int32 {
	if b {
		return 1
	}
	return 0
}
func bug20530(interface{}) {} //TODO remove when https://github.com/golang/go/issues/20530 is fixed.
func init()                { nzf32 *= -1; nzf64 *= -1 }

var inf = math.Inf(1)
var nzf32 float32 // -0.0
var nzf64 float64 // -0.0
func elem0(a *uint32, index uintptr) *uint32 {
	return (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + 4*index))
}
func postInc1(p *int32, d int32) int32 { v := *p; *p += d; return v }
func str(n int) *int8                  { return (*int8)(unsafe.Pointer(&strTab[n])) }
func wstr(n int) *int32                { return (*int32)(unsafe.Pointer(&strTab[n])) }

var strTab = []byte("test.db\x00:memory:\x00%s-journal\x00unable to initialize database\x0a\x00CREATE TABLE t1(x);\x00cannot create table t1: %d\x0a\x00Everything seems ok.\x0a\x00We hit an error.\x0a\x00Starting worker %d\x0a\x00%d: %d\x0a\x00PRAGMA synchronous=OFF\x00INSERT INTO t1 VALUES('bogus data')\x00Worker %d finished\x0a\x00")
