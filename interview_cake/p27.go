// Reverse Words
package main

import "fmt"

// reverseWords reverse the msg.
// ignoring the possible of unicode here, using bytes.
func reverseWords(msg []byte) {
	l := len(msg)
	wordStart := 0
	wordEnd := 0
	for i := 0; i < l/2; i++ {
		if isSpaceOrPunctuation(msg[i]) {
			wordEnd = i
			reverseWords(msg[wordStart:wordEnd])
			wordStart = i + 1
		}

		msg[i], msg[l-i-1] = msg[l-i-1], msg[i]
	}
}

// Extensible for further punctuation
func isSpaceOrPunctuation(b byte) bool {
	switch b {
	case 32:
		return true
	default:
		return false
	}
}

func reverseChars(chars []byte) {
}

var tests = []struct {
	msg  string
	want string
}{
	{
		msg:  "hello world",
		want: "world hello",
	},
	{
		msg:  "",
		want: "",
	},
	{
		msg:  "the eagle has landed",
		want: "landed has eagle the",
	},
}

func main() {
	for _, t := range tests {
		reverseWords([]byte(t.msg))
		got := string(t.msg)
		if got != t.want {
			fmt.Errorf("input %s: got %s, want %s", t.msg, got, t.want)
		}
	}
}
