package starlinghub

type ListDevices struct {
	Status  string   `json:"status"`
	Devices []Device `json:"devices"`
}

type ListDevice struct {
	Status     string `json:"status"`
	Properties Device `json:"properties"`
}

type Device struct {
	Type          string `json:"type"`
	ID            string `json:"id"`
	Where         string `json:"where"`
	Name          string `json:"name"`
	SerialNumber  string `json:"serialNumber"`
	StructureName string `json:"structureName"`
	FixtureType   string `json:"fixtureType,omitempty"`
	BatteryStatus string `json:"batteryStatus,omitempty"`
	ButtonPushed  bool   `json:"buttonPushed,omitempty"`
	ContactState  string `json:"contactState,omitempty"`
}
