package downloader

import "context"

type SFTPClient struct {
	Host string
}

func (s *SFTPClient) Download(ctx context.Context, dir string, remoteDir string) (err error) {
	// TODO
	return nil
}
