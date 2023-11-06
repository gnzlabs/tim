package secure

type Connection interface {
	PublicKey() *[32]byte
	SetPeerKey(*[32]byte) error
	SharedKey() *[32]byte
	Encrypt(string) (string, error)
	Decrypt(string) ([]byte, error)
}
