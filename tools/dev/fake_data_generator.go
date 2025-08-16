package dev

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-faker/faker/v4"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type FakeDataGeneratorParams = mcp.CallToolParamsFor[struct {
	DataType string `json:"dataType" jsonschema:"Type of fake data to generate; enum: [email, name, firstname, lastname, sentence, paragraph, word, uuid, uuid_digit, ipv4, ipv6, macaddress, url, domainname, username, password, phonenumber, tollfreephonenumber, e164phonenumber, latitude, longitude, unix_time, date, timestring, monthname, yearstring, dayofweek, dayofmonth, timestamp, century, timezone, timeperiod, cctype, ccnumber, currency, amountwithcurrency]"`
	Count    int    `json:"count,omitempty" jsonschema:"Optional: Number of fake data items to generate (default: 1)"`
}]

func FakeDataGenerator(
	ctx context.Context,
	session *mcp.ServerSession,
	params *FakeDataGeneratorParams,
) (*mcp.CallToolResultFor[any], error) {

	dataType := strings.ToLower(params.Arguments.DataType)
	count := params.Arguments.Count
	if count == 0 {
		count = 1
	}

	var generatedData []string

	for i := 0; i < count; i++ {
		switch dataType {
		case "email":
			generatedData = append(generatedData, faker.Email())
		case "name":
			generatedData = append(generatedData, faker.Name())
		case "firstname":
			generatedData = append(generatedData, faker.FirstName())
		case "lastname":
			generatedData = append(generatedData, faker.LastName())
		case "sentence":
			generatedData = append(generatedData, faker.Sentence())
		case "paragraph":
			generatedData = append(generatedData, faker.Paragraph())
		case "word":
			generatedData = append(generatedData, faker.Word())
		case "uuid":
			generatedData = append(generatedData, faker.UUIDHyphenated())
		case "uuid_digit":
			generatedData = append(generatedData, faker.UUIDDigit())
		case "ipv4":
			generatedData = append(generatedData, faker.IPv4())
		case "ipv6":
			generatedData = append(generatedData, faker.IPv6())
		case "macaddress":
			generatedData = append(generatedData, faker.MacAddress())
		case "url":
			generatedData = append(generatedData, faker.URL())
		case "domainname":
			generatedData = append(generatedData, faker.DomainName())
		case "username":
			generatedData = append(generatedData, faker.Username())
		case "password":
			generatedData = append(generatedData, faker.Password())
		case "phonenumber":
			generatedData = append(generatedData, faker.Phonenumber())
		case "tollfreephonenumber":
			generatedData = append(generatedData, faker.TollFreePhoneNumber())
		case "e164phonenumber":
			generatedData = append(generatedData, faker.E164PhoneNumber())
		case "latitude":
			generatedData = append(generatedData, fmt.Sprintf("%f", faker.Latitude()))
		case "longitude":
			generatedData = append(generatedData, fmt.Sprintf("%f", faker.Longitude()))
		case "unix_time":
			generatedData = append(generatedData, fmt.Sprintf("%d", faker.UnixTime()))
		case "date":
			generatedData = append(generatedData, faker.Date())
		case "timestring":
			generatedData = append(generatedData, faker.TimeString())
		case "monthname":
			generatedData = append(generatedData, faker.MonthName())
		case "yearstring":
			generatedData = append(generatedData, faker.YearString())
		case "dayofweek":
			generatedData = append(generatedData, faker.DayOfWeek())
		case "dayofmonth":
			generatedData = append(generatedData, faker.DayOfMonth())
		case "timestamp":
			generatedData = append(generatedData, faker.Timestamp())
		case "century":
			generatedData = append(generatedData, faker.Century())
		case "timezone":
			generatedData = append(generatedData, faker.Timezone())
		case "timeperiod":
			generatedData = append(generatedData, faker.Timeperiod())
		case "cctype":
			generatedData = append(generatedData, faker.CCType())
		case "ccnumber":
			generatedData = append(generatedData, faker.CCNumber())
		case "currency":
			generatedData = append(generatedData, faker.Currency())
		case "amountwithcurrency":
			generatedData = append(generatedData, faker.AmountWithCurrency())
		default:
			return nil, fmt.Errorf("unsupported data type: %s", dataType)
		}
	}

	return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{&mcp.TextContent{Text: strings.Join(generatedData, "\n")}},
		},
		nil
}
