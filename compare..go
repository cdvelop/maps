package maps

import (
	"fmt"
	"reflect"
)

// Función para comparar dos slices de mapas sin importar el orden
func AreSliceMapsIdentical(maps_a, maps_b []map[string]string) bool {
	// Verificar si los slices tienen la misma longitud
	if len(maps_a) != len(maps_b) {
		return false
	}

	// Verificar cada mapa individual en ambos slices
	for _, map_a := range maps_a {
		found := false
		for _, map_b := range maps_b {
			if reflect.DeepEqual(map_a, map_b) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func CompareMaps(a, b map[string]string) (result string) {

	if reflect.DeepEqual(a, b) {
		return
	}

	keysNotFound, differentValues := mapsDifferent(a, b)
	// keysNotFound, differentValues = mapsDifferent(b, a)

	if len(keysNotFound) > 0 {
		// result += fmt.Sprintf("%v", keysNotFound)
		result += strings{}.Join(keysNotFound, ", ")
	}
	if len(differentValues) > 0 {
		result += strings{}.Join(differentValues, ", ")
	}

	return result
}

// main_name ej: A
func mapsDifferent(a, b map[string]string) (keysNotFound, differentValues []string) {

	for keyA, valueA := range a {
		if valueB, exists := b[keyA]; !exists {
			keysNotFound = append(keysNotFound, mapKeyNotFound(keyA, "B"))
		} else if valueA != valueB {
			differentValues = append(differentValues, mapDifValue(keyA, valueA, valueB))
		}
	}

	// Check for keys in b that are not in a
	for keyB := range b {
		if _, exists := a[keyB]; !exists {
			keysNotFound = append(keysNotFound, mapKeyNotFound(keyB, "A"))
		}
	}
	return
}

func mapKeyNotFound(key, targetMap string) string {
	return fmt.Sprintf("-field: [%v] not found in map [%v]", key, targetMap)
}

func mapDifValue(key, valueA, valueB string) string {
	return fmt.Sprintf("-field: [%v] value a:[%v] != value b:[%v]", key, valueA, valueB)
}
