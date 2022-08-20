

# What can Go language do

> @author：Han Ru
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.

## 1. Why did we choose Go language

There may be many reasons for choosing the Go language. We have already introduced many of the features and advantages of the Go language in the previous documents. But the main reason should be based on the following two considerations:

1. Execution performance

   Shorten the response time of the API, and solve the problem of batch request access timeout. In the business scenario of Uwork, an API batch request often involves multiple calls to other interface services. In the previous PHP implementation mode, it is very difficult to make parallel calls, but serial processing cannot be fundamentally done. Improve processing performance. The GO language is different. The parallel processing of the API can be easily realized through the coroutine, and the processing efficiency can be maximized. Relying on Golang's high-performance HTTP Server to improve system throughput, from hundreds of PHP levels to thousands of miles or even tens of thousands.

2. Development efficiency

   The GO language is simple to use, efficient in code description, unified coding standards, and quick to get started. With a small amount of code, the framework can be standardized, and API business logic can be quickly constructed with a unified specification. It can quickly build various common components and public class libraries, further improve development efficiency, and realize mass production of functions in specific scenarios.

## 2. What can Go language do?

Since the release of version 1.0, the Go language has attracted the attention of many developers and has been widely used. The simplicity, efficiency, and concurrency of the Go language have attracted the participation of many traditional language developers, and the number is increasing.

In view of the characteristics of the Go language and the original intention of the design, the Go language as a server programming language is very suitable for processing logs, data packaging, virtual machine processing, file systems, distributed systems, database agents, etc.; in terms of network programming, the Go language is widely used in the Web Applications, API applications, download applications, etc. In addition, the Go language is also suitable for memory databases and cloud platforms. At present, many foreign cloud platforms are developed using Go.

-Server programming, if you used C or C++ to do those things, it is very suitable to use Go, such as processing logs, data packaging, virtual machine processing, file system, etc.
-Distributed systems, database agents, middleware, etc., such as Etcd.
-Network programming, this one is currently the most widely used, including web applications, API applications, download applications, and the net/http package built in Go basically implements all the network functions we usually use.
-Database operation
-Develop cloud platforms. Many foreign cloud platforms are currently using Go for development





## 3. Which companies or projects at home and abroad use Go language

After the release of Go, many companies, especially cloud computing companies, began to refactor their infrastructure with Go. Many of them directly adopted Go for development. Recently, Docker, which is in full swing, is developed with Go.

There are many open source projects developed in the Go language. Early Go language open source projects were only implemented by binding Go language and traditional projects with C language libraries, such as Qt, Sqlite, etc.; many later projects used Go language for re-native implementation. This process is simpler than other languages. This has also led to the emergence of a large number of native development projects using the Go language.



-Cloud computing infrastructure field

  Representative projects: docker, kubernetes, etcd, [consul](http://tonybai.com/2015/07/06/implement-distributed-services-registery-and-discovery-by-consul/), cloudflare CDN, Qiniu Cloud storage, etc.

-Basic software

  Representative projects: [tidb](https://github.com/pingcap/tidb), [influxdb](https://github.com/influxdata/influxdb), [cockroachdb](https://github.com/cockroachdb /cockroach) and so on.

-Microservices

  Representative projects: [go-kit](https://github.com/go-kit/kit), [micro](https://github.com/micro/micro), [typhon](https: //github.com/monzo), [bilibili](https://www.bilibili.com/), etc.

-Internet infrastructure

  Representative projects: [Ethereum](https://github.com/ethereum/go-ethereum), [hyperledger](https://github.com/hyperledger), etc.

  

> Some foreign companies that adopt Go, such as Google, Docker, Apple, Cloud Foundry, CloudFlare, Couchbase, CoreOS, Dropbox, MongoDB, AWS and other companies;
>
>Domestic companies that use Go to develop: such as Alibaba Cloud CDN, Baidu, Xiaomi, Qiniu, PingCAP, Huawei, Kingsoft, Cheetah Mobile, Ele.me and other companies.
>
>



### Docker

Docker is a virtualization technology at the operating system level. It can be isolated between the operating system and applications, and it can also be called a container. Docker can quickly run one or more instances on a physical server. A virtual packaging tool based on lxc can realize the formation of the PAAS platform. For example, start a CentOS operating system, and end after executing instructions on its internal command line, the whole process is as efficient as you are in the operating system.

Project link:

https://github.com/docker/docker





### go language

The early source code of Go language is written in C language and assembly language. Since Go 1.5, it is completely written in Go language itself. The source code of the Go language has great reference significance for understanding the underlying scheduling of the Go language. Readers who wish to have an in-depth understanding of the Go language are recommended to read it.

Project link:

https://github.com/golang/go



### Kubernetes

A container scheduling service built on Docker developed by Google, users can manage cloud container clusters through Kubernetes clusters.

Project link:

https://github.com/kubernetes/kubernetes



### etcd

A distributed and reliable KV storage system that can quickly perform cloud configuration.

Project link:

https://github.com/coreos/etcd



### beego

Beego is a Tornado framework similar to Python. It adopts the design idea of ​​RESTFul and is a very lightweight, highly scalable and high-performance web application framework written in Go language.

Project link:

https://github.com/astaxie/beego



### martini

A web framework for quickly building modular web applications.

Project link:

https://github.com/go-martini/martini



### codis

Excellent domestic distributed Redis solution.

Project link:

https://github.com/CodisLabs/codis



### delve

Go language

Powerful debugger, integrated by many integrated environments and editors.

Project link:

https://github.com/derekparker/delve



### Facebook

Facebook is also using it. For this reason, they have also established an open source organization facebookgo on Github. You can visit Facebook's open source projects through https://github.com/facebookgo, such as the famous smoothly upgraded grace.



### Uber





### Tencent

As a large domestic company, Tencent is still daring to try, especially Docker containerization. They have done docker-scale practice in 15 years. For details, please refer to http://www.infoq.com/cn/articles /tencent-millions-scale-docker-application-practice.





### Baidu

The currently known use of Baidu is in the operation and maintenance side, a BFE project of Baidu operation and maintenance, which is responsible for the access of front-end traffic. Their person in charge shared in 2016, you can check this http://www.infoq.com/cn/presentations/application-of-golang-in-baidu-frontend.

The second is Baidu's messaging system. Responsible for the development and maintenance of the company's mobile messaging system server.

### Jingdong

JD Cloud News Push System, Cloud Storage, and JD Mall all use Go for development.



### Millet

Xiaomi's support for Golang is nothing more than the open source of the operation and maintenance monitoring system, which is http://open-falcon.com/.

In addition, teams such as Xiaomi Interactive Entertainment, Xiaomi Mall, Xiaomi Video, and Xiaomi Ecochain are all using Golang.



### 360

360 uses Golang a lot, one is the open source log search system Poseidon, hosted on Github, https://github.com/Qihoo360/poseidon.

The 360 ​​push team is also using it, and they also wrote a blog post on the official Golang blog https://blog.golang.org/qihoo.





### Qiniu Cloud

Qiniu Cloud used nearly 500,000 lines of code to implement the entire product. Qiniu cloud storage product website: http://qiniu.com/. Online time: 2011-9-1. Application Scope: The entire product (including basic services, Web terminals, statistical platforms, various gadgets, etc.) Go code line percentage: 99.9% Daily PV: confidential



### 美团

Meituan background flow support program. Application scope: Support the background traffic of the master station (sort, recommendation, search, etc.), provide load balancing, cache, fault tolerance, conditional distribution, statistical operation indicators (qps, latency) and other functions.



### Didi

Basic service platform.



###Jinshan Weikan 

Application scope: service interface, background process service, message system, picture system



### Sogou

Sogou push system. The part of the Push system used to maintain the connection with the client.



### QOR-Modular e-commerce system 

-QOR official website: [QOR: E-commerce & CMS SDK written in Go](https://link.zhihu.com/?target=http%3A//getqor.com)
-github address: qor/qor · GitHub
-Application range: the entire product

### weico

Product name: Weico 3.0, all code on the server side is implemented in Go.



### Xian Xia Dao

-Product URL: [Xian Xia Dao official website-heart moving game] (Xian Xia Dao official website-heart moving game)
-Application: Game server (communication, logic, data storage)



### Quick game

-Website: [Quick Play Mini Games, Stand-alone Games, Web Games, Quick Play Games, Quick Play Game Boxes] (Quick Play Mini Games, Stand-alone Games, Web Games, Quick Play Games, Quick Play Game Boxes) 
-Application scope: real-time message system, user authentication, user conversation, unified statistical interface



### Shanda Cloud CDN

-Website: Shanda Cloud Computing
-Application scope: CDN scheduling system, distribution system, monitoring system, short domain name service, CDN internal open platform, operation report system and other small tools, etc.





### Bmob mobile back-end cloud service platform

-Product URL: Bmob mobile back-end cloud service platform
-Application scope: Restful API (using Beego), statistical analysis platform, common services such as sending emails, asynchronous queue processing, statistical user space and interface requests



### Group Strategy

-Website: [Group Strategy-Unified Team Communication, Work Efficiently] (Group Strategy-Unified Team Communication, Work Efficiently)
-Application range: the whole system



### BiddingX DSP advertising system

-Website: BiddingX_Professional DSP solution provider
-Application scope: bidding, exposure statistics, click to jump



### Neighborhoods

-Website: Homepage-Neighborhoods
-Application scope: background service



### Leanote

-Website: Leanote

### Bearychat

-Website: BearyChat

### Home Bean

-Website: Zhaidou.com-Build the most beautiful home by yourself, take Zhaidou as you like

### Whiteboard-Design Drawing Discussion Tool

-Website: Whiteboard

### Laboratory building

-Website: Laboratory Building-The first IT online education platform centered on experiments



### Sina Weibo

Middleware and flexible scheduling are written in Java and Go, and Weibo video transcoding and storage services are written in Go.



### IQIYI

VR back-end system middleware, HTTP interface on the VR side.



### Cheetah Mobile

News push



### NetEase

NetEase Honeycomb Container Public Cloud.



### Bilibili

Barrage



### Giant Network

The server of some mobile games.



### Today's headlines





Nsq: Nsq is a high-performance, high-availability message queuing system developed by the Go language. It has very high performance and can handle billions of messages every day;

Packer: used to generate image files for different platforms, such as VM, vbox, AWS, etc. The author is the author of vagrant

Skynet: Distributed scheduling framework

Doozer: Distributed synchronization tool, similar to ZooKeeper

Heka: mazila's open source log processing system

Cbfs: Couchbase's open source distributed file system

Tsuru: The open source PAAS platform has exactly the same functions as SAE

Groupcache: a caching system for Google download system written by the author of memcahe

God: A cache system similar to redis, but supports distribution and scalability

Gor: network traffic packet capture and replay tool



There are many more, such as Alibaba Middleware, Jumei Youpin, Gaosheng Holdings, Tantan, Douyu Live, Renrenche, AsiaInfo, Udesk, Fangfutong, Lucky Cat, Sany Group, Meishan, etc. The general choice is to choose the appropriate product system for your company, such as message push, monitoring, container, etc. Golang is particularly suitable for network concurrent services. This is his strength, so it is also preferred. For these items. As a large-scale project development language, Go language has been used in many large companies one after another, and even turned to Go development completely.



## Fourth, write at the end

Of course, whether a technology can be developed depends on three points. (The following views are quoted from https://www.cnblogs.com/qwangxiao/p/8318894.html)

-**Is there a better community. **Ecosystems like C, C++, Java, Python and JavaScript are very rich and popular. Especially communities where many commercial organizations participate are more popular, such as the Linux community.
-**Is there an industrial standard. **Like C, C++, Java, there are standardization organizations. Especially Java, which has also produced enterprise-level standards like J2EE in its architecture.
-**Is there one or more killer apps. **The killer applications of C, C++ and Java needless to say, even for PHP, it is not a good programming language, because it is the key technology in LAMP, the first killer solution in the Linux era. Therefore, it has also developed.

The above three points are very critical. New technologies only need to account for one or two of them to be very good, not to mention some technologies, such as Java, which account for all three points. Therefore, the development of Java is so it is good. Of course, in addition to the above three important points, there are other influencing factors, such as:

-**Whether the learning curve is low and whether it is fast to get started. **This is very important, and C++ is getting worse and worse at this point.
-**Is there a good development framework to improve development efficiency? **Such as: Java's Spring framework, C++'s STL, etc.
-**Whether there are one or more giant technology companies as backing. **For example: IBM, Sun behind Java and Linux...
-**Have you solved the pain points in software development? **For example: Java solves the memory management problems of C and C++.

Using these rulers to measure the Go language, we can clearly see:

-Go language is easy to learn;
-Go language solves the pain points of concurrent programming and writing low-level application development efficiency;
-Go language has Google, a world-class technology company behind;
-The killer application of the Go language is Docker, and the Docker ecosystem has completely burst in recent years.

Therefore, the future of the Go language is limitless. Of course, I personally think that Go may swallow a lot of C, C++, and Java projects. However, the main project that the Go language swallows should be the middle-level project, which is neither the very low-level nor the business layer.

> In other words, the Go language will not swallow the low-level C and C++ level, nor will it swallow the high-level projects such as the Java business layer. What the Go language can swallow must be projects on PaaS, such as some message caching middleware, service discovery, service agents, control systems, agents, log collection, etc. There are no complex business scenarios, nor can it reach the special bottom layer (such as operating systems). ) Software projects or tools at the middle platform layer. And C and C++ will be hit to the lower level, and Java will be hit to the upper business layer.

Okay, let's use the above ruler to measure the killer application Docker of the Go language. You will find that it is basically the same.

-Docker is easy to get started.
-Docker solves the environmental problems in operation and maintenance and the pain points of service scheduling.
-There are big companies in the Docker ecosystem to help. Such as Google.
-Docker produced the industry standard OCI.
-Docker's community and ecosystem have emerged like Java and Linux.
-...

Therefore, although Docker a few years ago had a lot of pits at the time, compared to these big factors, those small pits were not a problem. It just takes some time, and these small holes can be completely filled in the next 5-10 years.

Similarly, we can see that Kubernetes, as a key technology for service and container scheduling, will definitely be the final winner.

Finally, I want to talk about why we should enter these new technologies earlier, instead of waiting for these technologies to mature before entering. There are several reasons.

The development process of technology is very important. Because you can clearly see the development process of this new technology ecosystem. The biggest gain for us is not the technology itself, but the change of technology and the development of the industry.

From this, we have seen various very specific thoughts and ideas, which are more valuable than the technology itself. Because this not only allows us to rethink the technologies we have mastered and how to better solve existing problems, but it also allows me to see the future. Not only we have technical advantages, but this knowledge also gives us many more possibilities in our technical career.

These key new technologies will allow you to take advantage of the technology. These are very important for an individual or company that needs technical leadership.

A company or individual who can take advantage of technology will have greater influence than other companies or individuals. Once the future industry demand detonates, then the influence of the company or individual will form a relatively large moat, and can quickly generate economic benefits.



The application scope of Go has been expanding. Cloud computing, microservices, blockchain, and heavyweight projects written in Go are everywhere. The docker/kubernetes ecosystem, with hundreds/tens of millions of lines of code, basically dominates the cloud native application market. Last year's popular blockchain, Ethereum's geth, Bitcoin's btcd, and Lightning Network's lnd, were all developed in the Go language. Again, look at the ecology of various languages, maybe they are not as unbearable as you think. . . The Go language is indeed not "advanced" in design, but it is also another "pragmatic". In fact, go has been very popular both at home and abroad. Google is used a lot in foreign countries, and uber is also used. There is a famous Toutiao in China, which is well-known for hundreds of billions of visits per day. How many languages ​​do not have such a big application scenario in their lifetime.



Qianfeng Go language learning group: 784190273

Corresponding video address:

https://www.bilibili.com/video/av47467197

https://www.bilibili.com/video/av56018934/

Source code:

https://github.com/rubyhan1314/go_foundationhttps://www.bilibili.com/video/av47467197/?p=6

