package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-maintenance-item-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetMaintenanceItem(maintenancePlan, maintenanceItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Item":
			func() {
				c.Item(maintenancePlan, maintenanceItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Item(maintenancePlan, maintenanceItem string) {
	itemData, err := c.callMaintenanceItemSrvAPIRequirementItem("MaintenanceItem", maintenancePlan, maintenanceItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	callObjectsData, err := c.callToCallObjects(itemData[0].ToCallObjects)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(callObjectsData)

}

func (c *SAPAPICaller) callMaintenanceItemSrvAPIRequirementItem(api, maintenancePlan, maintenanceItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_MAINTENANCEITEM", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, maintenancePlan, maintenanceItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCallObjects(url string) ([]sap_api_output_formatter.ToCallObjects, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCallObjects(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, maintenancePlan, maintenanceItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("MaintenancePlan eq '%s' and MaintenanceItem eq '%s'", maintenancePlan, maintenanceItem))
	req.URL.RawQuery = params.Encode()
}
