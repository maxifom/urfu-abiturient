package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"urfu-abiturient-api/ent"
	"urfu-abiturient-api/ent/abituriententry"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		DBURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	client, err := ent.Open("postgres", DBURL)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger(), middleware.Recover(), middleware.Gzip())

	dictCache := NewDictionaryCache(GetDictFromDB(client))

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dict_cache", dictCache)
			return h(c)
		}
	})

	e.GET("/last_update_time", func(c echo.Context) error {
		return c.JSON(http.StatusOK, LastUpdateTimeResponse{LastUpdateTime: client.LastUpdated.GetX(context.TODO(), 1).LastUpdated.Format(time.RFC3339)})
	})
	e.POST("/query", func(c echo.Context) error {
		var request QueryRequest
		err := c.Bind(&request)
		if err != nil {
			return err
		}

		if request.Program == "" && len(request.Name) < 3 && request.Number == 0 {
			return echo.NewHTTPError(400, "program or name more that 3 symbols or number is required")
		}

		q := client.AbiturientEntry.Query().
			Order(ent.Desc(abituriententry.FieldSum))

		if len(request.Basis) > 0 {
			q = q.Where(abituriententry.BasisIn(request.Basis...))
		}

		if len(request.Form) > 0 {
			forms := make([]abituriententry.Form, 0, len(request.Form))
			for _, f := range request.Form {
				forms = append(forms, abituriententry.Form(f))
			}
			q = q.Where(abituriententry.FormIn(forms...))
		}

		if request.Sum.LTE != nil {
			q = q.Where(abituriententry.SumLTE(*request.Sum.LTE))
		}

		if request.Sum.GTE != nil {
			q = q.Where(abituriententry.SumGTE(*request.Sum.GTE))
		}

		if request.Program != "" {
			q = q.Where(abituriententry.Program(request.Program))
		}

		if request.StatementGiven {
			q = q.Where(abituriententry.StatementGiven(request.StatementGiven))
		}

		if request.OriginalGiven {
			q = q.Where(abituriententry.OriginalGiven(request.OriginalGiven))
		}

		if len(request.Type) > 0 {
			q = q.Where(abituriententry.TypeIn(request.Type...))
		}

		if len(request.Status) > 0 {
			q = q.Where(abituriententry.StatusIn(request.Status...))
		}

		if len(request.Name) > 0 {
			q = q.Where(abituriententry.NameContainsFold(request.Name))
		}

		if request.Number != 0 {
			q = q.Where(abituriententry.Number(request.Number))
		}

		entries := q.AllX(context.Background())

		return c.JSON(http.StatusOK, entries)
	})

	e.GET("/dictionary", func(c echo.Context) error {
		dict, needUpdate := c.Get("dict_cache").(*DictionaryCache).Get()
		if needUpdate {
			dict = GetDictFromDB(client)

			c.Get("dict_cache").(*DictionaryCache).Update(dict)
		}

		return c.JSON(http.StatusOK, dict)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
