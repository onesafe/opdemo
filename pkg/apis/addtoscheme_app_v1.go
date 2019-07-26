package apis

import (
	v1 "gitlab.4pd.io/wangyiping/opdemo/pkg/apis/app/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
}
