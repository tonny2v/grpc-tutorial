package main

//import (
//"context"
//"encoding/json"
//"fmt"
//"github.com/gorilla/mux"
//"google.golang.org/grpc"
//"log"
//"net/http"
//"strconv"
//"time"
//pb "tonny2v/proto"
//)

//func main() {
////	Connect to Add service
//conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
//if err != nil {
//log.Fatalf("Dial Failed: %v", err)
//}
//addClient := pb.NewAddServiceClient(conn)

//routes := mux.NewRouter()
//routes.HandleFunc("/add/{a}/{b}", func(w http.ResponseWriter, r *http.Request) {
//w.Header().Set("Content-Type", "application/json; charset=UFT-8")

//vars := mux.Vars(r)
//a, err := strconv.ParseUint(vars["a"], 10, 64)
//if err != nil {
//json.NewEncoder(w).Encode("Invalid parameter A")
//}
//b, err := strconv.ParseUint(vars["b"], 10, 64)
//if err != nil {
//json.NewEncoder(w).Encode("Invalid parameter B")
//}

//ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
//defer cancel()

//req := &pb.AddRequest{A: a, B: b}
//if resp, err := addClient.Compute(ctx, req); err == nil {
//msg := fmt.Sprintf("Summation is %d", resp.Result)
//json.NewEncoder(w).Encode(msg)
//} else {
//msg := fmt.Sprintf("Internal server error: %s", err.Error())
//json.NewEncoder(w).Encode(msg)
//}
//}).Methods("GET")

//fmt.Println("Application is running on : 8080 .....")
//http.ListenAndServe(":8080", routes)
/*}*/
