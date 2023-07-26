package server

import (
	"httpTempate/pkg/helper"
	"net/http"
	"net/url"
	"path"
)

func (s *Server) appealHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slogger := s.logger.With().Str("fn", "appealHandler").Logger()
		u, err := url.Parse(s.config.ApiTNG)
		if err != nil {
			slogger.Error().
				Str("url", s.config.ApiTNG).
				Msgf("unable to parse : %s", err.Error())
			helper.RespondMessage(w, r, http.StatusInternalServerError, "Opps Something Happened")
			return
		}
		authID := "1234"
		appealID := "11"
		if authID == "" || appealID == "" {
			slogger.Warn().
				Str("authID", authID).
				Str("appealID", appealID).
				Msg("Missing Data")
			helper.RespondMessage(w, r, http.StatusBadRequest, "Missing Data")
		}
		u.Path = path.Join(u.Path, "authorizations", authID)
		u.Path = path.Join(u.Path, "appeals", appealID)

		// appealDto := model.AppealBuilderDto{}
		// err := json.NewDecoder(req.Body).Decode(&appealDto)
		// appealJSON, err := json.Marshal(appealDto.Appeal)
		// req, err := http.NewRequest("PUT", u.String(), bytes.NewReader(appealJSON)))
		// req.Header.Add("Authorization", r.Header.Get("Authorization"))
		// req.Header.Set("Content-Type", "application/json")
		// resp, err := s.netClient.Do(req)
		// body, err := ioutil.ReadAll(resp.Body)
		// defer resp.Body.Close()
		// respondMessage(w, r, http.StatusOK, string(resp.Body))

		// JWTRequest
		// appealDto := model.AppealBuilderDto{}
		// jwt := &helper.JWTRequest {
		// 	URL: u.String(),
		// 	Method: http.MethodPut,
		// 	JWT: r.Header.Get("Authorization"),
		// 	Timeout: timeoutHttp * time.Second,
		// }
		// body, err := jwt(appealDto)
		// jwt := r.Header.Get("Authorization")
		// resp, err := jwtExecute(http.MethodPut, u.String(), jwt, timeoutHttp * time.Second)
		//

		helper.RespondMessage(w, r, http.StatusOK, u.String())
	}
}
