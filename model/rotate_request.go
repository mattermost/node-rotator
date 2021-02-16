package model

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// RotateClusterRequest specifies the parameters for a new cluster rotation.
type RotateClusterRequest struct {
	ClusterID            string `json:"clusterID,omitempty"`
	MaxScaling           int64  `json:"maxScaling,omitempty"`
	RotateMasters        bool   `json:"rotateMasters,omitempty"`
	RotateWorkers        bool   `json:"rotateWorkers,omitempty"`
	MaxDrainRetries      int64  `json:"maxDrainRetries,omitempty"`
	EvictGracePeriod     int64  `json:"evictGracePeriod,omitempty"`
	WaitBetweenRotations int64  `json:"waitBetweenRotations,omitempty"`
}

// NewRotateClusterRequestFromReader decodes the request and returns after validation and setting the defaults.
func NewRotateClusterRequestFromReader(reader io.Reader) (*RotateClusterRequest, error) {
	var rotateClusterRequest RotateClusterRequest
	err := json.NewDecoder(reader).Decode(&rotateClusterRequest)
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "failed to decode rotate cluster request")
	}

	err = rotateClusterRequest.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "rotate cluster request failed validation")
	}
	rotateClusterRequest.SetDefaults()

	return &rotateClusterRequest, nil
}

// Validate validates the values of a cluster rotate request.
func (request *RotateClusterRequest) Validate() error {
	if request.ClusterID == "" {
		return errors.Errorf("Cluster ID cannot be empty")
	}

	return nil
}

// SetDefaults sets the default values for a cluster provision request.
func (request *RotateClusterRequest) SetDefaults() {

}