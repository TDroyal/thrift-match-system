### this is my thrift-demo

#### 客户端go，服务器端为c++
- go下载官网：[官网](https://golang.google.cn/dl/)
- linux安装go参考教程：[参考教程](https://www.liwenzhou.com/posts/Go/install/)
- [c++ 开多线程服务器官网教程](https://github.com/apache/thrift/blob/master/tutorial/cpp/CppServer.cpp)

#### thrift在ubuntu上的安装过程：
```
wget https://dlcdn.apache.org/thrift/0.20.0/thrift-0.20.0.tar.gz  //或者直接去官网下载此压缩包
tar -zxvf thrift-0.20.0.tar.gz  //解压
```
#### 参考官网[官网](https://thrift.apache.org/docs/install/debian.html)
- The following command will install tools and libraries required to build and install the Apache Thrift compiler and C++ libraries on a Debian/Ubuntu Linux based system.
```
sudo apt-get install automake bison flex g++ git libboost-all-dev libevent-dev libssl-dev libtool make pkg-config
```
#### 参考官网[官网](https://thrift.apache.org/docs/BuildingFromSource)
```
./bootstrap.sh
./configure
make
make install
```

#### 检测安装结果
```
thrift -version
```


#### c++如何编译并链接
```
g++ -c *.cpp    //将所有的cpp文件编译为.o文件
g++ *.o -o main -lthrift -pthread   //动态链接加入thrift和thread,生成可执行文件main
./main         //启动匹配服务
```
- 上述链接过程可能报错：[参考文章](https://blog.csdn.net/vitaminc4/article/details/78707198)


