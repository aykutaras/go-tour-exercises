package gotourexercises

import (
    "strings"
)

func WordCount(s string) map[string]int {
    mapped := make(map[string]int)
    splitted := strings.Fields(s)
    for _, str := range splitted {
        mapped[str] += 1
    }
    
    return mapped
}