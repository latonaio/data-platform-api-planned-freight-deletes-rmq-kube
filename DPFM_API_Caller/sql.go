package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-planned-freight-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-planned-freight-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) Header(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE header.PlannedFreight = %d ", input.Header.PlannedFreight),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	header.PlannedFreight
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_freight_header_data as header 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

