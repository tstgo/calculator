/*
 * @Author: leoking
 * @Date: 2022-09-01 15:45:21
 * @LastEditTime: 2022-09-01 15:48:12
 * @LastEditors: leoking
 * @Description:
 */
package calculator

type Number interface {
	int | float32 | float64
}

/**
 * @description:
 * @return {*}
 */
func Add[T Number](a, b T) T {
	return a + b
}
