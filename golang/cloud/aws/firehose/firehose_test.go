package main



import(
	"testing"
	"context"
	
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/aws/request"
)


type mockFirehoseClient struct{
	firehoseiface.FirehoseAPI
}


func (m mockFirehoseClient) PutRecordWithContext(aws.Context, *firehose.PutRecordInput, ...request.Option) (*firehose.PutRecordOutput, error){


	// store to memory
	return &firehose.PutRecordOutput{}, nil
}


func TestProcessRecord(t *testing.T){
	mockSvc := &mockFirehoseClient{}

	if err := ProcessRecord(context.TODO(), mockSvc); err != nil{
		t.Errorf("ProcessRecord failed with error: %v", err)
	}

}