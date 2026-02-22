package main

import (
	"fmt"
)

func main() {
	test_input := "error reading go.mod: missing module declaration. To specify the module path:\n\n\tgo mod init <module_path>\n\nReplace <module_path> with the desired module path, which is typically the repository URL or a unique identifier for your project."
	freq := word_frequency_count(test_input)
	fmt.Println(freq)



	// test case for palindrome
	palindrome := is_palindrome("A man, a plan, a canal, Panama")
	fmt.Println(palindrome)
}
