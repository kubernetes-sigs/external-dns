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

func (r router) writeSentence(command []string) error {
	for _, word := range command {
		// log.Println(">>> ", word)
		err := r.writeWord(word)
		if err != nil {
			return err
		}
	}
	return r.writeWord("")
}

func (r router) writeWord(word string) error {
	err := r.writeLength(len(word))
	if err != nil {
		return err
	}
	_, err = r.connection.Write([]byte(word))
	return err
}

func (r router) writeLength(length int) error {
	var b []byte

	switch {
	case length < 0x80:
		b = append(b, byte(length))
	case length < 0x4000:
		length |= 0x8000
		b = append(b, byte((length>>8)&0xff))
		b = append(b, byte(length&0xff))
	case length < 0x200000:
		length |= 0xc00000
		b = append(b, byte((length>>16)&0xff))
		b = append(b, byte((length>>8)&0xff))
		b = append(b, byte(length&0xff))
	case length < 0x10000000:
		length |= 0xE0000000
		b = append(b, byte((length>>24)&0xff))
		b = append(b, byte((length>>16)&0xff))
		b = append(b, byte((length>>8)&0xff))
		b = append(b, byte(length&0xff))
	default:
		b = append(b, byte(0xf0))
		b = append(b, byte((length>>24)&0xff))
		b = append(b, byte((length>>16)&0xff))
		b = append(b, byte((length>>8)&0xff))
		b = append(b, byte(length&0xff))
	}

	_, err := r.connection.Write(b)
	return err
}
