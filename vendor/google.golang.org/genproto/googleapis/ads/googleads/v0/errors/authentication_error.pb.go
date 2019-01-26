// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/authentication_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible authentication errors.
type AuthenticationErrorEnum_AuthenticationError int32

const (
	// Enum unspecified.
	AuthenticationErrorEnum_UNSPECIFIED AuthenticationErrorEnum_AuthenticationError = 0
	// The received error code is not known in this version.
	AuthenticationErrorEnum_UNKNOWN AuthenticationErrorEnum_AuthenticationError = 1
	// Authentication of the request failed.
	AuthenticationErrorEnum_AUTHENTICATION_ERROR AuthenticationErrorEnum_AuthenticationError = 2
	// Client Customer Id is not a number.
	AuthenticationErrorEnum_CLIENT_CUSTOMER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 5
	// No customer found for the provided customer id.
	AuthenticationErrorEnum_CUSTOMER_NOT_FOUND AuthenticationErrorEnum_AuthenticationError = 8
	// Client's Google Account is deleted.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_DELETED AuthenticationErrorEnum_AuthenticationError = 9
	// Google account login token in the cookie is invalid.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 10
	// A problem occurred during Google account authentication.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_AUTHENTICATION_FAILED AuthenticationErrorEnum_AuthenticationError = 25
	// The user in the google account login token does not match the UserId in
	// the cookie.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH AuthenticationErrorEnum_AuthenticationError = 12
	// Login cookie is required for authentication.
	AuthenticationErrorEnum_LOGIN_COOKIE_REQUIRED AuthenticationErrorEnum_AuthenticationError = 13
	// User in the cookie is not a valid Ads user.
	AuthenticationErrorEnum_NOT_ADS_USER AuthenticationErrorEnum_AuthenticationError = 14
	// Oauth token in the header is not valid.
	AuthenticationErrorEnum_OAUTH_TOKEN_INVALID AuthenticationErrorEnum_AuthenticationError = 15
	// Oauth token in the header has expired.
	AuthenticationErrorEnum_OAUTH_TOKEN_EXPIRED AuthenticationErrorEnum_AuthenticationError = 16
	// Oauth token in the header has been disabled.
	AuthenticationErrorEnum_OAUTH_TOKEN_DISABLED AuthenticationErrorEnum_AuthenticationError = 17
	// Oauth token in the header has been revoked.
	AuthenticationErrorEnum_OAUTH_TOKEN_REVOKED AuthenticationErrorEnum_AuthenticationError = 18
	// Oauth token HTTP header is malformed.
	AuthenticationErrorEnum_OAUTH_TOKEN_HEADER_INVALID AuthenticationErrorEnum_AuthenticationError = 19
	// Login cookie is not valid.
	AuthenticationErrorEnum_LOGIN_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 20
	// User Id in the header is not a valid id.
	AuthenticationErrorEnum_USER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 22
	// An account administrator changed this account's authentication settings.
	// To access this Google Ads account, enable 2-Step Verification in your
	// Google account at https://www.google.com/landing/2step.
	AuthenticationErrorEnum_TWO_STEP_VERIFICATION_NOT_ENROLLED AuthenticationErrorEnum_AuthenticationError = 23
	// An account administrator changed this account's authentication settings.
	// To access this Google Ads account, enable Advanced Protection in your
	// Google account at https://landing.google.com/advancedprotection.
	AuthenticationErrorEnum_ADVANCED_PROTECTION_NOT_ENROLLED AuthenticationErrorEnum_AuthenticationError = 24
)

var AuthenticationErrorEnum_AuthenticationError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "AUTHENTICATION_ERROR",
	5:  "CLIENT_CUSTOMER_ID_INVALID",
	8:  "CUSTOMER_NOT_FOUND",
	9:  "GOOGLE_ACCOUNT_DELETED",
	10: "GOOGLE_ACCOUNT_COOKIE_INVALID",
	25: "GOOGLE_ACCOUNT_AUTHENTICATION_FAILED",
	12: "GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH",
	13: "LOGIN_COOKIE_REQUIRED",
	14: "NOT_ADS_USER",
	15: "OAUTH_TOKEN_INVALID",
	16: "OAUTH_TOKEN_EXPIRED",
	17: "OAUTH_TOKEN_DISABLED",
	18: "OAUTH_TOKEN_REVOKED",
	19: "OAUTH_TOKEN_HEADER_INVALID",
	20: "LOGIN_COOKIE_INVALID",
	22: "USER_ID_INVALID",
	23: "TWO_STEP_VERIFICATION_NOT_ENROLLED",
	24: "ADVANCED_PROTECTION_NOT_ENROLLED",
}
var AuthenticationErrorEnum_AuthenticationError_value = map[string]int32{
	"UNSPECIFIED":                               0,
	"UNKNOWN":                                   1,
	"AUTHENTICATION_ERROR":                      2,
	"CLIENT_CUSTOMER_ID_INVALID":                5,
	"CUSTOMER_NOT_FOUND":                        8,
	"GOOGLE_ACCOUNT_DELETED":                    9,
	"GOOGLE_ACCOUNT_COOKIE_INVALID":             10,
	"GOOGLE_ACCOUNT_AUTHENTICATION_FAILED":      25,
	"GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH": 12,
	"LOGIN_COOKIE_REQUIRED":                     13,
	"NOT_ADS_USER":                              14,
	"OAUTH_TOKEN_INVALID":                       15,
	"OAUTH_TOKEN_EXPIRED":                       16,
	"OAUTH_TOKEN_DISABLED":                      17,
	"OAUTH_TOKEN_REVOKED":                       18,
	"OAUTH_TOKEN_HEADER_INVALID":                19,
	"LOGIN_COOKIE_INVALID":                      20,
	"USER_ID_INVALID":                           22,
	"TWO_STEP_VERIFICATION_NOT_ENROLLED":        23,
	"ADVANCED_PROTECTION_NOT_ENROLLED":          24,
}

func (x AuthenticationErrorEnum_AuthenticationError) String() string {
	return proto.EnumName(AuthenticationErrorEnum_AuthenticationError_name, int32(x))
}
func (AuthenticationErrorEnum_AuthenticationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_authentication_error_8d31aef746d95263, []int{0, 0}
}

// Container for enum describing possible authentication errors.
type AuthenticationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationErrorEnum) Reset()         { *m = AuthenticationErrorEnum{} }
func (m *AuthenticationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AuthenticationErrorEnum) ProtoMessage()    {}
func (*AuthenticationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_error_8d31aef746d95263, []int{0}
}
func (m *AuthenticationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationErrorEnum.Unmarshal(m, b)
}
func (m *AuthenticationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationErrorEnum.Marshal(b, m, deterministic)
}
func (dst *AuthenticationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationErrorEnum.Merge(dst, src)
}
func (m *AuthenticationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AuthenticationErrorEnum.Size(m)
}
func (m *AuthenticationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AuthenticationErrorEnum)(nil), "google.ads.googleads.v0.errors.AuthenticationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.AuthenticationErrorEnum_AuthenticationError", AuthenticationErrorEnum_AuthenticationError_name, AuthenticationErrorEnum_AuthenticationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/authentication_error.proto", fileDescriptor_authentication_error_8d31aef746d95263)
}

var fileDescriptor_authentication_error_8d31aef746d95263 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x65, 0x65, 0x7c, 0x79, 0x83, 0x1a, 0x77, 0xb4, 0xdd, 0x24, 0x2a, 0xa8, 0x26, 0x04, 0x0f,
	0xa4, 0x95, 0x78, 0x42, 0x3c, 0xb9, 0xf1, 0x6d, 0x6b, 0x35, 0xb5, 0x83, 0xe3, 0x64, 0x08, 0x55,
	0xb2, 0xca, 0x5a, 0x85, 0x4a, 0x5b, 0x33, 0x35, 0xdd, 0x7e, 0x10, 0x8f, 0xfc, 0x14, 0x84, 0xf8,
	0x1d, 0xfc, 0x0c, 0xe4, 0x64, 0xa9, 0xba, 0x30, 0x78, 0xca, 0x4d, 0xce, 0x39, 0xf7, 0x1c, 0xdf,
	0x5c, 0xa3, 0xf7, 0x71, 0x92, 0xc4, 0x67, 0xf3, 0xce, 0x74, 0x96, 0x76, 0xf2, 0xd2, 0x56, 0x57,
	0xdd, 0xce, 0x7c, 0xb5, 0x4a, 0x56, 0x69, 0x67, 0x7a, 0xb9, 0xfe, 0x3a, 0x5f, 0xae, 0x17, 0xa7,
	0xd3, 0xf5, 0x22, 0x59, 0x9a, 0xec, 0xab, 0x73, 0xb1, 0x4a, 0xd6, 0x09, 0x69, 0xe5, 0x7c, 0x67,
	0x3a, 0x4b, 0x9d, 0x8d, 0xd4, 0xb9, 0xea, 0x3a, 0xb9, 0xb4, 0xfd, 0x7b, 0x17, 0x35, 0xe8, 0x0d,
	0x39, 0x58, 0x00, 0x96, 0x97, 0xe7, 0xed, 0x9f, 0xbb, 0xa8, 0x76, 0x0b, 0x46, 0xaa, 0x68, 0x2f,
	0x14, 0x81, 0x0f, 0x2e, 0xef, 0x73, 0x60, 0xf8, 0x0e, 0xd9, 0x43, 0x0f, 0x42, 0x31, 0x12, 0xf2,
	0x44, 0xe0, 0x1d, 0xd2, 0x44, 0x07, 0x34, 0xd4, 0x43, 0x10, 0x9a, 0xbb, 0x54, 0x73, 0x29, 0x0c,
	0x28, 0x25, 0x15, 0xae, 0x90, 0x16, 0x3a, 0x72, 0x3d, 0x0e, 0x42, 0x1b, 0x37, 0x0c, 0xb4, 0x1c,
	0x83, 0x32, 0x9c, 0x19, 0x2e, 0x22, 0xea, 0x71, 0x86, 0xef, 0x91, 0x3a, 0x22, 0x1b, 0x40, 0x48,
	0x6d, 0xfa, 0x32, 0x14, 0x0c, 0x3f, 0x24, 0x47, 0xa8, 0x3e, 0x90, 0x72, 0xe0, 0x81, 0xa1, 0xae,
	0x2b, 0x43, 0xa1, 0x0d, 0x03, 0x0f, 0x34, 0x30, 0xfc, 0x88, 0xbc, 0x44, 0xcf, 0x4b, 0x98, 0x2b,
	0xe5, 0x88, 0xc3, 0xa6, 0x2d, 0x22, 0xaf, 0xd1, 0x71, 0x89, 0x52, 0xca, 0xd7, 0xa7, 0xdc, 0x03,
	0x86, 0x0f, 0xc9, 0x5b, 0xf4, 0xa6, 0xc4, 0x0c, 0x03, 0x50, 0x86, 0x0a, 0x66, 0x28, 0x0b, 0xf2,
	0x97, 0x31, 0x0f, 0xc6, 0x54, 0xbb, 0x43, 0xbc, 0x4f, 0x0e, 0xd1, 0x33, 0x4f, 0x0e, 0xb8, 0x28,
	0x2c, 0x15, 0x7c, 0x0c, 0xb9, 0x02, 0x86, 0x1f, 0x13, 0x8c, 0xf6, 0xed, 0x09, 0x0a, 0x15, 0x7e,
	0x42, 0x1a, 0xa8, 0x26, 0xad, 0xaf, 0xd1, 0x72, 0x04, 0x62, 0x13, 0xaf, 0x5a, 0x06, 0xe0, 0x93,
	0x9f, 0xf5, 0xc0, 0x76, 0x90, 0xdb, 0x00, 0xe3, 0x01, 0xed, 0xd9, 0x9c, 0x4f, 0xcb, 0x12, 0x05,
	0x91, 0x1c, 0x01, 0xc3, 0xc4, 0x4e, 0x78, 0x1b, 0x18, 0x02, 0x65, 0x76, 0xc8, 0xd7, 0x5e, 0x35,
	0xdb, 0xf2, 0x46, 0xe2, 0x02, 0x39, 0x20, 0x35, 0x54, 0xcd, 0x8e, 0xb7, 0xf5, 0x43, 0xea, 0xe4,
	0x15, 0x6a, 0xeb, 0x13, 0x69, 0x02, 0x0d, 0xbe, 0x89, 0x40, 0xf1, 0x7e, 0x31, 0x31, 0x7b, 0x36,
	0x10, 0x4a, 0x7a, 0x36, 0x4f, 0x83, 0x1c, 0xa3, 0x17, 0x94, 0x45, 0x54, 0xb8, 0xc0, 0x8c, 0xaf,
	0xa4, 0x06, 0xf7, 0x6f, 0x56, 0xb3, 0xf7, 0x6b, 0x07, 0xb5, 0x4f, 0x93, 0x73, 0xe7, 0xff, 0x1b,
	0xd9, 0x6b, 0xde, 0xb2, 0x72, 0xbe, 0xdd, 0x65, 0x7f, 0xe7, 0x33, 0xbb, 0xd6, 0xc6, 0xc9, 0xd9,
	0x74, 0x19, 0x3b, 0xc9, 0x2a, 0xee, 0xc4, 0xf3, 0x65, 0xb6, 0xe9, 0xc5, 0xc5, 0xb8, 0x58, 0xa4,
	0xff, 0xba, 0x27, 0x1f, 0xf2, 0xc7, 0xb7, 0xca, 0xdd, 0x01, 0xa5, 0xdf, 0x2b, 0xad, 0x41, 0xde,
	0x8c, 0xce, 0x52, 0x27, 0x2f, 0x6d, 0x15, 0x75, 0x9d, 0xcc, 0x32, 0xfd, 0x51, 0x10, 0x26, 0x74,
	0x96, 0x4e, 0x36, 0x84, 0x49, 0xd4, 0x9d, 0xe4, 0x84, 0x2f, 0xf7, 0x33, 0xe3, 0x77, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x01, 0xf3, 0x55, 0x9f, 0x03, 0x00, 0x00,
}