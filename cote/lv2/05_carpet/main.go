package main

import "fmt"

/*
Leo는 카펫을 사러 갔다가 아래 그림과 같이 중앙에는 노란색으로 칠해져 있고 테두리 1줄은 갈색으로 칠해져 있는 격자 모양 카펫을 봤습니다.
Leo가 본 카펫에서 갈색 격자의 수 brown, 노란색 격자의 수 yellow가 매개변수로 주어질 때 카펫의 가로, 세로 크기를 순서대로 배열에 담아 return 하도록 solution 함수를 작성해주세요.

제한사항
갈색 격자의 수 brown은 8 이상 5,000 이하인 자연수입니다.
노란색 격자의 수 yellow는 1 이상 2,000,000 이하인 자연수입니다.
카펫의 가로 길이는 세로 길이와 같거나, 세로 길이보다 깁니다.

*/

type Tile struct {
	width  int
	height int
}

func solution(brown int, yellow int) []int {
	var b Tile
	y := Tile{
		width: 1,
	}

	for {
		if yellow%y.width != 0 {
			y.width++
			continue
		}

		y.height = yellow / y.width

		b.height = y.height + 2
		b.width = y.width + 2

		if b.height*b.width != brown+yellow {
			y.width++
			continue
		}

		if b.height > b.width {
			y.width++
			continue
		}

		return []int{b.width, b.height}
	}
}

func main() {
	// fmt.Println(solution(10, 2))  // [4,3]
	// fmt.Println(solution(8, 1))   // [3,3]
	// fmt.Println(solution(14, 6))  // [5,4]
	fmt.Println(solution(24, 24)) // [8,6]
}
