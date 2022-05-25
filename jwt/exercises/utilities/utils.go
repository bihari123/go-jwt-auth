package utilities

import "fmt"

func CheckError(funcName string, err error) (interface{},bool){
  if err!=nil{
    log.Println(fmt.Errorf("Error in %s: %w",funcName,err))
    return nil,true 
  }
  return nil,false 
}
