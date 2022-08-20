

# Go language environment construction

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## One, the official website of the Golang language

First, we log on to the official website of Golang: https://golang.org/

![guanwang2](img/guanwang2.png)



> Because of the relationship between Google and China, directly logging into Golang's official website requires overcoming the wall.



Of course, you can also log in to Golang's domestic website:
   

![WX20190403-095602](img/guanwang1.png)

## 2. Download



Golang is supported on three platforms: Mac, Windows and Linux. You can start from
   Download the installation package of the corresponding platform.

![xiazai1](img/xiazai1.png)



The website is not easy to visit in China, so you can visit the corresponding Chinese
   Or Go Language Chinese Network


   
     Download the installation software.

![xiazai2](img/xiazai2.png)



and

![xiazai3](img/xiazai3.png)



The latest version is Go1.12 released in February this year.



Mac OS
Download the osx installer from https://golang.org/dl/. Double-click to start the installation. As prompted, this should install Golang in /usr/local/go, and also add the folder /usr/local/go/bin to your PATH environment variable.

Windows
Download the MSI installer from https://golang.org/dl/. Double-click to start the installation and follow the prompts. This will install Golang:\Go in location c, and will also add the directory c:\Go\bin to your path environment variable.

Linux
Download the tar file from https://golang.org/dl/ and unzip it to /usr/local. Add /usr/local/go/bin to the PATH environment variable. This should be installed in linux.



```
The development kit is divided into an installation version and a compressed version. The installed version is specific to Mac and Windows, and their names are similar to:

-go1.12.1.darwin-amd64.pkg
-go1.12.1.windows-386.msi
-go1.12.1.windows-amd64.msi

The installation version, as the name implies, double-click to open it and an installation wizard will appear, allowing you to choose the installation path, helping you set up the environment and other information, which is more convenient and convenient.

The compressed version is a compressed file, which can be decompressed to get the contents, and their names are similar to:

-go1.12.1.darwin-amd64.tar.gz
-go1.12.1.linux-386.tar.gz
-go1.12.1.linux-amd64.tar.gz
-go1.12.1.windows-386.zip
-go1.12.1.windows-amd64.zip

After downloading the compressed version, we need to decompress it, and then move it to the path to be stored, and configure environment variables and other information. Compared with the installed version, it is more complicated, and there are more manual configurations.

```





## Three, install and configure environment variables

### 3.1 Linux system installation and configuration

For the Linux system, let's take Ubuntu as an example.

**Step 1: Download and install**

Go to the official website of go to download the go installation package, and it will be automatically downloaded to the download directory.

Open the terminal, go to the download directory, and check the installation package:

```shell
ruby@hanru:~$ cd download
ruby@hanru:~/download$ ls
```

![ubuntu1](img/ubuntu1.png)



Then we need to decompress the compressed package and copy it to the specified directory, so continue to execute the following command in the terminal:

```shell
ruby@ubuntu:~/download$ sudo tar -xzf go1.12.1.linux-amd64.tar.gz -C /usr/local
```

![ubuntu2](img/ubuntu2.png)



> Enter sudo, which means to execute the command as an administrator, and you need to enter a password

At this point, you will download the tar file from the go official website https://golang.org/dl/ and unzip it to the /usr/local directory. There will be a go folder in this directory.

You can enter this folder to view directly:

![ubuntu3](img/ubuntu3.png)



You can also view it through the terminal command, enter the following command:

```shell
ruby@hanru:~/download$ cd /usr/local
ruby@hanru:/usr/local$ ls
```

![ubuntu4](img/ubuntu4.png)



**Step 2: Configure environment variables**	

One: You need to install vim first.

Execute the following commands directly in the terminal:

```shell
ruby@ubuntu:~$ sudo apt-get install vim
```

Two: Edit the $HOME/.profile file

**A: Configure GOROOT first, which is the installation directory of go**

```shell
export GOROOT="/usr/local/go"
```



**B: Then configure GOPATH**

Gopath is where the Go project code is stored. This is a directory defined by ourselves. It is like the Workspace of other IDEs.

​ For Ubuntu systems, the Home/go directory is used as the gopath by default.

​ There are 3 subdirectories under this directory: src, pkg, bin

> The GO code must be in the workspace. The workspace is a directory that contains three subdirectories:
>
> ​ src ---- Each subdirectory in it is a package. Inside the package is the Go source code file
>
> ​ pkg ---- generated after compilation, the object file of the package
>
> ​ bin ---- The generated executable file.



```shell
export GOPATH=$HOME/go
```



**C: GOBIN**

```shell
export GOBIN=$GOROOT/bin
```



**D: Add to PATH**

We need to add GOBIN to the environment variable PATH. You can directly add the following content to $HOME/.profile through vi

```shell
export PATH=$PATH:$GOBIN
```





> Of course, you can also configure the bin directory of GO directly into the PATH:
>
> ```shell
> export PATH=$GOROOT/bin:$PATH
> ```
> 
>Equivalent to writing the above steps C and D together





Specific operation:


 1. First use the ls -a command to check whether there is a .profile file in the home directory. (Files beginning with. Are hidden files, use the -a command to view)
2. Enter directly in the terminal: vi $HOME/.profile
3. Enter i, slice to edit mode, copy the above content to the file, and save and exit.
>
> ​ After clicking the esc key,
>
> ​ :q!, force exit without saving
>
> ​ :wq, save and exit	



Three: Let the configuration file take effect immediately

Use the source command to make the configuration file effective

```shell
ruby@ubuntu:~$ source $HOME/.profile
```

Four: test installation

Version check

```shell
ruby@ubuntu:~$ go version
```

Check the configuration information of go

```shell
ruby@ubuntu:~$ go env
```



Additional extensions:

```
File saving after vi command

The English full name of vi in ​​Linux is Visual Interface.

Perform the following operations in the last line mode. [Press in command mode: switch to last line mode]
w
Save the file without exiting vi

w file
Save the file as file without exiting vi

w!
Force save, do not exit vi

wq
Keep the file and exit

wq!
Force save file and exit

q
Exit vi without saving the file

q!
Forcibly exit vi without saving the file

e!
Abandon all modifications and edit from the last saved file
```



### 3.2 mac system installation and configuration

Mac is divided into a compressed version and an installed version, both of which are 64-bit. The compressed version is similar to Linux, because both Mac and Linux are based on Unix, and the terminal is basically the same.

**A: Installation**

Find the downloaded pkg installation package: Generally, the downloaded files are in the download directory.

![anzhuang1](img/anzhuang1.png)



Double-click the pkg package and follow the instructions to install it successfully. 



Enter go version on the command line and get the version number of go, which means the installation is successful.

**Configure environment variables**

1. Open the terminal and enter cd ~ to enter the user's home directory; 
2. Enter the ls -all command to see if .bash_profile exists; 
3. There is both use vim .bash_profile to open the file; 
4. Enter i to enter vim editing mode; 
5. Enter the following code: 

```shell
export GOROOT=/usr/local/go
export GOPATH=/Users/ruby/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
```

Among them GOPATH: the root directory of daily development. GOBIN: is the bin directory under GOPATH.

6. Click ESC, and enter: wq to save and exit editing. You can enter vim.bash_profile to check whether the save is successful.

7. Enter source ~/.bash_profile to complete the configuration of golang environment variables. There is no prompt for successful configuration. 
8. Enter go env to view the configuration results.



### 3.3 Windows

**A, installation**

I won’t say anything more about the installation steps.

**B. Configure environment variables**

Note: If it is an msi installation file, the environment variables of the Go language will be automatically set.

My computer-right click "Properties"-"Advanced System Settings"-"Environment Variables"-"System Variables"

​ Assume that GO is installed in the root directory of the C drive

**New:**

-GOROOT: Go installation path (example: C:\Go)

-GOPATH: The path of the Go project (for example: E:\go). If there are more than one, add them separated by semicolons

  ![winhuanjing1](img/winhuanjing1.jpg)

**Revise:**

-Path: add in path: C:\Go\bin;%GOPATH%\bin;

  > You need to configure the executable directory in GOPATH to the environment variable, otherwise the third-party go tool you downloaded by yourself will not be available

  

  ![winhuanjing2](img/winhuanjing2.jpg)

  

> 1. The working directory is where we store the source code for development, and it corresponds to the environment variable GOPATH in Go. After this environment variable is specified, the files we generated by compiling the source code and so on will be placed in this directory. The configuration of the GOPATH environment variable refers to the installation of Go above, and configures it to the system variables under Windows.
> 2. There are mainly three directories under GOPATH: bin, pkg, src. The bin directory mainly stores executable files; the pkg directory stores compiled library files, mainly *.a files; the src directory mainly stores go source files





**C. Check if the installation and configuration are successful**

Use the shortcut key win+R, enter cmd, open the command line prompt, enter in the command line

```go
go env # View the configuration information of go
go version # View the version number of go
```





Qianfeng Go language learning group: 784190273

Corresponding video address:

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

Source code:

https://github.com/rubyhan1314/go_foundation
   

