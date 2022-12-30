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
 
func main() {
   conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())
 
   if err != nil {
       log.Fatalf("did not connect : %v", err)
   }
 
   defer conn.Close()
 
   c := pb.NewTodoServiceClient(conn)
 
   ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 
   defer cancel()
 
   filterTasks := []FilterTask{
       {Input: "Code review"},
       {Input: "Make YouTube Video"},
       {Input: "Go to the gym"},
       {Input: "Buy groceries"},
       {Input: "Meet with mentor"},
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