# func-util
`func-util` is a go language function library used to provide Java like streaming data processing.

# example

`Java Stream filter`
```Java
List<Integer> list = new ArrayList<>();
list.add(1)
list.add(2)
list.add(3)
list.add(4)
list = list.stream().filter(i -> i % 2 == 0).toList();
```

`Go func-util Filter`

```go
pendingTestSlice := make([]int, 0)
pendingTestSlice = append(pendingTestSlice, 1)
pendingTestSlice = append(pendingTestSlice, 2)
pendingTestSlice = append(pendingTestSlice, 3)
pendingTestSlice = append(pendingTestSlice, 4)

realResult := util.Filter(pendingTestSlice, func(i int) bool {
    return i%2 == 0
})
```