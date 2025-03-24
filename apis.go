package main

import (
	"fmt"
	"net/http"

	_ "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

type ApiServer struct {
	listenaddr string
	store storage
}

func NewApiServer(listenaddr string,store storage) *ApiServer{
	return &ApiServer{
		listenaddr: listenaddr,
		store: store,
	}
}
func (s *ApiServer) Run(){
	app:=fiber.New()
	app.Get("/",s.Handledefault)
	app.Post("/shorten",s.HandleURLShort)
	app.Get("/get/:url",s.HandleGetURL)
	app.Listen(s.listenaddr)
	fmt.Sprintln("JSON API running on port:%s",s.listenaddr)

}
func(s *ApiServer) Handledefault(c *fiber.Ctx) error {
	return c.JSON("GIVA is a fine jewellery brand of silver, gold and lab grown diamonds. We have grown to be the largest D2C jewellery brand in India that has gained people’s trust as a go-to choice for gifting. GIVA's journey shines even brighter! GIVA strives to bridge the gap between online ease and in-store delight with its omnichannel experience")
}
func(s *ApiServer) HandleURLShort(c *fiber.Ctx) error{
	req:=ShortUrlReq{}
	if err:= c.BodyParser(&req); err != nil{
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if req.LongURL==""{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Long URL required"})
	}
    _,err:=s.store.GetlongUrl(req.LongURL)
	if err==nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "URL already Present"})
	}
	url:= NewUrl(req.LongURL)
	s.store.CreateShortUrl(url)
	return c.JSON(url)
}
func(s *ApiServer) HandleGetURL(c *fiber.Ctx) error{
    shortCode:=c.Params("url")
	url,err:=s.store.GetUrl(shortCode)
	if err!=nil{
		return err
	}
	return c.Redirect(url.LongURL,http.StatusFound)

}
