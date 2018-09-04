package schema

import "time"

// Lun defines the schema of ai lun
type Server struct {
	UUID            string           `json:"uuid"`
	Name            string           `json:"name"`
	Pool            string           `json:"pool"`
	IOGroup         string           `json:"iogroup"`
	Created         time.Time        `json:"created"`
	Capacity        uint64           `json:"capacity"`
	CapacityUnit    string           `json:"capacityunit"`
	CapacitySavings string           `json:"capacitysavings"`
	Locked          bool             `json:"locked"`
	storagesystem   string           `json:"storagesystem"`
        Protection      LunProtection    `json:"protection"`
}

// LunProtection defines the schema of a server's resource protection.
type LunProtection struct {
	Delete  bool `json:"delete"`
	Rebuild bool `json:"rebuild"`
}

// LunGetResponse defines the schema of the response when
// retrieving a single lun.
type LunGetResponse struct {
	Server Server `json:"lun"`
}

// ServerListResponse defines the schema of the response when
// listing servers.
type LunListResponse struct {
	Luns []Lun `json:"luns"`
}

// LunCreateRequest defines the schema for the request to
// create a lun.
type LunCreateRequest struct {
	Name             string      `json:"name"`
	Pool             string      `json:"pool"`
	IOGroup          string      `json:"iogroup"`
	Capacity         uint64      `json:"capacity"`
	CapacityUnit     string      `json:"capacityunit"`
	CapacitySavings  string      `json:"capacitysavings"`
}

// LunCreateResponse defines the schema of the response when
// creating a lun.
type ServerCreateResponse struct {
	Lun          lun     `json:"lun"`
	Action       Action  `json:"action"`
}

// LunUpdateRequest defines the schema of the request to update a lun.
type LunUpdateRequest struct {
	Name string `json:"name,omitempty"`
}

// LunUpdateResponse defines the schema of the response when updating a lun 
type LunUpdateResponse struct {
	Lun Lun `json:"lun"`
}
