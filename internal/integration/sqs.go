package integration

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// NewQueue creates a new queue with pre-defined configuration.
//
// Pre-defined Fields Configuration
//  If the Queue should be FIFO:
//    - FIFO: true
//    - ContendBasedDeduplication: true
//    - DeduplicationScope: MessageGroup
func NewQueue(scope constructs.Construct, id string, isFifo bool, timeout *float64) awssqs.Queue {
	var props = &awssqs.QueueProps{}

	if isFifo {
		id += ".fifo"
		props.Fifo = ENABLED
		props.ContentBasedDeduplication = ENABLED
		props.DeduplicationScope = awssqs.DeduplicationScope_MESSAGE_GROUP
	}

	if timeout != nil {
		props.VisibilityTimeout = awscdk.Duration_Millis(timeout)
	}

	props.QueueName = jsii.String(id)
	props.RemovalPolicy = awscdk.RemovalPolicy_DESTROY

	return awssqs.NewQueue(scope, jsii.String(id), props)
}

// NewQueueWithDLQ creates a new queue with a deadleatter queue configured and with pre-defined configuration.
//
// Pre-defined Fields Configuration
//  If the Queue should be FIFO:
//    - FIFO: true
//    - ContendBasedDeduplication: true
//    - DeduplicationScope: MessageGroup
//  DeadLetterQueue:
//    - MaxReceiveCount: 5
func NewQueueWithDLQ(scope constructs.Construct, id string, isFifo bool, timeout *float64) awssqs.Queue {
	var (
		props = &awssqs.QueueProps{}
		dlqId = fmt.Sprintf("%s_dlq", id)
		dlq   = NewQueue(scope, dlqId, true, nil)
	)

	if isFifo {
		id += ".fifo"
		props.Fifo = ENABLED
		props.ContentBasedDeduplication = ENABLED
		props.DeduplicationScope = awssqs.DeduplicationScope_MESSAGE_GROUP
	}

	props.QueueName = jsii.String(id)
	props.DeadLetterQueue = NewDeadLetterQueue(dlq)
	props.RemovalPolicy = awscdk.RemovalPolicy_DESTROY
	props.VisibilityTimeout = awscdk.Duration_Millis(timeout)

	return awssqs.NewQueue(scope, jsii.String(id), props)
}

// NewDeadLetterQueue returns a pre-defined deadletter configuration.
//
// Pre-definend Field Configuration
//    - MaxReceiveCount: 5
func NewDeadLetterQueue(queue awssqs.IQueue) *awssqs.DeadLetterQueue {
	return &awssqs.DeadLetterQueue{
		Queue:           queue,
		MaxReceiveCount: jsii.Number(5),
	}
}
