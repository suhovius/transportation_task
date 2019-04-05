Here is request example:

curl -v -X POST -H 'Content-Type:application/json' -X POST -d '{"supply_list":[30,40,20],"demand_list":[20,30,30,10],"cost_table":[[2,3,2,4],[3,2,5,1],[4,3,2,6]]}' "http://localhost:8080/api/tasks/"
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /api/tasks/ HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type:application/json
> Content-Length: 99
>
* upload completely sent off: 99 out of 99 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Fri, 05 Apr 2019 04:22:12 GMT
< Content-Length: 948
<
* Connection #0 to host localhost left intact
{
  "UUID": "df1c02f4-588c-45f9-a633-86f0cb0bbe7d",
  "supply_list": [
    {
      "amount": 30.001,
      "potential": 0,
      "is_fake": false
    },
    {
      "amount": 40,
      "potential": -1,
      "is_fake": false
    },
    {
      "amount": 20,
      "potential": 0,
      "is_fake": false
    }
  ],
  "demand_list": [
    {
      "amount": 20.00025,
      "potential": 2,
      "is_fake": false
    },
    {
      "amount": 30.00025,
      "potential": 3,
      "is_fake": false
    },
    {
      "amount": 30.00025,
      "potential": 2,
      "is_fake": false
    },
    {
      "amount": 10.00025,
      "potential": 2,
      "is_fake": false
    }
  ],
  "table_cells": [
    [
      {
        "cost": 2,
        "delivery_amount": 20.00025
      },
      {
        "cost": 3,
        "delivery_amount": 0.0004999999999988347
      },
      {
        "cost": 2,
        "delivery_amount": 10.000250000000001
      },
      {
        "cost": 4,
        "delivery_amount": 0
      }
    ],
    [
      {
        "cost": 3,
        "delivery_amount": 0
      },
      {
        "cost": 2,
        "delivery_amount": 29.999750000000002
      },
      {
        "cost": 5,
        "delivery_amount": 0
      },
      {
        "cost": 1,
        "delivery_amount": 10.000249999999998
      }
    ],
    [
      {
        "cost": 4,
        "delivery_amount": 0
      },
      {
        "cost": 3,
        "delivery_amount": 0
      },
      {
        "cost": 2,
        "delivery_amount": 20
      },
      {
        "cost": 6,
        "delivery_amount": 0
      }
    ]
  ],
  "is_optimal_solution": true,
  "total_delivery_cost": 170.00225
}