/*
 * MIT License
 *
 * Copyright (c) 2020 Dan <https://github.com/delivrance>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package keygen

import "math/rand"

const (
	vowels            = "aeiou"
	consonants        = "bcdfghjklmnpqrstvwxyz"
	phoneticKeyLength = 10
)

type PhoneticKeygen struct {
	rand *rand.Rand
}

func NewPhoneticKeygen() *PhoneticKeygen {
	return &PhoneticKeygen{
		rand: newRand(),
	}
}

func (pk *PhoneticKeygen) GenerateKey() string {
	key := ""

	for i := 0; i < phoneticKeyLength; i++ {
		if i%2 == 0 {
			key += pk.getRandomFrom(consonants)
		} else {
			key += pk.getRandomFrom(vowels)
		}
	}

	return key
}

func (pk *PhoneticKeygen) getRandomFrom(s string) string {
	return string(s[pk.rand.Intn(len(s))])
}