// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keymanager.proto

package keymanager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import plugin "github.com/spiffe/spire/proto/common/plugin"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureRequest = plugin.ConfigureRequest

// ConfigureResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureResponse = plugin.ConfigureResponse

// GetPluginInfoRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoRequest = plugin.GetPluginInfoRequest

// GetPluginInfoResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoResponse = plugin.GetPluginInfoResponse

type KeyAlgorithm int32

const (
	KeyAlgorithm_UNSPECIFIED_KEY_ALGORITHM KeyAlgorithm = 0
	KeyAlgorithm_ECDSA_P256                KeyAlgorithm = 1
	KeyAlgorithm_ECDSA_P384                KeyAlgorithm = 2
)

var KeyAlgorithm_name = map[int32]string{
	0: "UNSPECIFIED_KEY_ALGORITHM",
	1: "ECDSA_P256",
	2: "ECDSA_P384",
}
var KeyAlgorithm_value = map[string]int32{
	"UNSPECIFIED_KEY_ALGORITHM": 0,
	"ECDSA_P256":                1,
	"ECDSA_P384":                2,
}

func (x KeyAlgorithm) String() string {
	return proto.EnumName(KeyAlgorithm_name, int32(x))
}
func (KeyAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{0}
}

type HashAlgorithm int32

const (
	HashAlgorithm_UNSPECIFIED_HASH_ALGORITHM HashAlgorithm = 0
	// These entries (and their values) line up with a subset of the go
	// crypto.Hash constants
	HashAlgorithm_SHA1       HashAlgorithm = 3
	HashAlgorithm_SHA224     HashAlgorithm = 4
	HashAlgorithm_SHA256     HashAlgorithm = 5
	HashAlgorithm_SHA384     HashAlgorithm = 6
	HashAlgorithm_SHA512     HashAlgorithm = 7
	HashAlgorithm_SHA3_224   HashAlgorithm = 10
	HashAlgorithm_SHA3_256   HashAlgorithm = 11
	HashAlgorithm_SHA3_384   HashAlgorithm = 12
	HashAlgorithm_SHA3_512   HashAlgorithm = 13
	HashAlgorithm_SHA512_224 HashAlgorithm = 14
	HashAlgorithm_SHA512_256 HashAlgorithm = 15
)

var HashAlgorithm_name = map[int32]string{
	0:  "UNSPECIFIED_HASH_ALGORITHM",
	3:  "SHA1",
	4:  "SHA224",
	5:  "SHA256",
	6:  "SHA384",
	7:  "SHA512",
	10: "SHA3_224",
	11: "SHA3_256",
	12: "SHA3_384",
	13: "SHA3_512",
	14: "SHA512_224",
	15: "SHA512_256",
}
var HashAlgorithm_value = map[string]int32{
	"UNSPECIFIED_HASH_ALGORITHM": 0,
	"SHA1":       3,
	"SHA224":     4,
	"SHA256":     5,
	"SHA384":     6,
	"SHA512":     7,
	"SHA3_224":   10,
	"SHA3_256":   11,
	"SHA3_384":   12,
	"SHA3_512":   13,
	"SHA512_224": 14,
	"SHA512_256": 15,
}

func (x HashAlgorithm) String() string {
	return proto.EnumName(HashAlgorithm_name, int32(x))
}
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{1}
}

type PublicKey struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Algorithm            KeyAlgorithm `protobuf:"varint,2,opt,name=algorithm,enum=spire.server.keymanager.KeyAlgorithm" json:"algorithm,omitempty"`
	PkixData             []byte       `protobuf:"bytes,3,opt,name=pkix_data,json=pkixData,proto3" json:"pkix_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{0}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (dst *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(dst, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PublicKey) GetAlgorithm() KeyAlgorithm {
	if m != nil {
		return m.Algorithm
	}
	return KeyAlgorithm_UNSPECIFIED_KEY_ALGORITHM
}

func (m *PublicKey) GetPkixData() []byte {
	if m != nil {
		return m.PkixData
	}
	return nil
}

type GenerateKeyRequest struct {
	KeyId                string       `protobuf:"bytes,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	KeyAlgorithm         KeyAlgorithm `protobuf:"varint,2,opt,name=key_algorithm,json=keyAlgorithm,enum=spire.server.keymanager.KeyAlgorithm" json:"key_algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GenerateKeyRequest) Reset()         { *m = GenerateKeyRequest{} }
func (m *GenerateKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeyRequest) ProtoMessage()    {}
func (*GenerateKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{1}
}
func (m *GenerateKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeyRequest.Unmarshal(m, b)
}
func (m *GenerateKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeyRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeyRequest.Merge(dst, src)
}
func (m *GenerateKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeyRequest.Size(m)
}
func (m *GenerateKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeyRequest proto.InternalMessageInfo

func (m *GenerateKeyRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *GenerateKeyRequest) GetKeyAlgorithm() KeyAlgorithm {
	if m != nil {
		return m.KeyAlgorithm
	}
	return KeyAlgorithm_UNSPECIFIED_KEY_ALGORITHM
}

type GenerateKeyResponse struct {
	PublicKey            *PublicKey `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GenerateKeyResponse) Reset()         { *m = GenerateKeyResponse{} }
func (m *GenerateKeyResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeyResponse) ProtoMessage()    {}
func (*GenerateKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{2}
}
func (m *GenerateKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeyResponse.Unmarshal(m, b)
}
func (m *GenerateKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeyResponse.Marshal(b, m, deterministic)
}
func (dst *GenerateKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeyResponse.Merge(dst, src)
}
func (m *GenerateKeyResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeyResponse.Size(m)
}
func (m *GenerateKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeyResponse proto.InternalMessageInfo

func (m *GenerateKeyResponse) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type GetPublicKeyRequest struct {
	KeyId                string   `protobuf:"bytes,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicKeyRequest) Reset()         { *m = GetPublicKeyRequest{} }
func (m *GetPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicKeyRequest) ProtoMessage()    {}
func (*GetPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{3}
}
func (m *GetPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicKeyRequest.Unmarshal(m, b)
}
func (m *GetPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicKeyRequest.Marshal(b, m, deterministic)
}
func (dst *GetPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicKeyRequest.Merge(dst, src)
}
func (m *GetPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GetPublicKeyRequest.Size(m)
}
func (m *GetPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicKeyRequest proto.InternalMessageInfo

func (m *GetPublicKeyRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

type GetPublicKeyResponse struct {
	PublicKey            *PublicKey `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetPublicKeyResponse) Reset()         { *m = GetPublicKeyResponse{} }
func (m *GetPublicKeyResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicKeyResponse) ProtoMessage()    {}
func (*GetPublicKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{4}
}
func (m *GetPublicKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicKeyResponse.Unmarshal(m, b)
}
func (m *GetPublicKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicKeyResponse.Marshal(b, m, deterministic)
}
func (dst *GetPublicKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicKeyResponse.Merge(dst, src)
}
func (m *GetPublicKeyResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicKeyResponse.Size(m)
}
func (m *GetPublicKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicKeyResponse proto.InternalMessageInfo

func (m *GetPublicKeyResponse) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type GetPublicKeysRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicKeysRequest) Reset()         { *m = GetPublicKeysRequest{} }
func (m *GetPublicKeysRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicKeysRequest) ProtoMessage()    {}
func (*GetPublicKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{5}
}
func (m *GetPublicKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicKeysRequest.Unmarshal(m, b)
}
func (m *GetPublicKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicKeysRequest.Marshal(b, m, deterministic)
}
func (dst *GetPublicKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicKeysRequest.Merge(dst, src)
}
func (m *GetPublicKeysRequest) XXX_Size() int {
	return xxx_messageInfo_GetPublicKeysRequest.Size(m)
}
func (m *GetPublicKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicKeysRequest proto.InternalMessageInfo

type GetPublicKeysResponse struct {
	PublicKeys           []*PublicKey `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetPublicKeysResponse) Reset()         { *m = GetPublicKeysResponse{} }
func (m *GetPublicKeysResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicKeysResponse) ProtoMessage()    {}
func (*GetPublicKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{6}
}
func (m *GetPublicKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicKeysResponse.Unmarshal(m, b)
}
func (m *GetPublicKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicKeysResponse.Marshal(b, m, deterministic)
}
func (dst *GetPublicKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicKeysResponse.Merge(dst, src)
}
func (m *GetPublicKeysResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicKeysResponse.Size(m)
}
func (m *GetPublicKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicKeysResponse proto.InternalMessageInfo

func (m *GetPublicKeysResponse) GetPublicKeys() []*PublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type SignDataRequest struct {
	KeyId                string        `protobuf:"bytes,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	HashAlgorithm        HashAlgorithm `protobuf:"varint,2,opt,name=hash_algorithm,json=hashAlgorithm,enum=spire.server.keymanager.HashAlgorithm" json:"hash_algorithm,omitempty"`
	Data                 []byte        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SignDataRequest) Reset()         { *m = SignDataRequest{} }
func (m *SignDataRequest) String() string { return proto.CompactTextString(m) }
func (*SignDataRequest) ProtoMessage()    {}
func (*SignDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{7}
}
func (m *SignDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignDataRequest.Unmarshal(m, b)
}
func (m *SignDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignDataRequest.Marshal(b, m, deterministic)
}
func (dst *SignDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignDataRequest.Merge(dst, src)
}
func (m *SignDataRequest) XXX_Size() int {
	return xxx_messageInfo_SignDataRequest.Size(m)
}
func (m *SignDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignDataRequest proto.InternalMessageInfo

func (m *SignDataRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SignDataRequest) GetHashAlgorithm() HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return HashAlgorithm_UNSPECIFIED_HASH_ALGORITHM
}

func (m *SignDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SignDataResponse struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignDataResponse) Reset()         { *m = SignDataResponse{} }
func (m *SignDataResponse) String() string { return proto.CompactTextString(m) }
func (*SignDataResponse) ProtoMessage()    {}
func (*SignDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keymanager_a61332f079830969, []int{8}
}
func (m *SignDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignDataResponse.Unmarshal(m, b)
}
func (m *SignDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignDataResponse.Marshal(b, m, deterministic)
}
func (dst *SignDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignDataResponse.Merge(dst, src)
}
func (m *SignDataResponse) XXX_Size() int {
	return xxx_messageInfo_SignDataResponse.Size(m)
}
func (m *SignDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignDataResponse proto.InternalMessageInfo

func (m *SignDataResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicKey)(nil), "spire.server.keymanager.PublicKey")
	proto.RegisterType((*GenerateKeyRequest)(nil), "spire.server.keymanager.GenerateKeyRequest")
	proto.RegisterType((*GenerateKeyResponse)(nil), "spire.server.keymanager.GenerateKeyResponse")
	proto.RegisterType((*GetPublicKeyRequest)(nil), "spire.server.keymanager.GetPublicKeyRequest")
	proto.RegisterType((*GetPublicKeyResponse)(nil), "spire.server.keymanager.GetPublicKeyResponse")
	proto.RegisterType((*GetPublicKeysRequest)(nil), "spire.server.keymanager.GetPublicKeysRequest")
	proto.RegisterType((*GetPublicKeysResponse)(nil), "spire.server.keymanager.GetPublicKeysResponse")
	proto.RegisterType((*SignDataRequest)(nil), "spire.server.keymanager.SignDataRequest")
	proto.RegisterType((*SignDataResponse)(nil), "spire.server.keymanager.SignDataResponse")
	proto.RegisterEnum("spire.server.keymanager.KeyAlgorithm", KeyAlgorithm_name, KeyAlgorithm_value)
	proto.RegisterEnum("spire.server.keymanager.HashAlgorithm", HashAlgorithm_name, HashAlgorithm_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KeyManager service

type KeyManagerClient interface {
	// Generates a new key
	GenerateKey(ctx context.Context, in *GenerateKeyRequest, opts ...grpc.CallOption) (*GenerateKeyResponse, error)
	// Get a public key by key id
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error)
	// Gets all public keys
	GetPublicKeys(ctx context.Context, in *GetPublicKeysRequest, opts ...grpc.CallOption) (*GetPublicKeysResponse, error)
	// Signs data with private key
	SignData(ctx context.Context, in *SignDataRequest, opts ...grpc.CallOption) (*SignDataResponse, error)
	// Applies the plugin configuration
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	// Returns the version and related metadata of the installed plugin
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type keyManagerClient struct {
	cc *grpc.ClientConn
}

func NewKeyManagerClient(cc *grpc.ClientConn) KeyManagerClient {
	return &keyManagerClient{cc}
}

func (c *keyManagerClient) GenerateKey(ctx context.Context, in *GenerateKeyRequest, opts ...grpc.CallOption) (*GenerateKeyResponse, error) {
	out := new(GenerateKeyResponse)
	err := grpc.Invoke(ctx, "/spire.server.keymanager.KeyManager/GenerateKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagerClient) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error) {
	out := new(GetPublicKeyResponse)
	err := grpc.Invoke(ctx, "/spire.server.keymanager.KeyManager/GetPublicKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagerClient) GetPublicKeys(ctx context.Context, in *GetPublicKeysRequest, opts ...grpc.CallOption) (*GetPublicKeysResponse, error) {
	out := new(GetPublicKeysResponse)
	err := grpc.Invoke(ctx, "/spire.server.keymanager.KeyManager/GetPublicKeys", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagerClient) SignData(ctx context.Context, in *SignDataRequest, opts ...grpc.CallOption) (*SignDataResponse, error) {
	out := new(SignDataResponse)
	err := grpc.Invoke(ctx, "/spire.server.keymanager.KeyManager/SignData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagerClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/spire.server.keymanager.KeyManager/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagerClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/spire.server.keymanager.KeyManager/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyManager service

type KeyManagerServer interface {
	// Generates a new key
	GenerateKey(context.Context, *GenerateKeyRequest) (*GenerateKeyResponse, error)
	// Get a public key by key id
	GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error)
	// Gets all public keys
	GetPublicKeys(context.Context, *GetPublicKeysRequest) (*GetPublicKeysResponse, error)
	// Signs data with private key
	SignData(context.Context, *SignDataRequest) (*SignDataResponse, error)
	// Applies the plugin configuration
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	// Returns the version and related metadata of the installed plugin
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

func RegisterKeyManagerServer(s *grpc.Server, srv KeyManagerServer) {
	s.RegisterService(&_KeyManager_serviceDesc, srv)
}

func _KeyManager_GenerateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagerServer).GenerateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.keymanager.KeyManager/GenerateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagerServer).GenerateKey(ctx, req.(*GenerateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManager_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagerServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.keymanager.KeyManager/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagerServer).GetPublicKey(ctx, req.(*GetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManager_GetPublicKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagerServer).GetPublicKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.keymanager.KeyManager/GetPublicKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagerServer).GetPublicKeys(ctx, req.(*GetPublicKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManager_SignData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagerServer).SignData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.keymanager.KeyManager/SignData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagerServer).SignData(ctx, req.(*SignDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManager_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagerServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.keymanager.KeyManager/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagerServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManager_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagerServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.keymanager.KeyManager/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagerServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.keymanager.KeyManager",
	HandlerType: (*KeyManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKey",
			Handler:    _KeyManager_GenerateKey_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _KeyManager_GetPublicKey_Handler,
		},
		{
			MethodName: "GetPublicKeys",
			Handler:    _KeyManager_GetPublicKeys_Handler,
		},
		{
			MethodName: "SignData",
			Handler:    _KeyManager_SignData_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _KeyManager_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _KeyManager_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keymanager.proto",
}

func init() { proto.RegisterFile("keymanager.proto", fileDescriptor_keymanager_a61332f079830969) }

var fileDescriptor_keymanager_a61332f079830969 = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x6f, 0xd2, 0x50,
	0x14, 0x5f, 0x19, 0x43, 0x7a, 0x28, 0xac, 0xb9, 0x3a, 0x45, 0xfc, 0x13, 0xd2, 0xc4, 0x85, 0xcd,
	0x59, 0x5c, 0x37, 0xc8, 0x5e, 0x2b, 0xc3, 0x81, 0x88, 0x92, 0xa2, 0x89, 0x5b, 0x4c, 0x9a, 0x6e,
	0x5c, 0xda, 0xa6, 0xa3, 0xad, 0xbd, 0x45, 0xed, 0x83, 0x9f, 0xc0, 0x4f, 0xe5, 0xb7, 0xf2, 0xd1,
	0xb4, 0xb4, 0xb4, 0xdd, 0x64, 0x6b, 0x8c, 0x4f, 0xdc, 0x73, 0xee, 0xef, 0xcf, 0xe9, 0x3d, 0xf7,
	0x70, 0x81, 0x35, 0xb0, 0x37, 0x53, 0x4c, 0x45, 0xc5, 0x0e, 0x6f, 0x3b, 0x96, 0x6b, 0xa1, 0x07,
	0xc4, 0xd6, 0x1d, 0xcc, 0x13, 0xec, 0x7c, 0xc5, 0x0e, 0x1f, 0x6f, 0xd7, 0x8e, 0x54, 0xdd, 0xd5,
	0xe6, 0xe7, 0xfc, 0x85, 0x35, 0x6b, 0x12, 0x5b, 0x9f, 0x4e, 0x71, 0x33, 0x80, 0x36, 0x03, 0x5e,
	0xf3, 0xc2, 0x9a, 0xcd, 0x2c, 0xb3, 0x69, 0x5f, 0xce, 0x55, 0x3d, 0xfa, 0x59, 0x48, 0x72, 0x3f,
	0x80, 0x1e, 0xcd, 0xcf, 0x2f, 0xf5, 0x8b, 0x01, 0xf6, 0x50, 0x05, 0x72, 0xfa, 0xa4, 0x4a, 0xd5,
	0xa9, 0x06, 0x2d, 0xe5, 0xf4, 0x09, 0xea, 0x00, 0xad, 0x5c, 0xaa, 0x96, 0xa3, 0xbb, 0xda, 0xac,
	0x9a, 0xab, 0x53, 0x8d, 0x8a, 0xf0, 0x8c, 0x5f, 0x51, 0x03, 0x3f, 0xc0, 0x9e, 0x18, 0x81, 0xa5,
	0x98, 0x87, 0x1e, 0x01, 0x6d, 0x1b, 0xfa, 0x77, 0x79, 0xa2, 0xb8, 0x4a, 0x75, 0xbd, 0x4e, 0x35,
	0x18, 0xa9, 0xe8, 0x27, 0x8e, 0x15, 0x57, 0xe1, 0xbe, 0x01, 0x3a, 0xc1, 0x26, 0x76, 0x14, 0x17,
	0x0f, 0xb0, 0x27, 0xe1, 0x2f, 0x73, 0x4c, 0x5c, 0xb4, 0x05, 0x05, 0x03, 0x7b, 0xf2, 0xb2, 0x96,
	0x0d, 0x03, 0x7b, 0xfd, 0x09, 0x7a, 0x03, 0x65, 0x3f, 0xfd, 0x8f, 0x25, 0x31, 0x46, 0x22, 0xe2,
	0x3e, 0xc1, 0xdd, 0x94, 0x31, 0xb1, 0x2d, 0x93, 0x60, 0x24, 0x02, 0xd8, 0xc1, 0x71, 0xc8, 0x06,
	0xf6, 0x02, 0xf7, 0x92, 0xc0, 0xad, 0xd4, 0x5f, 0x9e, 0x9c, 0x44, 0xdb, 0xd1, 0x92, 0xdb, 0xf3,
	0x95, 0xdd, 0x78, 0xeb, 0xc6, 0x6f, 0xe2, 0x4e, 0xe1, 0x5e, 0x1a, 0xfd, 0xff, 0x0a, 0xb9, 0x9f,
	0x96, 0x26, 0x61, 0x25, 0xdc, 0x67, 0xd8, 0xba, 0x92, 0x0f, 0x3d, 0x3b, 0x50, 0x8a, 0x3d, 0x49,
	0x95, 0xaa, 0xaf, 0x67, 0x34, 0x85, 0xa5, 0x29, 0xe1, 0x7e, 0x52, 0xb0, 0x39, 0xd6, 0x55, 0xd3,
	0x6f, 0xef, 0x2d, 0xfd, 0x1c, 0x42, 0x45, 0x53, 0x88, 0x76, 0xad, 0xa1, 0xdb, 0x2b, 0x2d, 0x7b,
	0x0a, 0xd1, 0xe2, 0x8e, 0x96, 0xb5, 0x64, 0x88, 0x10, 0xe4, 0x13, 0x77, 0x2c, 0x58, 0x73, 0x2f,
	0x81, 0x8d, 0x8b, 0x09, 0x3f, 0xf3, 0x31, 0xd0, 0x44, 0x57, 0x4d, 0xc5, 0x9d, 0x3b, 0x38, 0x28,
	0x88, 0x91, 0xe2, 0xc4, 0xee, 0x10, 0x98, 0xe4, 0xb5, 0x41, 0x4f, 0xe0, 0xe1, 0xc7, 0x77, 0xe3,
	0x51, 0xb7, 0xd3, 0x7f, 0xdd, 0xef, 0x1e, 0xcb, 0x83, 0xee, 0xa9, 0x2c, 0xbe, 0x3d, 0x79, 0x2f,
	0xf5, 0x3f, 0xf4, 0x86, 0xec, 0x1a, 0xaa, 0x00, 0x74, 0x3b, 0xc7, 0x63, 0x51, 0x1e, 0x09, 0xad,
	0x36, 0x4b, 0x25, 0xe2, 0x83, 0xa3, 0x43, 0x36, 0xb7, 0xfb, 0x8b, 0x82, 0x72, 0xaa, 0x6a, 0xf4,
	0x14, 0x6a, 0x49, 0xc1, 0x9e, 0x38, 0xee, 0xa5, 0x14, 0x8b, 0x90, 0x1f, 0xf7, 0xc4, 0x7d, 0x76,
	0x1d, 0x01, 0x14, 0xc6, 0x3d, 0x51, 0x10, 0x0e, 0xd9, 0x7c, 0xb4, 0x6e, 0xb5, 0xd9, 0x8d, 0x70,
	0xed, 0xeb, 0x17, 0xc2, 0x75, 0x6b, 0x5f, 0x60, 0xef, 0x20, 0x06, 0x8a, 0x7e, 0x5e, 0xf6, 0x19,
	0x10, 0x47, 0xad, 0x36, 0x5b, 0x5a, 0x46, 0x3e, 0x8b, 0x59, 0x46, 0x3e, 0xaf, 0xec, 0xd7, 0xbc,
	0xd0, 0x08, 0x98, 0x95, 0x64, 0xdc, 0x6a, 0xb3, 0x9b, 0xc2, 0xef, 0x3c, 0xc0, 0x00, 0x7b, 0xc3,
	0x45, 0x13, 0x90, 0x06, 0xa5, 0xc4, 0xe8, 0xa0, 0xe7, 0x2b, 0xbb, 0x75, 0x7d, 0xb2, 0x6b, 0x7b,
	0xd9, 0xc0, 0x61, 0xa7, 0x0c, 0x60, 0x92, 0x37, 0x15, 0xdd, 0xc4, 0xbe, 0x36, 0x71, 0xb5, 0x17,
	0x19, 0xd1, 0xa1, 0x99, 0x09, 0xe5, 0xd4, 0x58, 0xa0, 0x6c, 0xfc, 0x68, 0xac, 0x6a, 0x7c, 0x56,
	0x78, 0xe8, 0x27, 0x43, 0x31, 0xba, 0x9a, 0xa8, 0xb1, 0x92, 0x7b, 0x65, 0x94, 0x6a, 0x3b, 0x19,
	0x90, 0xa1, 0xc1, 0x19, 0xd0, 0x1d, 0xcb, 0x9c, 0xea, 0xea, 0xdc, 0xc1, 0x28, 0xfa, 0x93, 0x5c,
	0x3c, 0x05, 0x7c, 0xf8, 0x06, 0x2c, 0xf7, 0x23, 0xf9, 0xed, 0xdb, 0x60, 0xa1, 0xf6, 0x74, 0x71,
	0x58, 0xc1, 0x76, 0xdf, 0x9c, 0x5a, 0x68, 0xe7, 0xaf, 0xc4, 0x14, 0x26, 0xf2, 0xd8, 0xcd, 0x02,
	0x5d, 0xf8, 0xbc, 0x62, 0xce, 0x20, 0xfe, 0xc4, 0xd1, 0xda, 0x79, 0x21, 0x78, 0xb5, 0x0e, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x39, 0xa1, 0xb3, 0x1c, 0x07, 0x00, 0x00,
}
