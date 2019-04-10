In order to run tests use this command:

go test -v **/*_test.go

Here is request example:

curl -v -X POST -H 'Content-Type:application/json' -d '{"supply_list":[30,40,1520],"demand_list":[20,30,30,10],"cost_table":[[2,3,2,4],[3,2,5,1],[4,3,2,6]]}' "http://localhost:8080/api/tasks/"
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /api/tasks/ HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type:application/json
> Content-Length: 101
>
* upload completely sent off: 101 out of 101 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Sun, 07 Apr 2019 21:39:45 GMT
< Content-Length: 983
<
* Connection #0 to host localhost left intact
{
  "total_delivery_cost": 220,
  "UUID": "004af910-93a6-4746-9046-9d331da70800",
  "supply_list": [
    {
      "amount": 30,
      "potential": 0,
      "is_fake": false
    },
    {
      "amount": 40,
      "potential": 0,
      "is_fake": false
    },
    {
      "amount": 1520,
      "potential": 0,
      "is_fake": false
    }
  ],
  "demand_list": [
    {
      "amount": 20,
      "potential": 2,
      "is_fake": false
    },
    {
      "amount": 30,
      "potential": 0,
      "is_fake": false
    },
    {
      "amount": 30,
      "potential": 0,
      "is_fake": false
    },
    {
      "amount": 10,
      "potential": 0,
      "is_fake": false
    },
    {
      "amount": 1500,
      "potential": 0,
      "is_fake": true
    }
  ],
  "table_cells": [
    [
      {
        "cost": 2,
        "delivery_amount": 20
      },
      {
        "cost": 3,
        "delivery_amount": 0
      },
      {
        "cost": 2,
        "delivery_amount": 0
      },
      {
        "cost": 4,
        "delivery_amount": 0
      },
      {
        "cost": 0,
        "delivery_amount": 10
      }
    ],
    [
      {
        "cost": 3,
        "delivery_amount": 0
      },
      {
        "cost": 2,
        "delivery_amount": 30
      },
      {
        "cost": 5,
        "delivery_amount": 0
      },
      {
        "cost": 1,
        "delivery_amount": 0
      },
      {
        "cost": 0,
        "delivery_amount": 10
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
        "delivery_amount": 30
      },
      {
        "cost": 6,
        "delivery_amount": 10
      },
      {
        "cost": 0,
        "delivery_amount": 1480
      }
    ]
  ],
  "is_optimal_solution": true
}


Here is logging example:

INFO[0000] Starting server at port :8080

INFO[0014] Started 115fe1b9-10e8-4edf-80fe-ffbd88c15531 POST /api/tasks/ [::1]:62262 curl/7.54.0
INFO[0014] Received parameters: {"supply_list":[30,40,20],"demand_list":[20,30,30,10],"cost_table":[[2,3,2,4],[3,2,5,1],[4,3,2,9]]}  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Created Task UUID: 29b7e22b-3b77-4b71-be72-b541b01d0fdf  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=0      | X=0      | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=0      | X=0      | X=0      |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=0      | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Process Task UUID: 29b7e22b-3b77-4b71-be72-b541b01d0fdf  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Initial Preparations ===                  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #1 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Balancing                             request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=0      | X=0      | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=0      | X=0      | X=0      |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=0      | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Balancing: Task is already balanced. Skip balancing  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #2 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Apply Degeneracy Prevention                   request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=0      | X=0      | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=0      | X=0      | X=0      |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=0      | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Added 2.500000e-04 to each demand Amount. Added 1.000000e-03 to first supply Amount. Demand Amounts: [20.00025 30.00025 30.00025 10.00025] First supply Amount: 30.001000  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #3 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Calculate initial base plan with 'North West Corner' method  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=20     | X=0      |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=10     | X=10     |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Done 'North West Corner' base plan calculation  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Delivery Cost: 320                            request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Potentials Method. Iteration #1 ===       request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #1 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Initialize task inner state before current iteration start  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=20     | X=0      |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=10     | X=10     |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Reset potentials, grades and circuit data     request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #2 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform amount distribution check             request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=20     | X=0      |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=10     | X=10     |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Sums of delivery amounts by columns and rows match each other  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #3 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Degeneracy Check                      request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=20     | X=0      |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=10     | X=10     |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Solution is not Degenerate                    request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #4 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Calculate Potentials                          request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  | V[3]=13  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=20     | X=0      |
| U[1]=-1    | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=10     | X=10     |
| U[2]=-4    | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Potentials have been assigned to demand row and supply column  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #5 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Optimal Solution Check                request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  | V[3]=13  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=-4 | C=4 D=-9 |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=20     | X=0      |
| U[1]=-1    | -------- | -------- | -------- | min Δ    |
|            | C=3 D=2  | C=2 D=0  | C=5 D=0  | -------- |
|            |          |          |          | C=1      |
|            |          |          |          | D=-11    |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=10     | X=10     |
| U[2]=-4    | -------- | -------- | -------- | -------- |
|            | C=4 D=6  | C=3 D=4  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Not Optimal Solution. Min Negative Delta Cell: D[1][3]= -11  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #6 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Finding the circuit path and minimal theta vertex  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  | V[3]=13  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=-4 | C=4 D=-9 |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=20     | X=0  (+) |
| U[1]=-1    | -------- | -------- |  (-)     | min Δ    |
|            | C=3 D=2  | C=2 D=0  | -------- | -------- |
|            |          |          | C=5 D=0  | C=1      |
|            |          |          |          | D=-11    |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=10     | X=10     |
| U[2]=-4    | -------- | -------- |  (+)     |  (-)     |
|            | C=4 D=6  | C=3 D=4  | -------- | min θ    |
|            |          |          | C=2 D=0  | -------- |
|            |          |          |          | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Path: [{1 3} {2 3} {2 2} {1 2}]. Theta cell is at [2][3]  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #7 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Supply Redistribution                 request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  | V[3]=13  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=-4 | C=4 D=-9 |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=10     | X=10     |
| U[1]=-1    | -------- | -------- |  (-)     |  (+)     |
|            | C=3 D=2  | C=2 D=0  | -------- | min Δ    |
|            |          |          | C=5 D=0  | -------- |
|            |          |          |          | C=1      |
|            |          |          |          | D=-11    |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0  (-) |
| U[2]=-4    | -------- | -------- |  (+)     | min θ    |
|            | C=4 D=6  | C=3 D=4  | -------- | -------- |
|            |          |          | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Delivery amounts have been updated according to theta[2][3] value and signs (+) (-)  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Potentials Method. Iteration #2 ===       request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #1 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Initialize task inner state before current iteration start  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=10     | X=10     |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Reset potentials, grades and circuit data     request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #2 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform amount distribution check             request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=10     | X=10     |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Sums of delivery amounts by columns and rows match each other  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #3 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Degeneracy Check                      request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=10     | X=10     |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Solution is not Degenerate                    request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #4 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Calculate Potentials                          request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  |  V[3]=2  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=10     | X=10     |
| U[1]=-1    | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=-4    | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Potentials have been assigned to demand row and supply column  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #5 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Optimal Solution Check                request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  |  V[3]=2  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0      | X=0      |
| U[0]=0     | -------- | -------- | min Δ    | -------- |
|            | C=2 D=0  | C=3 D=0  | -------- | C=4 D=2  |
|            |          |          | C=2 D=-4 |          |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=10     | X=10     |
| U[1]=-1    | -------- | -------- | -------- | -------- |
|            | C=3 D=2  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=-4    | -------- | -------- | -------- | -------- |
|            | C=4 D=6  | C=3 D=4  | C=2 D=0  | C=9 D=11 |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Not Optimal Solution. Min Negative Delta Cell: D[0][2]= -4  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #6 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Finding the circuit path and minimal theta vertex  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  |  V[3]=2  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=10     | X=0  (+) | X=0      |
| U[0]=0     | -------- |  (-)     | min Δ    | -------- |
|            | C=2 D=0  | -------- | -------- | C=4 D=2  |
|            |          | C=3 D=0  | C=2 D=-4 |          |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=20     | X=10     | X=10     |
| U[1]=-1    | -------- |  (+)     |  (-)     | -------- |
|            | C=3 D=2  | -------- | min θ    | C=1 D=0  |
|            |          | C=2 D=0  | -------- |          |
|            |          |          | C=5 D=0  |          |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=-4    | -------- | -------- | -------- | -------- |
|            | C=4 D=6  | C=3 D=4  | C=2 D=0  | C=9 D=11 |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Path: [{0 2} {1 2} {1 1} {0 1}]. Theta cell is at [1][2]  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #7 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Supply Redistribution                 request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=6  |  V[3]=2  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=0  (-) | X=10     | X=0      |
| U[0]=0     | -------- | -------- |  (+)     | -------- |
|            | C=2 D=0  | C=3 D=0  | min Δ    | C=4 D=2  |
|            |          |          | -------- |          |
|            |          |          | C=2 D=-4 |          |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=30     | X=0  (-) | X=10     |
| U[1]=-1    | -------- |  (+)     | min θ    | -------- |
|            | C=3 D=2  | -------- | -------- | C=1 D=0  |
|            |          | C=2 D=0  | C=5 D=0  |          |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=-4    | -------- | -------- | -------- | -------- |
|            | C=4 D=6  | C=3 D=4  | C=2 D=0  | C=9 D=11 |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Delivery amounts have been updated according to theta[1][2] value and signs (+) (-)  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Potentials Method. Iteration #3 ===       request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #1 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Initialize task inner state before current iteration start  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=0      | X=10     | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=30     | X=0      | X=10     |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Reset potentials, grades and circuit data     request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #2 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform amount distribution check             request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=0      | X=10     | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=30     | X=0      | X=10     |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Sums of delivery amounts by columns and rows match each other  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #3 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Degeneracy Check                      request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=0  |  V[1]=0  |  V[2]=0  |  V[3]=0  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=0      | X=10     | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=30     | X=0      | X=10     |
| U[1]=0     | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Solution is not Degenerate                    request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #4 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Calculate Potentials                          request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=2  |  V[3]=2  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=0      | X=10     | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=0  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=30     | X=0      | X=10     |
| U[1]=-1    | -------- | -------- | -------- | -------- |
|            | C=3 D=0  | C=2 D=0  | C=5 D=0  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=0  | C=3 D=0  | C=2 D=0  | C=9 D=0  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Potentials have been assigned to demand row and supply column  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Step #5 ===                               request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Perform Optimal Solution Check                request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Current Task State Table:
+------------+----------+----------+----------+----------+
| → DEMAND → | B[0]=20  | B[1]=30  | B[2]=30  | B[3]=10  |
| ---------- |  V[0]=2  |  V[1]=3  |  V[2]=2  |  V[3]=2  |
| ↓ SUPPLY ↓ |          |          |          |          |
+------------+----------+----------+----------+----------+
| A[0]=30    | X=20     | X=0      | X=10     | X=0      |
| U[0]=0     | -------- | -------- | -------- | -------- |
|            | C=2 D=0  | C=3 D=0  | C=2 D=0  | C=4 D=2  |
+------------+----------+----------+----------+----------+
| A[1]=40    | X=0      | X=30     | X=0      | X=10     |
| U[1]=-1    | -------- | -------- | -------- | -------- |
|            | C=3 D=2  | C=2 D=0  | C=5 D=4  | C=1 D=0  |
+------------+----------+----------+----------+----------+
| A[2]=20    | X=0      | X=0      | X=20     | X=0      |
| U[2]=0     | -------- | -------- | -------- | -------- |
|            | C=4 D=2  | C=3 D=0  | C=2 D=0  | C=9 D=7  |
+------------+----------+----------+----------+----------+
  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Solution is optimal. Proccesing is Completed  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Delivery Cost: 170                            request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] === Caclulation took 11.186736ms and 3 iterations ===  request_id=115fe1b9-10e8-4edf-80fe-ffbd88c15531
INFO[0014] Finished 115fe1b9-10e8-4edf-80fe-ffbd88c15531 POST /api/tasks/ [::1]:62262 curl/7.54.0


Error logging example:

INFO[0010] Started 1bbb9979-dcac-40f7-a866-564973566fd9 GET /api/tasks/ [::1]:51369 curl/7.54.0
ERRO[0010] Invalid request method                        request_id=1bbb9979-dcac-40f7-a866-564973566fd9
INFO[0010] Finished 1bbb9979-dcac-40f7-a866-564973566fd9 GET /api/tasks/ [::1]:51369 curl/7.54.0

ERRO[0263] Task Solver: Calculation took 1m0.002204546s and exceded allowed limit of 1m0s  request_id=e4feda3e-eef1-41f2-9631-84cddaf0d91c
INFO[0263] Finished e4feda3e-eef1-41f2-9631-84cddaf0d91c POST /api/tasks/ [::1]:51321 curl/7.54.0

INFO[0310] Started c7748ba1-93e8-48e0-8a29-10fce0bf6580 POST /api/tasks/ [::1]:51335 curl/7.54.0
INFO[0310] Received parameters: {"supply_list":[30,40,20],"demand_list":[20,30,30,10],"cost_table":[]}  request_id=c7748ba1-93e8-48e0-8a29-10fce0bf6580
ERRO[0310] Params Validation Error: CostTable should have at least one row  request_id=c7748ba1-93e8-48e0-8a29-10fce0bf6580
INFO[0310] Finished c7748ba1-93e8-48e0-8a29-10fce0bf6580 POST /api/tasks/ [::1]:51335 curl/7.54.0

Error response example:

* upload completely sent off: 101 out of 101 bytes
< HTTP/1.1 500 Internal Server Error
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Fri, 05 Apr 2019 04:58:58 GMT
< Content-Length: 96
<
{"error_message":"Task Solver: Calculation took 23.451804ms and exceded allowed limit of 10ms"}
or
{"error_message":"Params Validation Error: CostTable should have at least one row"}
or
{"error_message":"Params Validation Error: DemandList size '1' and CostTable columns count '0' should be equal"}

and other error messages in the same format