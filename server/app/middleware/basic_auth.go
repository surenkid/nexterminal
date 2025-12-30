package middleware

import (
	"encoding/base64"
	"next-terminal/server/config"
	"strings"

	"github.com/labstack/echo/v4"
)

// BasicAuth 实现 HTTP Basic Authentication
// 在所有请求之前进行验证，包括静态资源
func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 如果未启用 Basic Auth，直接放行
		if !config.GlobalCfg.Server.BasicAuthEnable {
			return next(c)
		}

		// 如果用户名或密码为空，直接放行（未配置）
		if config.GlobalCfg.Server.BasicAuthUser == "" || config.GlobalCfg.Server.BasicAuthPass == "" {
			return next(c)
		}

		// 获取 Authorization header
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return c.String(401, "Unauthorized")
		}

		// 解析 Basic Auth
		if !strings.HasPrefix(auth, "Basic ") {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return c.String(401, "Unauthorized")
		}

		// 解码 Base64
		encoded := strings.TrimPrefix(auth, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return c.String(401, "Unauthorized")
		}

		// 分割用户名和密码
		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return c.String(401, "Unauthorized")
		}

		username := credentials[0]
		password := credentials[1]

		// 验证用户名和密码
		if username != config.GlobalCfg.Server.BasicAuthUser || password != config.GlobalCfg.Server.BasicAuthPass {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return c.String(401, "Unauthorized")
		}

		// 验证通过，继续处理请求
		return next(c)
	}
}

