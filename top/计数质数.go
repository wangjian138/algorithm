package main

// https://leetcode.cn/problems/count-primes/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func countPrimes(n int) int {
	// create boolean array with default value true
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	// loop through all numbers up to the square root of n
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			// loop through multiples of i and mark them as not prime
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	// count number of primes
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=204
func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
