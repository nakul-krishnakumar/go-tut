# GOLang Tutorial - For Loop

- GO doesn't have ```while loop``` , the only looping technique in GO is ```for loop```.
- ```while``` can be implemented as follows:
   ```go
   // while loop implementation
	i := 1
	for i <= 3  {
		fmt.Println(i)
		i += 1
	}

   ```

- ```for``` loop implementation
   ```go
   // for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
   ```

- ```for``` loop with ```range``` 
   ```go
   // for loop using range
	for i := range 3 {
		fmt.Println(i , " in range")
	}
   ```