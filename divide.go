/*
 * @Author: leoking
 * @Date: 2022-09-01 15:47:15
 * @LastEditTime: 2022-09-01 15:48:42
 * @LastEditors: leoking
 * @Description:
 */
package calculator

import "errors"

/**
 * @description:
 * @return {*}
 */
func Divide[T Number](a, b T) (T, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
