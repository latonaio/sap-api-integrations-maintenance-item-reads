# sap-api-integrations-maintenance-item-reads 
sap-api-integrations-maintenance-item-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で保全計画明細データ を取得するマイクロサービスです。    
sap-api-integrations-maintenance-item-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-maintenance-item-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_MAINTENANCEITEM_0001/overview  

## 動作環境  
sap-api-integrations-maintenance-item-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-maintenance-item-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。 

## 本レポジトリ が 対応する API サービス
sap-api-integrations-maintenance-item-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_MAINTENANCEITEM_0001/overview   
* APIサービス名(=baseURL): API_MAINTENANCEITEM  

## 本レポジトリ に 含まれる API名
sap-api-integrations-maintenance-item-reads には、次の API をコールするためのリソースが含まれています。  

* MaintenanceItem（保全計画 - 明細）※保全計画の詳細データを取得するために、ToCallObject、と合わせて利用されます。
* ToCallObject（保全計画 - 保全コール対象）

## API への 値入力条件 の 初期値
sap-api-integrations-maintenance-item-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.MaintenanceItem.MaintenancePlan（保全計画）
* inoutSDC.MaintenanceItem.MaintenanceItem（保全計画明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Item" が指定されています。    
  
```
	"api_schema": "MaintenanceItem",
	"accepter": ["Item"],
	"maintenance_plan": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "MaintenanceItem",
	"accepter": ["All"],
	"maintenance_plan": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetMaintenanceItem(maintenanceItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Item":
			func() {
				c.Item(maintenanceItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 保全計画　の　明細データ が取得された結果の JSON の例です。  
以下の項目のうち、"MaintenanceItem" ～ "to_MaintPlanCallObjects" は、/SAP_API_Output_Formatter/type.go 内 の Type Product {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-maintenance-item-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-maintenance-item-reads/SAP_API_Caller.(*SAPAPICaller).Item",
	"level": "INFO",
	"message": [
		{
			"MaintenanceItem": "1",
			"MaintenanceItemDescription": "Mechanical Inspection for Pump",
			"MaintenanceStrategy": "EM_01",
			"MaintenancePlanCategory": "PM",
			"MaintenancePlanCallObject": "",
			"MaintenancePlan": "1",
			"MaintenancePlanItemPosition": "1",
			"MaintenanceItemObjectList": "0",
			"FunctionalLocationLabelName": "1710-SPA-SAC-PLAR1-INLT-SCTV",
			"Equipment": "217100002",
			"Assembly": "",
			"AdditionalDeviceData": "",
			"TaskListType": "A",
			"TaskListGroup": "1",
			"TaskListGroupCounter": "1",
			"OperationSystemCondition": "",
			"NumberOfTaskListExecutions": "1",
			"MaintNotifTskIsAutomlyDtmnd": "",
			"MaintenancePlanningPlant": "1710",
			"MaintenancePlannerGroup": "930",
			"MaintenanceOrderType": "YBA2",
			"NotificationType": "",
			"MaintenanceActivityType": "YB4",
			"MainWorkCenter": "RES-0300",
			"MainWorkCenterPlant": "1710",
			"MaintPriority": "",
			"MaintPriorityType": "PM",
			"BusinessArea": "",
			"ImmediateReleaseIsBlocked": false,
			"Material": "",
			"SerialNumber": "",
			"ServiceDocumentType": "",
			"ServiceContract": "",
			"ServiceContractItem": "0",
			"ServiceOrderTemplate": "",
			"ServiceDocumentPriority": "0",
			"Product": "",
			"MaintenancePlant": "1710",
			"AssetLocation": "YB_1701",
			"AssetRoom": "",
			"PlantSection": "YOH",
			"WorkCenter": "",
			"ABCIndicator": "",
			"MaintObjectFreeDefinedAttrib": "",
			"CompanyCode": "1710",
			"MasterFixedAsset": "",
			"FixedAsset": "",
			"LocAcctAssgmtBusinessArea": "",
			"CostCenter": "17101301",
			"ControllingArea": "A000",
			"WBSElement": "",
			"SettlementOrder": "",
			"CycleSetSequence": "0",
			"StandingOrderNumber": "",
			"CreationDate": "/Date(1498435200000)/",
			"LastChangeDate": "/Date(1498435200000)/",
			"LastChangeDateTime": "",
			"to_MaintPlanCallObjects": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_MAINTENANCEITEM/MaintenanceItem('1')/to_MaintPlanCallObjects"
		}
	],
	"time": "2021-12-28T19:54:36.748885+09:00"
}
```
