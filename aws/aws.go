package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWS struct {
	Svc  *ec2.EC2
	Tags map[string]string
}

func NewAWS(key, secret, region string, tags map[string]string) *AWS {
	svc := ec2.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Region:      aws.String(region),
	})

	return &AWS{
		Svc:  svc,
		Tags: tags,
	}
}

// ListAddresses is fetch instance ip address from ec2
func (a *AWS) ListAddresses() (address []string, err error) {
	resp, err := a.Svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	address = make([]string, 0, len(resp.Reservations))

	for i := range resp.Reservations {
		for _, inst := range resp.Reservations[i].Instances {

			tm := make(map[string]string, len(inst.Tags))
			for _, tag := range inst.Tags {
				tm[*tag.Key] = *tag.Value
			}

			// 16:running
			code := *inst.State.Code

			for k, v := range a.Tags {
				if tv, ok := tm[k]; !ok || tv != v {
					continue
				}

				if code == 16 {
					fmt.Printf("- InstanceID: %+v\n", *inst.InstanceId)
					fmt.Printf("- PrivateDNSName: %+v\n", *inst.PrivateDnsName)
					fmt.Printf("- PrivateIPAddress: %+v\n", *inst.PrivateIpAddress)
					fmt.Printf("- Tags: %+v\n", tm)
					fmt.Println("")

					address = append(address, *inst.PrivateIpAddress)
				}
			}
		}
	}

	return
}
