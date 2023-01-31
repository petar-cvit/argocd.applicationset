// argoprojio
package argoprojio


// SignatureKey is the specification of a key required to verify commit signatures with.
type AppProjectSpecSignatureKeys struct {
	// The ID of the key in hexadecimal notation.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

