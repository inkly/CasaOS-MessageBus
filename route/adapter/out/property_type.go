package out

import (
	"github.com/inkly/CasaOS-MessageBus/codegen"
	"github.com/inkly/CasaOS-MessageBus/model"
)

func PropertyTypeAdapter(propertyType model.PropertyType) codegen.PropertyType {
	return codegen.PropertyType{
		Name: propertyType.Name,
	}
}
