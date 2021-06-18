package common

import st "github.com/srjnm/demoinfocs-golang/pkg/demoinfocs/sendtables"

func getInt(entity st.Entity, propName string) int {
	if entity == nil {
		return 0
	}

	return entity.PropertyValueMust(propName).IntVal
}

func getFloat(entity st.Entity, propName string) float32 {
	if entity == nil {
		return 0
	}

	return entity.PropertyValueMust(propName).FloatVal
}

func getString(entity st.Entity, propName string) string {
	if entity == nil {
		return ""
	}

	return entity.PropertyValueMust(propName).StringVal
}

func getBool(entity st.Entity, propName string) bool {
	if entity == nil {
		return false
	}

	return entity.PropertyValueMust(propName).BoolVal()
}
