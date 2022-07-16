package main

// Call: go run simple_z.go <pat> <string to search>

// Ex: go run simple_z.go "Qws?p" "Te/iLaHpQws?prfK3cp"

import (
	"os"
	"fmt"
	"errors"
)

type Z struct {
  Start int
  End int
}

// This returns the first match
func Simple_Z (pat string, str string) (*Z, error) {
  total := pat + "$" + str
  length := len(total)
  z_slice := make([]int, length, length)
  
  construct_z := func () {
    L := 0
    R := 0
    k := 0
    for i := 0; i < length; i++ {
      if i > R {
        L = i
        R = i
        for R < length && total[R-L] == total[R] {
          R++
        }
        z_slice[i] = R - L
        R--
      } else {
        k = i - L
        if z_slice[k] < R-i+1 {
          z_slice[i] = z_slice[k]
        } else {
          L = i
          for R < length && total[R-L] == total[R] {
            R++
          }
          z_slice[i] = R - L
          R--
        }
      }
    }
  }

  construct_z()
  
  for i := 0; i < length; i++ {
    if z_slice[i] == len(pat) {
    	n := i - 1

      return &Z{n-len(pat), n}, nil
    }
  }

  errMessage := "Pattern '"+ pat +"' not found."
  return nil, errors.New(errMessage)
}

func main () {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Please enter search pattern and string as args when you call this program.")
		return
	}
	pat := args[1] // pattern
	str := args[2] // string to search
	
	z, err := Simple_Z(pat, str)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Proof
	proof := pat == str[z.Start: z.End]
	if proof {
		fmt.Printf("Pattern '%s' found at slice [%d:%d]\n", pat, z.Start, z.End)
	} else {
		fmt.Println("simple_z is not returning accurate results.")
	}
}