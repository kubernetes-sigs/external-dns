package attributevalue

import (
	"fmt"

	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	streams "github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
)

// FromDynamoDBStreamsMap converts a map of Amazon DynamoDB Streams
// AttributeValues, and all nested members.
func FromDynamoDBStreamsMap(from map[string]streams.AttributeValue) (to map[string]ddb.AttributeValue, err error) {
	to = make(map[string]ddb.AttributeValue, len(from))
	for field, value := range from {
		to[field], err = FromDynamoDBStreams(value)
		if err != nil {
			return nil, err
		}
	}

	return to, nil
}

// FromDynamoDBStreamsList converts a slice of Amazon DynamoDB Streams
// AttributeValues, and all nested members.
func FromDynamoDBStreamsList(from []streams.AttributeValue) (to []ddb.AttributeValue, err error) {
	to = make([]ddb.AttributeValue, len(from))
	for i := 0; i < len(from); i++ {
		to[i], err = FromDynamoDBStreams(from[i])
		if err != nil {
			return nil, err
		}
	}

	return to, nil
}

// FromDynamoDBStreams converts an Amazon DynamoDB Streams AttributeValue, and
// all nested members.
func FromDynamoDBStreams(from streams.AttributeValue) (ddb.AttributeValue, error) {
	switch tv := from.(type) {
	case *streams.AttributeValueMemberNULL:
		return &ddb.AttributeValueMemberNULL{Value: tv.Value}, nil

	case *streams.AttributeValueMemberBOOL:
		return &ddb.AttributeValueMemberBOOL{Value: tv.Value}, nil

	case *streams.AttributeValueMemberB:
		return &ddb.AttributeValueMemberB{Value: tv.Value}, nil

	case *streams.AttributeValueMemberBS:
		bs := make([][]byte, len(tv.Value))
		for i := 0; i < len(tv.Value); i++ {
			bs[i] = append([]byte{}, tv.Value[i]...)
		}
		return &ddb.AttributeValueMemberBS{Value: bs}, nil

	case *streams.AttributeValueMemberN:
		return &ddb.AttributeValueMemberN{Value: tv.Value}, nil

	case *streams.AttributeValueMemberNS:
		return &ddb.AttributeValueMemberNS{Value: append([]string{}, tv.Value...)}, nil

	case *streams.AttributeValueMemberS:
		return &ddb.AttributeValueMemberS{Value: tv.Value}, nil

	case *streams.AttributeValueMemberSS:
		return &ddb.AttributeValueMemberSS{Value: append([]string{}, tv.Value...)}, nil

	case *streams.AttributeValueMemberL:
		values, err := FromDynamoDBStreamsList(tv.Value)
		if err != nil {
			return nil, err
		}
		return &ddb.AttributeValueMemberL{Value: values}, nil

	case *streams.AttributeValueMemberM:
		values, err := FromDynamoDBStreamsMap(tv.Value)
		if err != nil {
			return nil, err
		}
		return &ddb.AttributeValueMemberM{Value: values}, nil

	default:
		return nil, fmt.Errorf("unknown AttributeValue union member, %T", from)
	}
}
