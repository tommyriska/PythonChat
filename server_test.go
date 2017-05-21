package main

import "testing"

/* ***** EVERY TEST MUST BE RUN INDIVIDUALLY ***** */

/* test will check if the key array that is created
		on the server is the correct size (32 bytes) */
func TestCreateKey(t *testing.T) {

	key := createKey()
	var totalByte = 0

	for i := range key {
		fmt.Println(i)
		totalByte++
	}

	fmt.Println("Total byte i array --> ")
	fmt.Println(totalByte)

	if totalByte < 32 {
		t.Error("Expected 32, got: " totalByte)
	}
	if totalByte > 32 {
		t.Error("Expected 32. got: " totalByte)
	}
}

/* test will check if a string variable is altered after
 		method call (encrypted) */
func TestEncrypt(t *testing.T) {

	key := make([]byte, 32)
	var crypt string = "test crypto"
	text := encrypt(key, crypt)

	if text == crypt {
		t.Error("Expected new text, got same.")
	}
}

/* test will check if a string variable is altered after
		method call (decrypted) */
func TestDecrypt(t *testing.T) {

	key := make([]byte, 32)
	var crypt string = "eL7Abmpadmk8nAaVcpT-a6MziwEKcL5z3ifS54SsnxA="
	text := decrypt(key, crypt)

	if text == crypt {
		t.Error("Expected new text, got same")
	}
}
