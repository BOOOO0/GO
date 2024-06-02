// 크게 어렵지 않은 수학문제
package programmers

func solution(a int, b int, n int) int {
    var result int
    // go에서 while문 처럼 for문을 쓰는 방법
    for n >= a {
        result += (n / a) * b
        n = ((n / a) * b) + (n % a) 
        // fmt.Println(n, result)
    }
    
    return result
}