package blockchain

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn/resmgmtclient"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chmgmtclient"
	"time"
)

// 包括创建SDK需要的参数
type FabricSetup struct {
	//指定创建SDK时候所依赖的参数全部放在这个配置文件中，那么应用程序启动的时候。指定配置文件所在的路径。
	ConfigFile      string
	//通道名称
	ChannelID       string
	//是否被初始化成功
	Initialized     bool
	//通道配置文件，利用这个文件来创建通道
	ChannelConfig   string
	//组织管理员
	OrgAdmin        string
	//组织名称
	OrgName         string
	//资源管理客户端对象
	admin           resmgmtclient.ResourceMgmtClient
	//sdk对象
	sdk             *fabsdk.FabricSDK
}

//创建并初始化Fabric-SDK
func (t *FabricSetup) Initialize() error {
	fmt.Println("开始初始化SDK。。。")
	//判断SDK是否已经被初始化
	if t.Initialized {
		return fmt.Errorf("SDK已被实例化！")
	}
	//根据指定的SDK配置文件创建SDK对象
	sdk,err:=fabsdk.New(config.FromFile(t.ConfigFile))
	if err!=nil{
		return fmt.Errorf("SDK创建失败:%s",err)
	}
	//将获取到的SDK对象赋值到结构体中
	t.sdk=sdk
	/*
	1.根据指定的具有特权的用户创建用于管理通道的客户端API,t.sdk调用NewClient方法返回ChannelMgmt对象，NewClient里面必须指定当前
	组织具备最高管理权限的用户OrgAdmin,并指定该用户的组织名称
	*/
	chMgmtClient,err:=t.sdk.NewClient(fabsdk.WithUser(t.OrgAdmin),fabsdk.WithOrg(t.OrgName)).ChannelMgmt()
	if err!=nil{
		return fmt.Errorf("创建应用通道管理对象失败:%s",err)
	}
	//2.获取客户端的会话用户(目前只有session方法能够获取),session有下划线代表已过时，但是能用
	session,err:=t.sdk.NewClient(fabsdk.WithUser(t.OrgAdmin),fabsdk.WithOrg(t.OrgName)).Session()
	if err!=nil{
		return fmt.Errorf("获取会话用户失败 %s, %s: %s\n", t.OrgName, t.OrgAdmin, err)
	}
	orgAdminUser:=session
	//3.指定用于创建或更新通道的参数,获取到的是response对象，只有创建了通道后，我们才能安装链码，实例化链码，测试链码
	// 把创建的应用通道的参数对象必须赋值给一个变量，否则报错
	req:=chmgmtclient.SaveChannelRequest{ChannelID:t.ChannelID,ChannelConfig:t.ChannelConfig,SigningIdentity:orgAdminUser}
	//4.使用指定的参数创建或更新通道,把上面的req对象传递进去
	err =chMgmtClient.SaveChannel(req)
	if err!=nil{
		return fmt.Errorf("创建应用通道时发生错误:%s",err)
	}

		time.Sleep(time.Second*5)

	//创建一个用于管理系统资源的客户端API,有了这个用户，才能有权将peer加入通道，参数不需要指定OrgName了，因为是要管理当前网络所有资源，并不是特定的。只需要指定OrgAdmin就可以了
	t.admin,err=t.sdk.NewClient(fabsdk.WithUser(t.OrgAdmin)).ResourceMgmt()
	if err!=nil{
		return fmt.Errorf("创建资源管理客户端失败: %v\n", err)
	}

	// 将peer加入到指定的通道
	if err=t.admin.JoinChannel(t.ChannelID);err!=nil{
		return fmt.Errorf("peer节点加入通道失败: %v\n", err)
	}

	//设置SDK被初始化
	fmt.Println("初始化成功")
	t.Initialized=true
	return nil
}
