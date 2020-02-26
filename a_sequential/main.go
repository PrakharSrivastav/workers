package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	terms := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4}

	c := &http.Client{Timeout: time.Millisecond * 15000}

	for i,num := range terms{
		callApi(num,i,c)
	}
	log.Print(time.Now().Sub(start).Seconds())
}



func callApi(num,id int , c *http.Client) {
	baseURL := "https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/%d"

	ur := fmt.Sprintf(baseURL, num)
	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		//log.Printf("error creating a request for term %d :: error is %+v", num, err)
		return
	}
	res, err := c.Do(req)
	if err != nil {
		//log.Printf("error querying for term %d :: error is %+v", num, err)
		return
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		//log.Printf("error reading response body :: error is %+v", err)
		return
	}
	//log.Printf("%d  :: ok", id)
}