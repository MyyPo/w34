package auth

import (
	"context"
	"github.com/MyyPo/w34.Go/internal/adapters/psql"
	"testing"
)

func TestSignUp(t *testing.T) {
	t.Run("Saying heeey", func(t *testing.T) {
		rep := psql_adapters.NewPSQLRepository()
		impl := NewAuthServer(rep)
		got := impl.SignUp(context.Background())
		want := "Heeey"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
