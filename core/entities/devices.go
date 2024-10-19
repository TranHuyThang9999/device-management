package entities

type Devices struct {
	DeviceName  string   `json:"device_name"` // Tên thiết bị
	Quantity    int      `json:"quantity"`    //
	Description string   `json:"description"` // Mô tả thiết bị, có thể là null
	Files       []string `json:"files"`       //
}

type DeviceReqUpdate struct {
	ID          int64    `json:"id"`          // ID của thiết bị
	DeviceName  string   `json:"device_name"` // Tên thiết bị
	Quantity    int      `json:"quantity"`    //
	Description string   `json:"description"` // Mô tả thiết bị, có thể là null
	Files       []string `json:"files"`       //
}
type DevicesGetForUser struct {
	Id          int64    `json:"id"`
	DeviceName  string   `json:"device_name"`
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
	Url         []string `json:"url,omitempty"`
}
