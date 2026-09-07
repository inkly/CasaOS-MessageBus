package out

import (
	"time"

	"github.com/inkly/CasaOS-Common/utils"
	"github.com/inkly/CasaOS-MessageBus/codegen"
	"github.com/inkly/CasaOS-MessageBus/model"
)

func ActionAdapter(action model.Action) codegen.Action {
	return codegen.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  utils.Ptr(time.Unix(action.Timestamp, 0)),
	}
}
