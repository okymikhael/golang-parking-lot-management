# Readme
### Developer Note
1. To run program: `go run cmd/main.go`
2. To run test: 
    - `cd test`
    - `go test`
3. Developer go version: `go1.19.1 windows/amd64`

### Developer Explanation
List Endpoint with function:
#### 1. /create_parking_lot => StoreParkingLot
    This endpoint is for creating parking lot, there's 4 main function:
        1. Handler that only allow POST method
        2. Check if url path not wrong
        3. Handler slot is only can initialize one time
        4. Delete map init value and initialize slot value 
#### 2. /park => SetSlotParking
    This endpoint is for parking car, there's 7 main function:
        1. Handler that only allow POST method
        2. Check if url path not wrong
        3. Check if slot was initialize
        4. Check if car was parked
        5. Check if there's space for parking
        6. Check if parking area is full
        7. Insert car into slice
#### 3. /leave => EditSlotParking
    This endpoint is for delete car from parking lot, there's 3 main function:
        1. Handler that only allow POST method
        2. Check if url path not wrong
        3. Remove car from slice
#### 4. /status => GetStatus
    This endpoint is for show all data of parking lot, there's 2 main function:
        1. Handler that only allow GET method
        2. Parsing parking slot information into table
#### 5. /cars_registration_numbers/ =>  GetCarsRegisByColour
    This endpoint is for find car with color and return registration number, there's 4 main function:
        1. Handler that only allow GET method
        2. Check if url path not wrong
        3. Find car by color and return registation bumber
        4. Parsing parking slot information separate by comma
#### 6. /cars_slot => GetSlotByColour
    This endpoint is for find car with color black and return slot number, there's 4 main function:
        1. Handler that only allow GET method
        2. Check if url path not wrong
        3. Find car by color and return slot number
        4. Parsing parking slot information into slot number
#### 7. /slot_number => GetSlotByRegisNum
    This endpoint is for find car with registration number and return slot number, there's 4 main function:
        1. Handler that only allow GET method
        2. Check if url path not wrong
        3. Find car by color and return slot number
        4. Parsing parking slot information into slot number
#### 8. /bulk => StoreBulkData
    This endpoint is for running bulk data with formated text, there's 5 main function:
        1. Handler that only allow GET method
        2. Read txt from request body
        3. Set translator and missingMethod for data that have specific need
        4. Send request by http request then save response into tempResponse variable
        5. Parsing response information into one data
#### 9. /reset => ResetParkingLot
    This endpoint is for unit test and basically this endpoint for reset all data in parking lot without restart server, there's 2 main function:
        1. Handler that only allow POST method
        2. Delete max value if has value
