// Code generated by 'yaegi extract github.com/gotd/td/telegram/downloader'. DO NOT EDIT.

package yaegi

import (
	"context"
	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
	"reflect"
)

func init() {
	Symbols["github.com/gotd/td/telegram/downloader"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrHashMismatch": reflect.ValueOf(&downloader.ErrHashMismatch).Elem(),
		"NewDownloader":   reflect.ValueOf(downloader.NewDownloader),

		// type definitions
		"Builder":           reflect.ValueOf((*downloader.Builder)(nil)),
		"CDN":               reflect.ValueOf((*downloader.CDN)(nil)),
		"Client":            reflect.ValueOf((*downloader.Client)(nil)),
		"Downloader":        reflect.ValueOf((*downloader.Downloader)(nil)),
		"ExpiredTokenError": reflect.ValueOf((*downloader.ExpiredTokenError)(nil)),
		"RedirectError":     reflect.ValueOf((*downloader.RedirectError)(nil)),

		// interface wrapper definitions
		"_CDN":    reflect.ValueOf((*_github_com_gotd_td_telegram_downloader_CDN)(nil)),
		"_Client": reflect.ValueOf((*_github_com_gotd_td_telegram_downloader_Client)(nil)),
	}
}

// _github_com_gotd_td_telegram_downloader_CDN is an interface wrapper for CDN type
type _github_com_gotd_td_telegram_downloader_CDN struct {
	WUploadGetCDNFile func(ctx context.Context, request *tg.UploadGetCDNFileRequest) (tg.UploadCDNFileClass, error)
}

func (W _github_com_gotd_td_telegram_downloader_CDN) UploadGetCDNFile(ctx context.Context, request *tg.UploadGetCDNFileRequest) (tg.UploadCDNFileClass, error) {
	return W.WUploadGetCDNFile(ctx, request)
}

// _github_com_gotd_td_telegram_downloader_Client is an interface wrapper for Client type
type _github_com_gotd_td_telegram_downloader_Client struct {
	WUploadGetCDNFileHashes func(ctx context.Context, request *tg.UploadGetCDNFileHashesRequest) ([]tg.FileHash, error)
	WUploadGetFile          func(ctx context.Context, request *tg.UploadGetFileRequest) (tg.UploadFileClass, error)
	WUploadGetFileHashes    func(ctx context.Context, request *tg.UploadGetFileHashesRequest) ([]tg.FileHash, error)
	WUploadGetWebFile       func(ctx context.Context, request *tg.UploadGetWebFileRequest) (*tg.UploadWebFile, error)
	WUploadReuploadCDNFile  func(ctx context.Context, request *tg.UploadReuploadCDNFileRequest) ([]tg.FileHash, error)
}

func (W _github_com_gotd_td_telegram_downloader_Client) UploadGetCDNFileHashes(ctx context.Context, request *tg.UploadGetCDNFileHashesRequest) ([]tg.FileHash, error) {
	return W.WUploadGetCDNFileHashes(ctx, request)
}
func (W _github_com_gotd_td_telegram_downloader_Client) UploadGetFile(ctx context.Context, request *tg.UploadGetFileRequest) (tg.UploadFileClass, error) {
	return W.WUploadGetFile(ctx, request)
}
func (W _github_com_gotd_td_telegram_downloader_Client) UploadGetFileHashes(ctx context.Context, request *tg.UploadGetFileHashesRequest) ([]tg.FileHash, error) {
	return W.WUploadGetFileHashes(ctx, request)
}
func (W _github_com_gotd_td_telegram_downloader_Client) UploadGetWebFile(ctx context.Context, request *tg.UploadGetWebFileRequest) (*tg.UploadWebFile, error) {
	return W.WUploadGetWebFile(ctx, request)
}
func (W _github_com_gotd_td_telegram_downloader_Client) UploadReuploadCDNFile(ctx context.Context, request *tg.UploadReuploadCDNFileRequest) ([]tg.FileHash, error) {
	return W.WUploadReuploadCDNFile(ctx, request)
}