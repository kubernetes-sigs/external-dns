/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

func (r router) readSentence() ([]string, error) {
	reply := []string{}
	for {
		word := r.readWord()
		// log.Println("<<< ", word)
		if word == "" {
			break
		}
		reply = append(reply, word)
	}
	return reply, nil
}

func (r router) readWord() string {
	length := r.readLength()
	b := make([]byte, length)
	r.connection.Read(b)
	return string(b)
}

// returns byte value or 0 on error
func (r router) readByte() int {
	b := make([]byte, 1)
	_, err := r.connection.Read(b)
	if err != nil {
		return 0
	}
	return int(b[0])
}

func (r router) readLength() int {
	length := r.readByte()
	switch {
	case (length & 0x80) == 0x00:
		// one byte, done
	case (length & 0xc0) == 0x80:
		length &= ^0xc0
		length <<= 8
		length += r.readByte()
	case (length & 0xe0) == 0xc0:
		length &= ^0xe0
		length <<= 8
		length += r.readByte()
		length <<= 8
		length += r.readByte()
	case (length & 0xf0) == 0xe0:
		length &= ^0xf0
		length <<= 8
		length += r.readByte()
		length <<= 8
		length += r.readByte()
		length <<= 8
		length += r.readByte()
	case (length & 0xf8) == 0xf0:
		length = r.readByte()
		length <<= 8
		length += r.readByte()
		length <<= 8
		length += r.readByte()
		length <<= 8
		length += r.readByte()
	}
	return length
}
