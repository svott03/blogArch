package main
 
import (
   "context"
   "log"
   "time"
 
   pb "example.com/blogArch/proto"
   "google.golang.org/grpc"
)
 
const (
   ADDRESS = "localhost:50051"
)
 
type FilterTask struct {
   Input        string
}
 
// TODO add client file gateway
func main() {
   conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())
 
   if err != nil {
       log.Fatalf("did not connect : %v", err)
   }
 
   defer conn.Close()
 
   c := pb.NewTextFilterServiceClient(conn)
 
   ctx, cancel := context.WithTimeout(context.Background(), 80*time.Second)
 
   defer cancel()
 
   filterTasks := []FilterTask{
       {Input: "Buy groceries"},
       {Input: "Meet with mentor"},
			 {Input: "You are awful"},
			 {Input: "I don\\'t like that"},
   }
 
   for _, task := range filterTasks {
       res, err := c.CreateFilterOutput(ctx, &pb.FilterInput{Input: task.Input})

       if err != nil {
           log.Fatalf("could not create user: %v", err)
       }
 
       log.Printf(`
           Output : %s
       `, res.GetOutput())
   }
 
}