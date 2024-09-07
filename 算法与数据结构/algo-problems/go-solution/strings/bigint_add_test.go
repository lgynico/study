package strings

import (
	"fmt"
	"strings"
	"testing"
)

/**
描述
以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。

数据范围：s.length,t.length≤100000，字符串仅由'0'~‘9’构成

要求：时间复杂度O(n)

示例1
输入：
"1","99"
返回值：
"100"
说明：
1+99=100

示例2
输入：
"114514",""
返回值：
"114514"
*/

func TestBigintAdd(t *testing.T) {
	fmt.Println(bigint_add("1", "99"))
	fmt.Println(bigint_add("9", "999999999999999999999999999994"))
	rs := []rune{'1', '0', '9'}
	fmt.Println(string(rs))
	num := 65
	c := rune(num)
	fmt.Printf("%c", c)
}

func bigint_add(s, t string) string {
	if len(s) == 0 {
		return t
	}

	if len(t) == 0 {
		return s
	}

	var (
		result = strings.Builder{}
		carry  bool
		i      = len(s) - 1
		j      = len(t) - 1
		rs     = []rune(s)
		rt     = []rune(t)
	)

	for i >= 0 && j >= 0 {
		var d rune = (rs[i] - '0') + (rt[i] - '0')
		if carry {
			d += '1'
		}
		result.WriteRune(d % 10)
		carry = d >= 10
		i--
		j--

	}

	if i >= 0 {
		var d rune = rs[i] - '0'
		if carry {
			d += '1'
		}
		result.WriteRune(d % 10)
		carry = d >= 10
		i--
	}

	if j >= 0 {
		var d rune = rt[j] - '0'
		if carry {
			d += '1'
		}
		result.WriteRune(d % 10)
		carry = d >= 10
		j--
	}

	if carry {
		result.WriteRune('1')
	}

	return result.String()
}
