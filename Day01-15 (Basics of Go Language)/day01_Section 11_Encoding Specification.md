

# Coding Standards

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

This specification aims to provide a code specification guide for daily Go project development, facilitate the team to form a unified code style, and improve the readability, standardization and uniformity of the code. This specification will explain the naming conventions, annotation conventions, code styles and commonly used tools provided by the Go language. The specification refers to the style formulation of the official code of the go language.

## One, naming convention

Naming is a very important part of the code specification. A unified naming rule is conducive to improving the readability of the code. A good naming can get enough information just by naming.

Go starts with the letters a to Z or a to Z or an underscore when naming it, followed by zero or more letters, underscores, and numbers (0 to 9). Go does not allow the use of punctuation marks such as @, $, and% in naming. Go is a case-sensitive programming language. Therefore, Manpower and manpower are two different names.

> 1. When the name (including constants, variables, types, function names, structure fields, etc.) starts with an uppercase letter, such as: Group1, then the object using this form of identifier ** can be externally packaged code Used ** (the client program needs to import this package first), which is called export (like public in object-oriented languages);
> 2. **If the names start with a lowercase letter, they are not visible outside the package, but they are visible and available inside the entire package** (like private in object-oriented languages)



### 1. Package naming: package

Keep the package name consistent with the directory, try to use meaningful package names, short and meaningful, and try not to conflict with the standard library. The package name should be **lowercase** words, do not use underscores or mixed case.

```go
package demo

package main
```



### 2. File naming

Try to use meaningful file names, short and meaningful, should be **lowercase** words, use **underscore** to separate each word.

```go
my_test.go
```



### 3. Structure naming
-Adopt camel case nomenclature, the first letter is uppercase or lowercase according to access control

-The struct declaration and initialization format uses multiple lines, such as the following:

```go
// multi-line declaration
type User struct{
    Username string
    Email string
}

// Multi-line initialization
u := User{
    Username: "astaxie",
    Email: "astaxie@gmail.com",
}

```



### 4. Interface naming

-The basic naming rules and the above structure types
-The structure name of a single function is suffixed with "er", such as Reader, Writer.

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}

```



### 5. Variable naming

-Similar to the structure, variable names generally follow the camel case method, and the first letter is uppercase or lowercase according to the access control principle, but when encountering unique nouns, the following rules need to be followed: 
  -If the variable is private and the unique noun is the first word, use lowercase, such as apiClient
  -In other cases, the original wording of the term should be used, such as APIClient, repoID, UserID
  -Error example: UrlArray, should be written as urlArray or URLArray
-If the variable type is bool type, the name should start with Has, Is, Can or Allow

```go
var isExist bool
var hasConflict bool
var canManage bool
var allowGitHook bool
```



### 6, constant naming

Constants need to be composed of all capital letters, and use underscores to separate words

```go
const APP_VER = "1.0"
```


If it is an enumerated constant, you need to create the corresponding type first:

```go
type Scheme string

const (
    HTTP Scheme = "http"
    HTTPS Scheme = "https"
)

```



### 7, keywords

The following list shows the reserved words in Go. These reserved words cannot be used as constant or variable or any other identifier names.



![guanjianzi](http://7xtcwd.com1.z0.glb.clouddn.com/guanjianzi.jpg)



## Two, comments

Go provides C-style `/* */` block comments and C++-style `//` line comments. Line comments are the norm; block comments are mostly shown as package comments, but are useful in expressions or disable large amounts of code.

-Single-line comments are the most common form of comments, you can use single-line comments starting with // everywhere
-Multi-line comments are also called block comments. They all begin with /* and end with */, and cannot be nested. Multi-line comments are generally used for package document descriptions or commented into blocks of code snippets

The godoc tool that comes with the go language can generate documents based on comments, and generate a corresponding website (golang.org is directly generated using the godoc tool). The quality of the comments determines the quality of the generated documents. Every package should have a package comment, and a block comment before the package clause. For multi-file packages, the package comment only needs to exist in one file, any one is fine. The package review should introduce the package and provide information related to the entire package. It will first appear on the `godoc` page, and the detailed documentation below should be set.

How to write comments in detail
refer to:
   
    



### 1. Package comment

Every package should have a package comment, a block comment or a line comment before the package clause. If the package has multiple go files, they only need to appear in one go file (usually the file with the same name as the package). The package comment should contain the following basic information (please strictly follow this order, introduction, creator, creation time):

-Basic introduction of the package (package name, introduction)
-Creator, format: Creator: rtx name
-Creation time, format: Creation time: yyyyMMdd

For example, the comment example of the util package is as follows

```go
// util package, this package contains some constants shared by the project and encapsulates some shared functions in the project.
// Created by: hanru
// Creation time: 20190419
```



### 2. Structure (interface) comment

Each custom structure or interface should have a comment. The comment briefly introduces the structure and puts it on the previous line of the structure definition. The format is: structure name, structure description. At the same time, each member variable in the structure must have an explanation, which is placed after the member variable (note the alignment). Examples are as follows:

```go
// User, user object, defines the basic information of the user
type User struct{
    Username string // username
    Email string // Email
}
```



### 3. Function (method) comment

Each function or method (the function under the structure or interface is called the method) should have a comment. The comment of the function should include three aspects (written in strict accordance with this order):

-Brief description, format description: start with the function name, "," separates the description part
-Parameter list: one parameter per line, starting with the parameter name, "," separates the description part
-Return value: One return value per line

Examples are as follows:

```go
// NewtAttrModel, the factory method of the attribute data layer operation class
// Parameters:
// ctx: context information
// return value:
// Attribute operation class pointer
func NewAttrModel(ctx *common.Context) *AttrModel {
}
```



### 4. Code logic comments

For some key positions of the code logic, or some more complex logic, corresponding logic descriptions are needed to facilitate other developers to read the code. Examples are as follows:

```go
// Read attributes in batches from Redis, and record the id not read into an array, ready to read from DB
xxxxx
xxxxxxx
xxxxxxx
```



### 5. Comment style

Use Chinese comments uniformly, and strictly use spaces to separate Chinese and English characters. This is not only between Chinese and English, but also between English and Chinese punctuation. For example:

```go
// Read attributes in batches from Redis, and record the id not read into an array, ready to read from DB
```


The above Redis, id, DB and other Chinese characters are separated by spaces. 

* It is recommended to use single-line comments for all
* Like the code specification, single-line comments should not be too long and must not exceed 120 characters.



## Three, code style

### 1. Indentation and line break

-The indentation can be formatted directly with the gofmt tool (gofmt uses tab indentation);
-For line wrapping, the maximum length of a line should not exceed 120 characters. If it exceeds, please use line break to display, try to keep the format elegant.

We use the Goland development tool, and you can directly use the shortcut key: ctrl+alt+L.



### 2. The end of the statement

In Go language, there is no need to end with a colon similar to Java. The default line is a piece of data

If you plan to write multiple statements on the same line, they must use **;** 



### 3. Brackets and spaces

For brackets and spaces, you can also use the gofmt tool to format directly (go will force the opening brace to not wrap, and the newline will report a syntax error). Spaces should be left between all operators and operands.

```go
// Right way
if a> 0 {

} 

// wrong way
if a>0 // a, there should be a space between 0 and>
{// The left brace cannot be wrapped, and a syntax error will be reported

}

```



### 4. Import specification

In the case of multiple lines of import, goimports will automatically format it for you, but we still standardize some of the import specifications. If you introduce a package in a file, it is recommended to use the following format:

```go
import (
    "fmt"
)
```

If your package introduces three types of packages, standard library packages, program internal packages, and third-party packages, it is recommended to organize your package as follows:

```go
import (
    "encoding/json"
    "strings"

    "myproject/models"
    "myproject/controller"
    "myproject/utils"

    "github.com/astaxie/beego"
    "github.com/go-sql-driver/mysql"
)   
```

Import packages in an orderly manner. Different types are separated by spaces. The first is a real standard library, the second is a project package, and the third is a third-party package.

Do not use relative paths to import packages in the project:

```go
// This is a bad import
import "../net"

// This is the correct approach
import “github.com/repo/proj/src/net”
```

But if you are importing other packages in this project, it is best to use a relative path.

### 5. Error handling

-The principle of error handling is not to discard any calls that return err, do not use _ discard, you must handle all of them. Receive error, either return err, or use log to record
-Return as soon as possible: Once an error occurs, return immediately
-Try not to use panic unless you know what you are doing
-If the error description is in English, it must be lowercase, no punctuation ending is required
-Use an independent error stream for processing

```go
// wrong writing
if err != nil {
    // error handling
} else {
    // normal code
}

// correct writing
if err != nil {
    // error handling
    return // or continue, etc.
}
// normal code

```



### 6. Test

The unit test file name naming convention is example_test.go
The function name of the test case must start with Test, for example: TestExample
Each important function must first write a test case, and submit the test case together with the regular code to facilitate regression testing



## Four, commonly used tools

As mentioned above, the go language itself has made a lot of efforts in code standardization. Many restrictions are mandatory grammatical requirements. For example, the opening brace does not wrap, and the referenced package or the defined variable will report an error if it is not used. In addition, go still provides a lot of useful tools to help us standardize the code,

**gofmt**
Most of the format problems can be solved by gofmt. gofmt automatically formats the code to ensure that all go codes are consistent with the officially recommended format. Therefore, all format-related issues are subject to the results of gofmt.

**goimport**
We strongly recommend to use goimport, which adds automatic deletion and import of packages on the basis of gofmt.

```shell
go get golang.org/x/tools/cmd/goimports
```

**go vet**
The vet tool can help us statically analyze various problems in our source code, such as redundant code, the logic of returning in advance, and whether the tag of the struct meets the standard.

```shell
go get golang.org/x/tools/cmd/vet
```


Use as follows:

```go
go vet.
```



Qianfeng Go language learning group: 784190273

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation



   

