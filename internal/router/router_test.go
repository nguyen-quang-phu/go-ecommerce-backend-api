package router_test

import(
	"github.com/gin-gonic/gin"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/router"
	"testing"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := router.NewRouter()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
