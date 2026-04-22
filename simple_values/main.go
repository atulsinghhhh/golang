package main

// import "fmt"

/*

	integer
	string
	bool
	float do arthmetic ops

	variables in go

	constant in go
	const name string="golang"

	constant grouping
	const (

		port=5000,host='localhost
	)
		fmt.Print(port,host)

	loops in go


	switch statement

	type switch

	whoAmI := func(i interface{}){
		switch i:=i.(type){ // return type \

			case int:
				fmt.Println()

		} 
	}


	Arrays
			fixed size predictable
			constant table access
			memory optimazation


	// number sequence of specific length

	func main(){
		var nums[4]int

		fmt.Println(len(nums))
		nums[0] =1
		fmt.Println(nums[0])

		fmt.Println(nums)  //print
	}
	
	nums := [3]int{1,3,2}
	// 2d array

	nums :=[2][2]int ={{2,3},{3,4}}

	// slices dynamically array
	most used construct in go useful methods in array
	
	uninitialized slice is nil
	var nums[]int

	slice not nil at start
	var nums = make([]int,3,)

	capacity maxi no element can fit

	copy function
	var nums=make([]int,0,5)
	var nums1=make([]int,len(nums))

	nums.append(nums,2)

	copy(nums2,nums)

	slice operator

	var nums=[]int{1,3,3}

	fmt.Printl(nums[0:2])


	//  maps key value pair
	create mapp
	m:=make(map[string]string)

	if map key value didn't exist in map return 0 value

	// settings element
	m["name"]="golang"
	m["area"]="backend"

	fmt.Println(m["name"],m["area"])

	m:=mpp[string]int{"golang": 1,"cpp":2, "java":3}

	_,ok:=m["cpp"]

	if ok{
			fmt.Println("ok")
	}
	else fmt.Println("didn't ok")

	// range for iterating over data structure
	nums:=[]int {68,89,89}

	for i:=0;i<len(nums):i++{
		sum+=nums[i]
				fmt.Println(i)
	}

	_ ->index
	for _,num:=range nums{
		fmt.Println(num)
		sum+=num
	}
	
	fmt.Println(sum)

	for i,c:= range "golang" {
		fmt.Println(c)
	}

	function

	func add(a int,b int) int{
		return a+b
	}

	variadic function pass n paramaters
	fmt.Println(1,3,3,4,4,5,"hello")


	func sum(nums ...int) int {

	}

	Closures functions knows outer scops

	Pointers stores memory address

	by value passed
	func changeNUm(num *int){
		*num=5  dereference
		fmt.Println("In ChnageNum",num)
	}

	func main(){
		num:=1

		ChangeNum(&num)
		fmt.Println("memory address"&num)

		fmt.Println(num)
	}









*/

func Counter() func() int{
	var cnt int =0

	return func() int {
		cnt++
		return cnt
	}
}

func getlanguage() (string,string){
	return "golang", "cpp"
}

func sum(nums ...int) int {
	total:=0;

	for _, num:= range nums {
		total+=num
	}
	return total

	
}



// func main(){

// 	nu:= []int{3,4,5,9}
	
// 	fmt.Println(getlanguage())
// 	result:= sum(nu...)
// 	fmt.Println(result)

// 	// increment:=Counter()
// 	// fmt.Println(increment)

// }

// func main(){

// 	// var nums[]int
// 	var nums = make([]int,3)
// 	fmt.Println(cap(nums))

// 	fmt.Println(nums==nil)

// 	nums = append(nums,1)
// 	fmt.Println(cap(nums))
// 	nums = append(nums,4)
// 	fmt.Println(cap(nums))
// }

// func main(){

// 	// infer
// 	// var name string ="golang"
// 	var name="golang"

// 	// shorthnad synatax
// 	// name :="atul"

// 	var age int = 12

// 	var price float32 = 59.7
// 	fmt.Println(price)
// 	fmt.Println(age)
// 	fmt.Println(name)

// 	// for loops

// 	// while loop
// 	i:=1
// 	for i<=3 {
// 		fmt.Println(i)
// 		i++
// 	}

// 	// infinite loop
// 	// for {
// 	// 	fmt.Println("1")
// 	// }

// 	// for loop

// 	// for i:=0; i<3; i++ {
// 	// 	fmt.Println(i)
// 	// }
// }