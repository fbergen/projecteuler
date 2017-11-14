package main

import "fmt"

func next(i int) int {
  next := 0
  for i > 0 {
    d := i % 10
    next += d*d
    i = i / 10
  }
  return next;
}


func converges(i int, m map[int]int) int {
  in := i
  if in < 560 && m[in] != 0 {
    //fmt.Println("Cache hit!", m[in])
    return m[in]
  }
  if in == 1 || in == 89 {
    return in
  } else {
    a := converges(next(in), m)
    if in < 560 {
      m[in] = a
    }
    return a;
  }
}

func main() {
  m := make(map[int]int)
  counter := 0;
  for i := 1; i <= 10000000; i++ {
    c := converges(i, m);
    if c == 89 {
      counter++;
    }
  }
  fmt.Println("counter", counter);
}
