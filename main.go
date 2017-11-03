package main

import ("fmt"
  "sort"
  "encoding/json"
  "io/ioutil"
  "strings")

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
  //size_weight := 50
  color_weight := 20
  //style_weight := 30
  product_weight := 0
  user_data, _ := ioutil.ReadFile("user.json")
  product_data, _ := ioutil.ReadFile("product.json")

  var user1 User
  json.Unmarshal(user_data,&user1)

  var products Product
  json.Unmarshal(product_data,&products)
  for _,user := range user1{
    sort.Strings(user.PreferredSize)
    sort.Strings(user.PreferredColor)
    sort.Strings(user.PreferredStyle)

    //user_color := strings.Join(user.PreferredColor, " ")
    //user_size := strings.Join(user.PreferredSize, " ")
    //user_style := strings.Join(user.PreferredStyle, " ")

    for _, product := range products{
      fmt.Println(product.ProductTitle)
      fmt.Println(product.AvailableColor)
      /*sort.Strings(product.AvailableColor)
      sort.Strings(product.AvailableSize)
      sort.Strings(product.Productfeature)
      */

      /*
      fmt.Println(product.AvailableColor)
      if (sort.SearchStrings(product.AvailableColor, "Blue")) <len(product.AvailableColor){
        product_weight += color_weight
      }
      if (sort.SearchStrings(product.AvailableSize, user_size)) <len(product.AvailableSize) {
        product_weight += size_weight
      }
      if (sort.SearchStrings(product.Productfeature, user_style)) <len(product.Productfeature) {
        product_weight += style_weight
      }
      */
      if strings.Contains(strings.Join(product.AvailableColor, " "),"Red") {
        product_weight += color_weight
      }
      fmt.Println(product_weight)
      product_weight = 0

    }
  }



  /* strSlice := []string {"Texas","Washington","Montana","Alaska","Indiana","Ohio","Nevada"}
  search:= "Texas"
  sort.Strings(strSlice)
   pos := sort.SearchStrings(strSlice,search)
   fmt.Printf("Found %s at index %d in %v\n", search, pos, strSlice)
*/
}
