// food / 2 만큼 앞뒤로 추가 이부분 일단은 더하는걸로 해보고
// 근데 앞뒤로 그렇게 더하면 0을 중간에 넣어야 되는데 앞에서부터 1번 음식 2번 음식 각각 왼쪽 채울만큼 채우고
// 다 했으면 0 하고 나머지 추가하는 방향으로
// 1번 인덱스부터 마지막번까지 / 2 만큼 자기 인덱스를 더하고
// len - 1인 인덱스면 0 추가하고 끝내고
// 그 다음엔 역순으로 다시 반복문 아직 안익숙한 언어로 할라니까 어렵네 괜히 더 좋은 방법 있을지 떠오르지도 않고

// 구글링 하면서 찾은 내용
// writeString을 쓰면 속도가 더 빠르다.
// 이전 W 데브옵스 면접때 자바에서 문자열 추가할때 어떤게 더 빠른지 아냐는 질문 받아서 당황했었는데
// 그때도 단순히 String + String을 하면 느리기 때문에 쓰는 방법이 있는게 답이었다.
// 언어와 상관없이 문자열간 덧셈은 속도 저하의 원인인 듯 하다.

package programmers

import (
	"strconv"
	"strings"
)

func solution(food []int) string {

	var result strings.Builder
    
    for i := 1 ; i < len(food) ; i++ {
        for j := 0 ; j < (food[i] / 2) ; j++ {
            result.WriteString(strconv.Itoa(i))
        }
        if i == len(food) - 1 {
            result.WriteString("0")
        }
    }
    for i := len(food) - 1 ; i > 0 ; i-- {
        for j := 0 ; j < (food[i] / 2) ; j++ {
            result.WriteString(strconv.Itoa(i))
        }
    }
    return result.String()
}