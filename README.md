# gin项目结构怎么分？
相信这是不少小伙伴在刚接触gin都有的疑问。 gin的项目结构有很多，目前接触过的本质上上分为两种：
1. 面向过程。
2. 面向对象

**面向过程**
代码开发便捷。频繁的通过**包名+方法名**调用可能会在不注意的时候出现跨层调用的问题（service/controller可以直接跨server层调用到dao），这是所不希望的。


**面向对象**
面向对象调用方法大幅的降低了跨层调用的问题。但是如果采用Java框架形式的完全面向对象。需要为每个.go文件都创建一个对象。在调用其他.go文件的方法时，还需要再创建一个对象才能引用。这是比较繁琐的。


~~~java
public class ArticleManager extends AbstractPostsManager {
    @Resource
    private ArticleTypeRepository articleTypeRepository;
    @Resource
    private ArticleRepository articleRepository;
    @Resource
    private UserRepository userRepository;
    @Resource
    private UserFollowRepository userFollowRepository;
    @Resource
    private InformRepository informRepository;
    
    // 调用方法...
}
~~~

实习的时候发现部门的项目架构挺好用，于是就偷学了下来。
首先看一下项目结构
> ├─dao         持久层 数据库操作（mysql、redis）
> 
> ├─logger      日志
> 
> ├─middleware  中间件
> 
> ├─model       定义的模型
> 
> ├─pkg         公共包
> 
> │  ├─jwt
> 
> │  ├─response
> 
> │  └─settings
> 
> ├─server      业务层
> 
> └─service     服务器层（等同于控制层）

通过在 dao、server层定义结构体，在service上声明一个全局的对象，层层调用，以优化上面的问题。
注意：结构体属性都是**私有**
~~~ go
// dao
type Dao struct {
	db  *gorm.DB
	rdb *redis.Client
}

// server
type Server struct {
	dao    *dao.Dao
	single *singleflight.Group
	// kafka...
}

// service
var svc server.Server   
~~~

后续的调用都通过这个svc实现，具体的实现请看项目中的测试用例
