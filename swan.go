package swan

import (
	"net/url"

	"github.com/Dataman-Cloud/swan/src/types"
)

// Swan is the interface to the Swan API
type Swan interface {
	// -- APPLICATIONS ---
	// create an application in swan
	CreateApplication(version *types.Version) (*types.App, error)
	// get a list of applications from swan
	Applications(url.Values) ([]*types.App, error)
	// delete an application in swan
	DeleteApplication(appID string) error
	// get an applications from swan
	GetApplication(appID string) (*types.App, error)
	// update an application in swan
	UpdateApplication(appID string, version *types.Version) (*types.App, error)
	// proceed the rolling update
	ProceedUpdate(appID string, param *types.ProceedUpdateParam) error
	// updates slot's weight for one app
	UpdateWeights(appID string, param *types.UpdateWeightsParam) (*types.App, error)
	// cancel the rolling update
	CancelUpdate(appID string) error
	// scale up the app
	ScaleUp(appID string, param *types.ScaleUpParam) error
	// scale down the app
	ScaleDown(appID string, param *types.ScaleDownParam) error
	// get versions of the app
	GetAppVersions(appID string) ([]*types.Version, error)
	// get the app version
	GetAppVersion(appID, versionID string) (*types.Version, error)
	// get the app task
	GetAppTask(appID, taskIndex string) (*types.Task, error)

	// -- SUBSCRIPTIONS --
	AddEventsListener() (EventsChannel, error)

	// -- FRAMEWORK --
	GetFrameworkInfo() (*types.FrameworkInfo, error)

	// YAML parse
	ParseYAML(yaml string) (map[string][]string, error)

	// Compose Instance op
	RunCompose(req *types.ComposeRequest) (*types.ComposeInstance, error)
	ListComposeInstances(filter url.Values) ([]*types.ComposeInstance, error)
	GetComposeInstance(idOrName string) (*types.ComposeInstanceWrapper, error)
	RemoveComposeInstance(idOrName string) error
}
