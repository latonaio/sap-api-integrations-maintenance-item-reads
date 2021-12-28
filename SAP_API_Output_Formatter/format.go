package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-item-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			MaintenanceItem:              data.MaintenanceItem,
			MaintenanceItemDescription:   data.MaintenanceItemDescription,
			MaintenanceStrategy:          data.MaintenanceStrategy,
			MaintenancePlanCategory:      data.MaintenancePlanCategory,
			MaintenancePlanCallObject:    data.MaintenancePlanCallObject,
			MaintenancePlan:              data.MaintenancePlan,
			MaintenancePlanItemPosition:  data.MaintenancePlanItemPosition,
			MaintenanceItemObjectList:    data.MaintenanceItemObjectList,
			FunctionalLocationLabelName:  data.FunctionalLocationLabelName,
			Equipment:                    data.Equipment,
			Assembly:                     data.Assembly,
			AdditionalDeviceData:         data.AdditionalDeviceData,
			TaskListType:                 data.TaskListType,
			TaskListGroup:                data.TaskListGroup,
			TaskListGroupCounter:         data.TaskListGroupCounter,
			OperationSystemCondition:     data.OperationSystemCondition,
			NumberOfTaskListExecutions:   data.NumberOfTaskListExecutions,
			MaintNotifTskIsAutomlyDtmnd:  data.MaintNotifTskIsAutomlyDtmnd,
			MaintenancePlanningPlant:     data.MaintenancePlanningPlant,
			MaintenancePlannerGroup:      data.MaintenancePlannerGroup,
			MaintenanceOrderType:         data.MaintenanceOrderType,
			NotificationType:             data.NotificationType,
			MaintenanceActivityType:      data.MaintenanceActivityType,
			MainWorkCenter:               data.MainWorkCenter,
			MainWorkCenterPlant:          data.MainWorkCenterPlant,
			MaintPriority:                data.MaintPriority,
			MaintPriorityType:            data.MaintPriorityType,
			BusinessArea:                 data.BusinessArea,
			ImmediateReleaseIsBlocked:    data.ImmediateReleaseIsBlocked,
			Material:                     data.Material,
			SerialNumber:                 data.SerialNumber,
			ServiceDocumentType:          data.ServiceDocumentType,
			ServiceContract:              data.ServiceContract,
			ServiceContractItem:          data.ServiceContractItem,
			ServiceOrderTemplate:         data.ServiceOrderTemplate,
			ServiceDocumentPriority:      data.ServiceDocumentPriority,
			Product:                      data.Product,
			MaintenancePlant:             data.MaintenancePlant,
			AssetLocation:                data.AssetLocation,
			AssetRoom:                    data.AssetRoom,
			PlantSection:                 data.PlantSection,
			WorkCenter:                   data.WorkCenter,
			ABCIndicator:                 data.ABCIndicator,
			MaintObjectFreeDefinedAttrib: data.MaintObjectFreeDefinedAttrib,
			CompanyCode:                  data.CompanyCode,
			MasterFixedAsset:             data.MasterFixedAsset,
			FixedAsset:                   data.FixedAsset,
			LocAcctAssgmtBusinessArea:    data.LocAcctAssgmtBusinessArea,
			CostCenter:                   data.CostCenter,
			ControllingArea:              data.ControllingArea,
			WBSElement:                   data.WBSElement,
			SettlementOrder:              data.SettlementOrder,
			CycleSetSequence:             data.CycleSetSequence,
			StandingOrderNumber:          data.StandingOrderNumber,
			CreationDate:                 data.CreationDate,
			LastChangeDate:               data.LastChangeDate,
			LastChangeDateTime:           data.LastChangeDateTime,
			ToCallObjects:                data.ToCallObjects.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToCallObjects(raw []byte, l *logger.Logger) ([]ToCallObjects, error) {
	pm := &responses.ToCallObjects{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCallObjects. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toCallObjects := make([]ToCallObjects, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toCallObjects = append(toCallObjects, ToCallObjects{
			MaintenancePlan:              data.MaintenancePlan,
			MaintenancePlanCallNumber:    data.MaintenancePlanCallNumber,
			MaintenanceItem:              data.MaintenanceItem,
			MaintenanceOrder:             data.MaintenanceOrder,
			MaintenanceNotification:      data.MaintenanceNotification,
			ServiceOrder:                 data.ServiceOrder,
			MaintCallHorizonIsNotReached: data.MaintCallHorizonIsNotReached,
			SchedulingStatus:             data.SchedulingStatus,
			PlannedStartDate:             data.PlannedStartDate,
		})
	}

	return toCallObjects, nil
}
