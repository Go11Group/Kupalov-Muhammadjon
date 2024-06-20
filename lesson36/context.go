package main

import (
	"context"
	"fmt"
)

func main(){
	// AfterFunc 
	// WithDeadline
	// WithTimeoutCause
	// WithoutCancel
	// stp := context.AfterFunc(context.Background(), greet)
	// ctx, cancel := context.WithDeadlineCause(context.Background(), time.Now().Add(2*time.Second), context.DeadlineExceeded)

	defer cancel()
	greet(ctx)
	
}

func greet(ctx context.Context){
	select{
	case <- ctx.Done():
		fmt.Println("Bye Bye")
		return
	default:
		fmt.Println("Hello man")
	}
}