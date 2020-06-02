package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	m "ki-sdk/model"
	"log"
	"net/http"
)

func UploadSun(c *gin.Context) {

	res, err := upload(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"data":   res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   res,
		})
	}

}

// 序列化 数据
func Serialize2(c *gin.Context) (data *GoodsChaincode, err error) {

	if err := c.ShouldBindJSON(&data); err != nil {
		return data, err
	}
	return data, nil
}

func LoadSun(c *gin.Context) {

}
func upload(c *gin.Context) (res string, err error) {

	fmt.Println("数据上链")

	//解析数据

	data, err := Serialize2(c)
	if err != nil {
		log.Println("解析失败", err)
	}
	log.Print("解析成功")
	//todo  调用上链
	chaincode, err := uploadToChaincode(data.ChannelName, data.ChainCodeName, data.FunctionName, data.Data)
	// 调用上链
	if err != nil {
		log.Println("数据上链", err)
	}

	return chaincode, err
}

func uploadToChaincode(channelName string, chaincodeName string, functionName string, args []string) (result string, err error) {

	var peerList []string
	peerList = append(peerList, "peer0.org1.bookstore.com")
	peerList = append(peerList, "peer0.org2.bookstore.com")
	peerList = append(peerList, "peer1.org1.bookstore.com")
	peerList = append(peerList, "peer0.org2.bookstore.com")

	splicing, err := ArgsSplicing(args)

	request := channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         functionName,
		Args:        splicing,
	}
	response, err := m.App.SDK.Client.Execute(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peerList...),
	)
	if nil != err {
		return "", err
	} else {
		return string(response.Payload), nil
	}
}
