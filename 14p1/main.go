package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

/*
--- Day 14: One-Time Pad ---

In order to communicate securely with Santa while you're on this mission, you've been using a one-time pad that you generate using a pre-agreed algorithm. Unfortunately, you've run out of keys in your one-time pad, and so you need to generate some more.

To generate keys, you first get a stream of random data by taking the MD5 of a pre-arranged salt (your puzzle input) and an increasing integer index (starting with 0, and represented in decimal); the resulting MD5 hash should be represented as a string of lowercase hexadecimal digits.

However, not all of these MD5 hashes are keys, and you need 64 new keys for your one-time pad. A hash is a key only if:

It contains three of the same character in a row, like 777. Only consider the first such triplet in a hash.
One of the next 1000 hashes in the stream contains that same character five times in a row, like 77777.
Considering future hashes for five-of-a-kind sequences does not cause those hashes to be skipped; instead, regardless of whether the current hash is a key, always resume testing for keys starting with the very next hash.

For example, if the pre-arranged salt is abc:

The first index which produces a triple is 18, because the MD5 hash of abc18 contains ...cc38887a5.... However, index 18 does not count as a key for your one-time pad, because none of the next thousand hashes (index 19 through index 1018) contain 88888.
The next index which produces a triple is 39; the hash of abc39 contains eee. It is also the first key: one of the next thousand hashes (the one at index 816) contains eeeee.
None of the next six triples are keys, but the one after that, at index 92, is: it contains 999 and index 200 contains 99999.
Eventually, index 22728 meets all of the criteria to generate the 64th key.
So, using our example salt of abc, index 22728 produces the 64th key.

Given the actual salt in your puzzle input, what index produces your 64th one-time pad key?
*/

const INPUT = "jlmsuwbz"

type Char struct {
	char  byte
	index int
}

func NewChar(char byte, index int) Char {
	return Char{char, index}
}

type Chars []Char

func (c Chars) Len() int {
	return len(c)
}

func (c Chars) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Chars) Less(i, j int) bool {
	return c[i].index < c[j].index
}

func main() {
	fmt.Println(FindIndex())
}

func FindIndex() int {
	var index int
	chars := make(Chars, 0)
	charsMap := make(map[int]Char)

	for len(chars) < 64 || len(charsMap) > 0 {
		input := fmt.Sprintf("%s%d", INPUT, index)
		hashTmp := md5.Sum([]byte(input))
		hash := hex.EncodeToString(hashTmp[:])

		// Find 3's
		if len(chars) < 64 {
			for i := 0; i < len(hash)-2; i++ {
				c := hash[i]
				if c == hash[i+1] && c == hash[i+2] {
					charsMap[index] = NewChar(hash[i], index)
					break
				}
			}
		}

		for _, char := range charsMap {
			// skip currently found 3's
			if char.index == index {
				continue
			}
			// Remove invalid keys
			if index > char.index+1000 {
				delete(charsMap, char.index)
				continue
			}

			for i := 0; i < len(hash)-4; i++ {
				c := hash[i]
				if char.char == c && c == hash[i+1] && c == hash[i+2] &&
					c == hash[i+3] && c == hash[i+4] {
					chars = append(chars, char)
					delete(charsMap, char.index)
					break
				}
			}
		}

		index++
	}

	sort.Sort(chars)
	return chars[63].index
}
