// Copyright 2013-2015 Apcera Inc. All rights reserved.

package gssapi

/*
#include <gssapi/gssapi.h>

OM_uint32
wrap_gss_get_mic(void *fp,
	OM_uint32 * minor_status,
	const gss_ctx_id_t context_handle,
	gss_qop_t qop_req,
	const gss_buffer_t message_buffer,
	gss_buffer_t message_token)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_ctx_id_t,
		gss_qop_t,
		const gss_buffer_t,
		gss_buffer_t)
	) fp)(
		minor_status,
		context_handle,
		qop_req,
		message_buffer,
		message_token);
}

OM_uint32
wrap_gss_verify_mic(void *fp,
	OM_uint32 * minor_status,
	const gss_ctx_id_t context_handle,
	const gss_buffer_t message_buffer,
	const gss_buffer_t token_buffer,
	gss_qop_t * qop_state)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_ctx_id_t,
		const gss_buffer_t,
		const gss_buffer_t,
		gss_qop_t *)
	) fp)(
		minor_status,
		context_handle,
		message_buffer,
		token_buffer,
		qop_state);
}

OM_uint32
wrap_gss_wrap(void *fp,
	OM_uint32 * minor_status,
	const gss_ctx_id_t context_handle,
	int conf_req_flag,
	gss_qop_t qop_req,
	const gss_buffer_t input_message_buffer,
	int * conf_state,
	gss_buffer_t output_message_buffer)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_ctx_id_t,
		int,
		gss_qop_t,
		const gss_buffer_t,
		int *,
		gss_buffer_t)
	) fp)(
		minor_status,
		context_handle,
		conf_req_flag,
		qop_req,
		input_message_buffer,
		conf_state,
		output_message_buffer);
}

OM_uint32
wrap_gss_unwrap(void *fp,
	OM_uint32 * minor_status,
	const gss_ctx_id_t context_handle,
	const gss_buffer_t input_message_buffer,
	gss_buffer_t output_message_buffer,
	int * conf_state,
	gss_qop_t * qop_state)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_ctx_id_t,
		const gss_buffer_t,
		gss_buffer_t,
		int *,
		gss_qop_t *)
	) fp)(
		minor_status,
		context_handle,
		input_message_buffer,
		output_message_buffer,
		conf_state,
		qop_state);
}

*/
import "C"

// GetMIC implements gss_GetMIC API, as per https://tools.ietf.org/html/rfc2743#page-63.
// messageToken must be .Release()-ed by the caller.
func (ctx *CtxId) GetMIC(qopReq QOP, messageBuffer *Buffer) (
	messageToken *Buffer, err error) {

	min := C.OM_uint32(0)

	token, err := ctx.MakeBuffer(allocGSSAPI)
	if err != nil {
		return nil, err
	}

	maj := C.wrap_gss_get_mic(ctx.Fp_gss_get_mic,
		&min,
		ctx.C_gss_ctx_id_t,
		C.gss_qop_t(qopReq),
		messageBuffer.C_gss_buffer_t,
		token.C_gss_buffer_t)

	err = ctx.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// VerifyMIC implements gss_VerifyMIC API, as per https://tools.ietf.org/html/rfc2743#page-64.
func (ctx *CtxId) VerifyMIC(messageBuffer *Buffer, tokenBuffer *Buffer) (
	qopState QOP, err error) {

	min := C.OM_uint32(0)
	qop := C.gss_qop_t(0)

	maj := C.wrap_gss_verify_mic(ctx.Fp_gss_verify_mic,
		&min,
		ctx.C_gss_ctx_id_t,
		messageBuffer.C_gss_buffer_t,
		tokenBuffer.C_gss_buffer_t,
		&qop)

	err = ctx.stashLastStatus(maj, min)
	if err != nil {
		return 0, err
	}

	return QOP(qop), nil
}

// Wrap implements gss_wrap API, as per https://tools.ietf.org/html/rfc2743#page-65.
// outputMessageBuffer must be .Release()-ed by the caller
func (ctx *CtxId) Wrap(
	confReq bool, qopReq QOP, inputMessageBuffer *Buffer) (
	confState bool, outputMessageBuffer *Buffer, err error) {

	min := C.OM_uint32(0)

	encrypt := C.int(0)
	if confReq {
		encrypt = 1
	}

	outputMessageBuffer, err = ctx.MakeBuffer(allocGSSAPI)
	if err != nil {
		return false, nil, err
	}

	encrypted := C.int(0)

	maj := C.wrap_gss_wrap(ctx.Fp_gss_wrap,
		&min,
		ctx.C_gss_ctx_id_t,
		encrypt,
		C.gss_qop_t(qopReq),
		inputMessageBuffer.C_gss_buffer_t,
		&encrypted,
		outputMessageBuffer.C_gss_buffer_t)

	err = ctx.stashLastStatus(maj, min)
	if err != nil {
		return false, nil, err
	}

	return encrypted != 0,
		outputMessageBuffer,
		nil
}

// Unwrap implements gss_unwrap API, as per https://tools.ietf.org/html/rfc2743#page-66.
// outputMessageBuffer must be .Release()-ed by the caller
func (ctx *CtxId) Unwrap(
	inputMessageBuffer *Buffer) (
	outputMessageBuffer *Buffer, confState bool, qopState QOP, err error) {

	min := C.OM_uint32(0)

	outputMessageBuffer, err = ctx.MakeBuffer(allocGSSAPI)
	if err != nil {
		return nil, false, 0, err
	}

	encrypted := C.int(0)
	qop := C.gss_qop_t(0)

	maj := C.wrap_gss_unwrap(ctx.Fp_gss_unwrap,
		&min,
		ctx.C_gss_ctx_id_t,
		inputMessageBuffer.C_gss_buffer_t,
		outputMessageBuffer.C_gss_buffer_t,
		&encrypted,
		&qop)

	err = ctx.stashLastStatus(maj, min)
	if err != nil {
		return nil, false, 0, err
	}

	return outputMessageBuffer,
		encrypted != 0,
		QOP(qop),
		nil
}
