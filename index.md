# gin用户注册登录认证处理

为创新实践后端完善所使用的学习过程

## 相关技术说明

- **Golang：一门新的程序设计语言**
- **Gin：Golang的一种框架**
- **GORM：对于Golang语言友好的一种开发人员ORM库**
- **ORM：一种对象关系映射，用来将对象和数据库之间的映射的元数据，将面向对象语言程序中的对象自动持久化到关系数据库中。 本质上就是将数据从一种形式转换到另外一种形式。**
- **JSON WEB Token**（**JWT**，读作 [/dʒɒt/]），是一种基于JSON的、用于在网络上声明某种主张的令牌（token）。JWT通常由三部分组成: 头信息（header）, 消息体（payload）和签名（signature）。

## 程序说明

这个程序是在网上看到一些教程以后，经过很多次的试错，终于用gin+gorm写出来的一个关于用户登录，用户认证，处理的一段代码

## 编写过程

简单注册认证->实现简单的登录->将数据用Gorm存入数据库当中->使用JWT实现登录认证->统一封装请求格式

## 相关代码

然后对于数据进行验证，比如电话号码只能十一位

```go
if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"Telephone number must be 11 numbers",
		})
		return
	}
	if len(password) <=6 {
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"Password must more than 6 words",
		})
		return
	}
	if isTelephoneExist(db,telephone) {
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"User is Already exist Please login",
		})
		return
	}
```

如果用户名不为零，则随机生成十位用户名

```go
func RandomString(n int) string{
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
```

```go
	hasedPassword ,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"message":"Internal Server Error",
		})
	}
	newUser:=model.User_info{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}
	db.Create(&newUser)
	log.Println(name,password,telephone)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"massage":"Register Successful",
	})
}
```

接下来是登录界面

下来将用户输入的密码加密以后，存储在数据库当中

首先和注册一样，登录界面需要验证手机号是否在数据库当中存在

```go
	//判断手机号是否存在
	var user model.User_info
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"Unknown User",
		})
	}
```

接下来判断用户输入的密码是否正确，首先对密码用哈希值加密

```go
//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"密码错误",
		})
		return
	}
```

登录的时候，我们使用`jwt` 认证，代码如下

```go
func ReleaseToken(user model.User_info) (string,error) {
	expirationTime :=time.Now().Add(7*24*time.Hour)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "govzzp.cn",
			Subject: "user token",
		},

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString , err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString,nil
}
```

## 结果查看

这个是原来数据库当中的内容

![1](https://pic.govzzp.cn/img/sql1.JPG)

接下来我们首先注册一个用户

填写数据

![](https://pic.govzzp.cn/img/regeist.JPG)

然后查看结果

![2](https://pic.govzzp.cn/img/res.JPG)

若再用相同的手机号发起请求

![4](https://pic.govzzp.cn/img/restwice.JPG)

接下来就是登陆了，首先输入错误登录数据

![](https://pic.govzzp.cn/img/logindata.JPG)

![5](https://pic.govzzp.cn/img/loginerr.JPG)

接下来正确的登录

![](https://pic.govzzp.cn/img/loginsuuccess.JPG)
