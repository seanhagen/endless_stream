package game

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func stringSliceEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}

// luaMapToTable converts a Go map to a lua table
func luaMapToTable(m map[string]interface{}) *lua.LTable {
	// Main table pointer
	resultTable := &lua.LTable{}

	// Loop map
	for key, element := range m {

		switch element.(type) {
		case float64:
			resultTable.RawSetString(key, lua.LNumber(element.(float64)))
		case int64:
			resultTable.RawSetString(key, lua.LNumber(element.(int64)))
		case string:
			resultTable.RawSetString(key, lua.LString(element.(string)))
		case bool:
			resultTable.RawSetString(key, lua.LBool(element.(bool)))
		case []byte:
			resultTable.RawSetString(key, lua.LString(string(element.([]byte))))
		case map[string]interface{}:

			// Get table from map
			tble := luaMapToTable(element.(map[string]interface{}))

			resultTable.RawSetString(key, tble)

		case time.Time:
			resultTable.RawSetString(key, lua.LNumber(element.(time.Time).Unix()))

		case []map[string]interface{}:

			// Create slice table
			sliceTable := &lua.LTable{}

			// Loop element
			for _, s := range element.([]map[string]interface{}) {

				// Get table from map
				tble := luaMapToTable(s)

				sliceTable.Append(tble)
			}

			// Set slice table
			resultTable.RawSetString(key, sliceTable)

		case []interface{}:

			// Create slice table
			sliceTable := &lua.LTable{}

			// Loop interface slice
			for _, s := range element.([]interface{}) {

				// Switch interface type
				switch s.(type) {
				case map[string]interface{}:

					// Convert map to table
					t := luaMapToTable(s.(map[string]interface{}))

					// Append result
					sliceTable.Append(t)

				case float64:

					// Append result as number
					sliceTable.Append(lua.LNumber(s.(float64)))

				case string:

					// Append result as string
					sliceTable.Append(lua.LString(s.(string)))

				case bool:

					// Append result as bool
					sliceTable.Append(lua.LBool(s.(bool)))
				}
			}

			// Append to main table
			resultTable.RawSetString(key, sliceTable)
		}
	}

	return resultTable
}
