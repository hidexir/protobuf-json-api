package main

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/hidexir/protobuf-json-api/proto/pb"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/prefectures", func(c echo.Context) error {

		res := ToGetPrefecturesResponse()
		data, _ := proto.Marshal(&res)

		var hoge pb.GetUsersRes
		_ = proto.Unmarshal(data, &hoge)
		pp.Println(hoge)

		return c.Blob(http.StatusOK, "application/protobuf", data)
	})

	e.Start(":8090")
}

func ToGetPrefecturesResponse() (res pb.GetUsersRes) {
	user := []*pb.Users{
		{
			Id:   12,
			Name: "123",
		},
		{
			Id:   13,
			Name: "ksljfalksj",
		},
	}
	res = pb.GetUsersRes{
		Users: user,
	}
	return
}
