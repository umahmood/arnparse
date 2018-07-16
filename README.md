# ARN Parse

A Go Library which parses Amazon Resource Names (ARNs) into its individual 
components. So you can get useful information from the ARN such as partition, 
region, account id etc.

You can read more about ARNs via the documentation here [https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).

# Installation

> $ go get github.com/umahmood/arnparse

# Dependencies 

> testify | github.com/stretchr/testify/assert

# Go Version

> Tested with go version go1.10.3 darwin/amd64

# Usage
```
package main

import (
    "fmt"

    "github.com/umahmood/arnparse"
)

func main() {
    arn, err := arnparse.Parse("arn:aws:ec2:us-east-1:123456789012:vpc/vpc-fd580e98")
    if err != nil {
        // handle error
    }
    fmt.Println("Partition:", arn.Partition)
    fmt.Println("Service:" arn.Service)
    fmt.Println("Region:", arn.Region)
    fmt.Println("Account ID:", arn.AccountID)
    fmt.Println("Resource Type:", arn.ResourceType)
    fmt.Println("Resource:", arn.Resource)
}
```

Output:
```
Partition: aws
Service: ec2
Region: us-east-1
Account ID: 123456789012
Resource Type: vpc
Resource: vpc-fd580e98
```

# Documentation

> http://godoc.org/github.com/umahmood/arnparse

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
