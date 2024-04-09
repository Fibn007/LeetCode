import (
    "strings"
    "fmt"
)

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    if len(s) == 1 {
        return 1
    }
    charMap := make(map[byte]int)
    maxLength := 0
    start := 0

    for i := 0; i < len(s); i++ {
        if index, found := charMap[s[i]]; found && index >= start {
            start = index + 1
        }

        charMap[s[i]] = i
        
        if i-start+1 > maxLength {
            maxLength = i - start + 1
        }
    }
    return maxLength
}
