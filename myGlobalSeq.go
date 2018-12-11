/*
 * @Author: huqing
 * @Date: 2018-12-11 11:17:29
 * @Last Modified by: huqing
 * @Last Modified time: 2018-12-11 11:34:49
 */

package myGlobalSeq

import (
	"sync/atomic"
)

var myGlobal int64 = 0

func GetGlobalSeq() int64 {
	return atomic.AddInt64(&myGlobal, 1)
}

func init() {
	myGlobal = 0
}
