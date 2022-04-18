package mergesort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {

	testeSet := make([][]float32, 5)
	respSet := make([][]float32, 5)

	testeSet[0] = []float32{1, 7, 7, 9, 1, 8, 5, 0, 6, 0, 4}
	testeSet[1] = []float32{2, 4, 6, 0, 3, 5, 1}
	testeSet[2] = []float32{0.6046603, 0.9405091, 0.6645601, 0.4377142, 0.4246375, 0.68682307, 0.06563702, 0.15651925, 0.09696952, 0.30091187}
	testeSet[3] = []float32{0.12027272, 0.752847, 0.2279211, 0.014731542, 0.787094, 0.93680984, 0.59252065}
	testeSet[4] = []float32{0.41817072, 0.4446015, 0.33643034, 0.41354266, 0.22456239, 0.6501239, 0.59987456, 0.63542, 0.50003946, 0.119192526, 0.015621684, 0.038535707, 0.5887652, 0.30397257, 0.17745581, 0.4290025, 0.26011735, 0.42666996, 0.37603715, 0.6254644, 0.7470781, 0.095232844, 0.4197972, 0.4077865, 0.60282886, 0.7704953, 0.7643929, 0.17847776, 0.44926757, 0.50883955, 0.011586609, 0.18797536, 0.536431, 0.33205485, 0.6408248, 0.88653934, 0.99659747, 0.09452772, 0.6150095, 0.1014327, 0.3315563, 0.44181722, 0.18244326, 0.43098542, 0.9935248, 0.88696545, 0.8577459, 0.6975804, 0.28643784, 0.12147658, 0.76250285, 0.6664795, 0.7173761, 0.021818025, 0.9947064, 0.56416076, 0.12507954, 0.92262363, 0.25338045, 0.532607, 0.79232997, 0.32577395, 0.21792452, 0.6181618, 0.60865927, 0.86965406, 0.9199082, 0.7408588, 0.06631382, 0.57186973, 0.68083876, 0.9726245, 0.19538122, 0.3139331, 0.09786606, 0.02327837, 0.41597924, 0.5989567, 0.7881809, 0.387795, 0.32863715, 0.2912236, 0.8836703, 0.07242945, 0.74758387, 0.90391475, 0.9210973, 0.82449603, 0.40255097, 0.97081935, 0.20453806, 0.4421, 0.5819133, 0.93604946, 0.06707782, 0.5209518, 0.12930687, 0.43783313, 0.3981815, 0.74486893}

	respSet[0] = []float32{0, 0, 1, 1, 4, 5, 6, 7, 7, 8, 9}
	respSet[1] = []float32{0, 1, 2, 3, 4, 5, 6}
	respSet[2] = []float32{0.06563702, 0.09696952, 0.15651925, 0.30091187, 0.4246375, 0.4377142, 0.6046603, 0.6645601, 0.68682307, 0.9405091}
	respSet[3] = []float32{0.014731542, 0.12027272, 0.2279211, 0.59252065, 0.752847, 0.787094, 0.93680984}
	respSet[4] = []float32{0.011586609, 0.015621684, 0.021818025, 0.02327837, 0.038535707, 0.06631382, 0.06707782, 0.07242945, 0.09452772, 0.095232844, 0.09786606, 0.1014327, 0.119192526, 0.12147658, 0.12507954, 0.12930687, 0.17745581, 0.17847776, 0.18244326, 0.18797536, 0.19538122, 0.20453806, 0.21792452, 0.22456239, 0.25338045, 0.26011735, 0.28643784, 0.2912236, 0.30397257, 0.3139331, 0.32577395, 0.32863715, 0.3315563, 0.33205485, 0.33643034, 0.37603715, 0.387795, 0.3981815, 0.40255097, 0.4077865, 0.41354266, 0.41597924, 0.41817072, 0.4197972, 0.42666996, 0.4290025, 0.43098542, 0.43783313, 0.44181722, 0.4421, 0.4446015, 0.44926757, 0.50003946, 0.50883955, 0.5209518, 0.532607, 0.536431, 0.56416076, 0.57186973, 0.5819133, 0.5887652, 0.5989567, 0.59987456, 0.60282886, 0.60865927, 0.6150095, 0.6181618, 0.6254644, 0.63542, 0.6408248, 0.6501239, 0.6664795, 0.68083876, 0.6975804, 0.7173761, 0.7408588, 0.74486893, 0.7470781, 0.74758387, 0.76250285, 0.7643929, 0.7704953, 0.7881809, 0.79232997, 0.82449603, 0.8577459, 0.86965406, 0.8836703, 0.88653934, 0.88696545, 0.90391475, 0.9199082, 0.9210973, 0.92262363, 0.93604946, 0.97081935, 0.9726245, 0.9935248, 0.9947064, 0.99659747}

	for i, item := range testeSet {
		temp := item
		MergeSort(item)
		if !isSliceEqual(item, respSet[i]) {
			t.Error("For the set: ", temp, " was expect: ", respSet[0], " got: ", item)
		}
	}
}

func isSliceEqual(array1, array2 []float32) bool {
	if len(array1) == len(array2) {

		for i, item := range array1 {
			if item != array2[i] {
				return false
			}
		}
		return true

	} else {
		return false
	}
}
