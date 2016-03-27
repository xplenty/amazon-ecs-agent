// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Amazon EC2 Container Service (Amazon ECS) is a highly scalable, fast, container
// management service that makes it easy to run, stop, and manage Docker containers
// on a cluster of Amazon EC2 instances. Amazon ECS lets you launch and stop
// container-enabled applications with simple API calls, allows you to get the
// state of your cluster from a centralized service, and gives you access to
// many familiar Amazon EC2 features like security groups, Amazon EBS volumes,
// and IAM roles.
//
// You can use Amazon ECS to schedule the placement of containers across your
// cluster based on your resource needs, isolation policies, and availability
// requirements. Amazon EC2 Container Service eliminates the need for you to
// operate your own cluster management and configuration management systems
// or worry about scaling your management infrastructure.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ECS struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "ecs"

// New creates a new instance of the ECS client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ECS client from just a session.
//     svc := ecs.New(mySession)
//
//     // Create a ECS client with additional configuration
//     svc := ecs.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ECS {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ECS {
	svc := &ECS{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-11-13",
				JSONVersion:   "1.1",
				TargetPrefix:  "AmazonEC2ContainerServiceV20141113",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(jsonrpc.Build)
	svc.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ECS operation and runs any
// custom request initialization.
func (c *ECS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
