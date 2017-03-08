package swan

import (
	"github.com/Dataman-Cloud/swan/src/types"
)

// GetFrameworkInfo get the framework info
func (r *swanClient) GetFrameworkInfo() (*types.FrameworkInfo, error) {
	result := new(types.FrameworkInfo)
	if err := r.apiGet(APIVersion+"/framework/info", nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
