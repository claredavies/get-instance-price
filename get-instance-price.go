package get_instance_price

import (
	"errors"
    "strings"
	"github.com/claredavies/get-instance-price/models"
	"github.com/claredavies/get-instance-price/pkg/thirdparty/aws"
	"github.com/claredavies/get-instance-price/constants"
	"github.com/aws/aws-sdk-go/service/pricing"
)

var Version = "1.0"

var prices = []models.Price{
}

func getPriceByID(id string) (models.Price, error) {
    for _, p := range prices {
        if p.ID == id {
            return p, nil
        }
    }
    return models.Price{}, errors.New(constants.ErrMsgPriceNotFound)
}

func hasExistingInstanceTypeServiceType(serviceCode string,  instanceType string, prices []models.Price) bool {
    for _, price := range prices {
        if serviceCode == price.ServiceType &&
            strings.ToLower(instanceType) == strings.ToLower(price.InstanceType) {
            return true
        }
    }
    return false
}

func getPriceInstanceTypeServiceType(serviceCode string, instanceType string, prices []models.Price) models.Price {
    for _, price := range prices {
        if serviceCode == price.ServiceType &&
            strings.ToLower(instanceType) == strings.ToLower(price.InstanceType) {
            return price
        }
    }
    return models.Price{}
}

func getOnePriceViaQueryParams(serviceCode string, instanceType string) (models.Price, error) {
    if serviceCode == "" || instanceType == "" {
        return models.Price{}, constants.ErrQueryParameterMissing
    }

    onePrice, errRequestError := aws.FetchPricingDataFilter(constants.Region, serviceCode, constants.RegionCode, instanceType)
    if errRequestError != nil {
        return models.Price{}, constants.ErrNoMatchingResults
    }

    if len(onePrice) == 0 {
            return models.Price{}, constants.ErrNoMatchingResults
        } else if len(onePrice) > 1 {
            return models.Price{}, constants.ErrTooManyResultsReturned
        }

    return onePrice[0], errRequestError
}

func fetchJsonUnstructuredFilter(serviceCode string, instanceType string) (*pricing.GetProductsOutput, error) {
    if serviceCode == "" || instanceType == "" {
        return nil, constants.ErrQueryParameterMissing
    }

    jsonResult, err := aws.FetchPricingDataJsonFilter(constants.Region, serviceCode, constants.RegionCode, instanceType)

    if err != nil {
        return nil, err
    }

    return jsonResult, err
}

func GetPrice(serviceCode string, instanceType string) (models.Price, error) {
    if serviceCode == "" || instanceType == "" {
        return models.Price{}, constants.ErrQueryParameterMissing
    }

    alreadyExists := hasExistingInstanceTypeServiceType(serviceCode,  instanceType, prices)

    if alreadyExists == false {
        onePrice, err := getOnePriceViaQueryParams(serviceCode, instanceType)
        if err == nil {
           prices = append(prices, onePrice)
           return onePrice, nil
        } else {
            return models.Price{}, err
        }
    } else {
        price := getPriceInstanceTypeServiceType(serviceCode,  instanceType, prices)
        return price, nil
    }
}

func GetPriceStorage(serviceCode string, usageType string, regionCode string) (models.Price, error)  {
    if serviceCode == "" {
            return models.Price{}, constants.ErrQueryParameterMissing
        }

    onePrice, errRequestError := aws.FetchPricingDataStorage(constants.Region, serviceCode, regionCode, usageType)

    if errRequestError != nil {
            return models.Price{}, constants.ErrNoMatchingResults
    }

    if len(onePrice) == 0 {
            return models.Price{}, constants.ErrNoMatchingResults
        } else if len(onePrice) > 1 {
            return models.Price{}, constants.ErrTooManyResultsReturned
        }

    return onePrice[0], errRequestError
}

func GetPrices() []models.Price {
	return prices
}

func FetchJsonUnstructured(serviceCode string) (*pricing.GetProductsOutput, error)  {
    if serviceCode == "" {
        return nil, constants.ErrQueryParameterMissing
    }

    jsonResult, err := aws.FetchPricingDataJson(constants.Region, serviceCode)
    if err != nil {
        return nil, err
    }

    return jsonResult, err
}