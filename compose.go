package swan

import (
	"net/url"

	"github.com/Dataman-Cloud/swan/src/types"
)

func (r *swanClient) ParseYAML(yaml string) (map[string][]string, error) {
	ret := make(map[string][]string)
	req := map[string]string{"yaml": yaml}
	err := r.apiPost(APICompose+"/parse", req, &ret)
	return ret, err
}

func (r *swanClient) RunCompose(req *types.ComposeRequest) (*types.ComposeInstance, error) {
	ret := new(types.ComposeInstance)
	err := r.apiPost(APICompose, req, &ret)
	return ret, err
}

func (r *swanClient) ListComposeInstances(filter url.Values) ([]*types.ComposeInstance, error) {
	ret := make([]*types.ComposeInstance, 0, 0)
	err := r.apiGet(APICompose+"?"+filter.Encode(), nil, &ret)
	return ret, err
}

func (r *swanClient) GetComposeInstance(idOrName string) (*types.ComposeInstanceWrapper, error) {
	ret := new(types.ComposeInstanceWrapper)
	err := r.apiGet(APICompose+"/"+idOrName, nil, &ret)
	return ret, err
}

func (r *swanClient) RemoveComposeInstance(idOrName string) error {
	return r.apiDelete(APICompose+"/"+idOrName, nil, nil)
}
