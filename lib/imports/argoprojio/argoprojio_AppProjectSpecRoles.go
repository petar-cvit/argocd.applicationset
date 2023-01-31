// argoprojio
package argoprojio


// ProjectRole represents a role that has access to a project.
type AppProjectSpecRoles struct {
	// Name is a name for this role.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description is a description of the role.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Groups are a list of OIDC group claims bound to this role.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// JWTTokens are a list of generated JWT tokens bound to this role.
	JwtTokens *[]*AppProjectSpecRolesJwtTokens `field:"optional" json:"jwtTokens" yaml:"jwtTokens"`
	// Policies Stores a list of casbin formatted strings that define access policies for the role in the project.
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

