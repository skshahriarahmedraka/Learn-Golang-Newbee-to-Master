

# One, File file operation

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.



First of all, the file class is in the os package, encapsulating the underlying file descriptor and related information, and encapsulating the implementation of Read and Write.

## 1. FileInfo interface

The methods related to File information are defined in the FileInfo interface.

```go
type FileInfo interface {
	Name() string // base name of the file file name. Extension aa.txt
	Size() int64 // File size, 12540 bytes
	Mode() FileMode // File permissions -rw-rw-rw-
	ModTime() time.Time // Modification time 2018-04-13 16:30:53 +0800 CST
	IsDir() bool // Whether it is a folder
	Sys() interface() // Basic data source interface (can return nil)
}
```



## 2. Permission

As for the operation permission perm, it is only necessary to specify it when creating a file, and it can be set to 0 when you do not need to create a new file. Although the Go language sets many constants for perm permissions, it is customary to use numbers directly, such as 0666 (the specific meaning is consistent with the Unix system).

Access control:

```go
There are two file permission representation methods under linux, namely "symbol representation" and "octal representation".

(1) Symbolic representation:
---- --- ---
type owner group others
File permissions are allocated in this way. Read and write executables correspond to rwx. If you don’t have that permission, use-instead
(-File d directory|connect symbol)
For example: -rwxr-xr-x


(2) Octal representation:	
r ——> 004
w ——> 002
x ——> 001
-——> 000

0755
0777
0555
0444
0666
```



## 3. Open mode

File open mode:

```go
const (
    O_RDONLY int = syscall.O_RDONLY // Open the file in read-only mode
    O_WRONLY int = syscall.O_WRONLY // Open the file in write-only mode
    O_RDWR int = syscall.O_RDWR // Open the file in read-write mode
    O_APPEND int = syscall.O_APPEND // append data to the end of the file when writing
    O_CREATE int = syscall.O_CREAT // If it does not exist, a new file will be created
    O_EXCL int = syscall.O_EXCL // Used in conjunction with O_CREATE, the file must not exist
    O_SYNC int = syscall.O_SYNC // Open the file for synchronous I/O
    O_TRUNC int = syscall.O_TRUNC // If possible, clear the file when opening
)
```



## 4, File operation

```go
type File
//File represents an open file object.

func Create(name string) (file *File, err error)
//Create uses mode 0666 (anyone can read and write, not executable) to create a file named name, if the file already exists, it will be truncated (empty file). If successful, the returned file object can be used for I/O; the corresponding file descriptor has O_RDWR mode. If an error occurs, the underlying type of error is *PathError.

func Open(name string) (file *File, err error)
//Open opens a file for reading. If the operation is successful, the method of the returned file object can be used to read the data; the corresponding file descriptor has the O_RDONLY mode. If an error occurs, the underlying type of error is *PathError.

func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
//OpenFile is a more general file opening function, most callers use Open or Create instead of this function. It will use the specified options (such as O_RDONLY, etc.), the specified mode (such as 0666, etc.) to open the file with the specified name. If the operation is successful, the returned file object can be used for I/O. If an error occurs, the underlying type of error is *PathError.

func NewFile(fd uintptr, name string) *File
//NewFile creates a file using the given Unix file descriptor and name.

func Pipe() (r *File, w *File, err error)
//Pipe returns a pair of associated file objects. Reading from r will return the data written to w. This function will return two file objects and possible errors.

func (f *File) Name() string
//Name method returns the file name (provided to Open/Create and other methods).

func (f *File) Stat() (fi FileInfo, err error)
//Stat returns the FileInfo type value describing the file f. If an error occurs, the underlying type of error is *PathError.

func (f *File) Fd() uintptr
//Fd returns the integer type Unix file descriptor corresponding to the file f.

func (f *File) Chdir() error
//Chdir changes the current working directory to f, f must be a directory. If an error occurs, the underlying type of error is *PathError.

func (f *File) Chmod(mode FileMode) error
//Chmod modifies the mode of the file. If an error occurs, the underlying type of error is *PathError.

func (f *File) Chown(uid, gid int) error
//Chown modifies the user ID and group ID of the file. If an error occurs, the underlying type of error is *PathError.

func (f *File) Close() error
//Close closes the file f, so that the file cannot be used for reading or writing. It returns possible errors.


func (f *File) Readdir(n int) (fi []FileInfo, err error)
//Readdir reads the contents of the directory f and returns a []FileInfo with n members. These FileInfos are returned by Lstat in the order of directories. The next call to this function will return the information of the remaining unread content from the previous call. If n>0, the Readdir function returns a slice of up to n members. At this time, if Readdir returns an empty slice, it will return a non-nil error explaining the reason. If the end of the directory f is reached, the return value err will be io.EOF. If n<=0, the Readdir function returns a slice composed of FileInfo of all file objects remaining in the directory. At this point, if the Readdir call succeeds (read everything until the end), it will return the slice and the error value of nil. If an error is encountered before the end is reached, the slice composed of the previously successfully read FileInfo and the error will be returned.

func (f *File) Readdirnames(n int) (names []string, err error)
//Readdir reads the contents of the directory f and returns a []string with n members. The slice member is the name of the file object in the directory, in directory order. The next call to this function will return the information of the remaining unread content from the previous call. If n>0, the Readdir function returns a slice of up to n members. At this time, if Readdir returns an empty slice, it will return a non-nil error explaining the reason. If the end of the directory f is reached, the return value err will be io.EOF. If n<=0, the Readdir function returns a slice composed of the names of all remaining file objects in the directory. At this point, if the Readdir call succeeds (read everything until the end), it will return the slice and the error value of nil. If an error is encountered before the end is reached, the slice composed of the previously successfully read name and the error will be returned.

func (f *File) Truncate(size int64) error
//Truncate changes the size of the file, it does not change the current position of the I/O. If the file is truncated, the extra part will be discarded. If an error occurs, the underlying type of error is *PathError.



```



## 5. Sample code

File information: FileInfo

```go
package main

import (
	"os"
	"fmt"
)

func main() {
	/*
	FileInfo: file information
		interface
			Name(), file name
			Size(), file size, in bytes
			IsDir(), is it a directory
			ModTime(), modification time
			Mode(), permission

	 */
	fileInfo,err := os.Stat("/Users/ruby/Documents/pro/a/aa.txt")
	if err != nil{
		fmt.Println("err :",err)
		return
	}
	fmt.Printf("%T\n",fileInfo)
	//file name
	fmt.Println(fileInfo.Name())
	//File size
	fmt.Println(fileInfo.Size())
	//Is it a directory
	fmt.Println(fileInfo.IsDir()) //IsDirectory
	//Change the time
	fmt.Println(fileInfo.ModTime())
	//Permission
	fmt.Println(fileInfo.Mode()) //-rw-r--r--
}

```



operation result:

![fileyunxing1](img/fileyunxing1.png)



File operation example:

```go
package main

import (
	"fmt"
	"path/filepath"
	"path"
	"os"
)

func main() {
	/*
	File operations:
	1. Path:
		Relative path: relative
			ab.txt
			Relative to the current project
		Absolute path: absolute
			/Users/ruby/Documents/pro/a/aa.txt

		.Current directory
		..Up one level
	2. Create a folder, if the folder exists, the creation fails
		os.MkDir(), create a layer
		os.MkDirAll(), can create multiple layers

	3. Create a file, Create uses mode 0666 (anyone can read and write, not executable) to create a file named name, if the file already exists, it will be truncated (empty file)
		os.Create(), create a file

	4. Open the file: Let the current program establish a connection with the specified file
		os.Open(filename)
		os.OpenFile(filename,mode,perm)

	5. Close the file: The link between the program and the file is broken.
		file.Close()

	5. Delete files or directories: use with caution, use with caution, then use with caution
		os.Remove(), delete files and empty directories
		os.RemoveAll(), delete all
	 */
	 //1. Path
	 fileName1:="/Users/ruby/Documents/pro/a/aa.txt"
	 fileName2:="bb.txt"
	 fmt.Println(filepath.IsAbs(fileName1)) //true
	 fmt.Println(filepath.IsAbs(fileName2)) //false
	 fmt.Println(filepath.Abs(fileName1))
	 fmt.Println(filepath.Abs(fileName2)) // /Users/ruby/go/src/l_file/bb.txt

	 fmt.Println("Get parent directory:",path.Join(fileName1,".."))

	 //2. Create a directory
	 //err := os.Mkdir("/Users/ruby/Documents/pro/a/bb",os.ModePerm)
	 //if err != nil{
	 // fmt.Println("err:",err)
	 // return
	 //}
	 //fmt.Println("The folder was created successfully...")
	 //err :=os.MkdirAll("/Users/ruby/Documents/pro/a/cc/dd/ee",os.ModePerm)
	 //if err != nil{
	 // fmt.Println("err:",err)
	 // return
	 //}
	 //fmt.Println("Multi-layer folder created successfully")

	 //3. Create a file: Create uses mode 0666 (anyone can read and write, not execute) to create a file named name, if the file already exists, it will be truncated (empty file)
	 //file1,err :=os.Create("/Users/ruby/Documents/pro/a/ab.txt")
	 //if err != nil{
	 // fmt.Println("err:",err)
	 // return
	 //}
	 //fmt.Println(file1)

	 //file2,err := os.Create(fileName2)//Create a file with a relative path, based on the current project
	 //if err != nil{
	 // fmt.Println("err :",err)
	 // return
	 //}
	 //fmt.Println(file2)

	 //4. Open the file:
	 //file3 ,err := os.Open(fileName1) //read-only
	 //if err != nil{
	 // fmt.Println("err:",err)
	 // return
	 //}
	 //fmt.Println(file3)
	/*
	The first parameter: file name
	The second parameter: how to open the file
		const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR int = syscall.O_RDWR // open the file read-write.
		// The remaining values ​​may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT // create a new file if none exists.
		O_EXCL int = syscall.O_EXCL // used with O_CREATE, file must not exist.
		O_SYNC int = syscall.O_SYNC // open for synchronous I/O.
		O_TRUNC int = syscall.O_TRUNC // truncate regular writable file when opened.
	)
	The third parameter: file permissions: the file does not exist to create a file, you need to specify permissions
	 */
	 //file4,err := os.OpenFile(fileName1,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
	 //if err != nil{
	 // fmt.Println("err:",err)
	 // return
	 //}
	 //fmt.Println(file4)

	 //5 close the file,
	 //file4.Close()

	 //6. Delete files or folders:
	 //Delete Files
	//err := os.Remove("/Users/ruby/Documents/pro/a/aa.txt")
	//if err != nil{
	// fmt.Println("err:",err)
	// return
	//}
	//fmt.Println("The file was deleted successfully...")
	//Delete directory
	err := os.RemoveAll("/Users/ruby/Documents/pro/a/cc")
	if err != nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Println("Delete the directory successfully...")
}

```



# Two, I/O operation

I/O operations are also called input and output operations. Among them, I refers to Input, and O refers to Output, which is used to read or write data. In some languages, it is also called stream operation, which refers to the data communication channel.

The Golang standard library's abstraction of IO is very delicate, and each component can be combined at will, which can be used as a model of interface design.

## 1. io package

The io package provides a series of interfaces for I/O primitive operations. It mainly wraps some existing implementations, such as those in the os package, and abstracts these into practical functions and some other related interfaces.

Since these interfaces and primitive operations wrap low-level operations in different implementations, customers should not assume that they are safe for parallel execution.

The most important in the io package are two interfaces: Reader and Writer interfaces. First, let's introduce these two interfaces.

The definition of the Reader interface, the Read() method is used to read data.

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

```
Read reads len(p) bytes into p. It returns the number of bytes read n (0 <= n <= len(p)) and any errors encountered. Even if n <len(p) returned by Read, it will use all of p as temporary storage space during the call. If some data is available but less than len(p) bytes, Read will return what is available as usual instead of waiting for more.

When Read encounters an error or EOF condition after successfully reading n> 0 bytes, it will return the number of bytes read. It will return (non-nil) errors from the same call or return errors (and n == 0) from subsequent calls. An example of this general situation is that Reader will return a non-zero number of bytes at the end of the input stream. The possible return is either err == EOF or err == nil. In any case, the next Read should return 0, EOF.

The caller should always process bytes with n> 0 before considering the error err. Doing so can properly handle I/O errors after reading some bytes and allowing EOF behavior.

The implementation of Read prevents the return of a zero byte count and a nil error, and the caller should treat this situation as a no-op.
```





The definition of the Writer interface, the Write() method is used to write out data.

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

```
Write writes len(p) bytes from p to the basic data stream. It returns the number of bytes written from p n (0 <= n <= len(p)) and any errors encountered that caused the write to stop prematurely. If Write returns n <len(p), it must return a non-nil error. Write cannot modify the data of this slice, even if it is temporary.
```



The definition of the Seeker interface encapsulates the basic Seek method.

```go
type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
}
```

```
Read and write pointer used by Seeker to move data
Seek sets the pointer position for the next read and write operation, and each read and write operation starts from the pointer position
The meaning of whence:
	If whyce is 0: it means to move the pointer from the beginning of the data
	If whyce is 1: it means to move the pointer from the current pointer position of the data
	If whyce is 2: it means to move the pointer from the end of the data
offset is the offset of the pointer movement
	Return the pointer position after the movement and any errors encountered during the movement
```



 The definition of the ReaderFrom interface encapsulates the basic ReadFrom method.

```go
type ReaderFrom interface {
        ReadFrom(r Reader) (n int64, err error)
}
```

```
ReadFrom reads data from r to the data stream of the object
	Until r returns EOF or r read error occurs
	The return value n is the number of bytes read
	The return value err is the return value err of r
```



The definition of the WriterTo interface encapsulates the basic WriteTo method.

```go
type WriterTo interface {
        WriteTo(w Writer) (n int64, err error)
}
```

```
WriterTo writes the data stream of the object to w
	Until the data stream of the object is completely written or a write error is encountered
	The return value n is the number of bytes written
	The return value err is the return value err of w
```



Define the ReaderAt interface, the ReaderAt interface encapsulates the basic ReadAt method

```go
type ReaderAt interface {
        ReadAt(p []byte, off int64) (n int, err error)
}
```

```
ReadAt reads data from off of the object data stream to p
	Ignore the reading and writing pointer of the data, start reading from the starting position of the data offset off
	If the data stream of the object is only partially available, it is not enough to fill p
	Then ReadAt will wait for all data to be available and continue to write to p
	Return until p is filled
	At this point, ReadAt is more strict than Read
	Returns the number of bytes read n and the error encountered while reading
	If n <len(p), you need to return an err value to illustrate
	Why is p not filled (such as EOF)
	If n = len(p), and the data of the object is not all read, then
	err will return nil
	If n = len(p), and the data of the object has just been read, then
	err will return EOF or nil (not sure)
```



Define the WriterAt interface, the WriterAt interface encapsulates the basic WriteAt method

```go
type WriterAt interface {
        WriteAt(p []byte, off int64) (n int, err error)
}
```

```
WriteAt writes the data in p to the off position of the object data stream
	Ignore the read and write pointers of the data, start writing from the start position of the data offset off
	Returns the number of bytes written and errors encountered during writing
	If n <len(p), an err value must be returned to indicate
	Why p is not completely written
```





other. . .




## 2. File read and write

The file class is in the os package, encapsulating the underlying file descriptor and related information, and encapsulating the implementation of Read and Write.

```go

func (f *File) Read(b []byte) (n int, err error)
//The Read method reads up to len(b) bytes of data from f and writes to b. It returns the number of bytes read and any errors that may be encountered. The file termination flag is read 0 bytes and the return value err is io.EOF.

func (f *File) ReadAt(b []byte, off int64) (n int, err error)
//ReadAt reads len(b) byte data from the specified position (relative to the beginning of the file) and writes it to b. It returns the number of bytes read and any errors that may be encountered. When n
   
    
	 fmt.Println(n) //4
	 fmt.Println(bs) //[97 98 99 100]
	fmt.Println(string(bs)) //abcd

	// second read
	n,err = file.Read(bs)
	fmt.Println(err)//
    
      fmt.Println(n)//4 fmt.Println(bs) //[101 102 103 104] fmt.Println(string(bs)) //efgh //Read n for the third time, err = file.Read( bs) fmt.Println(err) //
     
       fmt.Println(n) //2 fmt.Println(bs) //[105 106 103 104] fmt.Println(string(bs)) //ijgh //The fourth time read n, err = file.Read( bs) fmt.Println(err) //EOF fmt.Println(n) //0 */ n := -1 for{ n,err = file.Read(bs) if n == 0 || err == io .EOF{ fmt.Println("Read to the end of the file, end the reading operation...") break} fmt.Println(string(bs[:n]))}} ``` Write out the data to the local file: ```go package main import ("os" "fmt" "log") func main() {/* Write out the data: */ fileName := "/Users/ruby/Documents/pro/a/ab.txt" //step1: open the file//step2: write out the data//step3: close the file//file,err := os.Open(fileName) file,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY |os.O_APPEND,os.ModePerm) if err != nil{ fmt.Println(err) return} defer file.Close() //Write out the data //bs :=[]byte{65,66,67,68 ,69,70}//A,B,C,D,E,F bs :=[] byte{97,98,99,100} //a,b,c,d //n,err := file.Write (bs) n,err := file.Write(bs[:2]) fmt.Println(n) HandleErr(err) file.WriteString("\n") //Write the string n directly, err = file.WriteString("HelloWorld") fmt.Println(n) HandleErr(err) file.WriteString( "\n") n,err =file.Write([]byte("today")) fmt.Println(n) HandleErr(err)} func HandleErr(err error){ if err != nil{ log.Fatal( err)}} ``` # 3. File copying is mainly a method of operating flow in the io package. Today, I will mainly learn copy. Just copy a file to another directory. Its principle is to read the data in the file from the source file through the program, and then write it out to the target file. ![copyfile](img/copyfile.png) ## 1. Method 1: The Read() and Write() methods under the io package are implemented. We can read while using the Read() and Write() methods under the io package Write, you can copy the file. This method reads the file in blocks, and the size of the block will also affect the performance of the program. ```go} /* The function of this function: realize the copy of the file, the return value is the total number of copies (bytes), error*/ func copyFile1(srcFile,destFile string)(int,error){ file1,err: =os.Open(srcFile) if err != nil{ return 0,err} file2,err:=os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm) if err !=nil{ return 0, err} defer file1.Close() defer file2.Close() //copy data bs:= make([]byte,1024,1024) n :=-1//The amount of data read total := 0 for {n,err = file1.Read(bs) if err == io.EOF || n = = 0{ fmt.Println("The copy is complete...") break }else if err !=nil{ fmt.Println("An error was reported...") return total,err} total += n file2.Write(bs[ :n])} return total,nil} ``` ## 2. Method 2: Implementation of the Copy() method under the io package We can also directly use the Copy() method under the io package. The sample code is as follows: ```go func copyFile2(srcFile, destFile string)(int64,error){ file1,err:=os.Open(srcFile) if err != nil{ return 0,err} file2,err:=os .OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm) if err !=nil{ return 0,err} defer file1.Close() defer file2.Close() return io.Copy(file2,file1)} ``` ### Extended content: In the io package (golang version 1.12), not only the Copy() method is provided, but there are two other public copy methods: CopyN(), CopyBuffer(). ```go Copy(dst,src) is to copy all src to dst. CopyN(dst,src,n) is to copy n bytes in src to dst. CopyBuffer (dst, src, buf) is to specify a buf buffer area to be completely copied with this size. ``` Their relationship is as follows:![20190316084535903](img/20190316084535903.jpg) (The picture comes from the network) As can be seen from the figure, no matter which copy method it is, it is ultimately implemented by the private method copyBuffer(). ```go func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {// If the reader has a WriteTo method, use it to do the copy. // Avoids an allocation and a copy. if wt, ok := src.(WriterTo); ok {return wt.WriteTo(dst)} // Similarly, if the writer has a ReadFrom method, use it to do the copy. if rt, ok := dst .(ReaderFrom); ok {return rt.ReadFrom(src)} if buf == nil {size := 32 * 1024 if l, ok := src.(*LimitedReader); ok && int64(size)> lN {if lN <1 {size = 1} else {size = int(lN)}} buf = make([]byte, size)} for {nr, er := src.Read(buf) if nr> 0 {nw, ew := dst.Write(buf[0:nr]) if nw> 0 {written += int64(nw)} if ew != nil {err = ew break} if nr!= nw {err = ErrShortWrite break}} if er != nil {if er != EOF {err = er} break}} return written, err} ``` From this part of the code, we can see that there are three main types of copying . 1. If the copied Reader (src) will try to be asserted as writerTo, if possible, directly call the writerTo method below. 2. If Writer (dst) will try to be asserted as ReadFrom, if possible, directly call the following readfrom Method 3. If there is no implementation, call the underlying read to implement the copy. Among them, there is such a piece of code: ```go if buf == nil {size := 32 * 1024 if l, ok := src.(*LimitedReader); ok && int64(size)> lN {if lN <1 { size = 1} else {size = int(lN)}} buf = make([]byte, size)} ``` This part mainly implements the processing of Copy and CopyN. Through the above call relationship diagram, we can see that after CopyN is called, it will convert Reader to LimiteReader. The difference is that if Copy, directly create a buffer with a default size of 32*1024 buf, if it is CopyN, it will first determine the number of bytes to be copied, if it is less than the default size, it will create a buf equal to the number of bytes to be copied. ## 3. Method 3: The third method of the ioutil package is to use the `ioutil.WriteFile()` and `ioutil.ReadFile()` in the ioutil package, but because of the use of one-time read files, one-time write files Therefore, this method is not suitable for large files and is prone to memory overflow. Sample code: ```go func copyFile3(srcFile, destFile string)(int,error){ input, err := ioutil.ReadFile(srcFile) if err != nil {fmt.Println(err) return 0,err} err = ioutil.WriteFile(destFile, input, 0644) if err != nil {fmt.Println("The operation failed:", destFile ) fmt.Println(err) return 0,err} return len(input),nil} ``` ## 4. In the end, let’s test these 3 kinds of copies that take time. The copied files are all the same mp4 file (400M), ![WX20190702-124039](img/WX20190702-124039.png) Code: ```go func main() {/* Copy file: */ //srcFile := "/home/ruby/ Document/pro/aa.txt" //destFile := "/home/ruby/Document/aa.txt" srcFile :="/Users/ruby/Documents/pro/a/001_Introduction to applet.mp4" destFile: ="001_Introduction to Mini Programs.mp4" total,err:=copyFile1(srcFile,destFile) fmt.Println(err) fmt.Println(total)} ``` The first type: Read() and Write( under the io package ) Direct read and write: The size of the slice we create to read the data directly affects the performance. ```go localhost:l_file ruby$ time go run demo05_copy.go The copy is complete. .input, 0644) if err != nil {fmt.Println("Operation failed:", destFile) fmt.Println(err) return 0,err} return len(input),nil} ``` ## 4. Final summary , Let’s test these 3 kinds of copies that take time, the copied files are all the same mp4 file (400M), ![WX20190702-124039](img/WX20190702-124039.png) Code: ```go func main () {/* Copy file: */ //srcFile := "/home/ruby/document/pro/aa.txt" //destFile := "/home/ruby/document/aa.txt" srcFile :=" /Users/ruby/Documents/pro/a/001_Introduction to applet.mp4" destFile:="001_Introduction to applet.mp4" total,err:=copyFile1(srcFile,destFile) fmt.Println(err) fmt. Println(total)} ``` The first type: Read() and Write() under the io package directly read and write: The size of the slice we create to read the data directly affects the performance. ```go localhost:l_file ruby$ time go run demo05_copy.go The copy is complete. .input, 0644) if err != nil {fmt.Println("Operation failed:", destFile) fmt.Println(err) return 0,err} return len(input),nil} ``` ## 4. Final summary , Let’s test these 3 kinds of copies that take time, the copied files are all the same mp4 file (400M), ![WX20190702-124039](img/WX20190702-124039.png) Code: ```go func main () {/* Copy file: */ //srcFile := "/home/ruby/document/pro/aa.txt" //destFile := "/home/ruby/document/aa.txt" srcFile :=" /Users/ruby/Documents/pro/a/001_Introduction to applet.mp4" destFile:="001_Introduction to applet.mp4" total,err:=copyFile1(srcFile,destFile) fmt.Println(err) fmt. Println(total)} ``` The first type: Read() and Write() under the io package directly read and write: The size of the slice we create to read the data directly affects the performance. ```go localhost:l_file ruby$ time go run demo05_copy.go The copy is complete. .[WX20190702-124039](img/WX20190702-124039.png) Code: ```go func main() {/* Copy file: */ //srcFile := "/home/ruby/document/pro/aa.txt "//destFile := "/home/ruby/document/aa.txt" srcFile :="/Users/ruby/Documents/pro/a/001_Introduction to applet.mp4" destFile:="001_Introduction to applet .mp4" total,err:=copyFile1(srcFile,destFile) fmt.Println(err) fmt.Println(total)} ``` The first type: Read() and Write() under the io package directly read and write: our own The size of the slice created to read the data directly affects performance. ```go localhost:l_file ruby$ time go run demo05_copy.go The copy is complete. .[WX20190702-124039](img/WX20190702-124039.png) Code: ```go func main() {/* Copy file: */ //srcFile := "/home/ruby/document/pro/aa.txt "//destFile := "/home/ruby/document/aa.txt" srcFile :="/Users/ruby/Documents/pro/a/001_Introduction to applet.mp4" destFile:="001_Introduction to applet .mp4" total,err:=copyFile1(srcFile,destFile) fmt.Println(err) fmt.Println(total)} ``` The first type: Read() and Write() under the io package directly read and write: our own The size of the slice created to read the data directly affects performance. ```go localhost:l_file ruby$ time go run demo05_copy.go The copy is complete. .Println(total)} ``` The first type: Read() and Write() under the io package directly read and write: The size of the slice we create to read the data directly affects the performance. ```go localhost:l_file ruby$ time go run demo05_copy.go The copy is complete. .Println(total)} ``` The first type: Read() and Write() under the io package directly read and write: The size of the slice we create to read the data directly affects the performance. ```go localhost:l_file ruby$ time go run demo05_copy.go The copy is complete. .
      
        401386819 real 0m7.911s user 0m2.900s sys 0m7.661s ``` Second: Copy() method under io package: ```go localhost:l_file ruby$ time go run demo05_copy.go 
       
         401386819 real 0m1.594s user 0m0.533s sys 0m1.136s ``` The third type: ioutil package ```go localhost:l_file ruby$ time go run demo05_copy.go 
        
          401386819 real 0m1.515s user 0m0.339s sys 0m0.625s ``` Operation result: ![WX20190702-124719](img/WX20190702-124719.png) These three methods, in terms of performance, no matter it is io.Copy( ) Is still the ioutil package, and the performance is pretty good. # 四、断点再传## 1. Seeker interface Seeker is an interface that wraps the basic Seek method. ```go type Seeker interface {Seek(offset int64, whence int) (int64, error)} ``` seek(offset,whence), set the position of the pointer cursor, read and write files randomly: The first parameter: offset The second parameter: how to set 0: seekStart means relative to the beginning of the file, 1: seekCurrent means relative to the current offset, 2: seek end means relative to the end. ```go const (SeekStart = 0 // seek relative to the origin of the file SeekCurrent = 1 // seek relative to the current offset SeekEnd = 2 // seek relative to the end) ``` Sample code: we want to read Take the aa.txt file in the local /Users/ruby/Documents/pro/a directory, the content of the file is: abcdefghij these characters. ![WX20190703-155441](img/WX20190703-155441.png) Sample code: ```go package main import ("os" "fmt" "io") func main() {/* seek(offset,whence),Set the position of the pointer cursor. The first parameter: offset. The second parameter: how to set 0: seekStart means relative to the beginning of the file, 1: seekCurrent means relative to the current offset, 2: seek end means relative to the end. const (SeekStart = 0 // seek relative to the origin of the file SeekCurrent = 1 // seek relative to the current offset SeekEnd = 2 // seek relative to the end) Read the file randomly: You can set the position of the pointer cursor*/ file,_:=os.OpenFile("/Users/ruby/Documents/pro/a/aa.txt",os.O_RDWR,0) defer file.Close() bs :=[]byte{0} file.Read (bs) fmt.Println(string(bs)) file.Seek(4,io.SeekStart) file.Read(bs) fmt.Println(string(bs)) file.Seek(2,0) //also SeekStart file .Read(bs) fmt.Println(string(bs)) file.Seek(3,io.SeekCurrent) file.Read(bs) fmt.Println(string(bs)) file.Seek(0,io.SeekEnd) file .WriteString("ABC")} ``` Operation result: ![WX20190703-155739](img/WX20190703-155739.png) Local file: ![WX20190703-155821](img/WX20190703-155821.png) ## 2. Resuming the transfer with a breakpoint. First, think about a few questions Q1: If the file you want to transfer is relatively large, is there a way to reduce the time consumption? Q2: If the program is interrupted due to various reasons during the file transfer process, does the file need to be restarted next time it is restarted? Q3: Do you support pause and resume when transferring files? Even if these two operations are distributed before and after the program process is killed. It can be realized by resuming the transmission through breakpoints, and different languages ​​have different implementation methods. Let’s take a look at how to implement the Seek() method in the Go language: Let’s talk about the idea first: If you want to achieve a resumable transmission, the main thing is to remember how much data has been transmitted last time, then we can create a temporary file to record that it has been transmitted. When the transfer is resumed, the amount of data that has been transferred last time is first read from the temporary file, and then set to the read and write positions through the Seek() method, and then continue to transfer the data. Sample code: ```go package main import ("fmt" "os" "strconv" "io") func main() {/* Breakpoint resume: file transfer: file copy /Users/ruby/Documents/pro/ Copy a/guliang.jpeg to guliang4.jpeg Idea: While copying, record the total amount copied*/ srcFile:="/Users/ruby/Documents/pro/a/guliang.jpeg" destFile:="guliang4.jpeg" tempFile:=destFile+"temp.txt" //fmt.Println(tempFile) file1,_:=os.Open(srcFile) file2,_:=os.OpenFile(destFile,os.O_CREATE|os.O_WRONLY,os.ModePerm ) file3,_:=os.OpenFile(tempFile,os.O_CREATE|os.O_RDWR,os.ModePerm) defer file1.Close() defer file2.Close() //1. Read the data in the temporary file, according to seek file3.Seek(0,io.SeekStart) bs:=make([]byte,100,100) n1 ,err:=file3.Read(bs) fmt.Println(n1) countStr:=string(bs[:n1]) fmt.Println(countStr) //count,_:=strconv.Atoi(countStr) count,_: =strconv.ParseInt(countStr,10,64) fmt.Println(count) //2. Set the offset for reading and writing file1.Seek(count,0) file2.Seek(count,0) data:=make( []byte,1024,1024) n2:=-1// The amount of data read n3:=-1// The amount of data written total :=int(count)//The total amount read for{ // 3. Read data n2, err=file1.Read(data) if err ==io.EOF{ fmt.Println("File copy completed...") file3.Close() os.Remove(tempFile) break} // Write the data to the target file n3,_=file2.Write(data[:n2]) total += n3 //Store the total amount of copy in a temporary file file3.Seek(0,io.SeekStart) file3.WriteString (strconv.Itoa(total)) //Pretend to be powered off//if total>8000{ // panic("Pretend to be powered off..., pretend...") //}}} ``` # Five , Bufio package> @author: Han Ru>Copyright: Beijing Qianfeng Internet Technology Co., Ltd. Life will not end, continue to go go go. . In the io operation, Go language also provides a bufio package, which can greatly improve the efficiency of file reading and writing. ## 1. The principle of bufio package bufio improves efficiency through buffering. The efficiency of the io operation itself is not low, the low is the frequent access to the files on the local disk. So bufio provides a buffer (allocate a piece of memory). Reading and writing are in the buffer first, and then the file is read and written to reduce the number of accesses to the local disk, thereby improving efficiency. Simply put, when you read the file into the buffer (memory) and then read it, you can avoid the io of the file system and improve the speed. In the same way, when performing a write operation, the file is first written into the buffer (memory), and then the buffer is written into the file system. After reading the above explanation, some people may be confused. Compare the content->file and content->buffer->file directly. The buffer seems to have no effect. In fact, the buffer is designed to store multiple writes, and the contents of the buffer are written into the file at the last breath. ![WX20190704-113648](img/WX20190704-113648.png) bufio encapsulates the io.Reader or io.Writer interface object, and creates another object that also implements the interface. The io.Reader or io.Writer interface implements the read() and write() methods. These two methods can be used for objects that implement this interface. Reader object bufio.Reader is a package of io.Reader in bufio, `go // Reader implements buffering for an io.Reader object. type Reader struct {buf []byte rd io.Reader // reader provided by the client r , w int // buf read and write positions err error lastByte int // last byte read for UnreadByte;-1 means invalid lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid} ``` bufio.Read(p []byte) is equivalent to reading the content of the size len(p), the idea is as follows: 1. When there is content in the cache area, fill all the contents of the cache area into p and clear the cache area 2. When there is no content in the cache area and len(p)>len(buf), that is, the content to be read is more than the cache area If it is bigger, just go to the file to read it 3. When there is no content in the buffer area and len(p)
         
          = len(b.buf) {// Large read, empty buffer. // Read directly into p to avoid copy. n, b.err = b.rd.Read(p) if n <0 {panic(errNegativeRead)} if n> 0 {b.lastByte = int(p[n-1]) b.lastRuneSize = -1} return n, b.readErr()} // One read. // Do not use b.fill, which will loop. br = 0 bw = 0 n, b.err = b.rd.Read(b.buf) if n <0 {panic(errNegativeRead)} if n == 0 {return 0, b.readErr()} bw += n} // copy as much as we can n = copy(p, b.buf[br:bw]) br += n b.lastByte = int(b.buf[br-1]) b.lastRuneSize = -1 return n, nil} ``` Explanation: The reader internally judges whether the contents of the buffer area have been read out by maintaining a r, w, that is, the position index for reading and writing. Writer object bufio.Writer is the encapsulation of io.Writer in bufio```go // Writer implements buffering for an io.Writer object. // If an error occurs writing to a Writer, no more data will be // accepted and all subsequent writes, and Flush, will return the error.// After all data has been written, the client should call the // Flush method to guarantee all data has been forwarded to // the underlying io.Writer. type Writer struct {err error buf []byte n int wr io.Writer } ``` bufio.Write(p []byte) The idea is as follows: 1. Determine whether the available capacity in buf can be put down p 2. If it can be put down, directly splice p to the back of buf, that is, put the content in the buffer 3. If the available capacity of the buffer is not enough to put it down, and the buffer is empty at this time, just write p directly to the file. 4. If the available capacity of the buffer is not enough to put it down, and the buffer has content at this time, use p Fill up the buffer, write all the contents of the buffer to the file, and empty the buffer 5. Determine whether the remaining content of p can be placed in the buffer, and if it can be placed (this time is the same as in step 1), put the content To buffer 6. If the remaining content of p is still larger than the buffer, (note that the buffer is empty at this time, the situation is the same as step 3), then write the remaining content of p directly to the file. The following is the source code ```go // Write writes the contents of p into the buffer. // It returns the number of bytes written. // If nn <len(p), it also returns an error explaining // why the write is short. func (b *Writer) Write(p []byte) (nn int, err error) {for len(p)> b.Available() && b.err == nil {var n int if b.Buffered() == 0 {// Large write,empty buffer. // Write directly from p to avoid copy. n, b.err = b.wr.Write(p)} else {n = copy(b.buf[bn:], p) bn += n b. Flush()} nn += np = p[n:]} if b.err != nil {return nn, b.err} n := copy(b.buf[bn:], p) bn += n nn += n return nn, nil} ``` Explanation: b.wr stores an io.writer object, which implements the Write() interface, so you can use b.wr.Write(p) to write the content of p document. b.flush() will write the contents of the buffer area to the file. After all writing is completed, because the buffer area will store the contents, you need to manually flush() to the file. b.Available() is the available capacity of buf, which is equal to len(buf)-n. The following figure explains one of the cases, that is, there is content in the buffer area, and the remaining p is greater than the buffer area! [WX20190704-122357](img/WX20190704-122357.png) ## 2, bufio package bufio package implements buffered I /O. It wraps an io.Reader or io.Writer interface object, creates another object that also implements the interface, and also provides buffering and some text I/O helper functions. bufio.Reader: bufio.Reader implements the following interfaces: io.Reader io.WriterTo io.ByteScanner io.RuneScanner ```go // NewReaderSize encapsulates rd into a buffered bufio.Reader object, // the buffer size is determined by size Specify (if it is less than 16, it will be set to 16). // If the base type of rd is bufio with sufficient cache.Reader type, directly convert // rd to base type and return. func NewReaderSize(rd io.Reader, size int) *Reader // NewReader is equivalent to NewReaderSize(rd, 4096) func NewReader(rd io.Reader) *Reader // Peek returns a slice of the cache, the slice refers to the first n in the cache Data of bytes, // this operation will not read the data, it is just a reference. The quoted data is valid before the next read operation //. If the slice length is less than n, an error message will be returned to explain the reason. // If n is greater than the total size of the buffer, ErrBufferFull is returned. func (b *Reader) Peek(n int) ([]byte, error) // Read reads data from b to p, and returns the number of bytes read and the error encountered. // If the cache is not empty, you can only read the data in the cache, and will not extract data from the underlying io.Reader //. If the cache is empty, then: // 1, len(p) >= cache size , Skip the cache, and read directly from the underlying io.Reader to p. // 2. len(p) <cache size, first read the data from the underlying io.Reader to the cache //, and then read from the cache to p. func (b *Reader) Read(p []byte) (n int, err error) // Buffered returns the length of the unread data in the buffer. func (b *Reader) Buffered() int // ReadBytes function is the same as ReadSlice, except that it returns a buffered copy. func (b *Reader) ReadBytes(delim byte) (line []byte,err error) // ReadString function is the same as ReadBytes, except that it returns a string. func (b *Reader) ReadString(delim byte) (line string, err error) ... ``` bufio.Writer: bufio.Writer implements the following interfaces: io.Writer io.ReaderFrom io.ByteWriter ```go / / NewWriterSize encapsulates wr into a buffered bufio.Writer object, // the buffer size is specified by size (if it is less than 4096, it will be set to 4096). // If the base type of wr is the bufio.Writer type with sufficient buffer, then directly convert // wr to the base type and return. func NewWriterSize(wr io.Writer, size int) *Writer // NewWriter is equivalent to NewWriterSize(wr, 4096) func NewWriter(wr io.Writer) *Writer // The function of WriteString is the same as Write, except that the string func is written (b *Writer) WriteString(s string) (int, error) // WriteRune writes the UTF-8 encoding of r to b, and returns the encoding length of r. func (b *Writer) WriteRune(r rune) (size int, err error) // Flush submits the data in the cache to the underlying io.Func (b *Writer) Flush() error in Writer // Available returns the length of the unused space in the buffer func (b *Writer) Available() int // Buffered returns the length of the uncommitted data in the buffer func (b * Writer) Buffered() int // Reset resets the underlying Writer of b to w, discards all data in the buffer, and resets // all flags and error messages. It is equivalent to creating a new bufio.Writer. func (b *Writer) Reset(w io.Writer) ... ``` ## 3. Example code to read data: ```go package main import ("os" "fmt" "bufio") func main( ) {/* bufio: efficient io read and write buffer cache io: input/output wrap the Reader and Write objects under the io package, and package with cache to improve the efficiency of reading and writing ReadBytes() ReadString() ReadLine() */ fileName:="/Users/ruby/Documents/pro/a/english.txt" file,err := os.Open(fileName) if err != nil{ fmt.Println(err) return} defer file.Close() //Create Reader object//b1 := bufio.NewReader(file) //1.Read(), efficient reading//p := make([]byte,1024) //n1,err := b1.Read (p) //fmt.Println(n1) //fmt.Println(string(p[:n1])) //2.ReadLine() //data,flag,err:= b1.ReadLine() //fmt.Println(flag) //fmt.Println(err) //fmt.Println(data) //fmt.Println(string(data)) //3.ReadString() // s1,err :=b1.ReadString('\n') // fmt.Println(err) // fmt.Println(s1) // // s1,err = b1.ReadString('\n') // fmt .Println(err) // fmt.Println(s1) // //s1,err = b1.ReadString('\n') //fmt.Println(err) //fmt.Println(s1) // // for{ // s1,err := b1.ReadString('\n') // if err == io.EOF{ // fmt.Println("Read completed...") // break //} // fmt.Println(s1) //) //4.ReadBytes() //data,err :=b1.ReadBytes('\n') //fmt.Println(err) //fmt.Println(string(data) ) //Scanner //s2 := "" //fmt.Scanln(&s2) //fmt.Println(s2) b2 := bufio.NewReader(os.Stdin) s2, _ := b2.ReadString('\n ') fmt.Println(s2)} ``` Local file: english.txt file content: ![WX20190704-172759](img/WX20190704-172759.png) Sample code for writing data: ```go package main import ("os" "fmt" "bufio") func main() {/* bufio: efficient io read and write buffer cache io: input/output will be under the io package Reader and Write objects are packaged, and packaged with a buffer to improve the efficiency of reading and writing func (b *Writer) Write(p []byte) (nn int, err error) func (b *Writer) WriteByte(c byte) error func (b *Writer) WriteRune(r rune) (size int, err error) func (b *Writer) WriteString(s string) (int, error) */ fileName := "/Users/ruby/Documents/pro/a/ cc.txt" file,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY,os.ModePerm) if err != nil{ fmt.Println(err) return} defer file.Close() w1 := bufio.NewWriter(file) //n,err := w1.WriteString("helloworld") //fmt.Println(err) //fmt.Println(n) //w1.Flush() //Flush the buffer for i:=1;i<=1000;i++{ w1.WriteString(fmt.Sprintf("%d:hello",i))} w1.Flush()} ``` # 6. In addition to the io package that can read and write data, the ioutil package also provides an auxiliary tool package called ioutil in the Go language. Although there are not many methods in it, they are all pretty easy to use. ```go import "io/ioutil" ``` The introduction of this package is only one sentence: Package ioutil implements some I/O utility functions. ## 1. Methods of the ioutil package Let's take a look at the methods inside: ```go // Discard is an io.Writer interface, calling its Write method will do nothing // and always return successfully. var Discard io.Writer = devNull(0) // ReadAll reads all the data in r and returns the read data and errors encountered. // If the reading is successful, err returns nil instead of EOF, because ReadAll is defined to read // all data, so EOF will not be treated as an error. func ReadAll(r io.Reader) ([]byte, error) // ReadFile reads all the data in the file and returns the data read and the error encountered. // If the read is successful, err returns nil instead of EOF func ReadFile(filename string) ([]byte, error) // WriteFile writes data to the file, and the file will be cleared before writing. // If the file does not exist, it will be created with the specified permissions. // Return the error encountered. func WriteFile(filename string, data []byte, perm os.FileMode) error // ReadDir reads all directories and files in the specified directory (excluding subdirectories). // Return a list of read file information and errors encountered. The list is sorted. func ReadDir(dirname string) ([]os.FileInfo,error) // NopCloser wraps r as a ReadCloser type, but the Close method does nothing. func NopCloser(r io.Reader) io.ReadCloser // TempFile creates a temporary file prefixed with prefix in the dir directory and opens it in read//write mode. Returns the file object created and the error encountered. // If dir is empty, the file will be created in the default temporary directory (see os.TempDir). Multiple // calls will create different temporary files. The caller can get the full path of the file through f.Name(). // The temporary file created by calling this function should be deleted by the caller. func TempFile(dir, prefix string) (f *os.File, err error) // The function of TempDir is the same as TempFile, except that it creates a directory and returns the full path of the directory. func TempDir(dir, prefix string) (name string, err error) ``` ## 2. Sample code: ```go package main import ("io/ioutil" "fmt" "os") func main() { /* ioutil package: ReadFile() WriteFile() ReadDir() .. */ //1. Read all the data in the file //fileName1 := "/Users/ruby/Documents/pro/a/aa.txt "//data, err := ioutil.ReadFile(fileName1) //fmt.Println(err) //fmt.Println(string(data)) //2.Write out data//fileName2:="/Users/ruby /Documents/pro/a/bbb.txt"//s1:="helloworld faces the sea and spring flowers bloom" //err:=ioutil.WriteFile(fileName2,[]byte(s1),0777) //fmt.Println(err) //3. //s2: ="qwertyuiopsdfghjklzxcvbnm" //r1:=strings.NewReader(s2) //data,_:=ioutil.ReadAll(r1) //fmt.Println(data) //4.ReadDir(), read a directory Sub-content: sub-files and sub-directories, but only one level //dirName:="/Users/ruby/Documents/pro/a" //fileInfos,_:=ioutil.ReadDir(dirName) //fmt.Println(len( fileInfos)) //for i:=0;i
          
         
        
       
      
     
    
   

