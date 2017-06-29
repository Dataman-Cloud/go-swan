package swan

import (
	"net/url"

	"github.com/Dataman-Cloud/swan/types"
)

// Swan is the interface to the Swan API
type Swan interface {
	// -- APPLICATIONS ---
	// create an application in swan
	CreateApplication(version *types.Version) (*CreateResponse, error)
	// get a list of applications from swan
	Applications(url.Values) ([]*types.Application, error)
	// delete an application in swan
	DeleteApplication(appID string) error
	// get an applications from swan
	GetApplication(appID string) (*types.Application, error)
	// CreateApplicationVersion
	CreateApplicationVersion(appID string, version *types.Version) (*CreateResponse, error)
	// update an application in swan
	UpdateApplication(appID string, update *types.UpdateBody) error
	// rollback application to previous version
	RollbackApplication(appID string) error

	// updates slot's weight for one app
	UpdateWeights(appID string, param *types.UpdateWeightsBody) error

	// scale the app
	ScaleApplication(appID string, param *types.ScaleBody) error

	// create app version
	CreateAppVersion(appID string, version *types.Version) (*types.Version, error)
	// get app versions
	GetAppVersions(appID string) ([]*types.Version, error)
	// get app version
	GetAppVersion(appID, versionID string) (*types.Version, error)

	// get the app tasks
	GetAppTasks(appID string) ([]*types.Task, error)
	// get the app task
	GetAppTask(appID, taskID string) (*types.Task, error)

	// -- SUBSCRIPTIONS --
	AddEventsListener() (EventsChannel, error)

	// YAML parse
	ParseYAML(yaml string) (map[string][]string, error)

	// Compose Instance op
	RunCompose(req *types.Compose) (string, error)
	ListCompose(filter url.Values) ([]*types.Compose, error)
	GetCompose(idOrName string) (*types.Compose, error)
	RemoveCompose(idOrName string) error
}
