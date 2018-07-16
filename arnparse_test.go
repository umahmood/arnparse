package arnparse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html

var testCases = []struct {
	testName string
	test     string
	err      error
	arn      *Arn
}{
	{
		testName: "Resource Type With Slash",
		test:     "arn:aws:ec2:us-east-1:123456789012:vpc/vpc-fd580e98",
		arn: &Arn{
			Partition:    "aws",
			Service:      "ec2",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "vpc",
			Resource:     "vpc-fd580e98",
		},
	},
	{
		testName: "Resource Type With Colon",
		test:     "arn:aws:codecommit:us-east-1:123456789012:MyDemoRepo",
		arn: &Arn{
			Partition:    "aws",
			Service:      "codecommit",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "",
			Resource:     "MyDemoRepo",
		},
	},
	{
		testName: "Resource Type With Multiple Colons",
		test:     "arn:aws:logs:us-east-1:123456789012:log-group:my-log-group*:log-stream:my-log-stream*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "logs",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "log-group",
			Resource:     "my-log-group*:log-stream:my-log-stream*",
		},
	},
	{
		testName: "No Resource Type",
		test:     "arn:aws:cloudwatch:us-east-1:123456789012:alarm:MyAlarmName",
		arn: &Arn{
			Partition:    "aws",
			Service:      "cloudwatch",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "alarm",
			Resource:     "MyAlarmName",
		},
	},
	{
		testName: "Resource With Single Slash",
		test:     "arn:aws:kinesisvideo:us-east-1:123456789012:stream/example-stream-name/0123456789012",
		arn: &Arn{
			Partition:    "aws",
			Service:      "kinesisvideo",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "stream",
			Resource:     "example-stream-name/0123456789012",
		},
	},
	{
		testName: "Resource With Multiple Slashes",
		test:     "arn:aws:macie:us-east-1:123456789012:trigger/example61b3df36bff1dafaf1aa304b0ef1a975/alert/example8780e9ca227f98dae37665c3fd22b585",
		arn: &Arn{
			Partition:    "aws",
			Service:      "macie",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "trigger",
			Resource:     "example61b3df36bff1dafaf1aa304b0ef1a975/alert/example8780e9ca227f98dae37665c3fd22b585",
		},
	},
	{
		testName: "No Region No Account ID",
		test:     "arn:aws:s3:::my_corporate_bucket",
		arn: &Arn{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "my_corporate_bucket",
		},
	},
	{
		testName: "Spaces",
		test:     "arn:aws:artifact:::report-package/Certifications and Attestations/SOC/*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "artifact",
			Region:       "",
			AccountID:    "",
			ResourceType: "report-package",
			Resource:     "Certifications and Attestations/SOC/*",
		},
	},
	{
		testName: "Wildcard",
		test:     "arn:aws:ec2:us-east-1:123456789012:instance/*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "ec2",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "instance",
			Resource:     "*",
		},
	},
	{
		testName: "Double Wildcard",
		test:     "arn:aws:events:us-east-1:*:*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "events",
			Region:       "us-east-1",
			AccountID:    "*",
			ResourceType: "",
			Resource:     "*",
		},
	},
	{
		testName: "No Prefix",
		test:     "something:aws:s3:::my_corporate_bucket",
		err:      ErrMalformedArn,
	},
	{
		testName: "Empty String",
		test:     "",
		err:      ErrMalformedArn,
	},
	{
		testName: "API Gateway 1",
		test:     "arn:aws:apigateway:us-east-1::a123456789012bc3de45678901f23a45:/test/mydemoresource/*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "apigateway",
			Region:       "us-east-1",
			AccountID:    "",
			ResourceType: "",
			Resource:     "a123456789012bc3de45678901f23a45:/test/mydemoresource/*",
		},
	},
	{
		testName: "API Gateway 2",
		test:     "arn:aws:execute-api:us-east-1:123456789012:8kjmp19d1h/*/*/*/*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "execute-api",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "",
			Resource:     "8kjmp19d1h/*/*/*/*",
		},
	},
	{
		testName: "SNS 1",
		test:     "arn:aws:sns:*:123456789012:my_corporate_topic",
		arn: &Arn{
			Partition:    "aws",
			Service:      "sns",
			Region:       "*",
			AccountID:    "123456789012",
			ResourceType: "",
			Resource:     "my_corporate_topic",
		},
	},
	{
		testName: "SNS 2",
		test:     "arn:aws:sns:us-east-1:123456789012:my_corporate_topic:02034b43-fefa-4e07-a5eb-3be56f8c54ce",
		arn: &Arn{
			Partition:    "aws",
			Service:      "sns",
			Region:       "us-east-1",
			AccountID:    "123456789012",
			ResourceType: "",
			Resource:     "my_corporate_topic:02034b43-fefa-4e07-a5eb-3be56f8c54ce",
		},
	},
	{
		testName: "S3 1",
		test:     "arn:aws:s3:::my_corporate_bucket/exampleobject.png",
		arn: &Arn{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "my_corporate_bucket/exampleobject.png",
		},
	},
	{
		testName: "S3 2",
		test:     "arn:aws:s3:::my_corporate_bucket/*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "my_corporate_bucket/*",
		},
	},
	{
		testName: "S3 3",
		test:     "arn:aws:s3:::my_corporate_bucket/Development/*",
		arn: &Arn{
			Partition:    "aws",
			Service:      "s3",
			Region:       "",
			AccountID:    "",
			ResourceType: "",
			Resource:     "my_corporate_bucket/Development/*",
		},
	},
}

func TestArns(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("test %s", tc.testName), func(t *testing.T) {
			arn, err := Parse(tc.test)
			assert.Equal(t, err, tc.err, "incorrect err field")
			if arn != nil {
				assert.Equal(t, arn.Partition, tc.arn.Partition, "incorrect partition field")
				assert.Equal(t, arn.Service, tc.arn.Service, "incorrect service field")
				assert.Equal(t, arn.Region, tc.arn.Region, "incorrect region field")
				assert.Equal(t, arn.AccountID, tc.arn.AccountID, "incorrect account id field")
				assert.Equal(t, arn.ResourceType, tc.arn.ResourceType, "incorrect resource type field")
				assert.Equal(t, arn.Resource, tc.arn.Resource, "incorrect resource field")
			}
		})
	}
}
