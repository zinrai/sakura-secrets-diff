package main

import "bytes"

func CompareSecrets(local []byte, remote string) bool {
	return bytes.Equal(local, []byte(remote))
}
