package in

import (
	"github.com/inkly/CasaOS-MessageBus/codegen"
	"github.com/inkly/CasaOS-MessageBus/model"
)

func ActionTypeAdapter(actionType codegen.ActionType) model.ActionType {
	propertyTypeList := make([]model.PropertyType, 0)
	for _, propertyType := range actionType.PropertyTypeList {
		propertyTypeList = append(propertyTypeList, PropertyTypeAdapter(propertyType))
	}

	return model.ActionType{
		SourceID:         actionType.SourceID,
		Name:             actionType.Name,
		PropertyTypeList: propertyTypeList,
	}
}
