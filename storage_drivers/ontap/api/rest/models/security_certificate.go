// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityCertificate security certificate
//
// swagger:model security_certificate
type SecurityCertificate struct {

	// links
	Links *SecurityCertificateLinks `json:"_links,omitempty"`

	// Provides the key identifier of the issuing CA certificate that signed the SSL certificate.
	// Example: 26:1F:C5:53:5B:D7:9E:E2:37:74:F4:F4:06:09:03:3D:EB:41:75:D7
	// Read Only: true
	AuthorityKeyIdentifier string `json:"authority_key_identifier,omitempty"`

	// Certificate authority
	// Read Only: true
	// Max Length: 256
	// Min Length: 1
	Ca string `json:"ca,omitempty"`

	// FQDN or custom common name. Provide on POST when creating a self-signed certificate.
	// Example: test.domain.com
	CommonName string `json:"common_name,omitempty"`

	// Certificate expiration time. Can be provided on POST if creating self-signed certificate. The expiration time range is between 1 day to 10 years.
	ExpiryTime string `json:"expiry_time,omitempty"`

	// Hashing function. Can be provided on POST when creating a self-signed certificate. Hash functions md5 and sha1 are not allowed on POST.
	// Enum: [sha1 sha256 md5 sha224 sha384 sha512]
	HashFunction *string `json:"hash_function,omitempty"`

	// Chain of intermediate Certificates in PEM format. Only valid in POST when installing a certificate.
	IntermediateCertificates []string `json:"intermediate_certificates,omitempty"`

	// Key size of requested Certificate in bits. One of 512, 1024, 1536, 2048, 3072. Can be provided on POST if creating self-signed certificate. Key size of 512 is not allowed on POST.
	KeySize *int64 `json:"key_size,omitempty"`

	// Certificate name. If not provided in POST, a unique name specific to the SVM is automatically generated.
	// Example: cert1
	Name string `json:"name,omitempty"`

	// Private key Certificate in PEM format. Only valid for create when installing a CA-signed certificate. This is not audited.
	// Example: -----BEGIN PRIVATE KEY----- MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAu1/a8f3G47cZ6pel Hd3aONMNkGJ8vSCH5QjicuDm92VtVwkAACEjIoZSLYlJvPD+odL+lFzVQSmkneW7 VCGqYQIDAQABAkAcfNpg6GCQxoneLOghvlUrRotNZGvqpUOEAvHK3X7AJhz5SU4V an36qvsAt5ghFMVM2iGvGaXbj0dAd+Jg64pxAiEA32Eh9mPtFSmZhTIUMeGcPmPk qIYCEuP8a/ZLmI9s4TsCIQDWvLQuvjSVfwPhi0TFAb5wqAET8X5LBFqtGX5QlUep EwIgFnqM02Gc4wtLoqa2d4qPkYu13+uUW9hLd4XSd6i/OS8CIQDT3elU+Rt+qIwW u0cFrVvNYSV3HNzDfS9N/IoxTagfewIgPvXADe5c2EWbhCUkhN+ZCf38AKewK9TW lQcDy4L+f14= -----END PRIVATE KEY-----
	PrivateKey string `json:"private_key,omitempty"`

	// Public key Certificate in PEM format. If this is not provided in POST, a self-signed certificate is created.
	// Example: -----BEGIN CERTIFICATE----- MIIBuzCCAWWgAwIBAgIIFTZBrqZwUUMwDQYJKoZIhvcNAQELBQAwHDENMAsGA1UE AxMEVEVTVDELMAkGA1UEBhMCVVMwHhcNMTgwNjA4MTgwOTAxWhcNMTkwNjA4MTgw OTAxWjAcMQ0wCwYDVQQDEwRURVNUMQswCQYDVQQGEwJVUzBcMA0GCSqGSIb3DQEB AQUAA0sAMEgCQQDaPvbqUJJFJ6NNTyK3Yb+ytSjJ9aa3yUmYTD9uMiP+6ycjxHWB e8u9z6yCHsW03ync+dnhE5c5z8wuDAY0fv15AgMBAAGjgYowgYcwDAYDVR0TBAUw AwEB/zALBgNVHQ8EBAMCAQYwHQYDVR0OBBYEFMJ7Ev/o/3+YNzYh5XNlqqjnw4zm MEsGA1UdIwREMEKAFMJ7Ev/o/3+YNzYh5XNlqqjnw4zmoSCkHjAcMQ0wCwYDVQQD EwRURVNUMQswCQYDVQQGEwJVU4IIFTZBrqZwUUMwDQYJKoZIhvcNAQELBQADQQAv DovYeyGNnknjGI+TVNX6nDbyzf7zUPqnri0KuvObEeybrbPW45sgsnT5dyeE/32U 9Yr6lklnkBtVBDTmLnrC -----END CERTIFICATE-----
	PublicCertificate string `json:"public_certificate,omitempty"`

	// scope
	Scope NetworkScopeReadonly `json:"scope,omitempty"`

	// Serial number of certificate.
	// Read Only: true
	// Max Length: 40
	// Min Length: 1
	SerialNumber string `json:"serial_number,omitempty"`

	// Provides the key identifier used to identify the public key in the SSL certificate.
	// Example: 26:1F:C5:53:5B:D7:9E:E2:37:74:F4:F4:06:09:03:3D:EB:41:75:D8
	// Read Only: true
	SubjectKeyIdentifier string `json:"subject_key_identifier,omitempty"`

	// svm
	Svm *SecurityCertificateSvm `json:"svm,omitempty"`

	// Type of Certificate. The following types are supported:
	// * client - a certificate and its private key used by an SSL client in ONTAP.
	// * server - a certificate and its private key used by an SSL server in ONTAP.
	// * client_ca - a Certificate Authority certificate used by an SSL server in ONTAP to verify an SSL client certificate.
	// * server_ca - a Certificate Authority certificate used by an SSL client in ONTAP to verify an SSL server certificate.
	// * root_ca - a self-signed certificate used by ONTAP to sign other certificates by acting as a Certificate Authority.
	//
	// Enum: [client server client_ca server_ca root_ca]
	Type string `json:"type,omitempty"`

	// Unique ID that identifies a certificate.
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this security certificate
func (m *SecurityCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificate) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityCertificate) validateCa(formats strfmt.Registry) error {
	if swag.IsZero(m.Ca) { // not required
		return nil
	}

	if err := validate.MinLength("ca", "body", m.Ca, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ca", "body", m.Ca, 256); err != nil {
		return err
	}

	return nil
}

var securityCertificateTypeHashFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha1","sha256","md5","sha224","sha384","sha512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityCertificateTypeHashFunctionPropEnum = append(securityCertificateTypeHashFunctionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// hash_function
	// HashFunction
	// sha1
	// END DEBUGGING
	// SecurityCertificateHashFunctionSha1 captures enum value "sha1"
	SecurityCertificateHashFunctionSha1 string = "sha1"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// hash_function
	// HashFunction
	// sha256
	// END DEBUGGING
	// SecurityCertificateHashFunctionSha256 captures enum value "sha256"
	SecurityCertificateHashFunctionSha256 string = "sha256"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// hash_function
	// HashFunction
	// md5
	// END DEBUGGING
	// SecurityCertificateHashFunctionMd5 captures enum value "md5"
	SecurityCertificateHashFunctionMd5 string = "md5"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// hash_function
	// HashFunction
	// sha224
	// END DEBUGGING
	// SecurityCertificateHashFunctionSha224 captures enum value "sha224"
	SecurityCertificateHashFunctionSha224 string = "sha224"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// hash_function
	// HashFunction
	// sha384
	// END DEBUGGING
	// SecurityCertificateHashFunctionSha384 captures enum value "sha384"
	SecurityCertificateHashFunctionSha384 string = "sha384"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// hash_function
	// HashFunction
	// sha512
	// END DEBUGGING
	// SecurityCertificateHashFunctionSha512 captures enum value "sha512"
	SecurityCertificateHashFunctionSha512 string = "sha512"
)

// prop value enum
func (m *SecurityCertificate) validateHashFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityCertificateTypeHashFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityCertificate) validateHashFunction(formats strfmt.Registry) error {
	if swag.IsZero(m.HashFunction) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashFunctionEnum("hash_function", "body", *m.HashFunction); err != nil {
		return err
	}

	return nil
}

func (m *SecurityCertificate) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *SecurityCertificate) validateSerialNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.SerialNumber) { // not required
		return nil
	}

	if err := validate.MinLength("serial_number", "body", m.SerialNumber, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("serial_number", "body", m.SerialNumber, 40); err != nil {
		return err
	}

	return nil
}

func (m *SecurityCertificate) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

var securityCertificateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client","server","client_ca","server_ca","root_ca"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityCertificateTypeTypePropEnum = append(securityCertificateTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// type
	// Type
	// client
	// END DEBUGGING
	// SecurityCertificateTypeClient captures enum value "client"
	SecurityCertificateTypeClient string = "client"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// type
	// Type
	// server
	// END DEBUGGING
	// SecurityCertificateTypeServer captures enum value "server"
	SecurityCertificateTypeServer string = "server"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// type
	// Type
	// client_ca
	// END DEBUGGING
	// SecurityCertificateTypeClientCa captures enum value "client_ca"
	SecurityCertificateTypeClientCa string = "client_ca"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// type
	// Type
	// server_ca
	// END DEBUGGING
	// SecurityCertificateTypeServerCa captures enum value "server_ca"
	SecurityCertificateTypeServerCa string = "server_ca"

	// BEGIN DEBUGGING
	// security_certificate
	// SecurityCertificate
	// type
	// Type
	// root_ca
	// END DEBUGGING
	// SecurityCertificateTypeRootCa captures enum value "root_ca"
	SecurityCertificateTypeRootCa string = "root_ca"
)

// prop value enum
func (m *SecurityCertificate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityCertificateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityCertificate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security certificate based on the context it is used
func (m *SecurityCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthorityKeyIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSerialNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubjectKeyIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificate) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityCertificate) contextValidateAuthorityKeyIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "authority_key_identifier", "body", string(m.AuthorityKeyIdentifier)); err != nil {
		return err
	}

	return nil
}

func (m *SecurityCertificate) contextValidateCa(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ca", "body", string(m.Ca)); err != nil {
		return err
	}

	return nil
}

func (m *SecurityCertificate) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Scope.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *SecurityCertificate) contextValidateSerialNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "serial_number", "body", string(m.SerialNumber)); err != nil {
		return err
	}

	return nil
}

func (m *SecurityCertificate) contextValidateSubjectKeyIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subject_key_identifier", "body", string(m.SubjectKeyIdentifier)); err != nil {
		return err
	}

	return nil
}

func (m *SecurityCertificate) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityCertificate) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityCertificate) UnmarshalBinary(b []byte) error {
	var res SecurityCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityCertificateLinks security certificate links
//
// swagger:model SecurityCertificateLinks
type SecurityCertificateLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this security certificate links
func (m *SecurityCertificateLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificateLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security certificate links based on the context it is used
func (m *SecurityCertificateLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificateLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityCertificateLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityCertificateLinks) UnmarshalBinary(b []byte) error {
	var res SecurityCertificateLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityCertificateSvm security certificate svm
//
// swagger:model SecurityCertificateSvm
type SecurityCertificateSvm struct {

	// links
	Links *SecurityCertificateSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this security certificate svm
func (m *SecurityCertificateSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificateSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security certificate svm based on the context it is used
func (m *SecurityCertificateSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificateSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityCertificateSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityCertificateSvm) UnmarshalBinary(b []byte) error {
	var res SecurityCertificateSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityCertificateSvmLinks security certificate svm links
//
// swagger:model SecurityCertificateSvmLinks
type SecurityCertificateSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this security certificate svm links
func (m *SecurityCertificateSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificateSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security certificate svm links based on the context it is used
func (m *SecurityCertificateSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCertificateSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityCertificateSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityCertificateSvmLinks) UnmarshalBinary(b []byte) error {
	var res SecurityCertificateSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
