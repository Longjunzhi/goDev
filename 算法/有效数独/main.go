package main

import (
	"fmt"
)

//请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
//数字1-9在每一行只能出现一次。
//数字1-9在每一列只能出现一次。
//数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2f9gg/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	res := isValidSudoku(board)
	fmt.Printf("res: %v", res)
}

func isValidSudoku(board [][]byte) bool {
	// 思路 暴力算法
	// 每一行判断, map存起来，判断是否有值
	for i := 0; i < 9; i++ {
		res := make(map[byte]byte, 0)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := res[board[i][j]]; ok {
				return false
			}
			res[board[i][j]] = '2'
		}
	}
	// 每一列判断
	for i := 0; i < 9; i++ {
		res1 := make(map[byte]byte, 0)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if _, ok := res1[board[j][i]]; ok {
				return false
			}
			res1[board[j][i]] = '2'
		}
	}
	// 每个3宫格判断
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res2 := make(map[byte]byte, 0)
			k := i * 3
			l := j * 3
			for left := 0; left < 3; left++ {
				for height := 0; height < 3; height++ {
					if board[k][l] == '.' {
						l++
						continue
					}
					if ini, ok := res2[board[k][l]]; ok {
						fmt.Printf("ini : %v", ini)
						return false
					}
					res2[board[k][l]] = '2'
					l++
				}
				k++
				l = j * 3
			}
		}
	}
	return true
}
