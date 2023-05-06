package middleware

import (
    "fmt"
    "time"

    "github.com/camry/g/glog"
    "github.com/labstack/echo/v4"
)

// EchoLogger 针对 echo 框架日志中间件。
func EchoLogger() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            start := time.Now()

            err := next(c)
            if err != nil {
                c.Error(err)
            }

            req := c.Request()
            res := c.Response()

            fields := []any{
                "remote_ip", c.RealIP(),
                "latency", time.Since(start).String(),
                "host", req.Host,
                "request", fmt.Sprintf("%s %s", req.Method, req.RequestURI),
                "status", res.Status,
                "size", res.Size,
                "user_agent", req.UserAgent(),
            }

            id := req.Header.Get(echo.HeaderXRequestID)
            if id == "" {
                id = res.Header().Get(echo.HeaderXRequestID)
                fields = append(fields, "request_id", id)
            }

            n := res.Status
            switch {
            case n >= 500:
                fields = append(fields, "msg", "Server error")
                glog.Errorw(fields...)
            case n >= 400:
                fields = append(fields, "msg", "Client error")
                glog.Warnw(fields...)
            case n >= 300:
                fields = append(fields, "msg", "Redirection")
                glog.Infow(fields...)
            default:
                fields = append(fields, "msg", "Success")
                glog.Infow(fields...)
            }

            return nil
        }
    }
}
