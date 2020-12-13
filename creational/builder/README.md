# Computer example
## Immutability
In this example our builder contains the `Car` as a member. It gives us the advantage of 
not writing all the car properties into the builder class. However, since we dont create a new car object
in the `build` method, it always returns the copy of the object with the old values as long as we dont create a new Builder instance

Changing the struct like it would be even worse.

````go
type ComputerBuilder struct {
	computer *Computer
}
```` 

Think that we build the struct `computerBuilder`. After that we change one property and build it again.
In this case you would change the original `computer` struct as it is a pointer. That causes the change of `Brand`
property of `computerBuilder` too.
  
```go
    computerBuilder := ComputerBuilder{}
    computer := computerBuilder.ChooseBrand("IBM").ChooseMonitor("14 Inch").ChooseProcessor("Intel").ChooseStorageSize(456).Build()
    
    c := computerBuilder.ChooseBrand("Dell").Build()
    fmt.Println(c.Brand)
    fmt.Println(c.Monitor) 
```

## Director
In this example, there is no director structure as its responsibility is delegated to the client. As you can see it, 
we ultimately only set some properties of a struct. So there is no logic in the construction process. So it also doesnt
make sense to create another layer for `Director` as combinations of the properties are quite a lot.    

 