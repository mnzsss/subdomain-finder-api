package handler

import (
	"net/http"
	"subdomain-finder-api/services"

	v "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := v.New()

	server.GET("/sub-domains", func(context *v.Context) {
		context.SetHeader("Content-Type", "application/json")

		domain := context.Query("domain")

		if domain == "" {
			context.JSON(400, v.H{
				"message": "domain is required",
			})
			return
		}

		output, err := services.Subfinder(domain)

		if err != nil {
			context.JSON(500, v.H{
				"message": "failed to enumerate subdomains",
				"error":   err.Error(),
			})
			return
		}

		context.JSON(200, v.H{
			"message":    "subdomains enumerated successfully",
			"subdomains": output,
		})
	})

	server.Handle(w, r)
}
