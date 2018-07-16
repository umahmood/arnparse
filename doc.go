/*Package arnparse parses Amazon Resource Names (ARNs) into its individual
components.

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
*/
package arnparse
