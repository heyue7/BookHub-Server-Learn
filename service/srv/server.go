package srv

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	s := &Server{
		Engine: gin.New(),
	}

	s.Engine.NoRoute(s.noRoute)
	s.Engine.NoMethod(s.noMethod)

	s.Use(s.cors, s.noAccess)

	return s
}

func (s *Server) Run(name, addr string) {
	server := endless.NewServer(addr, s.Engine)

	// 使用 goroutine 来捕捉退出信号并优雅地关闭服务器
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Server panic: %v", r)
			}
		}()

		// 创建一个信号通道，用于接收系统的中断信号（如 SIGINT, SIGTERM）
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, os.Kill)

		// 等待接收到退出信号
		sig := <-sigChan
		log.Printf("Received signal: %s, shutting down server...\n", sig)

		if err := server.Close(); err != nil {
			log.Printf("Error while shutting down server: %v\n", err)
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Server panic: %v", r)
			}
		}()

		log.Printf("Server %s starting on %s. Press any key to stop.\n", name, addr)
		
		var str string
		fmt.Scanf("%s", &str)

		if err := server.Close(); err != nil {
			log.Printf("Error while shutting down server: %v\n", err)
		}
	}()

	log.Println(server.ListenAndServe())
}

// 跨域
func (s *Server) cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Next()
}

// 禁止访问
func (s *Server) noAccess(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	if c.Request.URL.Path == "/favicon.ico" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}

// 找不到路由
func (s *Server) noRoute(c *gin.Context) {
	s.abortWithStatus40X(c, 404, "Page not found")
}

// 找不到方法
func (s *Server) noMethod(c *gin.Context) {
	s.abortWithStatus40X(c, 405, "Method not allowed")
}

func (s *Server) abortWithStatus40X(c *gin.Context, code int32, msg string) {
	resp := Error(code, msg, msg)
	c.Set("__response", resp)
	c.AbortWithStatusJSON(int(code), resp)
}

func (s *Server) abortWithStatus50X(c *gin.Context, code int32, msg string) {
	resp := Error(code, msg, msg)
	c.Set("__response", resp)
	c.AbortWithStatusJSON(500, resp)
}
