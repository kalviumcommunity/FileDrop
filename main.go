package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

type EmailCheckResult struct {
	Email     string   `json:"email"`
	Format    bool     `json:"format"`
	Domain    string   `json:"domain"`
	MX        bool     `json:"mx"`
	MXRecords []string `json:"mx_records"`
	SPF       bool     `json:"spf"`
	SPFRecord string   `json:"spf_record"`
	DMARC     bool     `json:"dmarc"`
	DMARCRecord string `json:"dmarc_record"`
	Validity  string   `json:"validity"`
}

func main() {
	router := gin.Default()
	router.GET("/check/:email", checkEmailHandler)
	log.Fatal(router.Run(":8080"))
}

func checkEmailHandler(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(400, gin.H{"error": "Email address is required"})
		return
	}

	result := CheckEmail(email)

	if !result.Format {
		c.JSON(400, gin.H{"error": "Enter a valid email address"})
		return
	}

	c.JSON(200, result)
}

func CheckEmail(email string) EmailCheckResult {
	result := EmailCheckResult{
		Email:  email,
		Format: validateEmailFormat(email),
	}

	if !result.Format {
		return result
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		result.Domain = ""
		result.Validity = "Invalid email address"
		return result
	}

	result.Domain = parts[1]

	mxRecords, err := lookupMX(result.Domain)
	if err != nil {
		log.Printf("Error looking up MX records for %s: %v\n", result.Domain, err)
	}

	result.MX = len(mxRecords) > 0
	result.MXRecords = mxRecords

	txtRecords, err := lookupTXT(result.Domain)
	if err != nil {
		log.Printf("Error looking up TXT records for %s: %v\n", result.Domain, err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			result.SPF = true
			result.SPFRecord = record
			break
		}
	}

	dmarcRecords, err := lookupTXT("_dmarc." + result.Domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for %s: %v\n", result.Domain, err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			result.DMARC = true
			result.DMARCRecord = record
			break
		}
	}

	result.Validity = determineValidity(result.MX, result.SPF, result.DMARC)

	return result
}

func validateEmailFormat(email string) bool {
	err := checkmail.ValidateFormat(email)
	return err == nil
}

func determineValidity(hasMX, hasSPF, hasDMARC bool) string {
	if hasMX && hasSPF && hasDMARC {
		return "Valid"
	} else if !hasMX && !hasSPF && !hasDMARC {
		return "Invalid"
	} else if (!hasMX && !hasSPF) || (!hasMX && !hasDMARC) || (!hasSPF && !hasDMARC) {
		return "Unlikely"
	} else {
		return "Likely"
	}
}

func lookupMX(domain string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mxRecords, err := net.DefaultResolver.LookupMX(ctx, domain)
	if err != nil {
		return nil, err
	}

	var records []string
	for _, mx := range mxRecords {
		records = append(records, mx.Host)
	}

	return records, nil
}

func lookupTXT(domain string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	txtRecords, err := net.DefaultResolver.LookupTXT(ctx, domain)
	if err != nil {
		return nil, err
	}

	return txtRecords, nil
}
