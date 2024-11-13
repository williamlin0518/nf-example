package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getChiikawaRoute() []Route {
	return []Route{
		{
			Name:    "Yee~ya~ha!",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: s.YeeYaHa,
			// Use
			// curl -X GET http://127.0.0.163:8000/chiikawa/ -w "\n"
		},
		{
			Name:    "Create Chiikawa Character",
			Method:  http.MethodPost,
			Pattern: "/print",
			APIFunc: s.PrintChiikawa,
			// Use
			// curl -X POST http://127.0.0.163:8000/chiikawa/print -w "\n"
		},
	}
}

func (s *Server) YeeYaHa(c *gin.Context) {
	c.String(http.StatusOK, "Yee~ya~ha!")
}

func (s *Server) PrintChiikawa(c *gin.Context) {
	c.String(http.StatusOK, "..............*,.............&.., /*...&... /*..................................\n" +
	"...............%.............&,.. .&..,%..  &...................,...............\n" +
	"................*.............%,.. @..(,.. %...................* ...............\n" +
	"................*.............,%.. .@.(/.. && ............ .%&&.................\n" +
	"...............//////*,. ..,#.,%@.. .#(*.  (&&&&/ . ....%,......../,............\n" +
	"......... .%  .  ..     %&,                        *&,     ... .... ..( ...# ...\n" +
	"/.. .. ,( ...      ..&.     @                  &       &              ... # ....\n" +
	" ..%%,             &      .(                     &       &.              *,(..  \n" +
	"     #           ,#       (   #@%        *&&%     .       #(           (. ....( \n" +
	"     ./          %           &&&&%       &&&&(             &         #.         \n" +
	"       #        @.      .*...         *     .*.//.,        (/      %            \n" +
	"        %       @.     ./,#..,.    %%&&&,   ,.*,.,..       /*    ,*             \n" +
	"         ,       %                   %*/                   &    .*.             \n" +
	"         .*.     .%                  ,/.                  %.       %            \n" +
	"           .,      %*                   ,%&#/*/%(        &  ,. ,*/((##,##%%%%#(/\n" +
	"%%%%%&&&%(/*,.     .*,&&             &,    ,&@(.   #&*%@        .               \n" +
	"..   ...               .@               %&(*/#&&%/  . &                         \n" +
	"                        &                             #,                   ..   ")
}
