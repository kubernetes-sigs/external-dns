package attributevalue

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithyjson "github.com/aws/smithy-go/encoding/json"
)

func awsAwsjson10DeserializeDocumentAttributeValue(v *types.AttributeValue, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var uv types.AttributeValue
loop:
	for key, value := range shape {
		if value == nil {
			continue
		}
		switch key {
		case "B":
			var mv []byte
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected BinaryAttributeValue to be []byte, got %T instead", value)
				}
				dv, err := base64.StdEncoding.DecodeString(jtv)
				if err != nil {
					return fmt.Errorf("failed to base64 decode BinaryAttributeValue, %w", err)
				}
				mv = dv
			}
			uv = &types.AttributeValueMemberB{Value: mv}
			break loop

		case "BOOL":
			var mv bool
			if value != nil {
				jtv, ok := value.(bool)
				if !ok {
					return fmt.Errorf("expected BooleanAttributeValue to be of type *bool, got %T instead", value)
				}
				mv = jtv
			}
			uv = &types.AttributeValueMemberBOOL{Value: mv}
			break loop

		case "BS":
			var mv [][]byte
			if err := awsAwsjson10DeserializeDocumentBinarySetAttributeValue(&mv, value); err != nil {
				return err
			}
			uv = &types.AttributeValueMemberBS{Value: mv}
			break loop

		case "L":
			var mv []types.AttributeValue
			if err := awsAwsjson10DeserializeDocumentListAttributeValue(&mv, value); err != nil {
				return err
			}
			uv = &types.AttributeValueMemberL{Value: mv}
			break loop

		case "M":
			var mv map[string]types.AttributeValue
			if err := awsAwsjson10DeserializeDocumentMapAttributeValue(&mv, value); err != nil {
				return err
			}
			uv = &types.AttributeValueMemberM{Value: mv}
			break loop

		case "N":
			var mv string
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected NumberAttributeValue to be of type string, got %T instead", value)
				}
				mv = jtv
			}
			uv = &types.AttributeValueMemberN{Value: mv}
			break loop

		case "NS":
			var mv []string
			if err := awsAwsjson10DeserializeDocumentNumberSetAttributeValue(&mv, value); err != nil {
				return err
			}
			uv = &types.AttributeValueMemberNS{Value: mv}
			break loop

		case "NULL":
			var mv bool
			if value != nil {
				jtv, ok := value.(bool)
				if !ok {
					return fmt.Errorf("expected NullAttributeValue to be of type *bool, got %T instead", value)
				}
				mv = jtv
			}
			uv = &types.AttributeValueMemberNULL{Value: mv}
			break loop

		case "S":
			var mv string
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected StringAttributeValue to be of type string, got %T instead", value)
				}
				mv = jtv
			}
			uv = &types.AttributeValueMemberS{Value: mv}
			break loop

		case "SS":
			var mv []string
			if err := awsAwsjson10DeserializeDocumentStringSetAttributeValue(&mv, value); err != nil {
				return err
			}
			uv = &types.AttributeValueMemberSS{Value: mv}
			break loop

		default:
			uv = &types.UnknownUnionMember{Tag: key}
			break loop

		}
	}
	*v = uv
	return nil
}

func awsAwsjson10DeserializeDocumentBinarySetAttributeValue(v *[][]byte, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv [][]byte
	if *v == nil {
		cv = [][]byte{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col []byte
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected BinaryAttributeValue to be []byte, got %T instead", value)
			}
			dv, err := base64.StdEncoding.DecodeString(jtv)
			if err != nil {
				return fmt.Errorf("failed to base64 decode BinaryAttributeValue, %w", err)
			}
			col = dv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10DeserializeDocumentListAttributeValue(v *[]types.AttributeValue, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []types.AttributeValue
	if *v == nil {
		cv = []types.AttributeValue{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col types.AttributeValue
		if err := awsAwsjson10DeserializeDocumentAttributeValue(&col, value); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10DeserializeDocumentMapAttributeValue(v *map[string]types.AttributeValue, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var mv map[string]types.AttributeValue
	if *v == nil {
		mv = map[string]types.AttributeValue{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal types.AttributeValue
		if err := awsAwsjson10DeserializeDocumentAttributeValue(&parsedVal, value); err != nil {
			return err
		}
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsAwsjson10DeserializeDocumentNumberSetAttributeValue(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []string
	if *v == nil {
		cv = []string{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected NumberAttributeValue to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10DeserializeDocumentStringSetAttributeValue(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []string
	if *v == nil {
		cv = []string{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected StringAttributeValue to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10SerializeDocumentAttributeValue(v types.AttributeValue, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.AttributeValueMemberB:
		av := object.Key("B")
		av.Base64EncodeBytes(uv.Value)

	case *types.AttributeValueMemberBOOL:
		av := object.Key("BOOL")
		av.Boolean(uv.Value)

	case *types.AttributeValueMemberBS:
		av := object.Key("BS")
		if err := awsAwsjson10SerializeDocumentBinarySetAttributeValue(uv.Value, av); err != nil {
			return err
		}

	case *types.AttributeValueMemberL:
		av := object.Key("L")
		if err := awsAwsjson10SerializeDocumentListAttributeValue(uv.Value, av); err != nil {
			return err
		}

	case *types.AttributeValueMemberM:
		av := object.Key("M")
		if err := awsAwsjson10SerializeDocumentMapAttributeValue(uv.Value, av); err != nil {
			return err
		}

	case *types.AttributeValueMemberN:
		av := object.Key("N")
		av.String(uv.Value)

	case *types.AttributeValueMemberNS:
		av := object.Key("NS")
		if err := awsAwsjson10SerializeDocumentNumberSetAttributeValue(uv.Value, av); err != nil {
			return err
		}

	case *types.AttributeValueMemberNULL:
		av := object.Key("NULL")
		av.Boolean(uv.Value)

	case *types.AttributeValueMemberS:
		av := object.Key("S")
		av.String(uv.Value)

	case *types.AttributeValueMemberSS:
		av := object.Key("SS")
		if err := awsAwsjson10SerializeDocumentStringSetAttributeValue(uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsAwsjson10SerializeDocumentBinarySetAttributeValue(v [][]byte, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		av.Base64EncodeBytes(v[i])
	}
	return nil
}

func awsAwsjson10SerializeDocumentListAttributeValue(v []types.AttributeValue, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsAwsjson10SerializeDocumentAttributeValue(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10SerializeDocumentMapAttributeValue(v map[string]types.AttributeValue, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			continue
		}
		if err := awsAwsjson10SerializeDocumentAttributeValue(v[key], om); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson10SerializeDocumentNumberSetAttributeValue(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson10SerializeDocumentStringSetAttributeValue(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

// UnmarshalListJSON decodes a JSON-encoded byte slice into an array of DynamoDB
// AttributeValues using the AWS JSON 1.0 deserializer.
func UnmarshalListJSON(data []byte) ([]types.AttributeValue, error) {
	var raw []interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	var res []types.AttributeValue

	err := awsAwsjson10DeserializeDocumentListAttributeValue(&res, raw)

	return res, err
}

// UnmarshalMapJSON decodes a JSON-encoded byte slice into a map of DynamoDB
// AttributeValues using the AWS JSON 1.0 deserializer.
func UnmarshalMapJSON(data []byte) (map[string]types.AttributeValue, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	var res map[string]types.AttributeValue

	err := awsAwsjson10DeserializeDocumentMapAttributeValue(&res, raw)

	return res, err
}

// UnmarshalJSON decodes a JSON-encoded byte slice into a single DynamoDB
// AttributeValue using the AWS JSON 1.0 deserializer.
func UnmarshalJSON(data []byte) (types.AttributeValue, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	var parsedVal types.AttributeValue
	mapVar := parsedVal
	if err := awsAwsjson10DeserializeDocumentAttributeValue(&mapVar, raw); err != nil {
		return nil, err
	}
	parsedVal = mapVar

	return parsedVal, nil
}

// MarshalListJSON encodes a slice of DynamoDB AttributeValues
// into a JSON-formatted byte slice using the AWS JSON 1.0 serializer.
func MarshalListJSON(in []types.AttributeValue) ([]byte, error) {
	jsonEncoder := smithyjson.NewEncoder()

	if err := awsAwsjson10SerializeDocumentListAttributeValue(in, jsonEncoder.Value); err != nil {
		return nil, err
	}

	return jsonEncoder.Bytes(), nil
}

// MarshalMapJSON encodes a map of DynamoDB AttributeValues
// into a JSON-formatted byte slice using the AWS JSON 1.0 serializer.
func MarshalMapJSON(in map[string]types.AttributeValue) ([]byte, error) {
	jsonEncoder := smithyjson.NewEncoder()

	if err := awsAwsjson10SerializeDocumentMapAttributeValue(in, jsonEncoder.Value); err != nil {
		return nil, err
	}

	return jsonEncoder.Bytes(), nil
}

// MarshalJSON encodes a single DynamoDB AttributeValues
// into a JSON-formatted byte slice using the AWS JSON 1.0 serializer.
func MarshalJSON(in types.AttributeValue) ([]byte, error) {
	jsonEncoder := smithyjson.NewEncoder()

	if err := awsAwsjson10SerializeDocumentAttributeValue(in, jsonEncoder.Value); err != nil {
		return nil, err
	}

	return jsonEncoder.Bytes(), nil
}
