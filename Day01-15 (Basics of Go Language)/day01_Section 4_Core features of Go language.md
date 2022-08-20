

# Core features of Go language

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

Go language, as the offspring of a programming language, stands on the shoulders of giants and absorbs the characteristics of some other programming languages.

The Go programming language is an open source project that makes programmers more productive. Go language has strong expressive ability, it is concise, clear and efficient. Thanks to its concurrency mechanism, programs written with it can use multi-core and networked computers very effectively, and its novel type system makes the program structure flexible and modular. Compiling Go code into machine code is not only very fast, but also has a convenient garbage collection mechanism and a powerful runtime reflection mechanism. It is a fast, statically typed compiled language, but it feels like a dynamically typed interpreted language. (Taken from the official website)



## One, thought

Less can be more

Dao to Jane, small but true

It's easy to make things complicated, but it's hard to make things simple

Profound engineering culture



## Two, core features

The reason why Go language is powerful is that it can always grasp the pain points of programmers in the development of the server, and solve problems in the most direct, simple, efficient, and stable way. Here we will not discuss the specific syntax of the GO language in depth, but will only introduce the key aspects of the language that are important for simplifying programming to experience the core features of Go.

### 2.1 Concurrent programming

Go language is much more concise than most languages ​​in terms of concurrent programming. This is one of its biggest highlights, and it is also an important bargaining chip for its entry into high-concurrency and high-performance scenarios in the future.

![bingfa1](img/bingfa1.jpg)

Different from traditional multi-process or multi-thread, the concurrent execution unit of golang is a kind of coroutine called goroutine.

Because locks are used in shared data scenarios, coupled with GC, its concurrency performance is sometimes not as good as the asynchronous multiplexing IO model, so compared to most languages, golang's concurrent programming is simpler than concurrent performance and has a selling point.



In today's multi-core era, the significance of concurrent programming is self-evident. Of course, many languages ​​support multi-threaded, multi-process programming, but unfortunately, it is not so easy and pleasant to implement and control. The difference in Golang is that the language level supports goroutine concurrency (coroutines are also called microthreads, which are lighter, less expensive, and have higher performance than threads), and they are very simple to operate. The language level provides keywords (go) Used to start coroutines, and thousands of coroutines can be started on the same machine. Coroutines are often understood as lightweight threads. A thread can contain multiple coroutines, sharing the heap but not the stack. Coroutines are generally scheduled explicitly by the application, and context switching does not need to be down to the kernel layer, which is much more efficient. Synchronous communication is generally not done between coroutines, and there are two types of communication between coroutines in golang: 1) Shared memory type, that is, using global variables + mutex locks to achieve data sharing; 2) Message passing type, that is, using a unique type Some channel mechanisms carry out asynchronous communication.

Comparing JAVA's multi-threading and GO's coroutine implementation, it is obviously more direct and simple. This is where the charm of GO lies. It solves problems in a simple and efficient way. The keyword go is perhaps the most important sign of the GO language.

**High concurrency is the biggest highlight of Golang language**



### 2.2 Memory Reclamation (GC)

From C to C++, from the perspective of program performance, these two languages ​​allow programmers to manage memory by themselves, including memory application and release. Because there is no garbage collection mechanism, C/C++ runs very fast, but with it comes programmers' careful consideration of memory usage. Because even a little carelessness may cause "memory leaks" to waste resources or "wild pointers" to crash programs, etc. Although C++11 later used the concept of smart pointers, programmers still need to use them carefully. Later, in order to improve the speed of program development and the robustness of the program, high-level languages ​​such as java and C# introduced the GC mechanism, that is, programmers no longer need to consider memory recovery, etc., but the language feature provides a garbage collector to recover memory. However, the following may be a reduction in the efficiency of the program.

The GC process is: first stop the world, scan all objects and judge them to be alive, mark the recyclable objects in a bitmap area, then immediately start the world, resume the service, and set up a special gorountine to reclaim the memory to the free list for recovery Use, not physically release. The physical release is performed periodically by a dedicated thread.

The bottleneck of GC is that all objects must be scanned every time to judge alive. The more objects to be collected, the slower the speed. An empirical value is that it takes 1ms to scan 10w objects, so try to use a solution with fewer objects. For example, we also consider linked list, map, slice, and array for storage. Each element of linked list and map is an object, while slice or array It is an object, so slice or array is good for GC.

GC performance may continue to be optimized as the version is continuously updated. This area has not been carefully investigated. There are HotSpot developers in the team, and they should learn from the design ideas of jvm gc, such as generational recycling, safepoint, etc.

-Memory is automatically reclaimed, no more developers need to manage memory
-Developers focus on business realization, reducing mental burden
-Only need new to allocate memory, no need to release




### 2.3 Memory allocation

In the initialization phase, a large memory area is directly allocated. The large memory is divided into blocks of various sizes and placed in different free lists. When the object allocates space, a memory block of appropriate size is taken from the free list. When memory is reclaimed, unused memory will be replayed back to the free list. Free memory will be merged according to a certain strategy to reduce fragmentation.



### 2.4 Compile

Compilation involves two issues: compilation speed and dependency management

Golang currently has two compilers, one is Gccgo based on GCC, and the other is a set of compilers (6g and 8g) for 64-bit x64 and 32-bit x86 computers.

In terms of dependency management, since most of the third-party open source libraries of golang are on github, you can use them by adding the corresponding github path to the import of the code, and the library will be downloaded to the pkg directory of the project by default.

In addition, the usage of all entities in the code will be checked by default when compiling. Any package or variable that is not used will fail to compile. This is the rigorous side of golang.



### 2.5 Network Programming

Since golang was born in the Internet era, it is inherently decentralized and distributed. One of the specific manifestations is to provide a rich and convenient network programming interface. For example, the socket uses net.Dial (based on tcp/udp, which encapsulates the traditional Connect, listen, accept and other interfaces), http use http.Get/Post(), rpc use client.Call('class_name.method_name', args, &reply), etc.

> High-performance HTTP Server



### 2.6 Functions with multiple return values

In C, C++, including some other high-level languages, multiple function return values ​​are not supported. But this function is really needed, so in the C language, the return value is generally defined as a structure or returned in the form of function parameter references. In the Go language, as a new type of language, a language with a strong target can of course not give up the satisfaction of this demand, so it is necessary to support multiple return values ​​of functions.

When defining a function, you can add (a, b, c) after the input parameter, which means that there will be 3 return values ​​a, b, and c. This feature is available in many languages, such as python.

This syntactic sugar feature is of practical significance. For example, we often require the interface to return a triple (errno, errmsg, data). In most languages ​​that only allow one return value, we can only put the triple into When a map or array is returned, the receiver must also write code to check that the return value contains triples. If multiple return values ​​are allowed, it is directly enforced at the function definition level to make the code more concise and safe.





### 2.7 Language Interactivity

Language interactivity refers to whether the language can interact with other languages, such as the ability to call libraries compiled by other languages.

Most of the C modules are directly reused in the Go language, here called Cgo. Cgo allows developers to mix and write C language codes, and then the Cgo tool can extract these mixed C codes and generate call packaging code for C functions. Developers can basically completely ignore how this boundary between the Go language and the C language is crossed.

 Golang can interact with C programs, but not C++. There are two alternatives: 1) first compile c++ into a dynamic library, and then call a piece of c code by go, which dynamically calls the dynamic library through the dlfcn library (remember to export LD_LIBRARY_PATH); 2) use swig (not played)

### 2.8 Exception handling

Golang does not support structured exception resolution methods such as try...catch, because it feels that it will increase the amount of code and be abused. No matter how small an exception is, it will be thrown. The exception handling methods advocated by golang are:

-Ordinary exception: The callee returns an error object, and the caller judges the error object.
-Severe exception: Refers to interrupt panic (such as dividing by 0), use defer...recover...panic mechanism to capture processing. Critical exceptions are generally automatically thrown inside golang, and do not require users to actively throw them, avoiding the situation where the traditional try...catch is written everywhere. Of course, users can also use panic('xxxx') to actively throw, but this will degenerate this set of mechanisms into a structured exception mechanism.



### 2.9 Some other interesting features

-Type inference: Type definition: Supports the syntax of `var abc = 10`, which makes golang look a bit like a dynamically typed language, but golang is actually strongly typed, and the previous definition will be automatically deduced to be of type int.

  > As a strongly typed language, implicit type conversion is not allowed. Remember a principle: Let everything be explicit.
  >
  > Simply put, Go is a static language that is written like a dynamic language and has the efficiency of dynamic language development.

-As long as a type implements all the methods of an interface, the interface can be implemented without explicit inheritance.

  > The Go programming specification recommends that each Interface provides only one or two methods. This makes the purpose of each interface very clear. In addition, Go's implicit derivation also makes us more flexible in organizing program architecture. When writing JAVA/C++ programs, we need to design the parent class/subclass/interface well at the beginning, because once there are changes later, it will be very painful to modify. But Go is different. When you find that some methods can be abstracted into an interface during the implementation process, you can directly define the interface and it is OK. Other codes do not need to be modified. The automatic derivation of the compiler will do it for you. Good everything.

-No circular reference: If b is imported in a.go, if b.go imports a, it will report import cycle not allowed. The advantage is that some potential programming hazards can be avoided. For example, func1() in a calls func2() in b. If func2() can also call func1(), it will lead to an infinite loop.

-Defer mechanism: In the Go language, the keyword defer is provided, which can be used to specify the logic body that needs to be delayed, that is, it will be executed before the function body returns or when panic appears. This mechanism is very suitable for aftermath logic processing, such as avoiding possible resource leakage problems as early as possible.

  It can be said that defer is another very important and practical language feature after goroutine and channel. The introduction of defer can simplify programming to a large extent, and it appears more natural in language description and greatly enhanced The readability of the code.

-The concept of "package": Just like python, put the code of the same function in a directory and call it a package. Packages can be referenced by other packages. The main package is used to generate executable files, and each program has only one main package. The main purpose of the package is to improve the reusability of the code. Other packages can be introduced through package.

-Programming specifications: The programming specifications of the GO language are mandatory to be integrated into the language, such as clearly specifying the placement of curly braces, mandatory requirements for one line, no import of unused packages, no definition of unused variables, and gofmt tool mandatory format Code and so on. The strange thing is that these have also aroused dissatisfaction among many programmers. Some people have published XX crimes in the GO language, and there is no lack of accusations against programming standards. You know, from the perspective of project management, any development team will formulate specific programming specifications for a specific language, especially for companies like Google. The designers of GO believe that instead of writing the specification in the document, it is better to force the integration in the language, which is more direct and makes use of team collaboration and project management.

-Cross-compilation: For example, you can develop and run applications running under Windows on a computer running Linux. This is the first programming language that fully supports UTF-8. This is not only reflected in the fact that it can handle strings encoded in UTF-8, but even its source code file format uses UTF-8 encoding. Go language has achieved true internationalization!









## Three, function

Here we talk about a short paragraph: (quoted from an article on the Internet, which one can’t be remembered specifically)

A long time ago, there was an IT company that had a tradition of allowing employees to have 20% free time to develop experimental projects. On a certain day in 2007, several of the company’s big cows were using C++ to develop some cumbersome but core tasks, mainly including huge distributed clusters. Daniel found it very annoying. Later, the C++ committee came to their company to give a speech and said C++ About 35 new features will be added. One of these big cows is named Rob Pike. After hearing this, 10,000 xxx floated in his heart, "Isn't there enough C++ features? Simplifying C++ should be more fulfilling, right". Since then, Rob Pike and several other big cows discussed how to solve this problem. After a while, Rob Pike said that we should develop a language by ourselves. The name is "go", which is very short and easy to spell. The other big cows just said yes, and then they found a whiteboard and wrote down what functions they hoped to have. In the following time, the big cows happily discussed the characteristics of the design of this language. After a long time, they decided to use the C language as a prototype and learn from some features of other languages ​​to liberate programmers, liberate themselves, and then In 2009, the go language was born.

The following are the features of Go listed by these big cows:

-Standard grammar (no symbol table is needed for parsing)
-Garbage collection (exclusive)
-Headless file
-Clear dependency
-No circular dependency
-Constants can only be numbers
-int and int32 are two types
-Letter case sets visibility
-Any type (type) has a method (not a type)
-No subtype inheritance (not subtype)
-Package level initialization and clear initialization sequence
-The file is compiled into a package
-Package-level globals presented in any order
-No numeric type conversion (constant plays an auxiliary role)
-Implicit implementation of the interface (no "implement" declaration)
-Embedding (will not be promoted to super class)
-The method is declared in accordance with the function (no special location requirements)
-Methods are functions
-The interface has only methods (no data)
-Methods match by name (not type)
-No constructor and destructor
-postincrement (such as ++i) is a state, not an expression
-No preincrement (i++) and predecrement
-Assignment is not an expression
-Definite assignment and calculation sequence in function calls (no "sequence point")
-No pointer arithmetic
-The memory is always initialized with zero value
-Local variable value is legal
-There is no "this" in the method
-Segmented stack
-No static and other types of comments
-No template
-Built-in string, slice and map
-Array boundary check



> ## The real body of Daniel
>
> The biggest brand is definitely Ken Thompson, the designer of B and C languages, the founder of Unix and Plan 9, and the winner of the Turing Award in 1983. This list also includes Unix core members Rob Pike (the father of the go language), java Robert Griesemer, the developer of the HotSpot virtual machine and js v8 engine, Brad Fitzpatrick, the author of Memcached, etc.





The picture in this article is from the Internet, invaded





Qianfeng Go language learning group: 784190273

Corresponding video address:

https://www.bilibili.com/video/av47467197

https://www.bilibili.com/video/av56018934/

Source code:

https://github.com/rubyhan1314/go_foundation

