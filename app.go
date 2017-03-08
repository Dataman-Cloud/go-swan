package swan

import (
	"net/url"

	"github.com/Dataman-Cloud/swan/src/types"
)

// CreateApplication creates a new application in Swan
// version:the structure holding the application configuration
func (r *swanClient) CreateApplication(version *types.Version) (*types.App, error) {
	result := new(types.App)
	if err := r.apiPost(APIApps, &version, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Applications retrieves an array of all the applications in swan
func (r *swanClient) Applications(v url.Values) ([]*types.App, error) {
	applications := new([]*types.App)
	err := r.apiGet(APIApps+"?"+v.Encode(), nil, applications)
	if err != nil {
		return nil, err
	}

	return *applications, nil
}

// DeleteApplication deletes an application in Swan
func (r *swanClient) DeleteApplication(appID string) error {
	if err := r.apiDelete(APIApps+"/"+appID, nil, nil); err != nil {
		return err
	}

	return nil
}

// GetApplication retrieves an application from Swan
func (r *swanClient) GetApplication(appID string) (*types.App, error) {
	result := new(types.App)
	if err := r.apiGet(APIApps+"/"+appID, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateApplication updates an application in Swan
// 		version:		the structure holding the application configuration
func (r *swanClient) UpdateApplication(appID string, version *types.Version) (*types.App, error) {
	result := new(types.App)
	if err := r.apiPut(APIApps+"/"+appID, &version, result); err != nil {
		return nil, err
	}

	return result, nil
}

// ProceedUpdate proceed the rolling update
func (r *swanClient) ProceedUpdate(appID string, param *types.ProceedUpdateParam) error {
	if err := r.apiPatch(APIApps+"/"+appID+"/proceed-update", &param, nil); err != nil {
		return err
	}

	return nil
}

// CancelUpdate canceled the rolling update
func (r *swanClient) CancelUpdate(appID string) error {
	if err := r.apiPatch(APIApps+"/"+appID+"/cancel-update", nil, nil); err != nil {
		return err
	}

	return nil
}

// ScaleUp
func (r *swanClient) ScaleUp(appID string, param *types.ScaleUpParam) error {
	if err := r.apiPatch(APIApps+"/"+appID+"/scale-up", &param, nil); err != nil {
		return err
	}

	return nil
}

// ScaleDown
func (r *swanClient) ScaleDown(appID string, param *types.ScaleDownParam) error {
	if err := r.apiPatch(APIApps+"/"+appID+"/scale-down", &param, nil); err != nil {
		return err
	}

	return nil
}

// GetAppVersions get all versions of the given application
func (r *swanClient) GetAppVersions(appID string) ([]*types.Version, error) {
	result := new([]*types.Version)
	if err := r.apiGet(APIApps+"/"+appID+"/versions", nil, result); err != nil {
		return nil, err
	}

	return *result, nil
}

// GetAppVersion get the given version of the given application
func (r *swanClient) GetAppVersion(appID, versionID string) (*types.Version, error) {
	result := new(types.Version)
	if err := r.apiGet(APIApps+"/"+appID+"/versions/"+versionID, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetAppTask get the given task of the given application
func (r *swanClient) GetAppTask(appID, taskIndex string) (*types.Task, error) {
	result := new(types.Task)
	if err := r.apiGet(APIApps+"/"+appID+"/tasks/"+taskIndex, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
