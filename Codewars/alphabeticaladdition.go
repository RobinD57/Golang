package kata

func AddLetters(letters []rune) rune {
    if len(letters) == 0 {
        return 'z'
    }
    sum := 0
    for _, l := range letters {
        sum += int(l - 'a' + 1)
    }
    return 'a' + rune(sum - 1) % 26
}
