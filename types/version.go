package types

import (
	"bytes"
	"github.com/radkomih/gosemble/scale"
)

type ApiItem struct {
	Name    [8]byte
	Version uint32
}

type VersionData struct {
	SpecName           []byte
	ImplName           []byte
	AuthoringVersion   uint32
	SpecVersion        uint32
	ImplVersion        uint32
	Apis               []ApiItem
	TransactionVersion uint32
	StateVersion       uint32
}

func (v *VersionData) Encode() ([]byte, error) {
	var buffer = bytes.Buffer{}
	var encoder = scale.Encoder{Writer: &buffer}

	encoder.EncodeByteSlice(v.SpecName)
	encoder.EncodeByteSlice(v.ImplName)
	encoder.EncodeUint32(v.AuthoringVersion)
	encoder.EncodeUint32(v.SpecVersion)
	encoder.EncodeUint32(v.ImplVersion)
	encoder.EncodeInt32(int32(len(v.Apis)))
	for _, apiItem := range v.Apis {
		encoder.EncodeByteSlice(apiItem.Name[:])
		encoder.EncodeUint32(apiItem.Version)
	}
	encoder.EncodeUint32(v.TransactionVersion)
	encoder.EncodeUint32(v.StateVersion)

	return buffer.Bytes(), nil
}

func (v *VersionData) Decode(enc []byte) error {
	//var data VersionData
	//
	//_, err := scale.Decode(enc, &data)
	//if err != nil {
	//	return err
	//}
	//
	//v.SpecName = data.SpecName
	//v.ImplName = data.ImplName
	//v.AuthoringVersion = data.AuthoringVersion
	//v.SpecVersion = data.SpecVersion
	//v.ImplVersion = data.ImplVersion
	//v.Apis = data.Apis
	//v.TransactionVersion = data.TransactionVersion
	//v.StateVersion = data.StateVersion

	return nil
}
