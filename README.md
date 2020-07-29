# Team2-Case-Study-1
Repo for Case Study 1

![Screenshot](assets/walk_through.png)

### Libraries Required (Automatically get installed while running the application)
1. To install gjson, use : `go get -u "github.com/tidwall/gjson"`
2. To install GIN, use : `go get -u "github.com/gin-gonic/gin"`
3. To install gRPC, use : `go get -u "google.golang.org/grpc"`

### To run this project, follow the below steps :-
- Clone this repo : `git clone github.com/shashijangra22/Team2-Case-Study-1`
- Initialise the DB with: `go run cmd/DB/main.go`
- Start the Application using command: `go run cmd/API/main.go`
- Start the gRPC Serve using command: `go run cmd/Server/main.go`

### Now, The webserver will run on localhost:9001.

To fetch different query :
- `"localhost:9001/api"` for HomePage
- `"localhost:9001/api/orders"` for fetching all orders
- `"localhost:9001/api/avg-price"` for average price of orders per customer
- `"localhost:9001/auth/top-buyers/:numBuyers"` for top-customers based on expenditure
- `"localhost:9001/auth/top-restaurants/:numRestau"` for top-restaurants based on its revenue
- `"localhost:9001/auth/new-order"` to place a new order

#### The url with `auth` in path will require `username` = `team2` & `password` = `xurde`