# Team2-Case-Study-1
repo for case study 1

## To run this project, follow the below steps :-
### 1. Clone this repo.
### 2. Change the GOPATH = pwd of Team2-Case-Study-1
### 3. To install gjson, use : go get -u "github.com/tidwall/gjson"
### 4. To install GIN, use : go get -u "github.com/gin-gonic/gin"
### 5. Now, run command in Team2-Case-Study-1/src -> go run main.go modified.csv

#### Now, The webserver will run on localhost:9001.

To fetch different query :-
### 1.  "localhost:9001/api/" for HomePage
### 2.  "localhost:9001/api/orders" for fetching all orders
### 3.  "localhost:9001/api/avg-price" for average price of orders per customer
### 3.  "localhost:9001/api/top-buyers/:numBuyers" for top-customers based on expenditure
### 4.  "localhost:9001/api/top-restaurants/:numRestau" for top-restaurants based on its revenue
### 5.  "localhost:9001/api/new-order" to place a new order