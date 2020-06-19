package kata
import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
  return s == MyString(strings.ToUpper(string(s)))
}
