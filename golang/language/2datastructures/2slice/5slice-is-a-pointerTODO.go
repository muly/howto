// slice is a pointer to array: TODO:
// case 1: pass slice to a function (as a value), change the value of an element in the slice, 
//        and check the index value after returning from the function
//        expected: the change made to an element in the slice persists outside the function
// case 2: pass slice to a function (as a value), 
//        change the slice by increasing the length using append, but with in the cap of the slice. 
//        check the slice content after the function returns
//        expected: slice changes will persist outside the function
// case 3: pass slice to a function (as a value), 
//        change the slice by increasing the length using append, but beyond the cap of the slice. 
//        check the slice content after the function returns
//        expected: slice changes will not perist outside the function
