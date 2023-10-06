package utils

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Response[T any] struct {
	Code      int     `json:"code"`
	Message   string  `json:"message"`
	Data      T       `json:"data"`
	Timestamp string  `json:"timestamp"`
	Elapsed   float64 `json:"elapsed"`
}

func GetBody[T any](r *http.Request) (T, error) {
	var data T
	rawData, err := io.ReadAll(r.Body)

	if len(rawData) == 0 {
		return *new(T), fmt.Errorf("no data")
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(r.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetTokenFromRequest(req *http.Request) (string, error) {
	auth := req.Header.Get("authorization")
	if auth != "" {
		cred := strings.Split(auth, " ")
		if len(cred) > 1 && cred[0] == "Bearer" {
			if cred[1] != "" {
				return cred[1], nil
			}
			return "", fmt.Errorf("no token")
		}
		return "", fmt.Errorf("wrong auth header format")
	}
	return "", fmt.Errorf("wo auth header")
}

func SendResponse[T any](code int, data T, w http.ResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(bytes)
}

func SendResponseWithMessage[T any](ctx context.Context, code int, message *string, data T, w http.ResponseWriter) {
	now := time.Now()
	start := ctx.Value("start").(time.Time)

	if message == nil {
		message = new(string)
	}

	bytes, err := json.Marshal(&Response[T]{
		Code:      code,
		Message:   *message,
		Data:      data,
		Timestamp: now.Format("2006-01-02 15:04:05"),
		Elapsed:   now.Sub(start).Seconds(),
	})

	if err != nil {
		SendCustomError(ctx, http.StatusInternalServerError, app_errors.New(err, "internal server error", nil), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(bytes)
}

func SendError(code int, err string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write([]byte("{\"error\": \"" + err + "\"}"))
}

func SendCustomError(ctx context.Context, code int, err app_errors.AppError, w http.ResponseWriter) {
	now := time.Now()
	start := ctx.Value("start").(time.Time)

	bytes, marshalErr := json.Marshal(&Response[app_errors.AppError]{
		Code:      code,
		Message:   err.UIMessage,
		Data:      err,
		Timestamp: now.Format("2006-01-02 15:04:05"),
		Elapsed:   now.Sub(start).Seconds(),
	})

	if marshalErr != nil {
		SendCustomError(ctx, http.StatusInternalServerError, app_errors.New(marshalErr, "internal server error", nil), w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(bytes)
}

func SendHtml(code int, data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	_, _ = w.Write(data)
}

func IsEqualClientIp(serviceIp, clientIp string) (bool, error) {
	ip := ""
	u, err := url.Parse(serviceIp)
	if err != nil {
		return false, err
	}
	if u.Hostname() == "localhost" {
		ip = "127.0.0.1"
	} else {
		ip = u.Hostname()
	}
	c, err := url.Parse(strings.Split(clientIp, ":")[0])
	if err != nil {
		return false, err
	}
	return ip == c.String(), nil
}

func Jsonify(rows *sql.Rows) []string {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	c := 0
	results := make(map[string]interface{})
	data := []string{}

	for rows.Next() {
		if c > 0 {
			data = append(data, ",")
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for i, value := range values {
			switch value.(type) {
			case nil:
				results[columns[i]] = nil

			case []byte:
				s := string(value.([]byte))
				x, err := strconv.Atoi(s)

				if err != nil {
					results[columns[i]] = s
				} else {
					results[columns[i]] = x
				}

			default:
				results[columns[i]] = value
			}
		}

		b, _ := json.Marshal(results)
		data = append(data, strings.TrimSpace(string(b)))
		c++
	}

	return data
}

func GetReqContextWithTimeout(r *http.Request, sec int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(sec)*time.Second)
	r = r.WithContext(ctx)
	return r.Context(), cancel
}
