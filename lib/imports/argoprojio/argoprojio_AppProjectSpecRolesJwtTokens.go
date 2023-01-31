// argoprojio
package argoprojio


// JWTToken holds the issuedAt and expiresAt values of a token.
type AppProjectSpecRolesJwtTokens struct {
	Iat *float64 `field:"required" json:"iat" yaml:"iat"`
	Exp *float64 `field:"optional" json:"exp" yaml:"exp"`
	Id *string `field:"optional" json:"id" yaml:"id"`
}

