package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

// These constant fields hold the valid lengths of
// strings that are formatted and unformatted.
var (
	FORMATTED_LENGTH   = 13
	UNFORMATTED_LENGTH = 10
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a phone number: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n") //removes the trailing '\n' otherwise throws the count off..

	var length = len(text)
	if length == FORMATTED_LENGTH {
		if isFormatted(text) {
			fmt.Println("Holy cow! Let's unformat it!")
			var unformatted = unformat(text)
			fmt.Println("Unformatted number: ", unformatted)
		}
	} else {
		fmt.Println("Length of", text, "is", length, "so that won't work... Let's try formatting it!")
		fmt.Println(format(text))
	}
}

/*
The isFormatted method accepts a string argument
and determines whether it is properly formatted as
a US telephone number in the following manner:
(XXX)XXX-XXXX
If the argument is properly formatted, the method
returns true, otherwise false.
*/
func isFormatted(str string) (bool) {

	valid := true //flag to indicate valid format

	if len(str) == FORMATTED_LENGTH &&
		str[0] == '(' &&
		str[4] == ')' &&
		str[8] == '-' {
		valid = true
	} else {
		valid = false
	}
	return valid
}

/*
The unformat method accepts a string containing
a telephone number formatted as: (XXX)XXX-XXXX.
If the argument is formatted in this way, the
method returns an unformatted string where the
parentheses and hyphen have been removed. Otherwise,
it returns the original argument.
*/
func unformat(str string) (string) {

	var phone *PhoneNumber //demonstrates goes implementation of structs which is key to it's built in concurrency.
	phone = new(PhoneNumber)
	phone.Number = str

	if isFormatted(phone.Number) {
		phone.Number = strings.Replace(phone.Number, "(", "", -1)
		phone.Number = strings.Replace(phone.Number, ")", "", -1)
		phone.Number = strings.Replace(phone.Number, "-", "", -1)
	}
	return phone.Number
}

/*
The format method formats a string as: (XXX)XXX-XXXX.
If the length of the argument is UNFORMATTED_LENGTH
the method returns the formatted string. Otherwise,
it returns the original argument.
*/
func format(str string) (string) {
	var phone *PhoneNumber
	phone = new(PhoneNumber)
	phone.Number = str

	if len(phone.Number) == UNFORMATTED_LENGTH {

		phone.Number = phone.Number[:0] + "(" + phone.Number[0:]
		phone.Number = phone.Number[:4] + ")" + phone.Number[4:]
		phone.Number = phone.Number[:8] + "-" + phone.Number[8:]

		return phone.Number

	} else {
		fmt.Println("Try entering a number with NO formatting at all.. ")
		return ""
	}
}

type PhoneNumber struct {
	Number string
}
