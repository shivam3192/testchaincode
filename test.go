package main

import (
        "errors"
        "fmt"
       // "strconv"
        "encoding/json"
        "github.com/hyperledger/fabric/core/chaincode/shim"
)

// CrowdFundChaincode implementation
type CrowdFundChaincode struct {
}
type Info struct {

        Qrcode []string   `json:"qrcode"`
        Count  []string   `json:"count"`


}
//
// Init creates the state variable with name "account" and stores the value
// from the incoming request into this variable. We now have a key/value pair
// for account --> accountValue.
//
func (t *CrowdFundChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        // State variable "account"
        // The value stored inside the state variable "account"
        
        // Any error to be reported back to the client
        var err error

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }

     //   information := Info{}
       // informationbyte, err := json.Marshal(information)
     if err!=nil {
                        return nil, err
                }
         record := Info{}
       // errrecordmarshal := json.Unmarshal(recordByte,&record);
        record.Qrcode=append(record.Qrcode,"aaaa");
        record.Count=append(record.Count,"aasss");
        newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
                err=stub.PutState("default",newrecordByte);
         if err!=nil {
                        return nil, err
                }



        return nil, nil
}

//
// Invoke retrieves the state variable "account" and increases it by the ammount
// specified in the incoming request. Then it stores the new value back, thus
// updating the ledger.
//
func (t *CrowdFundChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    
var account string

        var err error

        if len(args) != 3 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }
          account = args[0]

         recordByte, err := stub.GetState(account);
        fmt.Println(recordByte);
        if err != nil {

            return nil, err
        }
        record := Info{}
        if recordByte != nil {
        errrecordmarshal := json.Unmarshal(recordByte,&record);
        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }    
               
        }
       
            
        record.Qrcode = append(record.Qrcode,args[1]);
        record.Count = append(record.Count,args[2]);
        newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
        err =stub.PutState(account,newrecordByte);
        if err != nil {

            return nil, err;
        } 
        return nil, nil
}


//
// Query retrieves the state variable "account" and returns its current value
// in the response.
//
func (t *CrowdFundChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if function != "query" {
                return nil, errors.New("Invalid query function name. Expecting \"query\".")
        }

        // State variable "account"
       // var account string
        // Any error to be reported back to the client
        var err error

         if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting name of the state variable to query.")
        }

        // Read in the name of the state variable to be returned
     var   account = args[0]
     //   information:=Info{}
        // Get the current value of the state variable
        accountValueBytes ,err := stub.GetState(account)
        if err != nil {
              //  jsonResp := "{\"Error\":\"Failed to get state for " + account + "\"}"
                 return nil, err
        }
      /*  if accountValueBytes == nil {
                jsonResp := "{\"Error\":\"Nil amount for " + account + "\"}"
                return nil, errors.New(jsonResp)
        }
      errUnmarshal:=json.Unmarshal(accountValueBytes,&information)
       if errUnmarshal!=nil {
                return nil,errUnmarshal
        }
        jsonResp := "{\"Name\":\"" + account + "\",\"qrcode\":\"" + information.qrcode + "\",\"count\":\"" + information.count + "\"}"
        fmt.Printf("Query Response:%s\n", jsonResp)
        */
        return accountValueBytes, nil
}

func main() {
        err := shim.Start(new(CrowdFundChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}

