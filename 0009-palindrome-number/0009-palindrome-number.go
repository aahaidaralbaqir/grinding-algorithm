// 10201
func isPalindrome(x int) bool {
  str := strconv.Itoa(x)
  length := len(str)

  for i := length - 1; i >= 0; i--{
    if  fmt.Sprintf("%c", str[i]) != fmt.Sprintf("%c", str[(length - 1) - i]) {
      return false
    }
  }
  return true
}