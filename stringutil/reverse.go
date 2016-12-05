//packagestringutil contains utility functions for working with strings
package stringutil	//package this file is a part of

//@Reverse returns its argument string reversed
func Reverse(s string) string {	//creates the function Reverse, which takes one string argument, s, and has the
				//has the return type of string
	r:= []rune(s)  	//creates an array of runes (the Go equivalent of char, but using 32 bits rather than 8 due to UTF-8 encoding)
			//uses the characters present in string s to initialize
	for i, j:= 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {	//create variables i and j
		//i is assigned 0, j is assigned the (length of r)-1
		//if i is less than the (length of r)/2
		//assign in parallel i to the value of i+1, and j to the value of j-1
		r[i], r[j] = r[j], r[i]	//Assign in parallel r[i] to the value of r[j], and vice versa
	}
	return string(r)	//return the rune array r cast as a string
}
