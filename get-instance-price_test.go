package get_instance_price

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPriceEC2(t *testing.T) {
	price, err := GetPrice("AmazonEC2", "m5.xlarge")
	assert.NoError(t, err)
	assert.NotEmpty(t, price)
}

func TestGetPriceEMR(t *testing.T) {
	price, err := GetPrice("ElasticMapReduce", "m5.xlarge")
	assert.NoError(t, err)
	assert.NotEmpty(t, price)
}

func TestGetPriceStorageEC2(t *testing.T) {
	price, err := GetPriceStorage("AmazonEC2", "EU-EBS:VolumeUsage.st1", "eu-west-1")
	assert.NoError(t, err)
	assert.NotEmpty(t, price)
}