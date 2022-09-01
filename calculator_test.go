/*
 * @Author: leoking
 * @Date: 2022-09-01 17:12:14
 * @LastEditTime: 2022-09-01 17:16:29
 * @LastEditors: leoking
 * @Description:
 */

package calculator

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	t.Log(Multiply(10, 20))
	t.Log(Multiply(10.22, 37.0))
}

func TestSub(t *testing.T) {
	t.Log(Sub(10, 0))
}
