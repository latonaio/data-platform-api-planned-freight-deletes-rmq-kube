package requests

type Header struct {
	PlannedFreight							int		`json:PlannedFreight`
	PlannedFreightType						*string	`json:PlannedFreightType`
	FreightAgreement						int		`json:FreightAgreement`
	FreightAgreementItem					int		`json:FreightAgreementItem`
	FreightAgreementItemAvailableFreight	int		`json:FreightAgreementItemAvailableFreight`
	FreightType								string	`json:FreightType`
	FreightSpec								string	`json:FreightSpec`
	FreightCalendar							string	`json:FreightCalendar`
	PlannedFreightDepartureDate				string	`json:PlannedFreightDepartureDate`
	PlannedFreightDepartureTime				string	`json:PlannedFreightDepartureTime`
	PlannedFreightArrivalDate				string	`json:PlannedFreightArrivalDate`
	PlannedFreightArrivalTime				string	`json:PlannedFreightArrivalTime`
	LogisticsPartner						int		`json:LogisticsPartner`
	DeliverToParty							int		`json:DeliverToParty`
	DeliverToPlant							string	`json:DeliverToPlant`
	DeliverFromParty						int		`json:DeliverFromParty`
	DeliverFromPlant						string	`json:DeliverFromPlant`
	FRPArea									*string	`json:FRPArea`
	FRPController							*string	`json:FRPController`
	FreightCapacityWeight					*float32`json:FreightCapacityWeight`
	FreightCapacityWeightUnit				*string	`json:FreightCapacityWeightUnit`
	CreationDate							string	`json:CreationDate`
	CreationTime							string	`json:CreationTime`
	LastChangeDate							string	`json:LastChangeDate`
	LastChangeTime							string	`json:LastChangeTime`
	PlannedFreightLongText					*string	`json:PlannedFreightLongText`
	PlannedFreightIsFixed					*int    `json:PlannedFreightIsFixed`
	PlannedFreightIsReleased				*int	`json:PlannedFreightIsReleased`
	IsMarkedForDeletion						*int	`json:IsMarkedForDeletion`
}
