package devices

type ChargingStationLocation struct {
	Lat float64 `json:"lat,omitempty"`

	Lng float64 `json:"lng,omitempty"`
}

type ChargingStationModem struct {
	// OCPP - Optional. This contains the ICCID of the modem’s SIMcard.
	Iccid string `json:"iccid,omitempty"`
	// OCPP - Optional. This contains the IMSI of the modem’s SIM card.
	Imsi string `json:"imsi,omitempty"`
}

type ChargingStation struct {
	Id string `json:"id"`
	// OCPP - Optional. Vendor-specific device identifier.
	SerialNumber string `json:"serialNumber,omitempty"`
	// OCPP - Required. Defines the model of the device.
	Model string `json:"model"`
	// OCPP - Required. Identifies the vendor (not necessarily in a unique manner).
	VendorName string `json:"vendorName"`
	// OCPP - Optional. This contains the firmware version of the Charging Station.
	FirmwareVersion string `json:"firmwareVersion,omitempty"`

	Modem *ChargingStationModem `json:"modem,omitempty"`

	Location *ChargingStationLocation `json:"location,omitempty"`
	// Date and time of last BootNotification received. As defined by date-time - RFC3339
	LastBoot string `json:"lastBoot,omitempty"`
}
