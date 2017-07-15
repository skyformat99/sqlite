// Code generated by ccgo. DO NOT EDIT.

// threadtest1
//  /*
//  ** 2002 January 15
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
//  ** Testing the thread safety of SQLite is difficult because there are very
//  ** few places in the code that are even potentially unsafe, and those
//  ** places execute for very short periods of time.  So even if the library
//  ** is compiled with its mutexes disabled, it is likely to work correctly
//  ** in a multi-threaded program most of the time.
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
	Xstdout = unsafe.Pointer(uintptr(unsafe.Pointer(&X__stdfiles)) + 8)
}

var Xstderr unsafe.Pointer

func init() {
	Xstderr = unsafe.Pointer(uintptr(unsafe.Pointer(&X__stdfiles)) + 16)
}

func Xmain(tls *crt.TLS, _argc int32, _argv **int8) (r0 int32) {
	var _i, _n int32
	var _id uint64
	var _zFile, _4_zDb, _4_zJournal *int8
	var _2_zBuf, _6_zBuf [200]int8
	r0 = int32(0)
	if (_argc > int32(2)) && (crt.Xstrcmp(tls, *elem0(_argv, uintptr(1)), str(0)) == int32(0)) {
		_verbose = int32(1)
		bug20530(_verbose)
		_argc -= 1
		*(*uintptr)(unsafe.Pointer(&_argv)) += uintptr(8)
	}
	if (_argc < int32(2)) || (store1(&_n, crt.Xatoi(tls, *elem0(_argv, uintptr(1)))) < int32(1)) {
		_n = int32(10)
	}
	_i = int32(0)
_4:
	if _i >= _n {
		goto _7
	}
	crt.Xsprintf(tls, (*int8)(unsafe.Pointer(&_2_zBuf)), str(3), (_i+int32(1))/int32(2))
	crt.Xunlink(tls, (*int8)(unsafe.Pointer(&_2_zBuf)))
	_i += 1
	goto _4
_7:
	_i = int32(0)
_8:
	if _i >= _n {
		goto _11
	}
	_zFile = bin.Xsqlite3_mprintf(tls, str(13), (_i%int32(2))+int32(1), (_i+int32(2))/int32(2))
	if (_i % int32(2)) == int32(0) {
		_4_zDb = elem2(_zFile, uintptr(2))
		_4_zJournal = bin.Xsqlite3_mprintf(tls, str(26), unsafe.Pointer(_4_zDb))
		crt.Xunlink(tls, _4_zDb)
		crt.Xunlink(tls, _4_zJournal)
		crt.Xfree(tls, unsafe.Pointer(_4_zJournal))
	}
	crt.Xpthread_create(tls, &_id, nil, _worker_bee, unsafe.Pointer(_zFile))
	crt.Xpthread_detach(tls, _id)
	_i += 1
	goto _8
_11:
	crt.Xpthread_mutex_lock(tls, &Xlock)
_13:
	if Xthread_cnt > int32(0) {
		crt.Xpthread_cond_wait(tls, &Xsig, &Xlock)
		goto _13
	}
	crt.Xpthread_mutex_unlock(tls, &Xlock)
	_i = int32(0)
_15:
	if _i >= _n {
		goto _18
	}
	crt.Xsprintf(tls, (*int8)(unsafe.Pointer(&_6_zBuf)), str(3), (_i+int32(1))/int32(2))
	crt.Xunlink(tls, (*int8)(unsafe.Pointer(&_6_zBuf)))
	_i += 1
	goto _15
_18:
	return int32(0)

	_ = _2_zBuf
	_ = _6_zBuf
	panic(0)
}

// C comment
//  /*
//  ** Enable for tracing
//  */
var _verbose int32

func _worker_bee(tls *crt.TLS, _pArg unsafe.Pointer) (r0 unsafe.Pointer) {
	var _i, _cnt, _t int32
	var _zFilename, _azErr *int8
	var _db unsafe.Pointer
	var _az **int8
	var _4_z1, _4_z2 [30]int8
	_zFilename = (*int8)(_pArg)
	_t = crt.Xatoi(tls, _zFilename)
	crt.Xpthread_mutex_lock(tls, &Xlock)
	Xthread_cnt += 1
	crt.Xpthread_mutex_unlock(tls, &Xlock)
	crt.Xprintf(tls, str(37), unsafe.Pointer(_zFilename))
	crt.Xfflush(tls, (*crt.XFILE)(Xstdout))
	_cnt = int32(0)
_0:
	if _cnt >= int32(10) {
		goto _3
	}
	bin.Xsqlite3_open(tls, elem2(_zFilename, uintptr(2)), (**bin.Xsqlite3)(unsafe.Pointer(&_db)))
	if _db == nil {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstdout), str(48), unsafe.Pointer(_zFilename))
		_Exit(tls, int32(1))
	}
	bin.Xsqlite3_busy_handler(tls, (*bin.Xsqlite3)(_db), _db_is_locked, unsafe.Pointer(_zFilename))
	Xdb_execute(tls, _db, _zFilename, str(64), _t)
	_i = int32(1)
_5:
	if _i > int32(100) {
		goto _8
	}
	Xdb_execute(tls, _db, _zFilename, str(89), _t, _i, _i*int32(2), _i*_i)
	_i += 1
	goto _5
_8:
	_az = Xdb_query(tls, _db, _zFilename, str(123), _t)
	Xdb_check(tls, _zFilename, str(148), _az, unsafe.Pointer(str(156)), int32(0))
	_az = Xdb_query(tls, _db, _zFilename, str(160), _t)
	Xdb_check(tls, _zFilename, str(183), _az, unsafe.Pointer(str(190)), int32(0))
	Xdb_execute(tls, _db, _zFilename, str(196), _t)
	_az = Xdb_query(tls, _db, _zFilename, str(160), _t)
	Xdb_check(tls, _zFilename, str(223), _az, unsafe.Pointer(str(231)), int32(0))
	_i = int32(1)
_9:
	if _i > int32(50) {
		goto _12
	}
	_az = Xdb_query(tls, _db, _zFilename, str(236), _t, _i)
	crt.Xsprintf(tls, (*int8)(unsafe.Pointer(&_4_z1)), str(268), _i*int32(2))
	crt.Xsprintf(tls, (*int8)(unsafe.Pointer(&_4_z2)), str(268), _i*_i)
	Xdb_check(tls, _zFilename, str(271), _az, unsafe.Pointer(&_4_z1), unsafe.Pointer(&_4_z2), int32(0))
	_i += 1
	goto _9
_12:
	Xdb_execute(tls, _db, _zFilename, str(280), _t)
	bin.Xsqlite3_close(tls, (*bin.Xsqlite3)(_db))
	_cnt += 1
	goto _0
_3:
	crt.Xprintf(tls, str(296), unsafe.Pointer(_zFilename))
	crt.Xfflush(tls, (*crt.XFILE)(Xstdout))
	crt.Xpthread_mutex_lock(tls, &Xlock)
	Xthread_cnt -= 1
	if Xthread_cnt <= int32(0) {
		crt.Xpthread_cond_signal(tls, &Xsig)
	}
	crt.Xpthread_mutex_unlock(tls, &Xlock)
	return nil

	_ = _azErr
	_ = _4_z1
	_ = _4_z2
	panic(0)
}

var Xlock crt.Xpthread_mutex_t

var Xthread_cnt int32

// C comment
//  /*
//  ** Come here to die.
//  */
func _Exit(tls *crt.TLS, _rc int32) {
	crt.Xexit(tls, _rc)
}

// C comment
//  /*
//  ** When a lock occurs, yield.
//  */
func _db_is_locked(tls *crt.TLS, _NotUsed unsafe.Pointer, _iCount int32) (r0 int32) {
	if _verbose != 0 {
		crt.Xprintf(tls, str(305), _NotUsed, _iCount)
	}
	crt.Xusleep(tls, uint32(100))
	return bool2int(_iCount < int32(40000))
}

// C comment
//  /*
//  ** Execute an SQL statement.
//  */
func Xdb_execute(tls *crt.TLS, _db unsafe.Pointer, _zFile *int8, _zFormat *int8, args ...interface{}) {
	var _rc int32
	var _zSql, _zErrMsg *int8
	var _ap []interface{}
	_zErrMsg = nil
	_ap = args
	_zSql = bin.Xsqlite3_vmprintf(tls, _zFormat, _ap)
	_ap = nil
	if _verbose != 0 {
		crt.Xprintf(tls, str(318), unsafe.Pointer(_zFile), unsafe.Pointer(_zSql))
	}
_0:
	_rc = bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), _zSql, nil, nil, &_zErrMsg)
	if _rc == int32(5) {
		goto _0
	}
	if _verbose != 0 {
		crt.Xprintf(tls, str(331), unsafe.Pointer(_zFile), unsafe.Pointer(_zSql))
	}
	if _zErrMsg != nil {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstdout), str(344), unsafe.Pointer(_zFile), unsafe.Pointer(_zSql), unsafe.Pointer(_zErrMsg))
		crt.Xfree(tls, unsafe.Pointer(_zErrMsg))
		bin.Xsqlite3_free(tls, unsafe.Pointer(_zSql))
		_Exit(tls, int32(1))
	}
	bin.Xsqlite3_free(tls, unsafe.Pointer(_zSql))
}

// C comment
//  /*
//  ** Execute a query against the database.  NULL values are returned
//  ** as an empty string.  The list is terminated by a single NULL pointer.
//  */
func Xdb_query(tls *crt.TLS, _db unsafe.Pointer, _zFile *int8, _zFormat *int8, args ...interface{}) (r0 **int8) {
	var _rc int32
	var _zSql, _zErrMsg *int8
	var _ap []interface{}
	var _sResult TQueryResult
	_zErrMsg = nil
	_ap = args
	_zSql = bin.Xsqlite3_vmprintf(tls, _zFormat, _ap)
	_ap = nil
	crt.Xmemset(tls, unsafe.Pointer(&_sResult), int32(0), uint64(24))
	_sResult.XzFile = _zFile
	if _verbose != 0 {
		crt.Xprintf(tls, str(373), unsafe.Pointer(_zFile), unsafe.Pointer(_zSql))
	}
	_rc = bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), _zSql, _db_query_callback, unsafe.Pointer(&_sResult), &_zErrMsg)
	if _rc != int32(17) {
		goto _1
	}
	if _zErrMsg != nil {
		crt.Xfree(tls, unsafe.Pointer(_zErrMsg))
	}
	_rc = bin.Xsqlite3_exec(tls, (*bin.Xsqlite3)(_db), _zSql, _db_query_callback, unsafe.Pointer(&_sResult), &_zErrMsg)
_1:
	if _verbose != 0 {
		crt.Xprintf(tls, str(387), unsafe.Pointer(_zFile), unsafe.Pointer(_zSql))
	}
	if _zErrMsg != nil {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstdout), str(399), unsafe.Pointer(_zFile), unsafe.Pointer(_zSql), unsafe.Pointer(_zErrMsg))
		crt.Xfree(tls, unsafe.Pointer(_zErrMsg))
		crt.Xfree(tls, unsafe.Pointer(_zSql))
		_Exit(tls, int32(1))
	}
	bin.Xsqlite3_free(tls, unsafe.Pointer(_zSql))
	if _sResult.XazElem == nil {
		_db_query_callback(tls, unsafe.Pointer(&_sResult), int32(0), nil, nil)
	}
	*elem0(_sResult.XazElem, uintptr(_sResult.XnElem)) = nil
	return _sResult.XazElem
}

// C comment
//  /*
//  ** The callback function for db_query
//  */
func _db_query_callback(tls *crt.TLS, _pUser unsafe.Pointer, _nArg int32, _azArg **int8, _NotUsed **int8) (r0 int32) {
	var _i int32
	var _pResult *TQueryResult
	_pResult = (*TQueryResult)(_pUser)
	if (_pResult.XnElem + _nArg) < _pResult.XnAlloc {
		goto _0
	}
	if _pResult.XnAlloc == int32(0) {
		_pResult.XnAlloc = _nArg + int32(1)
		goto _2
	}
	_pResult.XnAlloc = ((_pResult.XnAlloc * int32(2)) + _nArg) + int32(1)
_2:
	_pResult.XazElem = (**int8)(crt.Xrealloc(tls, unsafe.Pointer(_pResult.XazElem), uint64(_pResult.XnAlloc)*uint64(8)))
	if _pResult.XazElem == nil {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstdout), str(426), unsafe.Pointer(_pResult.XzFile))
		return int32(1)
	}
_0:
	if _azArg == nil {
		return int32(0)
	}
	_i = int32(0)
_5:
	if _i >= _nArg {
		goto _8
	}
	*elem0(_pResult.XazElem, uintptr(postInc1(&_pResult.XnElem, 1))) = bin.Xsqlite3_mprintf(tls, str(445), unsafe.Pointer(func() *int8 {
		if (*elem0(_azArg, uintptr(_i))) != nil {
			return (*elem0(_azArg, uintptr(_i)))
		}
		return str(448)
	}()))
	_i += 1
	goto _5
_8:
	return int32(0)
}

// C comment
//  /*
//  ** Check results
//  */
func Xdb_check(tls *crt.TLS, _zFile *int8, _zMsg *int8, _az **int8, args ...interface{}) {
	var _i int32
	var _z *int8
	var _ap []interface{}
	_ap = args
	_i = int32(0)
_0:
	if store2(&_z, (*int8)(crt.VAPointer(&_ap))) == nil {
		goto _3
	}
	if ((*elem0(_az, uintptr(_i))) == nil) || (crt.Xstrcmp(tls, *elem0(_az, uintptr(_i)), _z) != int32(0)) {
		crt.Xfprintf(tls, (*crt.XFILE)(Xstdout), str(449), unsafe.Pointer(_zFile), unsafe.Pointer(_zMsg), _i+int32(1), unsafe.Pointer(*elem0(_az, uintptr(_i))))
		Xdb_query_free(tls, _az)
		_Exit(tls, int32(1))
	}
	_i += 1
	goto _0
_3:
	_ap = nil
	Xdb_query_free(tls, _az)
}

// C comment
//  /*
//  ** Free the results of a db_query() call.
//  */
func Xdb_query_free(tls *crt.TLS, _az **int8) {
	var _i int32
	_i = int32(0)
_0:
	if (*elem0(_az, uintptr(_i))) == nil {
		goto _3
	}
	bin.Xsqlite3_free(tls, unsafe.Pointer(*elem0(_az, uintptr(_i))))
	_i += 1
	goto _0
_3:
	crt.Xfree(tls, unsafe.Pointer(_az))
}

var Xsig crt.Xpthread_cond_t

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
func elem0(a **int8, index uintptr) **int8 {
	return (**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + 8*index))
}
func elem2(a *int8, index uintptr) *int8 {
	return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + 1*index))
}
func postInc1(p *int32, d int32) int32 { v := *p; *p += d; return v }
func store2(p **int8, v *int8) *int8   { *p = v; return v }
func store1(p *int32, v int32) int32   { *p = v; return v }

type TQueryResult struct {
	XzFile  *int8
	XnElem  int32
	XnAlloc int32
	XazElem **int8
}                       // t3 struct{*int8,int32,int32,**int8}
func str(n int) *int8   { return (*int8)(unsafe.Pointer(&strTab[n])) }
func wstr(n int) *int32 { return (*int32)(unsafe.Pointer(&strTab[n])) }

var strTab = []byte("-v\x00testdb-%d\x00%d.testdb-%d\x00%s-journal\x00%s: START\x0a\x00%s: can't open\x0a\x00CREATE TABLE t%d(a,b,c);\x00INSERT INTO t%d VALUES(%d,%d,%d);\x00SELECT count(*) FROM t%d\x00tX size\x00100\x00SELECT avg(b) FROM t%d\x00tX avg\x00101.0\x00DELETE FROM t%d WHERE a>50\x00tX avg2\x0051.0\x00SELECT b, c FROM t%d WHERE a=%d\x00%d\x00readback\x00DROP TABLE t%d;\x00%s: END\x0a\x00BUSY %s #%d\x0a\x00EXEC %s: %s\x0a\x00DONE %s: %s\x0a\x00%s: command failed: %s - %s\x0a\x00QUERY %s: %s\x0a\x00DONE %s %s\x0a\x00%s: query failed: %s - %s\x0a\x00%s: malloc failed\x0a\x00%s\x00\x00%s: %s: bad result in column %d: %s\x0a\x00")
