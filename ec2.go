package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ec2Client struct {
	ec2 *ec2.EC2
}

func newEC2(region string, sess *session.Session) *ec2Client {
	return &ec2Client{
		ec2: ec2.New(sess, aws.NewConfig().WithRegion(region)),
	}
}

func (e *ec2Client) describeInstance(id *string) (*ec2.Instance, error) {
	params := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			id,
		},
	}
	resp, err := e.ec2.DescribeInstances(params)
	if err != nil {
		return nil, err
	}
	return resp.Reservations[0].Instances[0], nil
}

type ec2Meta struct {
	metadata *ec2metadata.EC2Metadata
}

func NewEC2Metadata(sess *session.Session) *ec2Meta {
	return &ec2Meta{metadata: ec2metadata.New(sess)}
}

func (m *ec2Meta) region() (string, error) {
	return m.metadata.Region()
}
