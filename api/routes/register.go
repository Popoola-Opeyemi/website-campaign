package routes

import (
	"api/models"
	"api/utils"
	"net/http"

	"github.com/labstack/echo"
)

// Register ...
type Register struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (s *Register) Init(prefix string, env *utils.Enviroment) error {
	s.Env = env
	s.Prefix = prefix

	req := s.Env.Echo.Group(prefix)
	req.POST("/register", s.Register)
	req.POST("/quote", s.Quotee)

	return nil
}

// Register...
func (s *Register) Register(c echo.Context) error {
	res := &utils.Response{}

	register := models.Customers{}

	if err := c.Bind(&register); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	// validating if all parameters are set
	if err := c.Validate(register); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	register.Type = 1

	if err := register.Create(s.Env.DB); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, res)

}

// Register...
func (s *Register) Quotee(c echo.Context) error {
	res := &utils.Response{}

	register := models.Customers{}

	if err := c.Bind(&register); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	// validating if all parameters are set
	if err := c.Validate(register); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	register.Type = 2

	if err := register.Create(s.Env.DB); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, res)

}
