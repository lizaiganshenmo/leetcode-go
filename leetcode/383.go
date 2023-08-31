package leetcode

// https://leetcode.cn/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150

func canConstruct(ransomNote string, magazine string) bool {
	n := len(ransomNote)
	tmpMap := make(map[byte]int, n)
	for i := range ransomNote {
		if val, ok := tmpMap[ransomNote[i]]; ok {
			tmpMap[ransomNote[i]] = val + 1
		} else {
			tmpMap[ransomNote[i]] = 1
		}
	}

	for i := range magazine {
		if val, ok := tmpMap[magazine[i]]; ok && val > 0 {
			tmpMap[magazine[i]] = val - 1
			if val-1 == 0 {
				delete(tmpMap, magazine[i])
			}
		}
	}

	for range tmpMap {
		return false
	}

	return true

}
