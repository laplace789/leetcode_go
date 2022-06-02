package hashtable

import "math/rand"

type RandomizedSet struct {
	set map[int]int //map[key]index
	arr []int       //arr of val
}

func Constructor() RandomizedSet {
	randomizeSet := new(RandomizedSet)
	randomizeSet.set = make(map[int]int, 0)
	randomizeSet.arr = make([]int, 0)
	return *randomizeSet
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.set[val]
	if !exists {
		this.set[val] = len(this.arr)
		this.arr = append(this.arr, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.set[val]
	if !exists {
		return false
	} else {
		//get last insert val
		lastVal := this.arr[len(this.arr)-1]
		//modify lastVal index
		this.set[lastVal] = idx
		delete(this.set, val)
		//remove element
		this.arr[idx] = lastVal
		this.arr = this.arr[:len(this.arr)-1]
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	//don't use `reflect.getMapKeys`
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
