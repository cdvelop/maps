package maps

import "reflect"

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
