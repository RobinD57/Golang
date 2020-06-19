package kata
import "unicode/utf8"

func Solution(word string) string {
  size := len(word)
  buf := make([]byte, size)
  for start := 0; start < size; {
    r, n := utf8.DecodeRuneInString(word[start:])
    start += n
    utf8.EncodeRune(buf[size-start:], r)
    }
  return string(buf)
}
