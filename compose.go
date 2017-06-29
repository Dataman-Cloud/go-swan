package swan

import (
	"net/url"

	"github.com/Dataman-Cloud/swan/types"
)

func (r *swanClient) ParseYAML(yaml string) (map[string][]string, error) {
	ret := make(map[string][]string)
	req := map[string]string{"yaml": yaml}
	err := r.apiPost(APICompose+"/parse", req, &ret)
	return ret, err
}

func (r *swanClient) RunCompose(req *types.Compose) (string, error) {
	var result string
	err := r.apiPost(APICompose, req, result)
	return result, err
}

func (r *swanClient) ListCompose(filter url.Values) ([]*types.Compose, error) {
	ret := make([]*types.Compose, 0, 0)
	err := r.apiGet(APICompose+"?"+filter.Encode(), nil, &ret)
	return ret, err
}

func (r *swanClient) GetCompose(idOrName string) (*types.Compose, error) {
	ret := new(types.Compose)
	err := r.apiGet(APICompose+"/"+idOrName, nil, &ret)
	return ret, err
}

func (r *swanClient) RemoveCompose(idOrName string) error {
	return r.apiDelete(APICompose+"/"+idOrName, nil, nil)
}
