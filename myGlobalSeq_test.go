/*
 * @Author: huqing
 * @Date: 2018-12-11 11:32:46
 * @Last Modified by: huqing
 * @Last Modified time: 2018-12-11 13:07:53
 */

package myGlobalSeq_test

import (
	"fmt"
	"myself/myGlobalSeq"
	"testing"
)

func TestGetSeq(t *testing.T) {
	//myGlobal.Get
	var iResult int64 = 0
	for i := 0; i < 10; i++ {
		iResult += myGlobalSeq.GetGlobalSeq()
	}
	fmt.Println(iResult)
}
