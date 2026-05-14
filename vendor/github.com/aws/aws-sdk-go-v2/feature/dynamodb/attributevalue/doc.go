// Package attributevalue provides marshaling and unmarshaling utilities to
// convert between Go types and Amazon DynamoDB AttributeValues.
//
// These utilities allow you to marshal slices, maps, structs, and scalar
// values to and from AttributeValue type. These utilities make it
// easier to convert between AttributeValue and Go types when working with
// DynamoDB resources.
//
// This package only converts between Go types and DynamoDB AttributeValue. See
// the feature/dynamodbstreams/attributevalue package for converting to
// DynamoDBStreams AttributeValue types.
//
// # Converting AttributeValue between DynamoDB and DynamoDBStreams
//
// The FromDynamoStreamsDBMap, FromDynamoStreamsDBList, and FromDynamoDBStreams
// functions provide the conversion utilities to convert a DynamoDBStreams
// AttributeValue type to a DynamoDB AttributeValue type. Use these utilities
// when you need to convert the AttributeValue type between the two APIs.
//
// # AttributeValue Marshaling
//
// To marshal a Go type to an AttributeValue you can use the Marshal,
// MarshalList, and MarshalMap functions. The List and Map functions are
// specialized versions of the Marshal for serializing slices and maps of
// Attributevalues.
//
// The following example uses MarshalMap to convert a Go struct, Record to a
// AttributeValue. The AttributeValue value is then used as input to the
// PutItem operation call.
//
//	type Record struct {
//	    ID     string
//	    URLs   []string
//	}
//
//	//...
//
//	r := Record{
//	    ID:   "ABC123",
//	    URLs: []string{
//	        "https://example.com/first/link",
//	        "https://example.com/second/url",
//	    },
//	}
//	av, err := attributevalue.MarshalMap(r)
//	if err != nil {
//	    return fmt.Errorf("failed to marshal Record, %w", err)
//	}
//
//	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
//	    TableName: aws.String(myTableName),
//	    Item:      av,
//	})
//	if err != nil {
//	    return fmt.Errorf("failed to put Record, %w", err)
//	}
//
// # AttributeValue Unmarshaling
//
// To unmarshal an AttributeValue to a Go type you can use the Unmarshal,
// UnmarshalList, UnmarshalMap, and UnmarshalListOfMaps functions. The List and
// Map functions are specialized versions of the Unmarshal function for
// unmarshal slices and maps of Attributevalues.
//
// The following example will unmarshal Items result from the DynamoDB's
// Scan API operation. The Items returned will be unmarshaled into the slice of
// the Records struct.
//
//	type Record struct {
//	    ID     string
//	    URLs   []string
//	}
//
//	//...
//
//	result, err := client.Scan(context.Context(), &dynamodb.ScanInput{
//	    TableName: aws.String(myTableName),
//	})
//	if err != nil {
//	    return fmt.Errorf("failed to scan  table, %w", err)
//	}
//
//	var records []Record
//	err := attributevalue.UnmarshalListOfMaps(results.Items, &records)
//	if err != nil {
//	     return fmt.Errorf("failed to unmarshal Items, %w", err)
//	}
//
// # Struct tags
//
// The AttributeValue Marshal and Unmarshal functions support the `dynamodbav`
// struct tag by default. Additional tags can be enabled with the
// EncoderOptions and DecoderOptions, TagKey option.
//
// See the Marshal and Unmarshal function for information on how struct tags
// and fields are marshaled and unmarshaled.
package attributevalue
