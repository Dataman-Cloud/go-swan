package swan

import (
	"net/url"

	"github.com/Dataman-Cloud/swan/types"
)

type CreateResponse struct {
	ID string `json:"Id"`
}

// CreateApplication creates a new application in Swan
// version:the structure holding the application configuration
func (r *swanClient) CreateApplication(version *types.Version) (*CreateResponse, error) {
	result := new(CreateResponse)
	if err := r.apiPost(APIApps, &version, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Applications retrieves an array of all the applications in swan
func (r *swanClient) Applications(v url.Values) ([]*types.Application, error) {
	applications := new([]*types.Application)
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
func (r *swanClient) GetApplication(appID string) (*types.Application, error) {
	result := new(types.Application)
	if err := r.apiGet(APIApps+"/"+appID, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateVersion
func (r *swanClient) CreateApplicationVersion(appID string, version *types.Version) (*CreateResponse, error) {
	result := new(CreateResponse)
	if err := r.apiPost(APIApps+"/"+appID+"/versions", &version, result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateApplication updates an application in Swan
func (r *swanClient) UpdateApplication(appID string, update *types.UpdateBody) error {
	if err := r.apiPut(APIApps+"/"+appID, &update, nil); err != nil {
		return err
	}

	return nil
}

// Rollback app version to previous version
func (r *swanClient) RollbackApplication(appID string) error {
	if err := r.apiPost(APIApps+"/"+appID, nil, nil); err != nil {
		return err
	}

	return nil
}

// UpdateWeights updates slots' weight for one app
func (r *swanClient) UpdateWeights(appID string, param *types.UpdateWeightsBody) error {
	if err := r.apiPatch(APIApps+"/"+appID+"/weights", &param, nil); err != nil {
		return err
	}

	return nil
}

// ScaleApplication
func (r *swanClient) ScaleApplication(appID string, param *types.ScaleBody) error {
	if err := r.apiPatch(APIApps+"/"+appID, &param, nil); err != nil {
		return err
	}

	return nil
}

// CreateAppVersion create app version
func (r *swanClient) CreateAppVersion(appID string, version *types.Version) (*types.Version, error) {
	result := new(types.Version)
	if err := r.apiPost(APIApps+"/"+appID+"/versions", &version, result); err != nil {
		return nil, err
	}
	return result, nil
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

// GetAppTasks get the tasks of the given application
func (r *swanClient) GetAppTasks(appID string) ([]*types.Task, error) {
	tasks := new([]*types.Task)
	if err := r.apiGet(APIApps+"/"+appID+"/tasks", nil, tasks); err != nil {
		return nil, err
	}
	return *tasks, nil
}

// GetAppTask get the given task of the given application
func (r *swanClient) GetAppTask(appID, taskID string) (*types.Task, error) {
	result := new(types.Task)
	if err := r.apiGet(APIApps+"/"+appID+"/tasks/"+taskID, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
