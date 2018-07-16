package arnparse

import (
	"errors"
	"strings"
)

// ErrMalformedArn arn string is malformed
var ErrMalformedArn = errors.New("malformed amazon resource name")

// Arn represents an parsed arn string instance
type Arn struct {
	Partition    string
	Service      string
	Region       string
	AccountID    string
	ResourceType string
	Resource     string
}

// Parse ARN
func Parse(arn string) (*Arn, error) {
	if !strings.Contains(arn, "arn:") {
		return nil, ErrMalformedArn
	}
	var (
		arnObj = &Arn{}
		elems  = strings.SplitN(arn, ":", 6)
	)

	arnObj.Resource = elems[5]
	arnObj.Service = elems[2]

	in := func(target string, services []string) bool {
		for _, srv := range services {
			if target == srv {
				return true
			}
		}
		return false
	}

	if !in(arnObj.Service, []string{"s3", "sns", "apigateway", "execute-api"}) {
		if strings.Contains(arnObj.Resource, "/") {
			r := strings.SplitN(arnObj.Resource, "/", 2)
			arnObj.ResourceType = r[0]
			arnObj.Resource = r[1]
		} else if strings.Contains(arnObj.Resource, ":") {
			r := strings.SplitN(arnObj.Resource, ":", 2)
			arnObj.ResourceType = r[0]
			arnObj.Resource = r[1]
		}
	}

	arnObj.Partition = elems[1]
	arnObj.Region = elems[3]
	arnObj.AccountID = elems[4]
	return arnObj, nil
}
