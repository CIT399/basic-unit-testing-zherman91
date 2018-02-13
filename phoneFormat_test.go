package main

import "testing"

func TestFormatterToFormatPhoneNumbers(t *testing.T) {
	actualResult := PhoneNumberFormatter("8147303959")
	var expectedResult = "(814)730-3959"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestFormatterToUnFormatPhoneNumbers(t *testing.T) {
	actualResult := PhoneNumberFormatter("(814)730-3959")
	var expectedResult = "8147303959"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
/*
~\basic-unit-testing-zherman91>go test -v
=== RUN   TestFormatterToFormatPhoneNumbers
Length of 8147303959 is 10 so that won't work... Let's try formatting it!
(814)730-3959
--- PASS: TestFormatterToFormatPhoneNumbers (0.00s)
=== RUN   TestFormatterToUnFormatPhoneNumbers
Holy cow! Let's unformat it!
Unformatted number:  8147303959
--- PASS: TestFormatterToUnFormatPhoneNumbers (0.00s)
PASS
ok      _/C_/Users/Zac/Documents/School/Secure_Programming/basic-unit-testing-zherman91 0.099s
*/
