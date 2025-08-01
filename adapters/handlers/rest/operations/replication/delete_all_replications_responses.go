//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// DeleteAllReplicationsNoContentCode is the HTTP code returned for type DeleteAllReplicationsNoContent
const DeleteAllReplicationsNoContentCode int = 204

/*
DeleteAllReplicationsNoContent Replication operation registered successfully

swagger:response deleteAllReplicationsNoContent
*/
type DeleteAllReplicationsNoContent struct {
}

// NewDeleteAllReplicationsNoContent creates DeleteAllReplicationsNoContent with default headers values
func NewDeleteAllReplicationsNoContent() *DeleteAllReplicationsNoContent {

	return &DeleteAllReplicationsNoContent{}
}

// WriteResponse to the client
func (o *DeleteAllReplicationsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteAllReplicationsBadRequestCode is the HTTP code returned for type DeleteAllReplicationsBadRequest
const DeleteAllReplicationsBadRequestCode int = 400

/*
DeleteAllReplicationsBadRequest Malformed request.

swagger:response deleteAllReplicationsBadRequest
*/
type DeleteAllReplicationsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteAllReplicationsBadRequest creates DeleteAllReplicationsBadRequest with default headers values
func NewDeleteAllReplicationsBadRequest() *DeleteAllReplicationsBadRequest {

	return &DeleteAllReplicationsBadRequest{}
}

// WithPayload adds the payload to the delete all replications bad request response
func (o *DeleteAllReplicationsBadRequest) WithPayload(payload *models.ErrorResponse) *DeleteAllReplicationsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete all replications bad request response
func (o *DeleteAllReplicationsBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAllReplicationsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAllReplicationsUnauthorizedCode is the HTTP code returned for type DeleteAllReplicationsUnauthorized
const DeleteAllReplicationsUnauthorizedCode int = 401

/*
DeleteAllReplicationsUnauthorized Unauthorized or invalid credentials.

swagger:response deleteAllReplicationsUnauthorized
*/
type DeleteAllReplicationsUnauthorized struct {
}

// NewDeleteAllReplicationsUnauthorized creates DeleteAllReplicationsUnauthorized with default headers values
func NewDeleteAllReplicationsUnauthorized() *DeleteAllReplicationsUnauthorized {

	return &DeleteAllReplicationsUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAllReplicationsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteAllReplicationsForbiddenCode is the HTTP code returned for type DeleteAllReplicationsForbidden
const DeleteAllReplicationsForbiddenCode int = 403

/*
DeleteAllReplicationsForbidden Forbidden

swagger:response deleteAllReplicationsForbidden
*/
type DeleteAllReplicationsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteAllReplicationsForbidden creates DeleteAllReplicationsForbidden with default headers values
func NewDeleteAllReplicationsForbidden() *DeleteAllReplicationsForbidden {

	return &DeleteAllReplicationsForbidden{}
}

// WithPayload adds the payload to the delete all replications forbidden response
func (o *DeleteAllReplicationsForbidden) WithPayload(payload *models.ErrorResponse) *DeleteAllReplicationsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete all replications forbidden response
func (o *DeleteAllReplicationsForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAllReplicationsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAllReplicationsUnprocessableEntityCode is the HTTP code returned for type DeleteAllReplicationsUnprocessableEntity
const DeleteAllReplicationsUnprocessableEntityCode int = 422

/*
DeleteAllReplicationsUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous.

swagger:response deleteAllReplicationsUnprocessableEntity
*/
type DeleteAllReplicationsUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteAllReplicationsUnprocessableEntity creates DeleteAllReplicationsUnprocessableEntity with default headers values
func NewDeleteAllReplicationsUnprocessableEntity() *DeleteAllReplicationsUnprocessableEntity {

	return &DeleteAllReplicationsUnprocessableEntity{}
}

// WithPayload adds the payload to the delete all replications unprocessable entity response
func (o *DeleteAllReplicationsUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *DeleteAllReplicationsUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete all replications unprocessable entity response
func (o *DeleteAllReplicationsUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAllReplicationsUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAllReplicationsInternalServerErrorCode is the HTTP code returned for type DeleteAllReplicationsInternalServerError
const DeleteAllReplicationsInternalServerErrorCode int = 500

/*
DeleteAllReplicationsInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response deleteAllReplicationsInternalServerError
*/
type DeleteAllReplicationsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteAllReplicationsInternalServerError creates DeleteAllReplicationsInternalServerError with default headers values
func NewDeleteAllReplicationsInternalServerError() *DeleteAllReplicationsInternalServerError {

	return &DeleteAllReplicationsInternalServerError{}
}

// WithPayload adds the payload to the delete all replications internal server error response
func (o *DeleteAllReplicationsInternalServerError) WithPayload(payload *models.ErrorResponse) *DeleteAllReplicationsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete all replications internal server error response
func (o *DeleteAllReplicationsInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAllReplicationsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAllReplicationsNotImplementedCode is the HTTP code returned for type DeleteAllReplicationsNotImplemented
const DeleteAllReplicationsNotImplementedCode int = 501

/*
DeleteAllReplicationsNotImplemented Replica movement operations are disabled.

swagger:response deleteAllReplicationsNotImplemented
*/
type DeleteAllReplicationsNotImplemented struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteAllReplicationsNotImplemented creates DeleteAllReplicationsNotImplemented with default headers values
func NewDeleteAllReplicationsNotImplemented() *DeleteAllReplicationsNotImplemented {

	return &DeleteAllReplicationsNotImplemented{}
}

// WithPayload adds the payload to the delete all replications not implemented response
func (o *DeleteAllReplicationsNotImplemented) WithPayload(payload *models.ErrorResponse) *DeleteAllReplicationsNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete all replications not implemented response
func (o *DeleteAllReplicationsNotImplemented) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAllReplicationsNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
