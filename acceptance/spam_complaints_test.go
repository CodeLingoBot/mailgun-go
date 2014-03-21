// +build acceptance

package acceptance

import (
	"github.com/mailgun/mailgun-go"
	"testing"
)

func TestGetComplaints(t *testing.T) {
	domain := reqEnv(t, "MG_DOMAIN")
	apiKey := reqEnv(t, "MG_API_KEY")
	publicApiKey := reqEnv(t, "MG_PUBLIC_API_KEY")
	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	n, complaints, err := mg.GetComplaints(-1, -1)
	if err != nil {
		t.Fatal(err)
	}
	if len(complaints) != n {
		t.Fatalf("Expected %d complaints; got %d", n, len(complaints))
	}
}

func TestGetComplaintFromBazNoComplaint(t *testing.T) {
	domain := reqEnv(t, "MG_DOMAIN")
	apiKey := reqEnv(t, "MG_API_KEY")
	publicApiKey := reqEnv(t, "MG_PUBLIC_API_KEY")
	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	complaint, err := mg.GetSingleComplaint("baz@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if complaint.Count != 0 {
		t.Fatal("Expected baz@example.com to have zero complaints.")
	}
}
