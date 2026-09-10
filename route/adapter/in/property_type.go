package in

import (
	"github.com/ReCasaOS/CasaOS-MessageBus/codegen"
	"github.com/ReCasaOS/CasaOS-MessageBus/model"
)

func PropertyTypeAdapter(propertyType codegen.PropertyType) model.PropertyType {
	return model.PropertyType{
		Name: propertyType.Name,
	}
}
