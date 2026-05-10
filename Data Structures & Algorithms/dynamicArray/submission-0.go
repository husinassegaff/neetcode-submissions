type DynamicArray struct {
	data []int // buffer penyimpanan
	length int // jumlah elemen yang dipakai
	capacity int // ukuran buffer
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		data: make([]int, capacity),
		length: 0,
		capacity: capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.data[i] = n
}

func (da *DynamicArray) Pushback(n int) {

	if da.length == da.capacity {
		da.resize()
	}

	da.data[da.length] = n
	da.length++
}

func (da *DynamicArray) Popback() int {
	lastValue := da.data[da.length-1]
	da.length--
	return lastValue
}

func (da *DynamicArray) resize() {
	newData := make([]int, da.capacity * 2)
	copy(newData, da.data)
	da.data = newData
	da.capacity *= 2
}

func (da *DynamicArray) GetSize() int {
	return da.length
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
