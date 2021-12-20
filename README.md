# Set

The `Set` struct allows unique values storing of any type.

## Initializing a Set

```go
s := set.New() // Empty Set
s := set.New("hi", 45, Person{name: "Gopher", birth: 2009}, false)
```

## Adding items to a Set

```go
s := set.New("hi", 45, Person{name: "Gopher", birth: 2009}, false)
s.Add(1.6823)

fmt.Println(s)
// [hi 45 {Gopher 2009} false 1.6823]
```

## Deleting items from a Set

```go
s := set.New("hi", 45, Person{name: "Gopher", birth: 2009}, false)
s.Delete(45)

fmt.Println(s)
// [hi {Gopher 2009} false]
```

## Checking if an item exists

```go
s := set.New("hi", 45, Person{name: "Gopher", birth: 2009}, false)
exist, index := s.Has(Person{name: "Gopher", birth: 2009})
fmt.Println(exist, index)
// true 2

exist, index := s.Has("test")
fmt.Println(exist, index)
// false -1
```

## Iterating a Set

```go
s := set.New("hi", 45, Person{name: "Gopher", birth: 2009}, false)

for s.Next() {
    fmt.Println(s.Get())
}

// hi
// 45
// {Gopher 2009}
// false
```

## Reseting a Set

```go
s := set.New("hi", 45, Person{name: "Gopher", birth: 2009}, false)
s.Clear()
```
