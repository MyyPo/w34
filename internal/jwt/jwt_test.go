package jwt

import (
	"testing"
	"time"
	// "github.com/MyyPo/w34.Go/configs"
)

func TestJWT(t *testing.T) {
	// conf, err := configs.NewConfig("$APP/app/configs")
	// if err != nil {
	// 	t.Errorf("err: %v", err)
	// }

	// t.Logf("current config: %v", conf)

	jwtManager, err := NewJWTManager("../../configs/rsa", "../../configs/rsa.pub",
		"../../configs/refresh_rsa", "../../configs/refresh_rsa.pub",
		time.Minute*10, time.Hour*48)
	if err != nil {
		t.Errorf("err initializing jwtManager: %v", err)
	}

	t.Run("Create an access token, then validate it", func(t *testing.T) {
		var userID int32 = 11

		gotJWT, err := jwtManager.GenerateAccessToken(userID)
		if err != nil {
			t.Errorf("jwt error: %q", err)
		}

		strUserID := "11"

		gotClaims, err := jwtManager.ValidateJwtExtractClaims(gotJWT, jwtManager.AccessPublicSignature)
		if err != nil {
			t.Errorf("failed to validate valid token: %q", err)
		}

		sub := gotClaims.Subject
		if sub != strUserID {
			t.Errorf("issued sub: %v, want id: %v", sub, strUserID)
		}
		if tknType := gotClaims.TknType; tknType != "access" {
			t.Errorf("issued tkn_type: %s, want tkn_type: %s", tknType, "access")
		}

	})
	t.Run("Try to validate a token with invalid signing method", func(t *testing.T) {
		hs256Token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

		_, err := jwtManager.ValidateJwtExtractClaims(hs256Token, jwtManager.AccessPublicSignature)
		if err == nil {
			t.Errorf("invalid token validated")
		}
	})
	t.Run("Create and validate a refresh token", func(t *testing.T) {
		const userID int32 = 5

		gotJWT, err := jwtManager.GenerateRefreshToken(userID)
		if err != nil {
			t.Errorf("jwt error: %q", err)
		}

		gotClaims, err := jwtManager.ValidateJwtExtractClaims(gotJWT, jwtManager.RefreshPublicSignature)
		if err != nil {
			t.Errorf("failed to validate valid token: %q", err)
		}

		if tknType := gotClaims.TknType; tknType != "refresh" {
			t.Errorf("issued tkn_type: %v, want tkn_type: %s", tknType, "refresh")
		}
		sub := gotClaims.Subject

		strUserID := "5"

		if sub != strUserID {
			t.Errorf("issued sub: %s, want id: %s", sub, strUserID)
		}

	})
	t.Run("Try to validate an expired access token", func(t *testing.T) {
		const expiredAccessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a25fdHlwZSI6ImFjY2VzcyIsInN1YiI6Ijg4IiwiZXhwIjoxNjcwNjY5NzA5LCJuYmYiOjE2NzA2NjkxMDksImlhdCI6MTY3MDY2OTEwOX0.zkB_Wkp_SRylHexumXluzlV8dEs4pSELuMfaz77TwBAcLtmhEWPBK-87grM9h3iEwPDAEErBZ1Dt4UqqHoP2WpaLEPQPcq3ib4hhCII42laFxKxIV9OUpe-j8ua92nDK8LN3NIg514BZoozET4O-q_4VeVskeoiA3VQrT0cfa7xspq8cjKwvAsOrgd6BvuWXrLhH4bLUFit6aEVGnuLVjz0NZjRQKnnsstOIuFIeBARhLMnYe5ZSmG_NhGXP016-_TZEdcdE0PCAm1QT6FZplgtKqJQJqO_qwI9qTsKJe2a9-DRYcks2q1v-QZL5Cv39JQKtgrS9cpRD6WpLrRV9fpxJICWpPNIGa6ym_RKu9yM_Oaf1LRYrOdBHJS2gXH7DyDggmjTpf3dQz-PM-9F-mS1Rjrb9d-Kf5VIN-lKjYkh5MIjnJGG1HnvFazUOGFZPTe_qCJF-aTxrk5tJywRsxh2pVcystxgZCI0VdIwB1r076NO4P3TLGtGhJ45_ptEJ_uXIcyqhlKe_TbOC6KFKct4Hy3Ra3HiWled9SQ9xRDOgy46bwa9G_NZ8-_pLPIEPxPmDx7GgoJEPTVx2328w_x0lIp8-6qgJyx0bSHbCuqxy2P1YQaAkxYocEuUSYpgjKAilEtHSA9UxyNO3WHUDkdSxycscQo-bhHOmLV9jY_M"

		_, err := jwtManager.ValidateJwtExtractClaims(expiredAccessToken, jwtManager.AccessPublicSignature)
		if err == nil {
			t.Errorf("didn't throw an error for an expired access token")
		}
	})
}
