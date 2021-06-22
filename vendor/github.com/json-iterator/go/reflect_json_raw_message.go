package jsoniter

import (
	"encoding/json"
	"github.com/modern-go/reflect2"
	"unsafe"
)

var jsonRawMessageType = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var jsoniterRawMessageType = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func createEncoderOfJsonRawMessage(ctx *ctx, typ reflect2.Type) ValEncoder {
	if typ == jsonRawMessageType {
		return &jsonRawMessageCodec{}
	}
	if typ == jsoniterRawMessageType {
		return &jsoniterRawMessageCodec{}
	}
	return nil
}

func createDecoderOfJsonRawMessage(ctx *ctx, typ reflect2.Type) ValDecoder {
	if typ == jsonRawMessageType {
		return &jsonRawMessageCodec{}
	}
	if typ == jsoniterRawMessageType {
		return &jsoniterRawMessageCodec{}
	}
	return nil
}

type jsonRawMessageCodec struct {
}

func (codec *jsonRawMessageCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	if iter.ReadNil() {
		*((*json.RawMessage)(ptr)) = nil
	} else {
		*((*json.RawMessage)(ptr)) = iter.SkipAndReturnBytes()
	}
}

func (codec *jsonRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	if *((*json.RawMessage)(ptr)) == nil {
		stream.WriteNil()
	} else {
		stream.WriteRaw(string(*((*json.RawMessage)(ptr))))
	}
}

func (codec *jsonRawMessageCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(ptr))) == 0
}

type jsoniterRawMessageCodec struct {
}

func (codec *jsoniterRawMessageCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if iter.ReadNil() {
		*((*RawMessage)(ptr)) = nil
	} else {
		*((*RawMessage)(ptr)) = iter.SkipAndReturnBytes()
	}
}

func (codec *jsoniterRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	if *((*RawMessage)(ptr)) == nil {
		stream.WriteNil()
	} else {
		stream.WriteRaw(string(*((*RawMessage)(ptr))))
	}
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	*((*json.RawMessage)(ptr)) = json.RawMessage(iter.SkipAndReturnBytes())
||||||| parent of 5ce8c7613 (update vendored files)
	*((*json.RawMessage)(ptr)) = json.RawMessage(iter.SkipAndReturnBytes())
=======
	if iter.ReadNil() {
		*((*json.RawMessage)(ptr)) = nil
	} else {
		*((*json.RawMessage)(ptr)) = iter.SkipAndReturnBytes()
	}
>>>>>>> 5ce8c7613 (update vendored files)
}

func (codec *jsonRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	if *((*json.RawMessage)(ptr)) == nil {
		stream.WriteNil()
	} else {
		stream.WriteRaw(string(*((*json.RawMessage)(ptr))))
	}
}

func (codec *jsonRawMessageCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(ptr))) == 0
}

type jsoniterRawMessageCodec struct {
}

func (codec *jsoniterRawMessageCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if iter.ReadNil() {
		*((*RawMessage)(ptr)) = nil
	} else {
		*((*RawMessage)(ptr)) = iter.SkipAndReturnBytes()
	}
}

func (codec *jsoniterRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
<<<<<<< HEAD
	stream.WriteRaw(string(*((*RawMessage)(ptr))))
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	stream.WriteRaw(string(*((*RawMessage)(ptr))))
=======
	if *((*RawMessage)(ptr)) == nil {
		stream.WriteNil()
	} else {
		stream.WriteRaw(string(*((*RawMessage)(ptr))))
	}
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	*((*json.RawMessage)(ptr)) = json.RawMessage(iter.SkipAndReturnBytes())
}

func (codec *jsonRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteRaw(string(*((*json.RawMessage)(ptr))))
}

func (codec *jsonRawMessageCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(ptr))) == 0
}

type jsoniterRawMessageCodec struct {
}

func (codec *jsoniterRawMessageCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	*((*RawMessage)(ptr)) = RawMessage(iter.SkipAndReturnBytes())
}

func (codec *jsoniterRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	stream.WriteRaw(string(*((*RawMessage)(ptr))))
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

func (codec *jsoniterRawMessageCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(*((*RawMessage)(ptr))) == 0
}
