// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kedge/config/http/routes/adhoc.proto

/*
Package kedge_config_http_routes is a generated protocol buffer package.

It is generated from these files:
	kedge/config/http/routes/adhoc.proto
	kedge/config/http/routes/routes.proto

It has these top-level messages:
	Adhoc
	Route
*/
package kedge_config_http_routes

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Adhoc) Validate() error {
	if this.Port != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Port); err != nil {
			return go_proto_validators.FieldError("Port", err)
		}
	}
	return nil
}
func (this *Adhoc_Port) Validate() error {
	for _, item := range this.AllowedRanges {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("AllowedRanges", err)
			}
		}
	}
	return nil
}
func (this *Adhoc_Port_Range) Validate() error {
	return nil
}
