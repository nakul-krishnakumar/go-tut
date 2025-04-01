# GOLang Tutorial - Hello World

- ```main.go``` works like ```index.js``` ( doesn't always have to be ```main.go``` , can be named anything, naming it ```main``` is just a convention)
- ```main.go``` will have a ```main``` function is the starting point of execution, just like in C/C++
- GO uses a ```fmt``` package for basic functions like print, println
- To run a GO file, do 
   - ```go build main.go``` , this builds an executable file ```main```
   - now, do ```./main``` , this will run the executable file

- OR else, to run a GO file without making a exe everytime, do ```go run main.go```


<details open>
   <summary>

     GO doesn't require ```;```

   </summary>


   **No, you generally do *not* need to put semicolons (`;`) at the end of statements in Go.**

   Here's why:

   1.  **Automatic Semicolon Insertion:** The Go compiler automatically inserts semicolons at the end of most lines. It follows simple rules: if a line ends with something that looks like the end of a statement (like an identifier, a literal value, `return`, `break`, `continue`, `++`, `--`, `)`, `]`, or `}`), the compiler inserts a semicolon right before the newline character.

   2.  **Idiomatic Go:** Writing Go code *without* semicolons at the end of lines is the standard, idiomatic way. You'll almost never see them in typical Go codebases.

   **When *Do* You Use Semicolons?**

   There are a couple of specific situations where you *will* use semicolons:

   1.  **Separating multiple statements on a single line:** While strongly discouraged for readability, you *can* put multiple simple statements on one line by separating them with a semicolon.
      ```go
      x := 5; fmt.Println(x) // Works, but generally avoid this style
      ```
      It's much better to write:
      ```go
      x := 5
      fmt.Println(x)
      ```

   2.  **In `for` loop headers:** The classic `for` loop structure requires semicolons to separate the initialization, condition, and post-iteration statements.
      ```go
      // init ; condition ; post
      for i := 0; i < 10; i++ {
         fmt.Println(i)
      }
      ```
      Even if you omit parts, the semicolons might still be needed as separators:
      ```go
      sum := 1
      for ; sum < 1000; { // Semicolons needed to show empty init and post
         sum += sum
      }
      ```

   **In summary:**

   *   **Don't** put semicolons at the end of your lines. Let the compiler handle it.
   *   **Do** use semicolons only when separating multiple statements on a *single* line (rare and usually bad style) or within the header of a `for` loop.

   </details>
   