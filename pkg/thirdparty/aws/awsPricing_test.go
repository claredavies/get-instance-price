package aws

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFetchPricingDataJsonEMR(t *testing.T) {
	_, err := FetchPricingDataJson("us-east-1", "ElasticMapReduce")
	assert.NoError(t, err)
}

func TestFetchPricingDataJsonEC2(t *testing.T) {
	_, err := FetchPricingDataJson("us-east-1", "AmazonEC2")
	assert.NoError(t, err)
}

func TestFetchPricingDataFilterEMR(t *testing.T) {
	prices, err := FetchPricingDataFilter("us-east-1", "ElasticMapReduce", "eu-west-1", "C6g.12xlarge")
	assert.NoError(t, err)
	assert.NotEmpty(t, prices)
}

func TestFetchPricingDataFilterEC2(t *testing.T) {
	prices, err := FetchPricingDataFilter("us-east-1", "AmazonEC2", "eu-west-1", "C6g.12xlarge")
	assert.NoError(t, err)
	assert.NotEmpty(t, prices)
}

func TestFetchPricingDataStorageEC2(t *testing.T) {
	prices, err := FetchPricingDataStorage("us-east-1", "AmazonEC2", "eu-west-1", "EU-EBS:VolumeUsage.st1")
	assert.NoError(t, err)
	assert.NotEmpty(t, prices)
}

func TestFetchPricingDataStorageEMR(t *testing.T) {
	prices, err := FetchPricingDataStorage("us-east-1", "ElasticMapReduce", "eu-west-1", "EU-EBS:VolumeUsage.st1")
	assert.NoError(t, err)
	assert.Empty(t, prices)
}

