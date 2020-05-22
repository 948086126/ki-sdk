package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"

	m "ki-sdk/model"

	"github.com/gin-gonic/gin"
)

// 数据上链
func UpLoad(c *gin.Context) {
	//
	// 数据上链
	res, err := upLoad(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": "fail",
				"data":   res,
			})
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": "success",
				"data": res,
			})
	}
	return
}

// 数据上链
func upLoad(c *gin.Context) (result string, err error) {

	fmt.Println("数据上链操作")

	// 解析数据
	data, err := Serialize(c)
	if err != nil {
		log.Println("解析数据失败：", err)
	}

	//
	log.Println("序列化成功：", data)
	//数据上链
	result, err1 := UploadByChaincode(data.ChannelName, data.ChainCodeName, data.FunctionName, data.Data)
	// 调用上链
	if err != nil {
		log.Println("数据上链", err1)
	}

	return result, err
}

//gin// 数据查询
func Load(c *gin.Context) {

	fmt.Println("数据查询操作")
}

//
// 序列化 数据
func Serialize(c *gin.Context) (data *ChainDb, err error) {

	if err := c.ShouldBindJSON(&data); err != nil {
		return data, err
	}
	return data, nil
}

// 数据上链
func UploadByChaincode(channelName string, chaincodeName string, funcName string, args []string) (result string, err error) {
	//

	var peerlist []string
	// 02. 设置背书节点
	peerlist = append(peerlist, "peer0.org1.bookstore.com")
	peerlist = append(peerlist, "peer1.org2.bookstore.com")

	Pg, err := ArgsSplicing(args)

	//	拼接请求参数
	request := channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         funcName,
		Args:        Pg,
	}

	// 上链
	response, err := m.App.SDK.Client.Execute(request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peerlist...),
	)

	// 05. 判断错误
	if err != nil {
		//
		log.Println("err:", err)

		return "", err
	}
	// 06. 返回调用结果
	return string(response.Payload), nil

}

// 参数拼接
func ArgsSplicing(arg []string) (res [][]byte, err error) {
	//
	if len(arg) > 0 {
		for k, _ := range arg {
			res = append(res, []byte(arg[k]))
		}
		return res, nil
	} else {
		return nil, errors.New("arg is <= 0")
	}
}
