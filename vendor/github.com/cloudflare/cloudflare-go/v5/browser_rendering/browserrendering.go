// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// BrowserRenderingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserRenderingService] method instead.
type BrowserRenderingService struct {
	Options    []option.RequestOption
	Content    *ContentService
	PDF        *PDFService
	Scrape     *ScrapeService
	Screenshot *ScreenshotService
	Snapshot   *SnapshotService
	Json       *JsonService
	Links      *LinkService
	Markdown   *MarkdownService
}

// NewBrowserRenderingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBrowserRenderingService(opts ...option.RequestOption) (r *BrowserRenderingService) {
	r = &BrowserRenderingService{}
	r.Options = opts
	r.Content = NewContentService(opts...)
	r.PDF = NewPDFService(opts...)
	r.Scrape = NewScrapeService(opts...)
	r.Screenshot = NewScreenshotService(opts...)
	r.Snapshot = NewSnapshotService(opts...)
	r.Json = NewJsonService(opts...)
	r.Links = NewLinkService(opts...)
	r.Markdown = NewMarkdownService(opts...)
	return
}
