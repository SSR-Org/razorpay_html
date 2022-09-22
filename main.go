package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	OrderId string
}

func main() {
	http.HandleFunc("/", App)
	log.Fatal(http.ListenAndServe(":8098", nil))
}

func App(w http.ResponseWriter, r *http.Request) {

	client := razorpay.NewClient("rzp_test_7MB1e7rfTp35VG", "YHDKuRvW0xvBvJb706mYfgDZ")

	data := map[string]interface{}{
		"amount":          50000,
		"currency":        "INR",
		"receipt":         "some_receipt_id",
		"payment_capture": 1,
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Save Order Id from the body
	value := body["id"]

	str := value.(string)

	HomePageVars := PageVariables{ // store the order_id in a struct
		OrderId: str,
	}

	t, err := template.ParseFiles("app.html") // parse the html file app.html
	if err != nil {                           // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, HomePageVars) // execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) // log it
	}

}
