## CTL Labeling Algorithm

#### Usage
```
$ go run main.go formula.go transition.go
```

#### Output of example
```
CTL formula: E[¬c2Uc1]
== RESULT ==
0 [SATISFY]:     n1 n2 true ¬c2 E[¬c2Uc1]
1 [SATISFY]:     t1 n2 true ¬c2 E[¬c2Uc1]
2 [SATISFY]:     t1 t2 true ¬c2 E[¬c2Uc1]
3 [SATISFY]:     c1 t2 true ¬c2 E[¬c2Uc1]
4 [SATISFY]:     c1 n2 true ¬c2 E[¬c2Uc1]
5 [NOT SATISFY]: n1 t2 true ¬c2
6 [NOT SATISFY]: t1 t2 true ¬c2
7 [NOT SATISFY]: t1 c2 true
8 [NOT SATISFY]: n1 c2 true
```