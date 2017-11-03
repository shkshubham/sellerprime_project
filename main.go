package main

import ("fmt"
  "encoding/json"
  "io/ioutil"
  "net/http")
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

func main()  {
  size_weight := 50
  color_weight := 20
  //style_weight := 30
  product_weight := 0
  user_data, _ := ioutil.ReadFile("user.json")
  product_data, _ := ioutil.ReadFile("product.json")
  var final_product int
  var found_product []byte
  var user1 User
  json.Unmarshal(user_data,&user1)

  var products Product
  json.Unmarshal(product_data,&products)
  for _,user := range user1{

    for _, product := range products{

      if (same_check(product.AvailableColor,user.PreferredColor)){
        product_weight += color_weight
      }
      if (same_check(product.AvailableSize,user.PreferredSize)){
        product_weight += size_weight
      }
      if product_weight > final_product{
        final_product = product_weight
        found_product, _ = json.Marshal(product)
      }
      fmt.Println(product_weight)
      product_weight = 0
    }
  }

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(found_product))
})
  http.ListenAndServe(":8080", nil)
  fmt.Println(string(found_product))
}
