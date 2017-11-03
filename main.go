package main

import ("fmt" // for printing
  "encoding/json" // for encoding json to struct or we can se marshing and unmarshing
  "io/ioutil" //for reading and writing files
  "net/http") //for create server

//function to check to slice and if their strings matches then it will return true, otherwise it will return false
func same_check(slice1 []string, slice2 []string) bool {
	var diff []string
	same := false

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if found {
				same = true
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
  //fmt.Println(diff)
	return same
}

// User struct for user.json
type User []struct {
	UserID         string   `json:"userId"`
	UserName       string   `json:"userName"`
	UserGeo        string   `json:"userGeo"`
	PreferredSize  []string `json:"preferredSize"`
	PreferredColor []string `json:"preferredColor"`
	PreferredStyle []string `json:"preferredStyle"`
	UserInterest   []string `json:"userInterest"`
	UserAttribute  struct {
		Age      string `json:"age"`
		Height   string `json:"height"`
		BodyType string `json:"bodyType"`
		Color    string `json:"color"`
		Gender   string `json:"gender"`
	} `json:"userAttribute"`
}
// product struct for product.json
type Product []struct {
	ImageURLs        []string `json:"imageURLs"`
	ProductTitle     string   `json:"productTitle"`
	Price            int      `json:"price"`
	DiscountedPrice  int      `json:"discountedPrice"`
	DiscountPer      string   `json:"discountPer"`
	MarketplaceName  string   `json:"marketplaceName"`
	AffliateLink     string   `json:"affliateLink"`
	AvailableColor   []string `json:"availableColor"`
	AvailableSize    []string `json:"availableSize"`
	ProductID        string   `json:"productId"`
	Productfeature   []string `json:"productfeature"`
	ProductAttribute struct {
		SuitableFor string `json:"Suitable For"`
		Neck        string `json:"Neck"`
	} `json:"productAttribute"`
	Rating             string `json:"rating"`
	Reviews            int    `json:"reviews"`
	RecommendedProduct []int  `json:"recommendedProduct"`
}

//main function
func main()  {
  //------------------------weights----------------------------------
  size_weight := 50
  color_weight := 20
  //style_weight := 30
  product_weight := 0
  total_weight := 70 //total weight of product if the weight of product is = to this means we have found our relevant product
  //-------------------------- xx------------------------------------

  //-------------reading user and product.json-----------------------
  user_data, _ := ioutil.ReadFile("user.json")
  product_data, _ := ioutil.ReadFile("product.json")
  //---------------------------xx-------------------------------------

  var final_product int
  var found_product []byte

  //-----------------creating struct into variable-------------------
  var products Product
  var user1 User
  //---------------------------xx-----------------------------------


  //unmarshing json to struct variable
  json.Unmarshal(user_data,&user1)
  json.Unmarshal(product_data,&products)
  //----------------------------xx---------------------------------

  //for loop for range of users as here is 1 user so it will run once only
  for _,user := range user1{
  //for loop for range of product as here is 3 products so it will run 3 times
    for _, product := range products{

      //checking if slice of product.availableColor is = to user.PreferredColor if yes then we will get true as return
      if (same_check(product.AvailableColor,user.PreferredColor)){
        product_weight += color_weight //assigning color weight to product weight
      }
      //here checking size of product and user, if yes then it will get true as return
      if (same_check(product.AvailableSize,user.PreferredSize)){
        product_weight += size_weight
      }
      //checking if product weight is greater then final_product weight or not if yes then we assign product_weight to final_product
      if product_weight > final_product{
        final_product = product_weight
        found_product, _ = json.Marshal(product) //now we parse the struct data of more weight product to json
      }
      fmt.Println(product_weight)
      if product_weight == total_weight{
        break
      }
      product_weight = 0
    }
  }
  //------------------------------------xx----------------------------------------------

  //---------------------------starting http server-------------------------------------
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //setting header of / url as json
    w.Header().Set("Content-Type", "application/json")
    //writing the json data as string because it is of datatype byte and to write something on page then it should be string
    fmt.Fprintf(w, string(found_product))
})
  //listening to port 8080
  http.ListenAndServe(":8080", nil)
  //------------------------------------------xx----------------------------------------
}
