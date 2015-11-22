package main


import (
"fmt"
"net/http"
"encoding/json"
"io/ioutil"
"os"
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"log"
"strconv"
"bytes"

)

type Best_route struct {
    locations [20]string
    count int 
}

var arraying[10] Best_route
var request_id_counter int
var for_put[20] int                              //for keeping track of how many times the user called the put api 
                   //referencing of request-id

const(
  Address = "ds045054.mongolab.com:45054"
  Database ="location"
  Username ="ketkigawande"
  Password ="17@Ketki"
)

type Sandbox_uber struct {
    
    Status string `json:"status"`
}


type Request struct {
       Starting_from_location_id string `json:"starting_from_location_id"`
       Location_ids []string `json:"location_ids"`
      
}

type Request2 struct {
      Start_latitude string `json:"start_latitude"`
      Start_longitude string `json:"start_longitude"`
      End_latitude string `json:"end_latitude"`
      End_longitude string `json:"end_longitude"`
      Product_id string `json:"product_id"`
  
}

type Response struct {
    
    Id string `json:"id"`
    Status string `json:"status"`
    Starting_from_location_id string `json:"starting_from_location_id"`
    Best_route_location_ids []string `json:"best_route_location_ids"`
    Total_uber_costs float64 `json:"total_uber_costs"`
    Total_uber_duration float64 `json:"total_uber_duration"`
    Total_distance float64 `json:"total_distance"` 
   } 

type Response1 struct {
    
    Id string `json:"id"`
    Name string `json:"name"`
    Address string `json:"address"`
    City string `json:"city"`
    State string `json:"state"`
    Zip string `json:"zip"`
       Coordinate struct {
           Lat string `json:"lat"`
           Lng string `json:"lng"`
         } `json:"coordinate"`
   }  


   type Response2 struct {

     Id string `json:"id"`
     Status string `json:"status"`
     Starting_from_location_id string `json:"starting_from_location_id"`
     Next_destination_location_id string `json:"next_destination_location_id"`
     Best_route_location_ids []string `json:"best_route_location_ids"`
     Total_uber_costs float64 `json:"total_uber_costs"`
     Total_uber_duration float64 `json:"total_uber_duration"`
     Total_distance float64 `json:"total_distance"`
     Uber_wait_time_eta float64 `json:"uber_wait_time_eta"`

 }

   func Post(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
       var res Response
       
       var while_counter int
       var starting_node string 
       var starting_lat string
       var starting_lon string
       var final_array[20] string
       response_counter := []string{}                  //changes required over here 
       //var response_count int
       var final_array_counter int 
       var best_route_counter int 
       var start_id string
       var end_id string
       var fare float64
       var vel float64
       var miles float64

       reply1:=Response1{}
       reply2:=Response1{}
       reply3:=Response1{}
       reply4:=Response1{}
       reply5:=Response1{}

       request1:=Request{}
       //fmt.Println(fare);   // remember to remove this statement
       //fmt.Println(request_id_counter);

       json.NewDecoder(req.Body).Decode(&request1)   
       arraying[request_id_counter].locations[0]=request1.Starting_from_location_id
       best_route_counter=best_route_counter+1;
       starting_node= request1.Starting_from_location_id
       counter:=len(request1.Location_ids);
       arraying[request_id_counter].count=counter;
       
       
        var visited[20] int 
        var complete_route[20] string
        visited[0]=0;
        complete_route[0]=request1.Starting_from_location_id;
        final_array[0]=request1.Starting_from_location_id;

        final_array_counter=1;
        //fmt.Println("here");
        //fmt.Println("complete_route array");
        for j:=1;j<counter+1;j++ {
        complete_route[j]=request1.Location_ids[j-1];
        //fmt.Println(complete_route[j]);
        }
      
        abc:=&mgo.DialInfo{
        Addrs:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
       }
       session, err := mgo.DialWithInfo(abc)
       //session, err := mgo.Dial("127.0.0.1")
       if err != nil {
               panic(err)
       }
       defer session.Close()

       // Optional. Switch the session to a monotonic behavior.
       session.SetMode(mgo.Monotonic, true)


       for while_counter<counter {       //for loop for ending the condition when all nodes are traversed
             
                         var id_to_insert string;
                         var minimum float64;
                         var for_visited int;
             


                          c := session.DB("location").C("Details")
       
                          err = c.Find(bson.M{"id": starting_node}).One(&reply1)
                          if err != nil {
                          panic(err)
                          } 

                                              starting_lat= reply1.Coordinate.Lat
                                              starting_lon= reply1.Coordinate.Lng
                                              

                  for i:=1;i<counter+1;i++ {          // for loop for comparison thing
             
                                  var pricing float64;
                                  //var dura float64
              
                     
                                              c := session.DB("location").C("Details")
       
                                              err = c.Find(bson.M{"id": complete_route[i]}).One(&reply2)
                                              if err != nil {
                                              panic(err)
                                              }
                                              

                                              

                                          a := fmt.Sprint("https://sandbox-api.uber.com/v1/estimates/price?start_latitude=",starting_lat,"&start_longitude=",starting_lon,"&end_latitude=",reply2.Coordinate.Lat,"&end_longitude=",reply2.Coordinate.Lng,"&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT")
                                          //fmt.Println(a);
                                          response2, err := http.Get(a)
      
                                          if err != nil {
                                          fmt.Printf("%s", err)
                                          os.Exit(1)
                                          } else {
                                          defer response2.Body.Close()
                                          contents, err := ioutil.ReadAll(response2.Body)
                                          if err != nil {
                                          fmt.Printf("%s", err)
                                          os.Exit(1)
                                          }

                                         
                                                               var f interface{}
                                                               err=json.Unmarshal(contents, &f)
                                                               mRes := f.(map[string]interface{})["prices"]
                                                               mRes0 := mRes.([]interface{})[0]
                                                               mEst := mRes0.(map[string]interface{})["high_estimate"]
                                                               pricing= mEst.(float64)
                                                                                  
                                               if (minimum==0) {
                                               //fmt.Println("the very first condition");
                                              
                                                      if(visited[i]==0) {
                                                                            minimum = pricing;
                                                                            
                                                                            id_to_insert=complete_route[i];
                                                                            for_visited=i;
                                                                        }
                                                      //fmt.Println(minimum)
                                               
                                               } else if (pricing<minimum) {
                                                
                                                    //fmt.Println("Nako yeu ithe atta");
                                                                  if(visited[i]==0) {
                                                                                          minimum = pricing;
                                                                                          
                                                                                          id_to_insert=complete_route[i];
                                                                                          //fmt.Println("Hey you");
                                                                                          for_visited=i;
                                                                                     }

     
                                              }else {

                                                                                     
                                                
                                                        //fmt.Println("Do nothing");
                                                     }
    
                                                       
                            }//else ends

              }//inner for ends
                                     
                                     
                                     best_route_counter=best_route_counter+1;
                                     
                                    
                                     
                                     visited[for_visited]=1;
                                     //fmt.Println("oye");
                                     starting_node=id_to_insert;
                                     //fmt.Println("oye");
                                     final_array[final_array_counter]=id_to_insert;
                                     response_counter=append(response_counter,id_to_insert);
                                     
                                     final_array_counter++;
                                     while_counter=while_counter+1;
                                     
                                     
 }//outer for ends

                                                      arraying[request_id_counter].locations=final_array;
                                                      //fmt.Println(arraying[request_id_counter].locations);
                                                      request_id_counter=request_id_counter+1;


                                                      
                                                      

                                                      for k:=0;k<counter;k++ {
                                                              
                                                              start_id=final_array[k];
                                                              end_id=final_array[k+1];

                                                              d := session.DB("location").C("Details")
       
                                                              err = d.Find(bson.M{"id": start_id}).One(&reply3)
                                                              if err != nil {
                                                              panic(err)
                                                              }  

                                                              f := session.DB("location").C("Details")
       
                                                              err = f.Find(bson.M{"id": end_id}).One(&reply4)
                                                              if err != nil {
                                                              panic(err)
                                                              }    

                                                                actual := fmt.Sprint("https://sandbox-api.uber.com/v1/estimates/price?start_latitude=",reply3.Coordinate.Lat,"&start_longitude=",reply3.Coordinate.Lng,"&end_latitude=",reply4.Coordinate.Lat,"&end_longitude=",reply4.Coordinate.Lng,"&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT")
                                                                    //fmt.Println(actual);
                                                                    response3, err := http.Get(actual)
                                
                                                                    if err != nil {
                                                                    fmt.Printf("%s", err)
                                                                    os.Exit(1)
                                                                    } else {
                                                                    defer response3.Body.Close()
                                                                    content, err := ioutil.ReadAll(response3.Body)
                                                                    if err != nil {
                                                                    fmt.Printf("%s", err)
                                                                    os.Exit(1)
                                                                    }

                                         
                                                               var g interface{}
                                                               err=json.Unmarshal(content, &g)
                                                               gRes := g.(map[string]interface{})["prices"]
                                                               gRes0 := gRes.([]interface{})[0]
                                                               gEst := gRes0.(map[string]interface{})["high_estimate"]
                                                               gDur := gRes0.(map[string]interface{})["duration"]
                                                               gDis := gRes0.(map[string]interface{})["distance"]
                                                               vel= vel+gDur.(float64)
                                                               
                                                               fare = fare+gEst.(float64)
                                                               
                                                               miles=miles+gDis.(float64)
                                                               
                                                      
                                                                          }//else ends

                                                    }//for ends

                                                    
                                                      
                                        e := session.DB("location").C("Details")
       
                                        err = e.Find(bson.M{"id": final_array[0]}).One(&reply5)
                                        if err != nil {
                                        panic(err)
                                        } 
                                        
                                        actual1 := fmt.Sprint("https://sandbox-api.uber.com/v1/estimates/price?start_latitude=",reply4.Coordinate.Lat,"&start_longitude=",reply4.Coordinate.Lng,"&end_latitude=",reply5.Coordinate.Lat,"&end_longitude=",reply5.Coordinate.Lng,"&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT")
                                                                    //fmt.Println(actual1);
                                                                    response4, err := http.Get(actual1)
                                
                                                                    if err != nil {
                                                                    fmt.Printf("%s", err)
                                                                    os.Exit(1)
                                                                    } else {
                                                                    defer response4.Body.Close()
                                                                    contentt, err := ioutil.ReadAll(response4.Body)
                                                                    if err != nil {
                                                                    fmt.Printf("%s", err)
                                                                    os.Exit(1)
                                                                    }

                                         
                                                               var l interface{}
                                                               err=json.Unmarshal(contentt, &l)
                                                               lRes := l.(map[string]interface{})["prices"]
                                                               lRes0 := lRes.([]interface{})[0]
                                                               lEst := lRes0.(map[string]interface{})["high_estimate"]
                                                               lDur := lRes0.(map[string]interface{})["duration"]
                                                               lDis := lRes0.(map[string]interface{})["distance"]
                                                               vel= vel+lDur.(float64)
                                                               
                                                               fare = fare+lEst.(float64)
                                                               
                                                               miles=miles+lDis.(float64)
                                                               
                                                       
                                                                 }


                                                       res.Id=strconv.Itoa(request_id_counter-1)
                                                       res.Status="planning"
                                                       res.Starting_from_location_id = request1.Starting_from_location_id
                                                       res.Best_route_location_ids = response_counter                   //this needs to be changed 
                                                       res.Total_uber_costs= fare  //mEst.(float64)
                                                       res.Total_uber_duration= vel///something
                                                       res.Total_distance= miles



     pqr :=&mgo.DialInfo{
        Addrs:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
       }
       session2, err := mgo.DialWithInfo(pqr)
       //session, err := mgo.Dial("127.0.0.1")
       if err != nil {
               panic(err)
       }
       defer session2.Close()
     //fmt.Println("Id:", id1)
       session2.SetMode(mgo.Monotonic, true)

       d := session.DB("location").C("Tripplanning")
       //fmt.Println(d)
       err = d.Insert(res)
       if err != nil {
               log.Fatal(err)
       }

      uj, _ := json.Marshal(res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
   } 

   func Get(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
       reply:=Response{}
       id1:=p.ByName("id")
        
      abc:=&mgo.DialInfo{
        Addrs:[]string{Address},
        Database:Database,
        Username:Username,
        Password:Password,
       }
       session, err := mgo.DialWithInfo(abc)
       //session, err := mgo.Dial("127.0.0.1")
       if err != nil {
               panic(err)
       }
       defer session.Close()

       // Optional. Switch the session to a monotonic behavior.
       session.SetMode(mgo.Monotonic, true)

       c := session.DB("location").C("Tripplanning")
       //fmt.Println(c)

       err = c.Find(bson.M{"id": id1}).One(&reply)
       if err != nil {
       panic(err)
       }

      

      var res Response
      res.Id=id1
      res.Status=reply.Status
      res.Starting_from_location_id=reply.Starting_from_location_id
      res.Best_route_location_ids=reply.Best_route_location_ids
      res.Total_uber_costs=reply.Total_uber_costs
      res.Total_uber_duration=reply.Total_uber_duration
      res.Total_distance=reply.Total_distance

     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
   }

   

   
func Put(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
                 rep1:=Response1{}
                 rep2:=Response1{}
                 repl:=Response{}
                 var req Request2
                 var req2 Sandbox_uber
                 var eta_uber float64
                 var res Response2
                 //var traversed_end_lat string
                 //var traversed_end_lng string


                 i:=p.ByName("trip_id")
                 trip_id1, err := strconv.Atoi(i)
                            if err != nil {
                                // handle error
                                fmt.Println(err)
                                os.Exit(2)
                            }


                 //reply:=Response2{}
                 var localcounter int 
                 var start string
                 var end string
                               abc:=&mgo.DialInfo{
                                                  Addrs:[]string{Address},
                                                  Database:Database,
                                                  Username:Username,
                                                  Password:Password,
                                                 }
                                                 session, err := mgo.DialWithInfo(abc)
                                                 //session, err := mgo.Dial("127.0.0.1")
                                                 if err != nil {
                                                         panic(err)
                                                 }
                                                 defer session.Close()

                                                 // Optional. Switch the session to a monotonic behavior.
                                                 session.SetMode(mgo.Monotonic, true)





                                      e := session.DB("location").C("Tripplanning")
       
                                      err = e.Find(bson.M{"id": p.ByName("trip_id")}).One(&repl)
                                      if err != nil {
                                      panic(err)
                                      } 
                                      //fmt.Println(repl.Status)


                           
                 if for_put[trip_id1]<arraying[trip_id1].count {              

                 //fmt.Println("in If loop");
                 localcounter=for_put[trip_id1];
                 start = arraying[trip_id1].locations[localcounter];

                 end=arraying[trip_id1].locations[localcounter+1];
                 for_put[trip_id1]=for_put[trip_id1]+1;

                 

                 
                      c := session.DB("location").C("Details")
       
                                              err = c.Find(bson.M{"id": start}).One(&rep1)
                                              if err != nil {
                                              panic(err)
                                              }


                      d := session.DB("location").C("Details")
       
                                              err = d.Find(bson.M{"id": end}).One(&rep2)
                                              if err != nil {
                                              panic(err)
                                              }

                   
                   req.Start_latitude=rep1.Coordinate.Lat
                   req.Start_longitude=rep1.Coordinate.Lng
                   req.End_latitude=rep2.Coordinate.Lat
                   req.End_longitude=rep2.Coordinate.Lng

                   product_url := fmt.Sprint("https://sandbox-api.uber.com/v1/products?latitude=",rep1.Coordinate.Lat,"&longitude=",rep1.Coordinate.Lng,"&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT");
                                              
                                              response4, err := http.Get(product_url)
      
                                          if err != nil {
                                          fmt.Printf("%s", err)
                                          os.Exit(1)
                                          } else {
                                          defer response4.Body.Close()
                                          contentss, err := ioutil.ReadAll(response4.Body)
                                          if err != nil {
                                          fmt.Printf("%s", err)
                                          os.Exit(1)
                                          }
                                          var g interface{}
                                          err=json.Unmarshal(contentss, &g)
                                          //fmt.Println(g);
                                           pRes := g.(map[string]interface{})["products"]
                                           pRes0 := pRes.([]interface{})[0]
                                           pEst := pRes0.(map[string]interface{})["product_id"]
                                           //fmt.Println(pEst)
                                           req.Product_id=pEst.(string)
                                           }
                   //req.Product_id="04a497f5-380d-47f2-bf1b-ad4cfdcb51f2"
                   

                   //var jsonStr = []byte(`{"start_latitude":"37.334381","start_longitude":"-121.89432","end_latitude":"37.77703","end_longitude":"-122.419571","product_id":"04a497f5-380d-47f2-bf1b-ad4cfdcb51f2"}`)
                   jsonStr,_:=json.Marshal(req)

                   client := &http.Client{}
                   
                     //PUT request to get requestId

                    //urlStr:="https://sandbox-api.uber.com/v1/requests?product_id=a1111c8c-c720-46c3-8534-2fcdd730040d&start_latitude=37.791762&start_longitude=-122.706677&end_latitude=37.885114&end_longitude=-122.7066777&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT";
                   read, err := http.NewRequest("POST","https://sandbox-api.uber.com/v1/requests", bytes.NewBuffer(jsonStr)) // <-- URL-encoded payload
                           if err != nil {
                       panic(err)
                         }
           
                   read.Header.Set("Content-Type", "application/json")
                   //read.Header.Add("Content-Type", "application/x-www-form-urlencoded")
                   read.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZXMiOlsicmVxdWVzdCJdLCJzdWIiOiIyNmM5ZTFkMy1iYjk4LTQ4YTgtOGY2Zi0zMDc5YWZlNmYzMzEiLCJpc3MiOiJ1YmVyLXVzMSIsImp0aSI6IjA1NjU5OGRiLWI0YzgtNDdhMS05OTI4LTM0YTY1NmM2ZTUwNiIsImV4cCI6MTQ1MDU4MTE2MSwiaWF0IjoxNDQ3OTg5MTYwLCJ1YWN0IjoianFiM3dDYlNnQ0hTTWZmbk9hMXpvenpXRTdnT3pSIiwibmJmIjoxNDQ3OTg5MDcwLCJhdWQiOiJVQWJGUWtHTFFCNldFNUNhaGpNY2xzclcwZjJlaW9jQSJ9.aLjrrNeW3NQWumZ7-b16I_6w4NLp8uYwDSteF13_bydLt11FpYS_X2UH3ydsYGMKWMnOw0DKQlVOStyLY5RX3caqDDjP44T_gQV4KUORDYvmD2KZ8xrQxZQPycyCgMTMRQX06n3mnmBkgolntHMJ226I_XSTxWn_NVuYgpF4BEx0jv4zvzQAfngfaRGv_NQgrY4q2m1H2ZEXLhPR4P3Iy9S7N1utULO0A7VHde9jh81CxtU4jByYHsaDwj3TtYeHozWXdSH0vNEgwSJ4MbEkG466KA8JgR9l-PB5b58Ecw5Ex9H6D9iyjmd-hq66dgWncgdoFZ_Mg6Vi3aNWxpkdzw")
                   
                   //read.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

                   resp, err := client.Do(read)
                   //fmt.Println(resp.Status)
                   
                                    if err != nil {
                                      fmt.Printf("%s", err)
                                      os.Exit(1)
                                    } else {
                                      defer resp.Body.Close()
                                      contents, err := ioutil.ReadAll(resp.Body)
                                      if err != nil {
                                          fmt.Printf("%s", err)
                                          os.Exit(1)
                                      }
                                      //fmt.Printf("%s\n", string(contents))
                                      var f interface{}
                                      err=json.Unmarshal(contents, &f)
                                      mRes := f.(map[string]interface{})["eta"]
                                      mReqId := f.(map[string]interface{})["request_id"]
                                      //fmt.Println(mReqId)

                                       //PUT request to set status at sandbox to completed


                                     puturl :=fmt.Sprint("https://sandbox-api.uber.com/v1/sandbox/requests/"+mReqId.(string))
                                     req2.Status="completed"
                                     jsonStr1,_:=json.Marshal(req2)

                    //urlStr:="https://sandbox-api.uber.com/v1/requests?product_id=a1111c8c-c720-46c3-8534-2fcdd730040d&start_latitude=37.791762&start_longitude=-122.706677&end_latitude=37.885114&end_longitude=-122.7066777&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT";
                           read1, err := http.NewRequest("PUT",puturl, bytes.NewBuffer(jsonStr1)) // <-- URL-encoded payload
                           if err != nil {
                       panic(err)
                         }
           
                   read1.Header.Set("Content-Type", "application/json")
                   //read.Header.Add("Content-Type", "application/x-www-form-urlencoded")
                   read1.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZXMiOlsicmVxdWVzdCJdLCJzdWIiOiIyNmM5ZTFkMy1iYjk4LTQ4YTgtOGY2Zi0zMDc5YWZlNmYzMzEiLCJpc3MiOiJ1YmVyLXVzMSIsImp0aSI6IjA1NjU5OGRiLWI0YzgtNDdhMS05OTI4LTM0YTY1NmM2ZTUwNiIsImV4cCI6MTQ1MDU4MTE2MSwiaWF0IjoxNDQ3OTg5MTYwLCJ1YWN0IjoianFiM3dDYlNnQ0hTTWZmbk9hMXpvenpXRTdnT3pSIiwibmJmIjoxNDQ3OTg5MDcwLCJhdWQiOiJVQWJGUWtHTFFCNldFNUNhaGpNY2xzclcwZjJlaW9jQSJ9.aLjrrNeW3NQWumZ7-b16I_6w4NLp8uYwDSteF13_bydLt11FpYS_X2UH3ydsYGMKWMnOw0DKQlVOStyLY5RX3caqDDjP44T_gQV4KUORDYvmD2KZ8xrQxZQPycyCgMTMRQX06n3mnmBkgolntHMJ226I_XSTxWn_NVuYgpF4BEx0jv4zvzQAfngfaRGv_NQgrY4q2m1H2ZEXLhPR4P3Iy9S7N1utULO0A7VHde9jh81CxtU4jByYHsaDwj3TtYeHozWXdSH0vNEgwSJ4MbEkG466KA8JgR9l-PB5b58Ecw5Ex9H6D9iyjmd-hq66dgWncgdoFZ_Mg6Vi3aNWxpkdzw")
                   
                   //read.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
                   Client1:= &http.Client{}
                   resp1, _ := Client1.Do(read1)
                   //fmt.Println("StatusCode")
                   fmt.Println(resp1.StatusCode)




                                      eta_uber=mRes.(float64)                                 //eta 


        
                                     }//else ends


                                     

                                       
                                                        res.Id=p.ByName("trip_id")
                                                        res.Status="requesting"
                                                        res.Starting_from_location_id=start
                                                        res.Next_destination_location_id=end
                                                        res.Total_uber_costs=repl.Total_uber_costs
                                                        res.Total_uber_duration=repl.Total_uber_duration
                                                        res.Best_route_location_ids=repl.Best_route_location_ids
                                                        res.Total_distance=repl.Total_distance
                                                        res.Uber_wait_time_eta=eta_uber

                           } else if for_put[trip_id1]==arraying[trip_id1].count {

                                      //fmt.Println("In if else loop")
                                      res.Status="Completed"


                                      c := session.DB("location").C("Details")
       
                                              err = c.Find(bson.M{"id": arraying[trip_id1].locations[arraying[trip_id1].count]}).One(&rep1)
                                              if err != nil {
                                              panic(err)
                                              }


                                              d := session.DB("location").C("Details")
       
                                              err = d.Find(bson.M{"id": arraying[trip_id1].locations[0]}).One(&rep2)
                                              if err != nil {
                                              panic(err)
                                              }

                                               
                                               req.Start_latitude=rep1.Coordinate.Lat
                                               req.Start_longitude=rep1.Coordinate.Lng
                                               req.End_latitude=rep2.Coordinate.Lat
                                               req.End_longitude=rep2.Coordinate.Lng
                                               product_url := fmt.Sprint("https://sandbox-api.uber.com/v1/products?latitude=",rep1.Coordinate.Lat,"&longitude=",rep1.Coordinate.Lng,"&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT");
                                              
                                              response4, err := http.Get(product_url)
      
                                                if err != nil {
                                                fmt.Printf("%s", err)
                                                os.Exit(1)
                                                } else {
                                                defer response4.Body.Close()
                                                contentss, err := ioutil.ReadAll(response4.Body)
                                                if err != nil {
                                                fmt.Printf("%s", err)
                                                os.Exit(1)
                                                }
                                                var g interface{}
                                                err=json.Unmarshal(contentss, &g)
                                                //fmt.Println(g);
                                                 pRes := g.(map[string]interface{})["products"]
                                                 pRes0 := pRes.([]interface{})[0]
                                                 pEst := pRes0.(map[string]interface{})["product_id"]
                                                 //fmt.Println(pEst)
                                                 req.Product_id=pEst.(string)
                                                 }
                                               //req.Product_id="04a497f5-380d-47f2-bf1b-ad4cfdcb51f2"
                                               

                                               //POST request to get REQUEST_ID

                   //var jsonStr = []byte(`{"start_latitude":"37.334381","start_longitude":"-121.89432","end_latitude":"37.77703","end_longitude":"-122.419571","product_id":"04a497f5-380d-47f2-bf1b-ad4cfdcb51f2"}`)
                   jsonStr,_:=json.Marshal(req)

                   client := &http.Client{}
                    //urlStr:="https://sandbox-api.uber.com/v1/requests?product_id=a1111c8c-c720-46c3-8534-2fcdd730040d&start_latitude=37.791762&start_longitude=-122.706677&end_latitude=37.885114&end_longitude=-122.7066777&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT";
                   read, err := http.NewRequest("POST","https://sandbox-api.uber.com/v1/requests", bytes.NewBuffer(jsonStr)) // <-- URL-encoded payload
                           if err != nil {
                       panic(err)
                         }
           
                   read.Header.Set("Content-Type", "application/json")
                   //read.Header.Add("Content-Type", "application/x-www-form-urlencoded")
                   read.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZXMiOlsicmVxdWVzdCJdLCJzdWIiOiIyNmM5ZTFkMy1iYjk4LTQ4YTgtOGY2Zi0zMDc5YWZlNmYzMzEiLCJpc3MiOiJ1YmVyLXVzMSIsImp0aSI6IjA1NjU5OGRiLWI0YzgtNDdhMS05OTI4LTM0YTY1NmM2ZTUwNiIsImV4cCI6MTQ1MDU4MTE2MSwiaWF0IjoxNDQ3OTg5MTYwLCJ1YWN0IjoianFiM3dDYlNnQ0hTTWZmbk9hMXpvenpXRTdnT3pSIiwibmJmIjoxNDQ3OTg5MDcwLCJhdWQiOiJVQWJGUWtHTFFCNldFNUNhaGpNY2xzclcwZjJlaW9jQSJ9.aLjrrNeW3NQWumZ7-b16I_6w4NLp8uYwDSteF13_bydLt11FpYS_X2UH3ydsYGMKWMnOw0DKQlVOStyLY5RX3caqDDjP44T_gQV4KUORDYvmD2KZ8xrQxZQPycyCgMTMRQX06n3mnmBkgolntHMJ226I_XSTxWn_NVuYgpF4BEx0jv4zvzQAfngfaRGv_NQgrY4q2m1H2ZEXLhPR4P3Iy9S7N1utULO0A7VHde9jh81CxtU4jByYHsaDwj3TtYeHozWXdSH0vNEgwSJ4MbEkG466KA8JgR9l-PB5b58Ecw5Ex9H6D9iyjmd-hq66dgWncgdoFZ_Mg6Vi3aNWxpkdzw")
                   
                   //read.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

                   resp, err := client.Do(read)
                   //fmt.Println(resp.Status)
                   
                                    if err != nil {
                                      fmt.Printf("%s", err)
                                      os.Exit(1)
                                    } else {
                                      defer resp.Body.Close()
                                      contents, err := ioutil.ReadAll(resp.Body)
                                      if err != nil {
                                          fmt.Printf("%s", err)
                                          os.Exit(1)
                                      }
                                      //fmt.Printf("%s\n", string(contents))
                                      var f interface{}
                                      err=json.Unmarshal(contents, &f)
                                      mRes := f.(map[string]interface{})["eta"]

                                             mReqId := f.(map[string]interface{})["request_id"].(string)
                                              //fmt.Println(mReqId)

                                     //PUT REQUEST TO SET STATUS AT SANBOX TO COMPLETED FOR THAT REQUEST_ID

                                     puturl :=fmt.Sprint("https://sandbox-api.uber.com/v1/sandbox/requests/"+mReqId)
                                     req2.Status="completed"
                                     jsonStr1,_:=json.Marshal(req2)

                    //urlStr:="https://sandbox-api.uber.com/v1/requests?product_id=a1111c8c-c720-46c3-8534-2fcdd730040d&start_latitude=37.791762&start_longitude=-122.706677&end_latitude=37.885114&end_longitude=-122.7066777&server_token=m4v-9KBbXzZ9WxEcpBRSfC64JToSWir9mPi4fnKT";
                           read1, err := http.NewRequest("PUT",puturl, bytes.NewBuffer(jsonStr1)) // <-- URL-encoded payload
                           if err != nil {
                       panic(err)
                         }
           
                   read1.Header.Set("Content-Type", "application/json")
                   //read.Header.Add("Content-Type", "application/x-www-form-urlencoded")
                   read1.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZXMiOlsicmVxdWVzdCJdLCJzdWIiOiIyNmM5ZTFkMy1iYjk4LTQ4YTgtOGY2Zi0zMDc5YWZlNmYzMzEiLCJpc3MiOiJ1YmVyLXVzMSIsImp0aSI6IjA1NjU5OGRiLWI0YzgtNDdhMS05OTI4LTM0YTY1NmM2ZTUwNiIsImV4cCI6MTQ1MDU4MTE2MSwiaWF0IjoxNDQ3OTg5MTYwLCJ1YWN0IjoianFiM3dDYlNnQ0hTTWZmbk9hMXpvenpXRTdnT3pSIiwibmJmIjoxNDQ3OTg5MDcwLCJhdWQiOiJVQWJGUWtHTFFCNldFNUNhaGpNY2xzclcwZjJlaW9jQSJ9.aLjrrNeW3NQWumZ7-b16I_6w4NLp8uYwDSteF13_bydLt11FpYS_X2UH3ydsYGMKWMnOw0DKQlVOStyLY5RX3caqDDjP44T_gQV4KUORDYvmD2KZ8xrQxZQPycyCgMTMRQX06n3mnmBkgolntHMJ226I_XSTxWn_NVuYgpF4BEx0jv4zvzQAfngfaRGv_NQgrY4q2m1H2ZEXLhPR4P3Iy9S7N1utULO0A7VHde9jh81CxtU4jByYHsaDwj3TtYeHozWXdSH0vNEgwSJ4MbEkG466KA8JgR9l-PB5b58Ecw5Ex9H6D9iyjmd-hq66dgWncgdoFZ_Mg6Vi3aNWxpkdzw")
                   
                   //read.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
                   Client1:= &http.Client{}
                   resp1, _ := Client1.Do(read1)
                   //fmt.Println("StatusCode")
                   fmt.Println(resp1.StatusCode)




                                      eta_uber=mRes.(float64) 
                                    }

                                      res.Starting_from_location_id=arraying[trip_id1].locations[arraying[trip_id1].count]
                                      res.Next_destination_location_id=arraying[trip_id1].locations[0]
                                      res.Id=p.ByName("trip_id")
                                      
                                      res.Total_uber_costs=repl.Total_uber_costs
                                      res.Total_uber_duration=repl.Total_uber_duration
                                      res.Total_distance=repl.Total_distance
                                      res.Best_route_location_ids=repl.Best_route_location_ids
                                      res.Uber_wait_time_eta=eta_uber
                                      for_put[trip_id1]=for_put[trip_id1]+1;
                                      //fmt.Println(for_put[trip_id1]);




                           } else {


                                                              //fmt.Println("In else loop")
                                                              //fmt.Println(traversed_end_lat)
                                                              //fmt.Println(traversed_end_lng)
                                                              res.Status="Completed"
                                                              res.Starting_from_location_id=arraying[trip_id1].locations[0]
                                                              res.Next_destination_location_id=""
                                                              res.Id=p.ByName("trip_id")
                                                              res.Total_uber_costs=repl.Total_uber_costs
                                                              res.Total_uber_duration=repl.Total_uber_duration
                                                              res.Best_route_location_ids=repl.Best_route_location_ids
                                                              res.Total_distance=repl.Total_distance
                                                              res.Uber_wait_time_eta=eta_uber
                                                              //for_put[trip_id1]=0




                           }
                  

                                 //fmt.Println("Id:", id1)
                                  uj, _ := json.Marshal(res)
                                  w.Header().Set("Content-Type", "application/json")
                                  w.WriteHeader(201)
                                  fmt.Fprintf(w, "\n\n\n\n%s", uj)
                                 //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

      
  }



   
func main() {
   request_id_counter=1 
   
   r := httprouter.New()

    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: r,
    }
    r.POST("/trips",Post)
    r.GET("/trips/:id",Get)
    r.PUT("/trips/:trip_id/request",Put)
    server.ListenAndServe()


 }

