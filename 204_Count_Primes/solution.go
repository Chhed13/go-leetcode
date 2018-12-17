package _204_Count_Primes

/* https://leetcode.com/problems/count-primes/

Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

/* Description
Erothosphenus sieve
- build bool array of n long
- implement counter for primes
- fill index arrays on each prime with multiplication of primes to true
- return count
Time complexity: O(nlogn). Space complexity: O(n)
*/


func countPrimes(n int) int {
    if n <= 1 {
        return 0
    }

    checkedPrimes := make([]bool,n)
    count := 0
    for i:= 2; i < n; i++{
        if !checkedPrimes[i] {
            count++
            for j := 1; j*i < n; j++ {
                checkedPrimes[i*j] = true
            }
        }
    }

    return count
}