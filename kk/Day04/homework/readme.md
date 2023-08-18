1. 整理知识
2. TODO LIST
   1. 认证
      1. 密码 =》 内置到程序当中
      2. 密码 =》 md5 + salt盐
      3. 打开程序
         1. 认证： 密码
         2. password + salt =》 hashed =》 程序内置的hash值比较
         3. 认证成功可以进行后续操作，如果连续输入密码超过3次失败直接退出程序。
      4. 输入密码： fmt.Scan()
         1. 为了不显示密码，使用一个第三方包，gopass
   2. sort
      1. 查询 输入输入排序方式（name：名字， startTime：开始时间） 升序排列

      第三方包：tablewriter
      |ID | 名字 | xxx | xxx|
      |xxx| xxx | xxx | xxx|