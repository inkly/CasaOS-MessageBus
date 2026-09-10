package out

import (
	"time"

	"github.com/ReCasaOS/CasaOS-Common/utils"
	"github.com/ReCasaOS/CasaOS-MessageBus/codegen"
	"github.com/ReCasaOS/CasaOS-MessageBus/model"
)

func ActionAdapter(action model.Action) codegen.Action {
	return codegen.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  utils.Ptr(time.Unix(action.Timestamp, 0)),
	}
}
