package route

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Code-Hex/battery-server/battery"
	"github.com/labstack/echo"
)

type BTInfo struct {
	Percent  int
	IsPowerd bool
}

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "🍣")
}

// Stream
func ShowBattery(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	for {
		var bt BTInfo
		var err error
		bt.Percent, bt.IsPowerd, err = battery.BatteryInfo()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if err := json.NewEncoder(c.Response()).Encode(bt); err != nil {
			return err
		}
		c.Response().(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}

	return nil
}
