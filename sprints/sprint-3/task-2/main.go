package main

func scoreOfString(s string) int {
    result := 0

    for i := 0; i < len(s) - 1; i++ {
        if(int(s[i]) - int(s[i + 1])) > 0  {
            result += int(s[i]) - int(s[i + 1])
        } else {
            result += int(s[i + 1]) - int(s[i])
        }

    }

    return result
}