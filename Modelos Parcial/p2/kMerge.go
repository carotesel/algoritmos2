// k merge de arr de arr de int. menor a mayor

type Item struct{
	num int
	pos int // pos en array chico
	ind int // pos en array grande
}

func cmp(a, b Item) int{
	return a.num - b.num
}

func kMerge(arr [][]int, k int) []int{

	heap := CrearHeap([]int, cmp)
	res := make([]int, 0, len(arr[0]) * k)

	k := len(arr)

	for i:=0; i<k; i++{
		heap.Encolar(Item{num:arr[i][0], pos:0, ind:i})
	}

	for !heap.EstaVacia(){
		min := heap.Desencolar()
		res = append(res, min.num)

		siguiente := min.pos + 1

		for siguiente < len(arr[min.ind]){
			heap.Encolar(Item{num:arr[min.ind][siguiente], pos:siguiente, ind:min.ind})
		}
	}

	return res
}