package secure

type Connection interface {
	PublicKey() [32]byte
	SetPeerKey(*[32]byte) error
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}
