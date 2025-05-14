# Packages in GO

### WHY USE PACKAGES?
- Two packages can have functions with same name but different functionality
- Better structuring and organizing
- 

## WHAT IS A MODULE?
- A module is a collection of related Go packages with versioning.

### HOW TO INITIALIZE A GO MODULE?
-   ```bash
      go mod init github.com/nakul-krishnakumar/<proj-name>
    ```
-  It is prefered to use github repo link as the module name
-  Usually 1 PROJECT = 1 MODULE

## HOW TO INSTALL THIRD-PARTY PACKAGES?
- Just like ```npm install``` , go uses ```go get```
- To see available third-party packages, visit https://pkg.go.dev/
- eg: 

      go get github.com/gin-gonic/gin

- To remove unused packages, do:

      go mod tidy

   - This command also installs packages which are imported in files but not yet installed

- ```go.sum``` is like the ```package-lock.json``` file


## FUNCTION SCOPE IN GO PACKAGES

| Function Name | Visibility           | Example                                 |
| ------------- | -------------------- | --------------------------------------- |
| `Add()`       | Exported (public)    | Accessible outside its package          |
| `add()`       | Unexported (private) | Only accessible within the same package |

✅ In Go, capitalization determines visibility. There’s no public or private keyword.

Exact same scoping rules apply to variables, constants, types, and even structs in Go. Let me explain:
