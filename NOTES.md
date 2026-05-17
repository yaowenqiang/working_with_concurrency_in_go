> go mod init first-example
> go test -v .
> go test main2_test.go -v

Mutex = "mutual exclusion"
> go run -race main3.go

The Dining Philosophers(哲学家就餐问题)

+ A classic computer problem introduced by Dijkstra in 1965
+ Five Philosophers live in a house together, and they always dine together at the same table, sitting in the same place
+ They always eat a apecial kind of spaghetti which requires tow forks 
+ There are two forks next to each plate, which menas that no two neighours can be eating at the same thime


Channels 

+ A menas of allowing communication to and from a GoRoutine
+ Channels can be buffered, or unbuffered
+ Once you're done with a channel, you must close it
+ Channels typically only accepted a given type of interface

The Sleeping Barber

+ A classic computer science problem introduced by Dijkstra in 1965
+ A barber goes to work in a barbershop with a waiting room with a fixed number of seats
+ If no one is in the waiting room, the barber goes to sleep
+ When a client shows up, if there are no seats available, he or she leaves.
+ If there is a seat available, and the barber is sleeping, the client wakes the baber up and gets a hair cut
+ if the barber is busy, the client takes a seat and waits his or her turn
+ Once the shop closes, no more clients are allowed in, but the barber has to stay until eveyone who is waiting gets a hair cut



