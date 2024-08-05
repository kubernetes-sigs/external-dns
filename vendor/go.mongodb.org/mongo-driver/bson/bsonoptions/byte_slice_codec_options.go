// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package bsonoptions

// ByteSliceCodecOptions represents all possible options for byte slice encoding and decoding.
<<<<<<< HEAD
type ByteSliceCodecOptions struct {
	EncodeNilAsEmpty *bool // Specifies if a nil byte slice should encode as an empty binary instead of null. Defaults to false.
}

// ByteSliceCodec creates a new *ByteSliceCodecOptions
func ByteSliceCodec() *ByteSliceCodecOptions {
	return &ByteSliceCodecOptions{}
}

// SetEncodeNilAsEmpty specifies  if a nil byte slice should encode as an empty binary instead of null. Defaults to false.
func (bs *ByteSliceCodecOptions) SetEncodeNilAsEmpty(b bool) *ByteSliceCodecOptions {
	bs.EncodeNilAsEmpty = &b
	return bs
}

// MergeByteSliceCodecOptions combines the given *ByteSliceCodecOptions into a single *ByteSliceCodecOptions in a last one wins fashion.
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
//
// Deprecated: Use the bson.Encoder and bson.Decoder configuration methods to set the desired BSON marshal
// and unmarshal behavior instead.
type ByteSliceCodecOptions struct {
	EncodeNilAsEmpty *bool // Specifies if a nil byte slice should encode as an empty binary instead of null. Defaults to false.
}

// ByteSliceCodec creates a new *ByteSliceCodecOptions
//
// Deprecated: Use the bson.Encoder and bson.Decoder configuration methods to set the desired BSON marshal
// and unmarshal behavior instead.
func ByteSliceCodec() *ByteSliceCodecOptions {
	return &ByteSliceCodecOptions{}
}

// SetEncodeNilAsEmpty specifies  if a nil byte slice should encode as an empty binary instead of null. Defaults to false.
//
// Deprecated: Use [go.mongodb.org/mongo-driver/bson.Encoder.NilByteSliceAsEmpty] instead.
func (bs *ByteSliceCodecOptions) SetEncodeNilAsEmpty(b bool) *ByteSliceCodecOptions {
	bs.EncodeNilAsEmpty = &b
	return bs
}

// MergeByteSliceCodecOptions combines the given *ByteSliceCodecOptions into a single *ByteSliceCodecOptions in a last one wins fashion.
//
// Deprecated: Merging options structs will not be supported in Go Driver 2.0. Users should create a
// single options struct instead.
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func MergeByteSliceCodecOptions(opts ...*ByteSliceCodecOptions) *ByteSliceCodecOptions {
	bs := ByteSliceCodec()
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.EncodeNilAsEmpty != nil {
			bs.EncodeNilAsEmpty = opt.EncodeNilAsEmpty
		}
	}

	return bs
}
