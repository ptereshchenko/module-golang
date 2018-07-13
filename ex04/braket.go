package brackets

func Bracket(str string) (bool, error) {
  var a, b, c bool  = true, true, true

  for i := 0; i < len(str); i++ {
	if str[i] == '(' {
      a = false
	  for j := i + 1; j < len(str); j += 2 {
		if str[j] == ')' {
		  a = true
		  break
		}
	  }
	}
	if str[i] == '[' {
	  b = false
	  for j := i + 1; j < len(str); j += 2 {
		if str[j] == ']' {
		  b = true
		  break
		}
      }
	}
	if str[i] == '{' {
	  c = false
	  for j := i + 1; j < len(str); j += 2 {
		if str[j] == '}' {
		  c = true
		  break
		}
	  }
	}
  }
  if a && b && c {
    return true, nil
  }
  return false, nil
}
